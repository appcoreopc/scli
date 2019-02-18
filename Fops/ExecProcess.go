package Fops

import (
	"fmt"
	"log"
	"os/exec"
)

type ExecProcess struct {
}

func (ep *ExecProcess) ExecuteCommand(commandPath, arg string) bool {

	fmt.Println("command executed ", commandPath, arg)
	cmdResult := exec.Command(commandPath, arg)

	out, err := cmdResult.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with '%s'\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))

	return false
}
