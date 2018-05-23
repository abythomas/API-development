package main 
 import (
    "fmt"
    "net/http"
    "log"
    "html" 
    "encoding/json"
    "github.com/gorilla/mux"
    "strconv"
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
    json.NewEncoder(w).Encode(articles)
    fmt.Println("Endpoint Hit: returnAllArticles")
}

func returnArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["key"]
    articles := Articles{
        Article{Title: "Sample 1", Desc: "Sample Article Description 1", Content: "Sample Article Content 1"},
        Article{Title: "Sample 2", Desc: "Sample Article Description 2", Content: "Sample Article Content 2"},
    }    
    fmt.Fprintf(w, "Number is %q\n",key)
    sel_article,err :=strconv.Atoi(key)
    if(err!=nil){
    	fmt.Println("Error")
    }
    json.NewEncoder(w).Encode(articles[sel_article-1])
    fmt.Println("Endpoint Hit: returnAllArticles")
}

func handleRequest(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",homepage)
	router.HandleFunc("/articles/{key}",returnArticle)
	router.HandleFunc("/articles",returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",router))
}

func main(){
	handleRequest()
}
