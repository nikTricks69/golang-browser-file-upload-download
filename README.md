# golang-http-file-upload-download

 A simple example of an HTTP upload and download in Go 

 Start with

 ```
 go run main.go -p 8180 -f c://temp

-p port to listen to (8080 default)
-f the local folder to use for upload and download (default .)
 ```
 
 http://localhost:8080/upload

 # Docker

 docker build -t simple-http-file-server .

 docker run --rm  -p 8182:8182 -v c://temp:/shares  simple-http-file-server

