package demo

import (
	"fmt"
	"os/exec"
)

func Test1() {
	cmd := exec.Command("cmd", "-c", "sleep 5;ls")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
