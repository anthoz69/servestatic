package servestatic

import (
	"net/http"
	"os"
)

type serveStatic struct {
	http.FileSystem
}

// New is create handler for serve static.
func New(file string) http.Handler {
	return http.FileServer(&serveStatic{http.Dir(file)})
}

func (fs *serveStatic) Open(name string) (http.File, error) {
	file, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, os.ErrNotExist
	}
	return file, nil
}
