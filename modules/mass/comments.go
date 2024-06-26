package mass

import (
	"fmt"
	"os/exec"
	"strings"
)

func GenerateAIComments(prompt string) (string, error) {
	cmd := exec.Command("python", "modules/mass/gg.py", prompt)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	outputString := strings.TrimSpace(string(output))
	return outputString, nil
}
