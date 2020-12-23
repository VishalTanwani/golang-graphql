package gql

import (
	"errors"
	"fmt"
	"graphql/app/db"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/graphql-go/graphql"
)

var key = []byte("wertyuiodfghjkcvbnm")

func generateJWT(name, password, email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["Name"] = name
	claim["Password"] = password
	claim["Email"] = email
	claim["Role"] = role
	claim["exp"] = time.Now().Add(time.Minute * 60).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return tokenString, nil

}

func verifyToken(tkn string) (bool, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error")
		}
		return key, nil
	})

	if err != nil {
		panic(err)
	}

	if token.Valid {
		return true, nil
	}
	return false, errors.New("User not Found")

}

func verifyAdminToken(tkn string) (bool, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error")
		}
		return key, nil
	})

	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims["Role"] == "ADMIN" {
		return true, nil
	}
	return false, errors.New("User not Found")

}

func verifyPMToken(tkn string) (bool, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error")
		}
		return key, nil
	})

	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims["Role"] == "PM" {
		return true, nil
	}
	return false, errors.New("User not Found")

}

//QueryResolver is struct
type QueryResolver struct{}

//GetAllUsers is a function wich return all users
func (q *QueryResolver) GetAllUsers(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyToken(token); ok {
		allUsers, err := db.LoadAllUsers()
		return allUsers, err
	}
	return nil, errors.New("UserNotFond")
}

