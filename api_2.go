package main 
 import (
    "fmt"
    "net/http"
    "log"
    "html" 
    "encoding/json"
    "github.com/gorilla/mux"
    )

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!!\n")
	fmt.Fprintf(w, "The entered path is, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Endpoint hit : homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    articles := Articles{
        Article{Title: "Sample 1", Desc: "Sample Article Description 1", Content: "Sample Article Content 1"},
        Article{Title: "Sample 2", Desc: "Sample Article Description 2", Content: "Sample Article Content 2"},
    }    
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(articles)
}

func handleRequest(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",homepage)
	router.HandleFunc("/article",returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",router))
}

func main(){
	handleRequest()
}
