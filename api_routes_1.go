package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "homepage",
        "GET",
        "/",
        homepage,
    },
    Route{
        "returnAllArticles",
        "GET",
        "/articles",
        returnAllArticles,
    },
    Route{
        "returnArticle",
        "GET",
        "/articles/{key}",
        returnArticle,
    },
    Route{
        "ArticleCreate",
        "POST",
        "/articles",
        ArticleCreate,
    },
}