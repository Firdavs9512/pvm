package php

type DarwinManager struct {
}

func (m *DarwinManager) Install(version string) error {
	return nil
}

func (m *DarwinManager) Uninstall(version string) error {
	return nil
}

func (m *DarwinManager) GetVersion() (string, error) {
	return "", nil
}

func (m *DarwinManager) GetInstalledVersions() ([]string, error) {
	return []string{}, nil
}

func (m *DarwinManager) Versions() ([]string, error) {
	return []string{}, nil
}
