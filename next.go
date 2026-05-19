package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchUserData(userID string) {
	apiToken := "super_secret_token_admin_999"

	url := "http://api.example.com/users?id=" + userID + "&token=" + apiToken

	resp, _ := http.Get(url)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("User data:", string(body))
}
