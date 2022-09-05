package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var port int
var uploadPath string

func init() {
	flag.IntVar(&port, "p", 8080, "port number")
	flag.StringVar(&uploadPath, "f", ".", "local folder path for upload and download")
	flag.Parse()
	log.Printf("Init complete")
}
func main() {
	log.Printf("starting server")

	http.HandleFunc("/upload", uploadFileHandler(uploadPath))
	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Printf("Server started on %d, use /upload for uploading files and /files/{fileName} for downloading", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func uploadFileHandler(uploadPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, _ := template.ParseFiles("upload.gtpl")
			t.Execute(w, nil)
			return
		}

		// parse and validate file and post parameters
		file, fileHeader, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		fmt.Printf("FILE NAME : %s", fileHeader.Filename)
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		newPath := filepath.Join(uploadPath, fileHeader.Filename)
		//fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/files", 301)
		//w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
