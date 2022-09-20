package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		w.Write([]byte("Please specify the \"user\" param"))
		return
	}
	reposUrl := "https://api.github.com/users/" + user + "/repos?type=owner"

	sortBy := r.URL.Query().Get("sortBy")
	if sortBy != "" {
		reposUrl += "&sort=" + sortBy
	}

	sortDirection := r.URL.Query().Get("sortDirection")
	if sortDirection != "" {
		reposUrl += "&direction=" + sortDirection
	}

	pageLength := r.URL.Query().Get("pageLength")
	if pageLength != "" {
		reposUrl += "&per_page=" + pageLength
	}

	pageNumber := r.URL.Query().Get("pageNumber")
	if pageNumber != "" {
		reposUrl += "&page=" + pageNumber
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", reposUrl, nil)
	if err != nil {
		fmt.Println("Error constructing request: ", err)
		w.Write([]byte("Something went wrong"))
		return
	}

	token := os.Getenv("GITHUB_API_TOKEN")
	req.Header.Set("authorization", fmt.Sprintf("token %s", token))

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error fetching repos: ", err)
		w.Write([]byte("Something went wrong"))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing repos body: ", err)
		w.Write([]byte("Something went wrong"))
	}

	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
