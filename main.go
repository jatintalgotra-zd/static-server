package main

import (
	"net/http"
	"path/filepath"

	"gofr.dev/pkg/gofr"
)

const defaultStaticFilePath = `./static`

func main() {
	app := gofr.New()

	staticFilePath := app.Config.GetOrDefault("STATIC_DIR_PATH", defaultStaticFilePath)

	app.UseMiddleware(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			filePath := filepath.Join(staticFilePath, "index.html")
			http.ServeFile(w, r, filePath)
		})
	})

	app.AddStaticFiles("/", staticFilePath)

	app.Run()
}
