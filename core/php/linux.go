package php

import (
	"os/exec"
	"strings"
)

type LinuxManager struct {
}

func (m *LinuxManager) Install(version string) error {
	cmd := exec.Command("apt-get", "install", "-y", "php"+version)
	return cmd.Run()
}

func (m *LinuxManager) Uninstall(version string) error {
	return nil
}

func (m *LinuxManager) GetVersion() (string, error) {
	cmd := exec.Command("php", "-v")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	// PHP versiyasini parse qilish
	version := strings.Split(string(output), "\n")[0]
	return version, nil
}

func (m *LinuxManager) GetInstalledVersions() ([]string, error) {
	return []string{}, nil
}

func (m *LinuxManager) Versions() ([]string, error) {
	return []string{}, nil
}
