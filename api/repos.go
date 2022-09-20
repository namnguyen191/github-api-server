package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	const NamReposUrl = "https://api.github.com/users/namnguyen191/repos"
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", NamReposUrl, nil)
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
