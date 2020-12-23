package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Issue is a structure
type Issue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	AssigneeID  int    `json:"assigneeID"`
	ReporterID  int    `json:"reporterID"`
	Status      string `json:"status"`
	Project     int    `json:"project"`
}

//IssuePrint is a structure
type IssuePrint struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	AssigneeID  int             `json:"assigneeID"`
	ReporterID  int             `json:"reporterID"`
	Status      string          `json:"status"`
	Project     int             `json:"projectID"`
	Assignee    *UserPrint      `json:"assignee"`
	Reporter    *UserPrint      `json:"reporter"`
	Watchers    *[]WatcherPrint `json:"watchers"`
	Comments    *[]CommentPrint `json:"comments"`
	Logs        *[]TimeLogPrint `json:"logs"`
}

//StartAddDataForIssues will get data from file add into db
func StartAddDataForIssues() {
	var issues []Issue
	data, err := ioutil.ReadFile("../dump_data/issues.json")
	checkErr(err)
	err = json.Unmarshal(data, &issues)
	checkErr(err)
	fmt.Println("data extracted from file")
	fmt.Println("# Inserting values")
	var lastInsertID int
	for _, value := range issues {
		err := DB.QueryRow("INSERT INTO issues(title,description,type,assignee,reporter,status,project) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;", value.Title, value.Description, value.Type, value.AssigneeID, value.ReporterID, value.Status, value.Project).Scan(&lastInsertID)
		if err != nil {
			panic(err)
		}
		fmt.Println("last inserted id =", lastInsertID)
	}

}

//LoadAllIssues retunr all issues from db
func LoadAllIssues() ([]IssuePrint, error) {
	row, err := DB.Query("SELECT id,title,description,type,assignee,reporter,status,project FROM issues")
	checkErr(err)
	var issues []IssuePrint
	for row.Next() {
		var u IssuePrint
		row.Scan(&u.ID, &u.Title, &u.Description, &u.Type, &u.AssigneeID, &u.ReporterID, &u.Status, &u.Project)
		u.Assignee = GetUserByID(u.AssigneeID)
		u.Reporter = GetUserByID(u.ReporterID)
		u.Watchers = GetWatcherByIssueID(u.ID)
		u.Comments = GetCommentByIssueID(u.ID)
		u.Logs = LoadAllLogByIssueID(u.ID)
		issues = append(issues, u)
	}
	return issues, nil
}

//GetAllIssuesByID give all issues by id
func GetAllIssuesByID(id int) *[]IssuePrint {
	row, err := DB.Query("SELECT id,title,description,type,assignee,reporter,status,project FROM issues WHERE project = $1", id)
	checkErr(err)
	var issues []IssuePrint
	for row.Next() {
		var u IssuePrint
		row.Scan(&u.ID, &u.Title, &u.Description, &u.Type, &u.AssigneeID, &u.ReporterID, &u.Status, &u.Project)
		u.Assignee = GetUserByID(u.AssigneeID)
		u.Reporter = GetUserByID(u.ReporterID)
		u.Watchers = GetWatcherByIssueID(u.ID)
		u.Comments = GetCommentByIssueID(u.ID)
		u.Logs = LoadAllLogByIssueID(u.ID)
		issues = append(issues, u)
	}
	return &issues
}

//CreateIssue create new issues
func CreateIssue(title, description, Type, status string, assignee, reporter, project int) (*IssuePrint, error) {
	var lastInsertID int
	err := DB.QueryRow("INSERT INTO issues(title,description,type,assignee,reporter,status,project) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;", title, description, Type, assignee, reporter, status, project).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	fmt.Println("last inserted id =", lastInsertID)
	issue := IssuePrint{
		ID:          lastInsertID,
		Title:       title,
		Description: description,
		Type:        Type,
		Status:      status,
		AssigneeID:  assignee,
		ReporterID:  reporter,
		Project:     project,
	}
	issue.Assignee = GetUserByID(issue.AssigneeID)
	issue.Reporter = GetUserByID(issue.ReporterID)
	issue.Watchers = GetWatcherByIssueID(issue.ID)
	issue.Comments = GetCommentByIssueID(issue.ID)
	issue.Logs = LoadAllLogByIssueID(issue.ID)
	_, err = AddWatcher(assignee, lastInsertID)
	_, err = AddWatcher(reporter, lastInsertID)
	return &issue, nil
}

//UpdateIssue create new issues
func UpdateIssue(title, description, Type, status string, id int) (*IssuePrint, error) {
	_, err := DB.Query("UPDATE issues SET title = $1, description = $2, type = $3,status = $4 WHERE id = $5", title, description, Type, status, id)
	if err != nil {
		return nil, err
	}
	issue := IssuePrint{
		ID:          id,
		Title:       title,
		Description: description,
		Type:        Type,
		Status:      status,
	}
	issue.Assignee = GetUserByID(issue.AssigneeID)
	issue.Reporter = GetUserByID(issue.ReporterID)
	issue.Watchers = GetWatcherByIssueID(issue.ID)
	issue.Comments = GetCommentByIssueID(issue.ID)
	issue.Logs = LoadAllLogByIssueID(issue.ID)
	// SendMailToWatchers() // this will send the mails to all watchers that issue is updated
	return &issue, nil
}

//DeleteIssue for delete projetc
func DeleteIssue(id int) ([]IssuePrint, error) {
	_, err := DB.Query("DELETE FROM issues WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return LoadAllIssues()
}
