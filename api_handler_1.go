package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "html"
    "io/ioutil"
    "io"
)

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!!\n")
	fmt.Fprintf(w, "The entered path is, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Endpoint hit : homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){   
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(Articles); err != nil {
        panic(err)
    }
    fmt.Println("Endpoint Hit: returnAllArticles")
}

func returnArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["key"]    
    fmt.Fprintf(w, "Number is %q\n",key)
    sel_article,err :=strconv.Atoi(key)
    if(err!=nil){
    	fmt.Println("Error")
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(Articles[sel_article-1]); err != nil {
        panic(err)
    }
    fmt.Println("Endpoint Hit: returnArticle")
}

func ArticleCreate(w http.ResponseWriter, r *http.Request) {
    var article Article
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &article); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) 
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    a := RepoCreate(article)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(a); err != nil {
        panic(err)
    }
}