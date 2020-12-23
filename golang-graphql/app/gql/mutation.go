package gql

import "github.com/graphql-go/graphql"

//RootMutation is struct
type RootMutation struct {
	Mutation *graphql.Object
}

//NewRootMutation root of all mutation
func NewRootMutation() *RootMutation {
	resolver := MutationResolver{}

	rootMutation := RootMutation{
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"createNewProject": &graphql.Field{
						Type:        Project,
						Description: "create new project",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"userID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.CreateNewProject,
					},
					"updateProject": &graphql.Field{
						Type:        Project,
						Description: "update project",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.UpdateProject,
					},
					"deleteProject": &graphql.Field{
						Type:        graphql.NewList(Project),
						Description: "delete project",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteProject,
					},
					"createNewIssue": &graphql.Field{
						Type:        Project,
						Description: "create new issue",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"type": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"assignee": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"reporter": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"status": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"project": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.CreateNewIssue,
					},
					"updateIssue": &graphql.Field{
						Type:        Project,
						Description: "update issue",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"type": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"status": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.UpdateIssue,
					},
					"deleteIssue": &graphql.Field{
						Type:        graphql.NewList(Issue),
						Description: "delete issue",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteIssue,
					},
					"signUp": &graphql.Field{
						Type:        User,
						Description: "create new User",
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"email": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"role": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.SingUp,
					},
					"addWatcher": &graphql.Field{
						Type:        Watcher,
						Description: "create new watcher",
						Args: graphql.FieldConfigArgument{
							"issueID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"userID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.AddNewWatcher,
					},
					"deleteWatcher": &graphql.Field{
						Type:        graphql.NewList(Watcher),
						Description: "delete watcher",
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteWatcher,
					},
					"addComment": &graphql.Field{
						Type:        Comment,
						Description: "add new comment",
						Args: graphql.FieldConfigArgument{
							"text": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"issueID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"userID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.AddNewComment,
					},
					"deleteComment": &graphql.Field{
						Type:        graphql.NewList(Comment),
						Description: "delete comment",
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteComment,
					},
					"addLog": &graphql.Field{
						Type:        TimeLog,
						Description: "add new log",
						Args: graphql.FieldConfigArgument{
							"field": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"previousValue": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"newValue": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"issueID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.AddNewLog,
					},
					"deleteLog": &graphql.Field{
						Type:        graphql.NewList(TimeLog),
						Description: "delete log",
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteLog,
					},
				},
			},
		),
	}
	return &rootMutation
}
