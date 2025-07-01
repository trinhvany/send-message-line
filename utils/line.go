package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"small_demo_go/system"
)

func PushLineMessage(body string) error {
	req, err := http.NewRequest("POST", system.LinePushURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + system.Secret_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Request sent, status:", resp.StatusCode)
	return nil
}
