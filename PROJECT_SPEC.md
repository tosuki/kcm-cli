# Project Spec: KCM-CLI (KDE Config Manager)

## 1. Visão Geral
Ferramenta CLI interativa em Go para gerenciar perfis de configuração do KDE Plasma (Snapshots). Permite capturar, listar, aplicar e deletar estados de customização do desktop.

## 2. Tech Stack & Dependências
- **Linguagem:** Go 1.21+
- **CLI Framework:** Cobra, Viper
- **TUI/Interface:** Bubble Tea, Lip Gloss, Bubbles
- **Requisitos de Sistema:** `kreadconfig5`, `plasmashell` (KDE Plasma 5/6).

## 3. Mapeamento de Configurações KDE (Crítico)
O KCM-CLI monitora e preserva os seguintes caminhos:

### Arquivos de Configuração (`~/.config/`)
- `kdeglobals`: Tema global, cores e ícones.
- `kwinrc`: Decoração de janelas e efeitos.
- `kcminputrc`: Mouse e teclado.
- `plasma-org.kde.plasma.desktop-appletsrc`: Layout de painéis e widgets.
- `plasmarc` / `plasmashellrc`: Configurações gerais do shell.
- `klassy/`: Configurações completas da decoração de janelas Klassy (pasta recursiva).

### Assets Locais (Cópia com Symlinks)
Para garantir fidelidade total, o script preserva links simbólicos:
- `~/.local/share/icons/`: Ícones instalados localmente.
- `~/.icons/`: Cursores e ícones (formato legado).
- `~/.local/share/plasma/look-and-feel/`: Temas de layout.
- `~/.local/share/aurorae/`: Temas de decoração de janela.
- `~/.local/share/color-schemes/`: Esquemas de cores.
- **Wallpaper**: Identificado via config e copiado fisicamente para `share/wallpaper.*`.

## 4. Estrutura de Armazenamento
Diretório: `~/.local/share/kcm-cli/profiles/[nome-do-perfil]/`
- `metadata.json`: Metadados do sistema e temas usados.
- `config/`: Espelhamento da pasta de configurações.
- `share/`: Espelhamento dos assets locais.

## 5. Fluxo de Operação
- **Save**: Conta arquivos, captura configurações via `kreadconfig5`, copia arquivos e assets (com barra de progresso).
- **Apply**: Cria rollback automático, restaura arquivos/assets e reinicia o `plasmashell`.
- **UI**: Interface interativa para gerenciar snapshots (Novo, Aplicar, Deletar).
