package main

import (
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
    "encoding/json"
    "strings"
    "path"
    "log"
    "os"
    "fmt"
)

var db *sql.DB

func init() {
    connStr := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?connect_timeout=10&sslmode=disable",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/css/", serveStatic("", "text/css"))
    http.HandleFunc("/js/", serveStatic("", "application/javascript"))
    http.HandleFunc("/", serveStatic("/index.html", "text/html"))
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.HandleFunc("/api/asteroids_scores", handleAsteroidsAPI)
    log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
    defer db.Close()
}

func serveStatic(filename, contentType string) func(http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        var filepath string
        if filename == "" {
            filepath = path.Join("client/dist/", req.URL.Path)
        } else {
            filepath = path.Join("client/dist/", filename)
        }

        res.Header().Set("Content-Type", contentType)
        http.ServeFile(res, req, filepath)
    }
}

func handleAsteroidsAPI(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        getAsteroidsScores(res, req)
    case "POST":
        postAsteroidsScore(res, req)
    }
}

func getAsteroidsScores(res http.ResponseWriter, req *http.Request) {
    queryText := "SELECT score, name FROM asteroids_scores ORDER BY score DESC"
    var queryValues []interface{}

    if limit := req.URL.Query().Get("limit"); limit != "" {
        queryText += " LIMIT $1"
        queryValues = append(queryValues, limit)
    }

    rows, err := db.Query(queryText, queryValues...)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var (
        responseData []map[string]interface{}
        score int
        name string
    )
    for rows.Next() {
        if err := rows.Scan(&score, &name); err != nil {
            log.Fatal(err)
        }

        record := map[string]interface{} {
            "score": score,
            "name": name,
        }

        responseData = append(responseData, record)
    }

    res.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(res).Encode(responseData)
    if err != nil {
        log.Fatal(err)
    }
}

func postAsteroidsScore(res http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var data struct {
        Score int
        Name string
    }
    err := decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }

    if len(data.Name) > 3 {
        res.WriteHeader(http.StatusBadRequest)
        res.Header().Set("Content-Type", "application/json")
        msg := `{"error": "Name cannot exceed 3 characters"}`
        res.Write([]byte(msg))
        return
    }

    queryText := "INSERT INTO asteroids_scores VALUES ($1, $2, DEFAULT)"
    queryValues := []interface{} { data.Score, strings.ToUpper(data.Name) }

    _, err = db.Exec(queryText, queryValues...)
    if err != nil {
        log.Fatal(err)
    }

    res.WriteHeader(http.StatusNoContent)
}
