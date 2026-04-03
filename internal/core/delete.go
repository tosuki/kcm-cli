package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// DeleteProfile remove permanentemente um perfil e seus arquivos.
func DeleteProfile(profileName string) error {
	homeDir, _ := os.UserHomeDir()
	profileDir := filepath.Join(homeDir, ".local/share/kcm-cli/profiles", profileName)

	if _, err := os.Stat(profileDir); os.IsNotExist(err) {
		return fmt.Errorf("perfil '%s' não encontrado", profileName)
	}

	return os.RemoveAll(profileDir)
}
