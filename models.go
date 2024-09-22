package main

type User struct {
    ID       int
    Username string
    Password string
}

type Task struct {
    ID       int
    UserID   int
    Title    string
    Completed bool
}
