package main
import (
    "io"
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func Demo(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Gorilla!\n")
}

func PathDemo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    io.WriteString(w, "Gorilla!\n")
    fmt.Fprintf(w, "Path: %v \n", vars["file_path"])
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/exec/{file_path:.*}", Demo).Methods("GET")
    r.HandleFunc("/exec/{file_path:.*}", PathDemo).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", r))
}
