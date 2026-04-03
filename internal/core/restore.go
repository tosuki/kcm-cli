package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// ApplySnapshot restaura um perfil de configuração e reinicia o shell do KDE.
func ApplySnapshot(profileName string) error {
	homeDir, _ := os.UserHomeDir()
	profileDir := filepath.Join(homeDir, ".local/share/kcm-cli/profiles", profileName)

	if _, err := os.Stat(profileDir); os.IsNotExist(err) {
		return fmt.Errorf("perfil '%s' não encontrado", profileName)
	}

	// 0. Criar Rollback automático do estado atual
	rollbackName := fmt.Sprintf("auto-rollback-%s", time.Now().Format("20060102-150405"))
	fmt.Printf("Criando rollback de segurança: %s\n", rollbackName)
	if err := SaveSnapshot(rollbackName); err != nil {
		return fmt.Errorf("falha ao criar rollback: %w", err)
	}

	// 1. Restaurar arquivos ~/.config
	configSourceDir := filepath.Join(profileDir, "config")
	entries, err := os.ReadDir(configSourceDir)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório de config: %w", err)
	}

	for _, entry := range entries {
		src := filepath.Join(configSourceDir, entry.Name())
		dst := filepath.Join(homeDir, ".config", entry.Name())
		
		info, _ := os.Stat(src)
		if info.IsDir() {
			_ = copyDir(src, dst)
		} else {
			if err := copyAny(src, dst); err != nil {
				fmt.Printf("! Aviso: falha ao restaurar %s: %v\n", entry.Name(), err)
			}
		}
	}

	// 1.5 Restaurar assets locais
	shareSourceDir := filepath.Join(profileDir, "share")
	if _, err := os.Stat(shareSourceDir); err == nil {
		shareEntries, _ := os.ReadDir(shareSourceDir)
		for _, entry := range shareEntries {
			src := filepath.Join(shareSourceDir, entry.Name())
			
			// Se for .icons (legado), restaura na home, senão na .local/share
			var dst string
			if entry.Name() == ".icons" {
				dst = filepath.Join(homeDir, ".icons")
			} else {
				dst = filepath.Join(homeDir, ".local/share", entry.Name())
			}

			if err := copyDir(src, dst); err != nil {
				fmt.Printf("! Aviso: falha ao restaurar asset %s: %v\n", entry.Name(), err)
			}
		}
	}

	// 2. Reiniciar o Plasma Shell (Importante para aplicar painéis e widgets)
	fmt.Println("Reiniciando o Shell do Plasma...")
	// Tenta encerrar o plasmashell
	_ = exec.Command("kquitapp5", "plasmashell").Run()
	
	// Inicia novamente em background
	cmd := exec.Command("kstart5", "plasmashell")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("falha ao reiniciar plasmashell: %w", err)
	}

	fmt.Printf("✓ Snapshot '%s' aplicado com sucesso.\n", profileName)
	return nil
}
