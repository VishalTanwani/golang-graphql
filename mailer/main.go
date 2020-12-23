package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
	"time"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

type user struct {
	ID    string
	Email string
	Name  string
	Role  string
}
type issue struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Assignee    user
	Reporter    user
}

func main() {
	jsonData := map[string]string{
		"query": `
			{
				issues
				{
					id
					title
					type
					description
					status
					assignee
					{
						id
						email
						name
						role
					}
					reporter
					{
						name
					}
				}
			}
        `,
	}
	var issues []issue
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "http://localhost:5000/graphql", bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var at int
	for i, c := range string(data) {
		if c == '[' {
			at = i
			break
		}
	}
	s := string(data)[at : len(string(data))-3]
	err = json.Unmarshal([]byte(s), &issues)
	for _, v := range issues {
		if v.Status == "OPEN" || v.Status == "INPROGRESS" {
			message := "IssueID :- " + strconv.Itoa(v.ID) + " IssueTitle :- " + v.Title + " IssueDescription :- " + v.Description + " ReporterName :- " + v.Reporter.Name
			sendMail(v.Assignee.Email, message)
		}
	}
}

func sendMail(toUser, messageUser string) {
	// Sender data.
	from := "myuser@gmail.com"
	password := "MySecretPassword"
	// Receiver email address.
	to := []string{toUser}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte(messageUser)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
