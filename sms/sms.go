package proovl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendSMS(user, token, route, from, to, text string) (string, error) {
	url := "https://www.proovl.com/api/send.php"

	postData := []byte(fmt.Sprintf("user=%s&token=%s&route=%s&from=%s&to=%s&text=%s", user, token, route, from, to, text))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	result := string(body)
	response := parseResponse(result)
	if response[0] == "Error" {
		return fmt.Sprintf("Error message: %s", response[1]), nil
	} else {
		return fmt.Sprintf("Message ID: %s; Status: %s", response[1], response[0]), nil
	}
}

func BulkSMS(user, token, route, from, to, text string) (string, error) {
	url := "https://www.proovl.com/api/send.php"

	toNumbers := strings.Split(to, ",")
	for _, number := range toNumbers {
		postData := []byte(fmt.Sprintf("user=%s&token=%s&route=%s&from=%s&to=%s&text=%s", user, token, route, from, number, text))
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
		if err != nil {
			return "", fmt.Errorf("Error creating request: %v", err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("Error sending request: %v", err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("Error reading response body: %v", err)
		}

		result := string(body)
		response := parseResponse(result)
		if response[0] == "Error" {
			return fmt.Sprintf("Error message: %s", response[1]), nil
		} else {
			fmt.Printf("Message ID: %s; Status: %s\n", response[1], response[0])
		}
	}
	return "Bulk SMS sent successfully", nil
}

func parseResponse(response string) []string {
	return strings.Split(response, ";")
}
