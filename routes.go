package main

import (
    "net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
    } else {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
    } else {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    }
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
    case http.MethodPost:
    default:
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    }
}
