# KCM-CLI (KDE Config Manager)

KCM-CLI é uma ferramenta de terminal interativa escrita em Go para gerenciar perfis de configuração (snapshots) do KDE Plasma. Ele permite salvar todo o seu visual (temas, ícones, painéis, cursores, Klassy, wallpapers) e alternar entre eles instantaneamente.

## 🚀 Funcionalidades

- **Snapshots Completos**: Salva configurações (`~/.config`) e assets locais (`~/.local/share/icons`, `~/.icons`, temas Aurorae, etc.).
- **Interface Interativa (TUI)**: Navegue, aplique e delete perfis visualmente.
- **Rollback Automático**: Sempre cria um backup de segurança antes de aplicar um novo perfil.
- **Suporte a Symlinks**: Preserva links simbólicos em temas de ícones complexos (como Win11/WhiteSur).
- **Suporte ao Klassy**: Faz backup completo das configurações do Klassy Window Decoration.
- **Backup de Wallpaper**: Identifica e copia o wallpaper atual para o snapshot.

## 📦 Instalação (Arch Linux)

### 1. Dependências de Sistema
O KCM-CLI utiliza ferramentas nativas do KDE para ler e aplicar as configurações. No Arch Linux, instale:
```bash
sudo pacman -S go kde-cli-tools plasma-workspace
```

### 2. Clonar e Compilar
1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/winkde.git
   cd winkde
   ```
2. Baixe as dependências do Go:
   ```bash
   go mod tidy
   ```
3. Compile o binário:
   ```bash
   go build -o kcm ./cmd/kcm/*.go
   ```
4. (Opcional) Mova para o seu PATH:
   ```bash
   sudo mv kcm /usr/local/bin/
   ```

## 🛠️ Como Usar

### Modo Interativo (Recomendado)
A maneira mais fácil de usar o KCM-CLI é através da interface visual:
```bash
kcm ui
```
- **n**: Criar novo snapshot do estado atual.
- **ENTER**: Aplicar o snapshot selecionado.
- **d / x**: Deletar o snapshot selecionado.
- **q / ESC**: Sair.

### Comandos CLI
Você também pode usar comandos diretos:
- **Salvar**: `kcm save "Meu Tema Dark"`
- **Listar**: `kcm list`
- **Aplicar**: `kcm apply "Meu Tema Light"`

## 📂 Onde os backups são salvos?
Os perfis ficam armazenados em:
`~/.local/share/kcm-cli/profiles/`

---
Desenvolvido para facilitar a vida de quem gosta de customizar o KDE Plasma sem medo de perder suas configurações.
