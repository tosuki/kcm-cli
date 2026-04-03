package core

import (
	"encoding/json"
	"kcm-cli/pkg/config"
	"os"
	"path/filepath"
)

// ListProfiles retorna uma lista de todos os perfis salvos e seus metadados.
func ListProfiles() ([]config.ProfileMetadata, error) {
	homeDir, _ := os.UserHomeDir()
	profilesBaseDir := filepath.Join(homeDir, ".local/share/kcm-cli/profiles")

	entries, err := os.ReadDir(profilesBaseDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []config.ProfileMetadata{}, nil
		}
		return nil, err
	}

	var profiles []config.ProfileMetadata
	for _, entry := range entries {
		if entry.IsDir() {
			metaPath := filepath.Join(profilesBaseDir, entry.Name(), "metadata.json")
			metaFile, err := os.ReadFile(metaPath)
			if err != nil {
				continue
			}

			var meta config.ProfileMetadata
			if err := json.Unmarshal(metaFile, &meta); err == nil {
				profiles = append(profiles, meta)
			}
		}
	}

	return profiles, nil
}
