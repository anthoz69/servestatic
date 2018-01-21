package serverstatic

import (
	"net/http"
	"os"
)

type servestatic struct {
	http.FileSystem
}

// New is create handler for serve static.
func New(file string) http.Handler {
	return http.FileServer(&servestatic{http.Dir(file)})
}

func (fs *servestatic) Open(name string) (http.File, error) {
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
