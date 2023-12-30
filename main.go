package main

import (
	"fmt"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w,"404 Not Found",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"Method Not Supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Println(w, "ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful\n")
	name:= r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Address = %s\n",address)
}

func main(){
	fmt.Println("Go Server")
	fileServer := http.FileServer(http.Dir("./Static"))

	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("Server Started At Port 8080")
	if err:= http.ListenAndServe(":8080",nil); err !=nil {
		log.Fatal(err)
	}
}