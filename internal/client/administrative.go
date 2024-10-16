package client

import (
	"context"
	"io"
	"net/http"
)

func RegisterPlayer(ctx context.Context) (string, error) {
	newRequest, err := http.NewRequest("GET", "http://localhost:7000/register", nil)
	if err != nil {
		return "", nil
	}
	client := http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
