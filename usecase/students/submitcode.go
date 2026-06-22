package student

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	StudentService "CWDLaunchPad/service/studentservice"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func CodeSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Decode payload safely by value to avoid nil pointer panic
	var submitCode *dto.SubmissionPayload
	if err := json.NewDecoder(r.Body).Decode(&submitCode); err != nil {
		fmt.Println("error decoding the submission payload")
		http.Error(w, "error decoding the submission payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Fetch Question Metadata to discover Class and Method signatures
	question, err := AdminService.GetQuestions(submitCode.QuestionID)
	if err != nil {
		fmt.Println("error fetching the question details")
		http.Error(w, "error fetching question data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Set up an isolated temporary workspace directory on the host machine
	tmpDir, err := os.MkdirTemp("", "java-runner-*")
	if err != nil {
		http.Error(w, "failed to create sandbox directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpDir) // Clean directory up when execution finishes

	// 5. Build dynamic Java orchestration main wrapper
	inputParsing, methodArgs := generateInputParsing(question[0].ParameterTypes, question[0].ParameterNames)

	printStatement := "System.out.println(result);"
	if strings.HasSuffix(question[0].ReturnType, "[]") {
		printStatement = "System.out.println(Arrays.toString(result));"
	}

	javaWrapperTemplate := `
		import java.util.*;

		public class Main {
			public static void main(String[] args) {
				Scanner sc = new Scanner(System.in);
				
				%s
				%s solver = new %s();
				%s result = solver.%s(%s);
				%s
			}
		}
		%s
		`
	fullJavaCode := fmt.Sprintf(
		javaWrapperTemplate,
		inputParsing,
		question[0].ClassName, question[0].ClassName,
		question[0].ReturnType, question[0].MethodName, methodArgs,
		printStatement,
		submitCode.SourceCode, // Appends student code as a non-public class at the bottom
	)

	// Write code to a local workspace file
	javaFilePath := filepath.Join(tmpDir, "Main.java")
	if err := os.WriteFile(javaFilePath, []byte(fullJavaCode), 0644); err != nil {
		http.Error(w, "failed to write source file", http.StatusInternalServerError)
		return
	}

	// 6. Pre-compile code using 'javac' once to verify syntax before passing to Docker
	compileCmd := exec.Command("javac", "--release", "17", javaFilePath)
	var compileErr bytes.Buffer
	compileCmd.Stderr = &compileErr
	if err := compileCmd.Run(); err != nil {
		http.Error(w, "Compilation Error:\n"+compileErr.String(), http.StatusBadRequest)
		return
	}

	passedTestcasesCount := 0

	var testCaseErr string
	// 7. Iterate and evaluate test cases sequentially
	for _, t := range question[0].TestCases {
		// Strict Time Limit Exceeded (TLE) protection: 3 seconds per testcase
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		cmd := exec.CommandContext(ctx, "docker", "run", "-i", "--rm",
			"-v", fmt.Sprintf("%s:/app", tmpDir),
			"-w", "/app",
			"eclipse-temurin:17-jdk",
			"java", "Main",
		)

		// Pipe input from Database directly through standard input stream
		cmd.Stdin = strings.NewReader(t.InputData)

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		cancel() // Release context resources immediately

		if err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Printf("Testcase %d: Time Limit Exceeded (TLE)\n", t.ID)
				continue
			}
			fmt.Printf("Runtime Error on testcase %d: %s\n", t.ID, stderr.String())
			continue
		}

		// Sanitize whitespaces/newlines on both sides for exact string verification
		actualOutput := strings.TrimSpace(stdout.String())
		expectedOutput := strings.TrimSpace(t.ExpectedData)

		if actualOutput == expectedOutput {
			passedTestcasesCount++
		} else {
			testCaseErr = fmt.Sprintf("Testcase %d Failed. Expected: %q, Got: %q\n", t.ID, expectedOutput, actualOutput)
		}
	}

	response := &dto.SubmitCodeResult{}

	submitCode.PassedTestcases = passedTestcasesCount
	submitCode.TotalTestcases = len(question[0].TestCases)

	// Store the submission information in DB for later results use
	if err := StudentService.InsertSubmission(submitCode); err != nil {
		http.Error(w, "error inserting submission"+err.Error(), http.StatusBadRequest)
		return
	}

	// 8. Construct response payload
	response.PassedTestcases = passedTestcasesCount
	response.TotalTestcases = len(question[0].TestCases)
	response.Message = "code executed successfully"
	response.StatusCode = http.StatusOK
	response.Error = testCaseErr

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// generateInputParsing creates Java Scanner code based on the parameter metadata
func generateInputParsing(paramTypes string, paramNames string) (string, string) {
	if paramTypes == "" || paramNames == "" {
		return "", ""
	}

	types := strings.Split(paramTypes, ",")
	names := strings.Split(paramNames, ",")

	var parsingCode strings.Builder
	var args []string

	for i, t := range types {
		if i >= len(names) {
			break
		}
		typeName := strings.TrimSpace(t)
		varName := strings.TrimSpace(names[i])
		args = append(args, varName)

		switch typeName {
		case "int":
			parsingCode.WriteString(fmt.Sprintf("        int %s = sc.nextInt();\n", varName))
		case "double":
			parsingCode.WriteString(fmt.Sprintf("        double %s = sc.nextDouble();\n", varName))
		case "String":
			parsingCode.WriteString(fmt.Sprintf("        String %s = sc.next();\n", varName))
		case "int[]":
			// Input format requirement: Size of array first, followed by elements
			parsingCode.WriteString(fmt.Sprintf("        int n_%s = sc.nextInt();\n", varName))
			parsingCode.WriteString(fmt.Sprintf("        int[] %s = new int[n_%s];\n", varName, varName))
			parsingCode.WriteString(fmt.Sprintf("        for(int i=0; i<n_%s; i++) { %s[i] = sc.nextInt(); }\n", varName, varName))
		default:
			parsingCode.WriteString(fmt.Sprintf("        String %s = sc.next();\n", varName))
		}
	}

	return parsingCode.String(), strings.Join(args, ", ")
}
