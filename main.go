package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hello there!")
}


func main(){

	http.HandleFunc("/hello", hello)
	fmt.Println("Server up and running on port 8080!")
	http.ListenAndServe(":8080", nil)
	
	
}