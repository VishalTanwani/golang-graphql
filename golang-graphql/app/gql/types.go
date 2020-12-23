package gql

import "github.com/graphql-go/graphql"

//User is user type
var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

//Project Type
var Project = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Project",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"createdby": &graphql.Field{
				Type: graphql.Int,
			},
			"Owner": &graphql.Field{
				Type: User,
			},
			"issues": &graphql.Field{
				Type: graphql.NewList(Issue),
			},
		},
	},
)

//Issue Type
var Issue = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Issue",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"assigneeID": &graphql.Field{
				Type: graphql.Int,
			},
			"reporterID": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"project": &graphql.Field{
				Type: graphql.Int,
			},
			"assignee": &graphql.Field{
				Type: User,
			},
			"reporter": &graphql.Field{
				Type: User,
			},
			"watchers": &graphql.Field{
				Type: graphql.NewList(Watcher),
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(Comment),
			},
			"logs": &graphql.Field{
				Type: graphql.NewList(TimeLog),
			},
		},
	},
)

//Watcher Type
var Watcher = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Watcher",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"issueID": &graphql.Field{
				Type: graphql.Int,
			},
			"userID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"user": &graphql.Field{
				Type: User,
			},
		},
	},
)

//Comment Type
var Comment = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"issueID": &graphql.Field{
				Type: graphql.Int,
			},
			"userID": &graphql.Field{
				Type: graphql.Int,
			},
			"user": &graphql.Field{
				Type: User,
			},
		},
	},
)

//TimeLog Type
var TimeLog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TimeLog",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"field": &graphql.Field{
				Type: graphql.String,
			},
			"timeStamp": &graphql.Field{
				Type: graphql.String,
			},
			"previousValue": &graphql.Field{
				Type: graphql.String,
			},
			"newValue": &graphql.Field{
				Type: graphql.String,
			},
			"issueID": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
