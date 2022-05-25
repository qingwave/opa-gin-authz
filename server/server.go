package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/rego"
	"go.uber.org/zap"
)

type Server struct {
	engine *gin.Engine
	opa    *rego.PreparedEvalQuery
	logger *zap.Logger
}

func New(opa *rego.PreparedEvalQuery, logger *zap.Logger) *Server {
	e := gin.Default()

	s := &Server{
		engine: e,
		opa:    opa,
		logger: logger,
	}

	s.routers()

	return s
}

func (s *Server) Run() error {
	return s.engine.Run(":8080")
}

func (s *Server) routers() {
	s.engine.Use(WithOPA(s.opa, s.logger))

	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello")
	})

	s.engine.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, mockUsers)
	})

	s.engine.GET("/api/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		var user *User
		for i := range mockUsers {
			if mockUsers[i].Name == name {
				user = &mockUsers[i]
				break
			}
		}
		c.JSON(200, gin.H{
			"data": user,
		})
	})

	s.engine.POST("/api/users", func(c *gin.Context) {
		user := User{}
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		mockUsers = append(mockUsers, user)

		c.JSON(200, nil)
	})
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	mockUsers = []User{
		{"alice", "alice@some.com"},
		{"bob", "bob@some.com"},
	}
)
