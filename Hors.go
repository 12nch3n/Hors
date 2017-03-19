package main
import (
    "io"
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/gorhill/cronexpr"
    "time"
    "encoding/json"
)

func Demo(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Gorilla!\n")
}

func PathDemo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    io.WriteString(w, "Gorilla!\n")
    fmt.Fprintf(w, "Path: %v \n", vars["file_path"])
}

type CRON_JOB struct{
    Cron        string `json:"cron"`
    ExecuteFile string `json:"execute_file"`
    Args        string `json:"args"`
}

func CronRegDemo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var cron CRON_JOB
    switch r.Method {
    case "POST":
        io.WriteString(w, "Gorilla!\n")
        decoder := json.NewDecoder(r.Body)
        decoder.Decode(&cron)
        cron.ExecuteFile = vars["file_path"]
        break
    case "GET":
        cron = CRON_JOB{Cron:"* */2 * * *", ExecuteFile:"demo_file.sh", Args: "arg1 arg2"}
        break
    }
    expr := cronexpr.MustParse(cron.Cron)
    nextTime := expr.Next(time.Now())
    fmt.Fprintf(w, "Path: %v \t Cron: %v \t Next Time: %v\n",
    vars["file_path"], cron.Cron, nextTime)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/refer/{file_path:.*}", Demo).Methods("GET")
    r.HandleFunc("/exec/{file_path:.*}", PathDemo).Methods("POST")
    r.HandleFunc("/cron/{file_path:.*}", CronRegDemo)
    r.HandleFunc("/cron/{file_path:.*}", CronRegDemo).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", r))
}
