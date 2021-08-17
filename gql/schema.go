package gql

import (
	"log"
	"net/http"
	"portfolio/model"
	"portfolio/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"contents": &graphql.Field{
			Type: graphql.String,
		},
		"category": &graphql.Field{
			Type: graphql.String,
		},
		"pub_date": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var PostInPostsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PostInPostsType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"contents": &graphql.Field{
			Type: graphql.String,
		},
		"category": &graphql.Field{
			Type: graphql.String,
		},
		"pub_date": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var PostsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Posts",
	Fields: graphql.Fields{
		"category": &graphql.Field{
			Type: graphql.String,
		},
		"current": &graphql.Field{
			Type: graphql.Int,
		},
		"first": &graphql.Field{
			Type: graphql.Int,
		},
		"last": &graphql.Field{
			Type: graphql.Int,
		},
		"next": &graphql.Field{
			Type: graphql.Int,
		},
		"previous": &graphql.Field{
			Type: graphql.Int,
		},
		"page_size": &graphql.Field{
			Type: graphql.Int,
		},
		"results": &graphql.Field{
			Type: graphql.NewList(PostInPostsType),
		},
	},
})

var PostFields = &graphql.Field{
	Type:        PostType,
	Description: "get post",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			h := service.Blog{}
			post, err := h.GetPost(strconv.Itoa(id))
			if err != nil {
				return model.BlogappPost{}, nil
			}
			return post, nil
		}
		return model.BlogappPost{}, nil
	},
}

var PostsFields = &graphql.Field{
	Type:        PostsType,
	Description: "get posts",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"category": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		page, ok := p.Args["page"].(int)
		category, ok := p.Args["category"].(string)
		if ok {
			h := service.Blog{}
			posts, err := h.GetPosts(page, category)
			if err != nil {
				return model.Posts{}, nil
			}

			res := map[string]interface{}{
				"next":      posts.NextPage,
				"previous":  posts.PrevPage,
				"category":  category,
				"results":   posts.Records,
				"current":   page,
				"total":     posts.TotalPage - 1,
				"page_size": 5,
				"first":     1,
				"last":      posts.TotalPage,
			}

			return res, nil
		}
		return model.Posts{}, nil
	},
}

var Schema = graphql.SchemaConfig{
	Query: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "BlogQuery",
			Fields: graphql.Fields{
				"getPost":  PostFields,
				"getPosts": PostsFields,
			},
		},
	),
}

func Response(c *gin.Context) {
	var p model.PostData

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schema, err := graphql.NewSchema(Schema)
	if err != nil {
		log.Fatalf("failed to get schema, error: %v", err)
	}

	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})

	c.JSON(http.StatusOK, result)
}
