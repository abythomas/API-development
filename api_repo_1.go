package main

import "fmt"

var currentId int
var article Article

func init() {
    RepoCreate(Article{Title: "Lost in time",Desc: "cfygvbhnjkm;ba",Content :"dvtwvfyuwfvqwfvqwfvqwvfuvfyqwfuv"})
    RepoCreate(Article{Title: "Hopeless",Desc :"acshdabkjslkncajbc", Content: "weibfyewvfqvfoqvwfq8owfyqw8"})
}

func RepoFindArticle(id int) Article {
    for _, a := range Articles {
        if a.Id == id {
            return a
        }
    }
    return Article{}
}

func RepoCreate(a Article) Article {
    currentId += 1
    a.Id = currentId
    Articles = append(Articles, a)
    return a
}

func RepoDestroyTodo(id int) error {
    for i, a := range Articles {
        if a.Id == id {
            Articles = append(Articles[:i], Articles[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find article with id of %d to delete", id)
}