package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const ApiURL string = "https://api.spacetraders.io/v2"

func CallApi[T any](url string, method string, body []byte) (*T, error) {
	client := &http.Client{}
	uri := ApiURL + url
	// Create the request
	var body_to_send *bytes.Buffer = nil
	if body != nil {
		body_to_send = bytes.NewBuffer(body)
	}
	req, newreq_err := http.NewRequest(method, uri, body_to_send)
	if newreq_err != nil {
		return nil, newreq_err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN"))
	// Make the request
	res, call_err := client.Do(req)
	if call_err != nil {
		fmt.Println("API call error", call_err)
		return nil, call_err
	}
	// Handle response
	res_body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Error while reading body", error)
		return nil, error
	}
	res.Body.Close()
	var return_body T
	err := json.Unmarshal(res_body, &return_body)
	if err != nil {
		fmt.Println("Error while parsing body", err)
		return nil, err
	}
	return &return_body, nil
}
