package services

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Firdavs9512/pvm/core/utils"
)

type Brew struct {
}

const latestVersion = "php@8.3"

var limitedVersions = []string{
	"php@8.0",
	"php@7.4",
	"php@7.3",
	"php@7.2",
	"php@7.1",
}

var supportedVersions = []string{
	latestVersion,
	"php@8.2",
	"php@8.1",
	"php@8.0",
	"php@7.4",
	"php@7.3",
	"php@7.2",
	"php@7.1",
}

const BREW_PREFIX = "/opt/homebrew"

type BrewManager struct {
}

var shell = utils.NewShell("sh")

func (m *BrewManager) Installed(formula string) (bool, error) {
	output, err := shell.Run(fmt.Sprintf("brew list --formula %s", formula))
	if err != nil {
		return false, err
	}

	// If output starts with "Error: No such keg" then formula is not installed
	if !strings.HasPrefix(output, "Error: No") {
		return true, nil
	}

	details := make(map[string]interface{})

	if formulae, ok := details["formulae"].([]interface{}); ok && len(formulae) > 0 {
		if formula, ok := formulae[0].(map[string]interface{}); ok {
			if installed, ok := formula["installed"]; ok {
				return installed != nil, nil
			}
		}
	}

	if casks, ok := details["casks"].([]interface{}); ok && len(casks) > 0 {
		if cask, ok := casks[0].(map[string]interface{}); ok {
			if installed, ok := cask["installed"]; ok {
				return installed != nil, nil
			}
		}
	}

	return false, nil
}

// SupportedPhpVersions returns the list of PHP versions that are supported by brew
func (m *BrewManager) SupportedPhpVersions() []string {
	return supportedVersions
}

// LimitedPhpVersions returns the list of PHP versions that are supported by brew but not by pvm
func (m *BrewManager) LimitedPhpVersions() []string {
	return limitedVersions
}

// RestartServices restarts the given services
func (m *BrewManager) RestartServices(services []string) error {
	for _, service := range services {
		// Stop service
		shell.Run(fmt.Sprintf("brew services stop %s", service))
		// Start service
		shell.Run(fmt.Sprintf("brew services start %s", service))
	}

	return nil
}

// StopServices stops the given services
func (m *BrewManager) StopServices(services []string) error {
	for _, service := range services {
		// Stop service
		shell.Run(fmt.Sprintf("brew services stop %s", service))
	}

	return nil
}

// StartServices starts the given services
func (m *BrewManager) StartServices(services []string) error {
	for _, service := range services {
		shell.Run(fmt.Sprintf("brew services start %s", service))
	}

	return nil
}

// HasLinkedPhp determines if php is currently linked.
func (m *BrewManager) HasLinkedPhp() bool {
	_, err := os.Lstat(fmt.Sprintf("%s/bin/php", BREW_PREFIX))
	return err == nil
}

// LinkedPhpVersion determines which version of PHP is linked in Homebrew.
func (m *BrewManager) LinkedPhpVersion() (string, error) {
	if !m.HasLinkedPhp() {
		return "", errors.New("no php version linked")
	}

	// TODO: Implement this
	return "", nil
}
