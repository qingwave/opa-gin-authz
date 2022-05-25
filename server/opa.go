package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/rego"
	"go.uber.org/zap"
)

// WithOPA is a gin middleware wrap the OPA authentication
func WithOPA(opa *rego.PreparedEvalQuery, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		groups := c.QueryArray("groups")
		input := map[string]interface{}{
			"method": c.Request.Method,
			"path":   c.Request.RequestURI,
			"subject": map[string]interface{}{
				"user":  user,
				"group": groups,
			},
		}
		logger.Info(fmt.Sprintf("start opa middleware %s, %#v", c.Request.URL.String(), input))
		res, err := opa.Eval(context.TODO(), rego.EvalInput(input))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		defer logger.Info(fmt.Sprintf("opa result: %v, %#v", res.Allowed(), res))

		if !res.Allowed() {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "forbidden",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
