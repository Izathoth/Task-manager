package main

import (
    "log"
    "net/http"
)

func main() {
    initDatabase()
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/tasks", tasksHandler)
    log.Println("Servidor iniciado na porta :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
