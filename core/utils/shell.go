package utils

import "os/exec"

type Shell struct {
	shell string
}

func NewShell(shell string) *Shell {
	return &Shell{shell: shell}
}

func (s *Shell) Run(command string) (string, error) {
	output, err := exec.Command(s.shell, "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
