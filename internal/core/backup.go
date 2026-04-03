package core

import (
	"encoding/json"
	"fmt"
	"kcm-cli/internal/kde"
	"kcm-cli/pkg/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveSnapshot captura o estado atual do KDE Plasma para um perfil nomeado.
func SaveSnapshot(profileName string) error {
	homeDir, _ := os.UserHomeDir()
	profilesDir := filepath.Join(homeDir, ".local/share/kcm-cli/profiles", profileName)

	if err := os.MkdirAll(filepath.Join(profilesDir, "config"), 0755); err != nil {
		return fmt.Errorf("falha ao criar pasta de config: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(profilesDir, "share"), 0755); err != nil {
		return fmt.Errorf("falha ao criar pasta de share: %w", err)
	}

	// 1. Coletar metadados
	globalTheme, _ := kde.ReadConfig("kdeglobals", "KDE", "LookAndFeelPackage")
	iconTheme, _ := kde.ReadConfig("kdeglobals", "Icons", "Theme")
	plasmaVer, _ := kde.GetPlasmaVersion()

	metadata := config.ProfileMetadata{
		Name:          profileName,
		CreatedAt:     time.Now(),
		PlasmaVersion: plasmaVer,
		GlobalTheme:   globalTheme,
		IconTheme:     iconTheme,
	}

	metaPath := filepath.Join(profilesDir, "metadata.json")
	metaFile, err := os.Create(metaPath)
	if err != nil {
		return fmt.Errorf("falha ao criar metadata.json: %w", err)
	}
	defer metaFile.Close()

	if err := json.NewEncoder(metaFile).Encode(metadata); err != nil {
		return fmt.Errorf("falha ao salvar metadata.json: %w", err)
	}

	// 2. Copiar arquivos de configuração
	configFiles := []string{
		"kdeglobals",
		"kwinrc",
		"kcminputrc",
		"plasmarc",
		"plasmashellrc",
		"plasma-org.kde.plasma.desktop-appletsrc",
	}

	for _, file := range configFiles {
		src := filepath.Join(homeDir, ".config", file)
		dst := filepath.Join(profilesDir, "config", file)
		if err := copyAny(src, dst); err != nil {
			fmt.Printf("! Aviso: falha ao copiar %s: %v\n", file, err)
		}
	}

	// 2.1 Copiar configurações específicas (Klassy, etc)
	specialConfigs := []string{
		"klassy",
	}
	for _, conf := range specialConfigs {
		src := filepath.Join(homeDir, ".config", conf)
		if _, err := os.Stat(src); err == nil {
			dst := filepath.Join(profilesDir, "config", conf)
			_ = copyDir(src, dst)
		}
	}

	// 3. Copiar assets locais
	localAssets := []struct {
		base string
		name string
	}{
		{base: ".local/share", name: "icons"},
		{base: ".local/share", name: "plasma/look-and-feel"},
		{base: ".local/share", name: "aurorae"},
		{base: ".local/share", name: "color-schemes"},
		{base: "", name: ".icons"}, // Cursors legados
	}

	for _, asset := range localAssets {
		src := filepath.Join(homeDir, asset.base, asset.name)
		if _, err := os.Stat(src); err == nil {
			dst := filepath.Join(profilesDir, "share", asset.name)
			fmt.Printf("Copiando asset: %s...\n", asset.name)
			if err := copyDir(src, dst); err != nil {
				fmt.Printf("! Aviso: falha ao copiar asset %s: %v\n", asset.name, err)
			}
		}
	}

	// 4. Copiar Wallpaper
	wallpaperPath, _ := kde.ReadConfig("plasma-org.kde.plasma.desktop-appletsrc", "Containments", "Image")
	if wallpaperPath != "" && strings.HasPrefix(wallpaperPath, "file://") {
		wallpaperPath = strings.TrimPrefix(wallpaperPath, "file://")
		if _, err := os.Stat(wallpaperPath); err == nil {
			dst := filepath.Join(profilesDir, "share", "wallpaper"+filepath.Ext(wallpaperPath))
			_ = copyFile(wallpaperPath, dst)
			fmt.Printf("✓ Wallpaper copiado: %s\n", filepath.Base(wallpaperPath))
		}
	}

	fmt.Printf("\n✓ Snapshot '%s' capturado com sucesso.\n", profileName)
	return nil
}
