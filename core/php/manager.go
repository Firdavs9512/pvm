package php

import (
	"fmt"
	"runtime"
)

type Manager interface {
	Install(version string) error
	Uninstall(version string) error
	GetVersion() (string, error)
	GetInstalledVersions() ([]string, error)
	Versions() ([]string, error)
}

func NewManager() (Manager, error) {
	switch runtime.GOOS {
	case "darwin":
		return &DarwinManager{}, nil
	case "windows":
		return &WindowsManager{}, nil
	case "linux":
		return &LinuxManager{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
