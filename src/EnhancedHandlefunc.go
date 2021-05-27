package main

import (
	"net/http"
	"log"
	"html/template"
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
	tpl, err:=template.ParseFiles("defaultpage.gohtml")
	if err!=nil{
		log.Println(err)
	}
	exerr:=tpl.Execute(resp, nil)
	if exerr!=nil{
		log.Println(exerr)
	}
	//resp.Write([]byte("hi, this default page\n"))
}

func dogpage(resp http.ResponseWriter, req *http.Request){
	tpl, err:=template.ParseFiles("dogpage.gohtml")
	if err!=nil{
		log.Println(err)
	}
	exerr:=tpl.Execute(resp, nil)
	if exerr!=nil{
		log.Println(exerr)
	}
	//resp.Write([]byte("hi, we are at dog page\n"))
}

func mypages(resp http.ResponseWriter, req *http.Request){
	tpl, err:=template.ParseFiles("mypage.gohtml")
	if err!=nil{
		log.Println(err)
	}
	data:="Nippun Bajaj"
	exerr:=tpl.Execute(resp, data)
	if exerr!=nil{
		log.Println(exerr)
	}
	//resp.Write([]byte("hi, i am Nippun\n"))
}