package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	for {
		currentTime := <-ticker
		executeCommand(currentTime)
	}
}

func executeCommand(currentTime time.Time) {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v", err)
	}

	wrappedOutput := wrapOutput(string(output), currentTime)

	fmt.Println(wrappedOutput)
}

func wrapOutput(psOutput string, timeOfExecution time.Time) string {
	currentTime := timeOfExecution.String()

	firstLine := fmt.Sprintf("********%v*******\n", currentTime)
	buffer := bytes.NewBufferString(firstLine)

	buffer.WriteString(psOutput)
	buffer.WriteString("\n")

	return buffer.String()
}
