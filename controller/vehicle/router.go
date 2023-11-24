package vehicle

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/group/api/v1")
	return g
}
