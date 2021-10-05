package rect_r

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sod-lol/filter-rect/core/repos/rect_repo"
)

//SetUpUrlRoutes for URL shortner app
func SetUpUrlRoutes(ctx context.Context, g *gin.RouterGroup) {

	rect_repo, hasUrlRepo := rect_repo.GetRectangleRepoFromContex(ctx)
	if !hasUrlRepo {
		logrus.Fatalf("[FATAL] context does not have URL Repo")
	}

	// g.POST("/submit", url_c.HandlerShortURLRequest(urlRepo))
	// g.GET("/:shortUrl", url_c.HandlerRedirect(urlRepo))
	// g.GET("/info/:shortUrl", url_c.HandlerShortURLInformation(urlRepo))
}
