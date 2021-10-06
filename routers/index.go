package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sod-lol/filter-rect/routers/rect_r"
)

//InitRoutes
//setup router of rect app
func InitRoutes(ctx context.Context, g *gin.RouterGroup) {

	rect_r.SetUpUrlRoutes(ctx, g.Group(""))
}
