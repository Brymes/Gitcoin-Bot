package gitcoin

import (
	"Brymes-Gitcoin-Bot/utils"
	"io"
	"log"
	"net/http"
)

func makeRequest(url string, logger *log.Logger) []byte {
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	utils.OnErrorPanic(err, "Error Making Request", logger)

	res, err := client.Do(req)
	utils.OnErrorPanic(err, "Error Making Request", logger)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	utils.OnErrorPanic(err, "Error Reading Response", logger)

	return body
}
