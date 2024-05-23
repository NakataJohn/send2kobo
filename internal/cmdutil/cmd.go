package cmdutil

import (
	"bytes"
	"os/exec"
)

func RunCmd(cmdStr string) (string, error) {
	var stderr bytes.Buffer
	var out bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	} else {
		return out.String(), err
	}
}
