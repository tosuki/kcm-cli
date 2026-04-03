package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func copyDir(src, dst string) error {
	var files []string
	_ = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		// Não adicionamos diretórios na lista de "arquivos para copiar conteúdo", 
		// mas garantimos que eles existam no destino.
		if info.IsDir() {
			relPath, _ := filepath.Rel(src, path)
			targetPath := filepath.Join(dst, relPath)
			_ = os.MkdirAll(targetPath, info.Mode())
			return nil
		}
		files = append(files, path)
		return nil
	})

	total := len(files)
	if total == 0 {
		return nil
	}

	for i, path := range files {
		relPath, _ := filepath.Rel(src, path)
		targetPath := filepath.Join(dst, relPath)
		
		if err := copyAny(path, targetPath); err != nil {
			fmt.Printf("\n! Erro ao copiar %s: %v\n", path, err)
		}

		percent := float64(i+1) / float64(total) * 100
		fmt.Printf("\r  Processando: [%-20s] %.0f%% (%d/%d)", 
			strings.Repeat("=", int(percent/5)), percent, i+1, total)
	}
	fmt.Println()
	return nil
}

// copyAny lida com arquivos regulares e links simbólicos
func copyAny(src, dst string) error {
	info, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Se for Link Simbólico, recria o link em vez de copiar o conteúdo
	if info.Mode()&os.ModeSymlink != 0 {
		linkTarget, err := os.Readlink(src)
		if err != nil {
			return err
		}
		// Remove o destino se já existir para evitar erro
		_ = os.Remove(dst)
		return os.Symlink(linkTarget, dst)
	}

	return copyFile(src, dst)
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
