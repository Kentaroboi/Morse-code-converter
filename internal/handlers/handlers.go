package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Index give to user our html form from index.html file.
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Upload here we parse form, get file from form, read it, convert data with service,
// write result to local file and give back the converted string.
func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// make file name from current time and keep extension of source file
	ext := filepath.Ext(header.Filename)
	safeTime := safeFileName(time.Now().UTC().String())
	name := safeTime + ext

	if err := os.WriteFile(name, []byte(result), 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}

// safeFileName change bad for Windows characters (mostly ":" from time) to "-".
func safeFileName(s string) string {
	replacer := strings.NewReplacer(":", "-")
	return replacer.Replace(s)
}
