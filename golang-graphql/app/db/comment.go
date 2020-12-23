package db

import "fmt"

//CommentPrint is structure for watcher
type CommentPrint struct {
	ID      int
	Text    string
	UserID  int
	IssueID int
	User    *UserPrint
	Issue   *IssuePrint
}

//AddComment is to add watcher in db
func AddComment(text string, userID, issueID int) (*WatcherPrint, error) {
	var lastInsertID int
	err := DB.QueryRow("INSERT INTO comments(text,issueid,usersid) VALUES($1,$2,$3) returning id;", text, issueID, userID).Scan(&lastInsertID)
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

//LoadAllComments give all watchers
func LoadAllComments() ([]CommentPrint, error) {
	row, err := DB.Query("SELECT id,text,issueid,usersid FROM comments")
	checkErr(err)
	data := []CommentPrint{}
	for row.Next() {
		var u CommentPrint
		row.Scan(&u.ID, &u.Text, &u.IssueID, &u.UserID)
		u.User = GetUserByID(u.UserID)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return data, nil
}

//DeleteComment will delete watcher
func DeleteComment(id int) ([]CommentPrint, error) {
	row, err := DB.Query("DELETE FROM comments WHERE id = $1", id)
	checkErr(err)
	defer row.Close()
	return LoadAllComments()
}

//GetCommentByIssueID will return watchers by issue id
func GetCommentByIssueID(id int) *[]CommentPrint {
	row, err := DB.Query("SELECT id,text,issueid,usersid FROM comments WHERE issueid=$1", id)
	checkErr(err)
	data := []CommentPrint{}
	for row.Next() {
		var u CommentPrint
		row.Scan(&u.ID, &u.Text, &u.IssueID, &u.UserID)
		u.User = GetUserByID(u.UserID)
		// fmt.Println(u)
		data = append(data, u)
	}
	defer row.Close()
	return &data
}
