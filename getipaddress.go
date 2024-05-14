package ipdance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getIpAddressInfo() (*IPInfo, error) {

	url := "https://api.ip.dance/"

	userAgent := "golang-ip-dance-package-0.1"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	req.Header.Set("User-Agent", userAgent)

	// Add the Content-Type application/json header
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		description := fmt.Sprintf("Received bad status code %s.", resp.Status)
		return nil, errors.New(description)
	}

	var jsonData []byte
	jsonData, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var response IPInfo
	json.Unmarshal(jsonData, &response)
	return &response, nil

}
