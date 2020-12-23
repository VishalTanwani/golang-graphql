package gql

import "github.com/graphql-go/graphql"

//RootQuery struct
type RootQuery struct {
	Query *graphql.Object
}

//NewRootQuery is the where all query are defined
func NewRootQuery() *RootQuery {

	resolver := QueryResolver{}

	rootQuery := RootQuery{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type:        graphql.NewList(User),
						Description: "Get All User Details",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.GetAllUsers,
					},
					"issues": &graphql.Field{
						Type:        graphql.NewList(Issue),
						Description: "Get All Issues Details",
						// Args: graphql.FieldConfigArgument{
						// 	"token": &graphql.ArgumentConfig{
						// 		Type: graphql.NewNonNull(graphql.String),
						// 	},
						// },
						Resolve: resolver.GetAllIssues,
					},
					"projects": &graphql.Field{
						Type:        graphql.NewList(Project),
						Description: "Get Projects Details",
						Args: graphql.FieldConfigArgument{
							"token": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.GetAllProjects,
					},
					"userById": &graphql.Field{
						Type:        User,
						Description: "get user by id",
						Args: graphql.FieldConfigArgument{
							"userID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.GetUserByID,
					},
					"projectByName": &graphql.Field{
						Type:        Project,
						Description: "get project by name",
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.GetProjectByName,
					},
					"watchers": &graphql.Field{
						Type:        graphql.NewList(Watcher),
						Description: "Get All watchers Details",
						Resolve:     resolver.GetAllWatchers,
					},
					"comments": &graphql.Field{
						Type:        graphql.NewList(Comment),
						Description: "Get All comments Details",
						Resolve:     resolver.GetAllComments,
					},
					"logs": &graphql.Field{
						Type:        graphql.NewList(Comment),
						Description: "Get All comments Details",
						Resolve:     resolver.GetAllLogs,
					},
				},
			},
		),
	}
	return &rootQuery
}
