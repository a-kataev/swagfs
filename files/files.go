package files

import (
	"embed"
	"io/fs"
)

//go:embed dist/*.js dist/*.js.map dist/*.css dist/*.css.map dist/*.html dist/*.png
var distFS embed.FS

func DistFS() fs.FS {
	f, _ := fs.Sub(distFS, "dist")

	return f
}

//go:embed config/swagger-initializer.js
var configFile []byte

const ConfigName = "swagger-initializer.js"

func ConfigData() []byte {
	return configFile
}
