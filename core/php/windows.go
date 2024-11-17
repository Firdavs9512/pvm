package php

type WindowsManager struct {
}

func (m *WindowsManager) Install(version string) error {
	return nil
}

func (m *WindowsManager) Uninstall(version string) error {
	return nil
}

func (m *WindowsManager) GetVersion() (string, error) {
	return "", nil
}

func (m *WindowsManager) GetInstalledVersions() ([]string, error) {
	return []string{}, nil
}

func (m *WindowsManager) Versions() ([]string, error) {
	return []string{}, nil
}
