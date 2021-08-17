package gql

import (
	"fmt"
	"log"
	"net/http"
	"portfolio/model"
	"portfolio/service"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var ConditionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Condition",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"mental": &graphql.Field{
			Type: graphql.Float,
		},
		"physical": &graphql.Field{
			Type: graphql.Float,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var ConditionField = &graphql.Field{
	Type:        graphql.NewList(ConditionType),
	Description: "get condition",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"start": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
		"end": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
	},
	// 現段階でユーザーごとの集計はフロントエンドで頑張ってる
	// そもそもいらなそうな感じもするけど、一応入れておきたいのと、期間ごとで集計できるようにするのはもう欲しいかも
	// 線形回帰で次の予測をするのはフロントエンドから別で投げるか、こっちで生成するか
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username, ok := p.Args["username"].(string)
		start, ok := p.Args["strat"]
		end, ok := p.Args["end"]
		if ok {
			m := service.MaririntonCondition{}
			fmt.Println(username, start, end)          // エラー回避
			condition, err := m.GetCondition(username) // TODO: 引数を追加する
			if err != nil {
				return model.Condition{}, nil
			}
			return condition, nil
		}
		return model.Condition{}, nil
	},
}

// TODO: mmutation 処理の追加
// まだバグだらけで雑音入りまくってるから
var ConditionSchema = graphql.SchemaConfig{
	Query: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ConditionQuery",
			Fields: graphql.Fields{
				"getCondition": ConditionField,
			},
		},
	),
}

type ConditionData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func ConditionResponse(c *gin.Context) {
	var conditionData ConditionData

	if err := c.ShouldBindJSON(&conditionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schema, err := graphql.NewSchema(ConditionSchema)
	if err != nil {
		log.Fatalf("failed to get schema, error: %v", err)
	}

	result := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         schema,
		RequestString:  conditionData.Query,
		VariableValues: conditionData.Variables,
		OperationName:  conditionData.Operation,
	})

	c.JSON(http.StatusOK, result)
}
