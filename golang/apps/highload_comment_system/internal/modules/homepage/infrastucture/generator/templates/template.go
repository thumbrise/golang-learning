package templates

import (
	"embed"
	_ "embed"
)

//go:embed template.html
var FS embed.FS
