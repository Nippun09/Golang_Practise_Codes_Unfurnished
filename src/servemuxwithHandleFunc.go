

package main

import (
	"net/http"
	"log"
	"html/template"
)

func main() {
	

	servemux:=http.NewServeMux()

	servemux.HandleFunc("/", http.FileServer("C:/Users/Documents/go_workspace/src/web_teplate_prac"))

	LASerr:=http.ListenAndServe(":8080", servemux)
	if(LASerr!=nil){
		log.Fatalln(LASerr)
	}
}