package main

import (
	"net/http"
	"log"
)

func main() {

	servemux:=http.NewServeMux()

	servemux.HandleFunc("/", defaultpage)

	servemux.HandleFunc("/dog", dogpage)

	servemux.HandleFunc("/me/", mypages)

	LASerr:=http.ListenAndServe(":8080", servemux)
	if(LASerr!=nil){
		log.Fatalln(LASerr)
	}
}

func defaultpage(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("hi, this default page\n"))
}

func dogpage(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("hi, we are at dog page\n"))
}

func mypages(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("hi, i am Nippun\n"))
}