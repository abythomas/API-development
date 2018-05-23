package main 
 import (
    "fmt"
    "net/http"
    "log"
    "html"
    )

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!!\n")
	fmt.Fprintf(w, "The entered path is, %q", html.EscapeString(r.URL.Path))     /* Because router is not used, localhost:8081/<anything> will give /<anything> as path*/
	fmt.Println("Endpoint hit : homepage")
}

func handleRequest(){
	http.HandleFunc("/",homepage)                   /*assigns homepage function to localhost:8081/hey */
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	handleRequest()
}
