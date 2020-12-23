package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Project is a structure
type Project struct {
	Name      string `json:"name"`
	CreatedBy int    `json:"created_by"`
	// Owner      *User
	// Issue      []Issue
}

//ProjectPrint is a structure
type ProjectPrint struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	CreatedBy int           `json:"created_by"`
	Owner     *UserPrint    `json:"Owner,omitempty"`
	Issues    *[]IssuePrint `json:"issues"`
}

//StartAddDataForProjects will get data from file add into db
func StartAddDataForProjects() {
	var projects []Project
	data, err := ioutil.ReadFile("../dump_data/projects.json")
	checkErr(err)
	err = json.Unmarshal(data, &projects)
	checkErr(err)
	fmt.Println("data extracted from file")
	fmt.Println("# Inserting values")
	var lastInsertID int
	for _, value := range projects {
		err := DB.QueryRow("INSERT INTO projects(name,created_by) VALUES($1,$2) returning id;", value.Name, value.CreatedBy).Scan(&lastInsertID)
		if err != nil {
			panic(err)
		}
		fmt.Println("last inserted id =", lastInsertID)
	}

}

//LoadAllProjects retunr all projects from db
func LoadAllProjects() ([]ProjectPrint, error) {
	row, err := DB.Query("SELECT id,name,created_by FROM projects")
	var projects []ProjectPrint
	checkErr(err)
	for row.Next() {
		var u ProjectPrint
		row.Scan(&u.ID, &u.Name, &u.CreatedBy)
		u.Issues = GetAllIssuesByID(u.ID)
		u.Owner = GetUserByID(u.CreatedBy)
		projects = append(projects, u)
	}
	return projects, nil
}

//GetProjectsByID return all projects by id
func GetProjectsByID(id int) *[]ProjectPrint {
	row, err := DB.Query("SELECT id,name,created_by FROM users WHERE created_by = $1", id)
	checkErr(err)
	var projects []ProjectPrint
	for row.Next() {
		var u ProjectPrint
		row.Scan(&u.ID, &u.Name, &u.CreatedBy)
		projects = append(projects, u)
	}
	defer row.Close()
	return &projects
}

//GetProjectByID return  project by id
func GetProjectByID(id int) *ProjectPrint {
	row, err := DB.Query("SELECT id,name FROM users WHERE id = $1", id)
	checkErr(err)
	var u ProjectPrint
	for row.Next() {
		row.Scan(&u.ID, &u.Name)
	}
	u.Issues = GetAllIssuesByID(u.ID)
	u.Owner = GetUserByID(u.CreatedBy)
	fmt.Println(u)
	defer row.Close()
	return &u
}

//CreateProject to create new project
func CreateProject(userID int, name string) (*ProjectPrint, error) {
	var lastInsertID int

	err := DB.QueryRow("INSERT INTO projects(name,created_by) VALUES($1,$2) returning id;", name, userID).Scan(&lastInsertID)
	if err != nil {
		return nil, err
	}
	fmt.Println("last inserted id =", lastInsertID)
	project := ProjectPrint{
		ID:        lastInsertID,
		Name:      name,
		CreatedBy: userID,
	}
	project.Issues = GetAllIssuesByID(project.ID)
	project.Owner = GetUserByID(project.CreatedBy)
	return &project, nil
}

//UpdateProject to create new project
func UpdateProject(name string, id int) (*ProjectPrint, error) {

	_, err := DB.Query("UPDATE projects SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return nil, err
	}
	row, err := DB.Query("SELECT id,name,created_by FROM project WHERE id = $1", id)
	checkErr(err)
	var u ProjectPrint
	for row.Next() {
		row.Scan(&u.ID, &u.Name, &u.CreatedBy)
	}
	u.Issues = GetAllIssuesByID(u.ID)
	u.Owner = GetUserByID(u.CreatedBy)
	return &u, nil
}

//DeleteProject for delete projetc
func DeleteProject(id int) ([]ProjectPrint, error) {
	_, err := DB.Query("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return LoadAllProjects()
}
