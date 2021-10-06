package rect_r

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sod-lol/filter-rect/controllers/rect_c"
	"github.com/sod-lol/filter-rect/core/repos/rect_repo"
)

//SetUpUrlRoutes for Rect filtering app
func SetUpUrlRoutes(ctx context.Context, g *gin.RouterGroup) {

	rect_repo, hasUrlRepo := rect_repo.GetRectangleRepoFromContex(ctx)
	if !hasUrlRepo {
		logrus.Fatalf("[FATAL] context does not have Rect Repo")
	}

	g.POST("/", rect_c.HandleAddRectangles(rect_repo))
	g.GET("/", rect_c.HandleGetAllRectangles(rect_repo))
}
