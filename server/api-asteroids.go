package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func handleAsteroidsAPI(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)

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

	responseData := make([]map[string]interface{}, 0)
	var (
		score int
		name  string
	)
	for rows.Next() {
		if err := rows.Scan(&score, &name); err != nil {
			log.Fatal(err)
		}

		record := map[string]interface{}{
			"score": score,
			"name":  name,
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
		Name  string
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
	queryValues := []interface{}{data.Score, strings.ToUpper(data.Name)}

	_, err = db.Exec(queryText, queryValues...)
	if err != nil {
		log.Fatal(err)
	}

	res.WriteHeader(http.StatusNoContent)
}

func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "GET, POST")
	(*res).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
