package bashCommands

import "os/exec"

type BashCommands struct {
	command string
}

func (b BashCommands) runCommand() (bool, error) {
	cmd := exec.Command(b.command)
	_, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return true, nil
}
