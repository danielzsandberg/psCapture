package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v", err)
	}

	wrappedOutput := wrapOutput(string(output))

	fmt.Println(wrappedOutput)
}

func wrapOutput(psOutput string) string {
	currentTime := time.Now().String()

	firstLine := fmt.Sprintf("********%v*******\n", currentTime)
	buffer := bytes.NewBufferString(firstLine)

	buffer.WriteString(psOutput)
	buffer.WriteString("\n")

	return buffer.String()
}
