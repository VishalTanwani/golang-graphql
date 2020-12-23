package db

import (
	"fmt"
	"time"
)

//TimeLogPrint is sturct
type TimeLogPrint struct {
	ID            int
	Field         string
	TimeStamp     string
	PreviousValue string
	NewValue      string
	IssueID       int
}

//LoadAllLog will return all logs in issue
func LoadAllLog() ([]TimeLogPrint, error) {
	row, err := DB.Query("SELECT id,field,timeStamp,previousValue,newValue,issueID FROM time_log")
	checkErr(err)
	data := []TimeLogPrint{}
	for row.Next() {
		var u TimeLogPrint
		row.Scan(&u.ID, &u.Field, &u.IssueID, &u.NewValue, &u.PreviousValue, &u.TimeStamp)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return data, nil
}

//LoadAllLogByIssueID will return all logs in issue
func LoadAllLogByIssueID(id int) *[]TimeLogPrint {
	row, err := DB.Query("SELECT id,field,timeStamp,previousValue,newValue,issueID FROM time_log where issueID = $1", id)
	checkErr(err)
	data := []TimeLogPrint{}
	for row.Next() {
		var u TimeLogPrint
		row.Scan(&u.ID, &u.Field, &u.IssueID, &u.NewValue, &u.PreviousValue, &u.TimeStamp)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return &data
}

//AddTimeLog add log in issue
func AddTimeLog(Field, PreviousValue, NewValue string, IssueID int) (*TimeLogPrint, error) {
	var lastInsertID int
	err := DB.QueryRow("INSERT INTO time_log(field,timeStamp,previousValue,newValue,issueID) VALUES($1,$2,$3,$4,$5) returning id;", Field, time.Now().Unix(), PreviousValue, NewValue, IssueID).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
	data := TimeLogPrint{
		ID:            lastInsertID,
		Field:         Field,
		TimeStamp:     time.Now().String(),
		PreviousValue: PreviousValue,
		NewValue:      NewValue,
		IssueID:       IssueID,
	}
	return &data, nil

}

//DeleteTimeLog add log in issue
func DeleteTimeLog(id int) ([]TimeLogPrint, error) {
	row, err := DB.Query("DELETE FROM comments WHERE id = $1", id)
	checkErr(err)
	defer row.Close()
	return LoadAllLog()

}
