package db

import (
	"fmt"
	"net/smtp"
)

//WatcherPrint is structure for watcher
type WatcherPrint struct {
	ID      int
	IssueID int
	UserID  int
	User    *UserPrint
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

//AddWatcher is to add watcher in db
func AddWatcher(userID, issueID int) (*WatcherPrint, error) {
	var lastInsertID int
	err := DB.QueryRow("INSERT INTO watchers(issueid,usersid) VALUES($1,$2) returning id;", issueID, userID).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	fmt.Println("last inserted id =", lastInsertID)

	u := WatcherPrint{
		ID:      lastInsertID,
		IssueID: issueID,
		UserID:  userID,
	}

	u.User = GetUserByID(userID)
	return &u, nil
}

//LoadAllWatchers give all watchers
func LoadAllWatchers() ([]WatcherPrint, error) {
	row, err := DB.Query("SELECT id,issueid,usersid FROM watchers")
	checkErr(err)
	data := []WatcherPrint{}
	for row.Next() {
		var u WatcherPrint
		row.Scan(&u.ID, &u.IssueID, &u.UserID)
		u.User = GetUserByID(u.UserID)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return data, nil
}

//DeleteWatcher will delete watcher
func DeleteWatcher(id int) ([]WatcherPrint, error) {
	row, err := DB.Query("DELETE FROM watchers WHERE id = $1", id)
	checkErr(err)
	defer row.Close()
	return LoadAllWatchers()
}

//GetWatcherByIssueID will return watchers by issue id
func GetWatcherByIssueID(id int) *[]WatcherPrint {
	row, err := DB.Query("SELECT id,issueid,usersid FROM watchers WHERE issueid=$1", id)
	checkErr(err)
	data := []WatcherPrint{}
	for row.Next() {
		var u WatcherPrint
		row.Scan(&u.ID, &u.IssueID, &u.UserID)
		u.User = GetUserByID(u.UserID)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return &data
}

//SendMailToWatchers send mail to watchers
func SendMailToWatchers(id int) {
	// Sender data.
	from := "myuser@gmail.com"
	password := "MySecretPassword"
	// Receiver email address.
	to := []string{}
	data := *GetWatcherByIssueID(id)
	for _, value := range data {
		user := GetUserByID(value.UserID)
		to = append(to, user.Email)
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte("there is an new update in the issue")
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
