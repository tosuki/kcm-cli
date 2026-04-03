package config

import "time"

type ProfileMetadata struct {
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	PlasmaVersion string    `json:"plasma_version"`
	GlobalTheme   string    `json:"global_theme"`
	IconTheme     string    `json:"icon_theme"`
	HasCustomWall bool      `json:"has_custom_wallpaper"`
}
