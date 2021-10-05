package rect_c

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sod-lol/filter-rect/core/models/rect_model"
	"github.com/sod-lol/filter-rect/core/repos/rect_repo"
)

type rectangleReq struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type rectanglesPostReq struct {
	Main   rectangleReq   `binding:"required" json:"main"`
	Inputs []rectangleReq `binding:"required" json:"Inputs"`
}

//POST
func HandleAddRectangles(rectRepo rect_repo.RectangleRepo) gin.HandlerFunc {
	return func(c *gin.Context) {

		var rectReq rectanglesPostReq

		if err := c.ShouldBindJSON(&rectReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := c.Request.Context()
		mainRect := rect_model.CreateRectangle(rectReq.Main.X, rectReq.Main.Y, rectReq.Main.Width, rectReq.Main.Height)

		var validRects []*rect_model.Rectangle
		for _, input := range rectReq.Inputs {
			inputRect := rect_model.CreateRectangle(input.X, input.Y, input.Width, input.Height)
			hasOverlap := IsRectangleOverlap(mainRect, inputRect)

			if hasOverlap {
				validRects = append(validRects, inputRect)
			}

		}

		for _, rect := range validRects {
			err := rectRepo.SaveRectangle(ctx, rect)
			logrus.Errorf("[ERROR] cannot save rectangle to database. error: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "try again later",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{})

	}
}

//GET
func HandleGetAllRectangles(rectRepo rect_repo.RectangleRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		allrects := rectRepo.GetAllRect(ctx)

		c.JSON(http.StatusOK, allrects)
	}
}