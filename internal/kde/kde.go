package kde

import (
	"os/exec"
	"strings"
)

// ReadConfig executa o comando kreadconfig5 para ler uma configuração específica.
func ReadConfig(file, group, key string) (string, error) {
	args := []string{"--group", group, "--key", key}
	if file != "" {
		args = append([]string{"--file", file}, args...)
	}

	cmd := exec.Command("kreadconfig5", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// GetPlasmaVersion tenta obter a versão do Plasma via plasmashell.
func GetPlasmaVersion() (string, error) {
	cmd := exec.Command("plasmashell", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "unknown", err
	}

	return strings.TrimSpace(string(output)), nil
}