//GetUserByID gone return user by id
func (q *QueryResolver) GetUserByID(p graphql.ResolveParams) (interface{}, error) {
	userID := p.Args["userID"].(int)
	allUsers, err := db.LoadAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range allUsers {
		if user.ID == userID {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

//GetAllProjects is a function wich return all projects
func (q *QueryResolver) GetAllProjects(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyToken(token); ok {
		allProjects, err := db.LoadAllProjects()
		return allProjects, err
	}
	return nil, errors.New("UserNotFond")
}

// GetProjectByName retunr project by name
func (q *QueryResolver) GetProjectByName(p graphql.ResolveParams) (interface{}, error) {
	allProjects, err := db.LoadAllProjects()
	name := p.Args["name"].(string)
	if err != nil {
		return nil, err
	}
	for _, project := range allProjects {
		if project.Name == name {
			return project, nil
		}
	}
	return nil, err
}

//GetAllIssues is a function wich return all issues
func (q *QueryResolver) GetAllIssues(p graphql.ResolveParams) (interface{}, error) {
	// token := p.Args["token"].(string)
	// if ok, _ := verifyToken(token); ok {
	allIssues, err := db.LoadAllIssues()
	return allIssues, err
	// }
	// return nil, errors.New("UserNotFond")

}

//GetAllWatchers is a function wich return all users
func (q *QueryResolver) GetAllWatchers(p graphql.ResolveParams) (interface{}, error) {
	// token := p.Args["token"].(string)
	// if ok, _ := verifyToken(token); ok {
	allWatchers, err := db.LoadAllWatchers()
	return allWatchers, err
	// }
	// return nil, errors.New("UserNotFond")
}

//GetAllComments is a function wich return all users
func (q *QueryResolver) GetAllComments(p graphql.ResolveParams) (interface{}, error) {
	// token := p.Args["token"].(string)
	// if ok, _ := verifyToken(token); ok {
	allComments, err := db.LoadAllComments()
	return allComments, err
	// }
	// return nil, errors.New("UserNotFond")
}

//GetAllLogs is a function wich return all users
func (q *QueryResolver) GetAllLogs(p graphql.ResolveParams) (interface{}, error) {
	// token := p.Args["token"].(string)
	// if ok, _ := verifyToken(token); ok {
	allLogs, err := db.LoadAllLog()
	return allLogs, err
	// }
	// return nil, errors.New("UserNotFond")
}

//MutationResolver is struct
type MutationResolver struct{}

//CreateNewProject is mutation for creating new project
func (q *MutationResolver) CreateNewProject(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyAdminToken(token); ok {
		userID := p.Args["userID"].(int)
		name := p.Args["name"].(string)

		return db.CreateProject(userID, name)
	}
	return nil, errors.New("UserNotFond")

}

//UpdateProject is mutation for update project
func (q *MutationResolver) UpdateProject(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyAdminToken(token); ok {
		id := p.Args["id"].(int)
		name := p.Args["name"].(string)

		return db.UpdateProject(name, id)
	}
	return nil, errors.New("UserNotFond")

}

//DeleteProject is mutation for delete project
func (q *MutationResolver) DeleteProject(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyAdminToken(token); ok {
		id := p.Args["id"].(int)

		return db.DeleteProject(id)
	}
	return nil, errors.New("UserNotFond")

}

//CreateNewIssue is mutation for creating new project
func (q *MutationResolver) CreateNewIssue(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyPMToken(token); ok {
		title := p.Args["title"].(string)
		description := p.Args["description"].(string)
		Type := p.Args["type"].(string)
		assignee := p.Args["assignee"].(int)
		reporter := p.Args["reporter"].(int)
		project := p.Args["project"].(int)
		status := p.Args["status"].(string)

		return db.CreateIssue(title, description, Type, status, assignee, reporter, project)
	}
	return nil, errors.New("UserNotFond")

}

//UpdateIssue is mutation for creating new project
func (q *MutationResolver) UpdateIssue(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyPMToken(token); ok {
		id := p.Args["id"].(int)
		title := p.Args["title"].(string)
		description := p.Args["description"].(string)
		Type := p.Args["type"].(string)
		status := p.Args["status"].(string)

		return db.UpdateIssue(title, description, Type, status, id)
	}
	return nil, errors.New("UserNotFond")

}

//DeleteIssue is mutation for delete project
func (q *MutationResolver) DeleteIssue(p graphql.ResolveParams) (interface{}, error) {
	token := p.Args["token"].(string)
	if ok, _ := verifyPMToken(token); ok {
		id := p.Args["id"].(int)

		return db.DeleteIssue(id)
	}
	return nil, errors.New("UserNotFond")

}

//SingUp will register new user
func (q *MutationResolver) SingUp(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"].(string)
	password := p.Args["password"].(string)
	email := p.Args["email"].(string)
	role := p.Args["role"].(string)

	token, err := generateJWT(name, password, email, role)
	if err != nil {
		panic(err)
	}

	fmt.Println(token)

	return db.CreateUser(name, password, email, role)
}

//AddNewWatcher add new watcher
func (q *MutationResolver) AddNewWatcher(p graphql.ResolveParams) (interface{}, error) {
	issueID := p.Args["issueID"].(int)
	userID := p.Args["userID"].(int)

	return db.AddWatcher(issueID, userID)
}

//DeleteWatcher delete watcher
func (q *MutationResolver) DeleteWatcher(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["issueID"].(int)

	return db.DeleteWatcher(id)
}

//AddNewComment add new watcher
func (q *MutationResolver) AddNewComment(p graphql.ResolveParams) (interface{}, error) {
	text := p.Args["text"].(string)
	issueID := p.Args["issueID"].(int)
	userID := p.Args["userID"].(int)

	return db.AddComment(text, issueID, userID)
}

//DeleteComment delete watcher
func (q *MutationResolver) DeleteComment(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["issueID"].(int)

	return db.DeleteComment(id)
}

//AddNewLog add new watcher
func (q *MutationResolver) AddNewLog(p graphql.ResolveParams) (interface{}, error) {
	field := p.Args["field"].(string)
	previousValue := p.Args["previousValue"].(string)
	newValue := p.Args["newValue"].(string)
	issueID := p.Args["issueID"].(int)

	return db.AddTimeLog(field, previousValue, newValue, issueID)
}

//DeleteLog delete watcher
func (q *MutationResolver) DeleteLog(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["issueID"].(int)

	return db.DeleteTimeLog(id)
}
