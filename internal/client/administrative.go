package client

import (
	"context"
	"io"
	"net/http"
)

func RegisterPlayer(ctx context.Context) (string, error) {
	resp, err := http.Get("http://127.0.0.1:7000/register")
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
