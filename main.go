package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
 if err:= r.ParseForm();err!=nil{
 fmt.Fprintf(w,"ParseForm() err: %v", err)
 return
 }

 fmt.Fprintf(w, "Post request success !!\n\n")

 name:=r.FormValue("name")
 email:=r.FormValue("email")

 fmt.Fprintf(w,"Name:%s\n", name)
 fmt.Fprintf(w, "Email:%s\n", email)

}

func helloHandler(w http.ResponseWriter, r *http.Request){
if(r.URL.Path!= "/hello"){
	http.Error(w,"404 Not Found",http.StatusNotFound)
	return
}
if(r.Method!="GET"){
	http.Error(w,"Method not supported", http.StatusNotFound)
	return
}

fmt.Fprintf(w,"Hello from the webserver")

}

func main(){
fileServer:= http.FileServer(http.Dir("./static")) //variable to store the directory of our static folder

http.Handle("/", fileServer)
http.HandleFunc("/form", formHandler) //formHandler is another function
http.HandleFunc("/hello", helloHandler) //helloHandler is again another function

fmt.Printf("Starting server at port 8080\n")
if err:=http.ListenAndServe(":8080", nil); err!=nil{
	log.Fatal(err)
}
}