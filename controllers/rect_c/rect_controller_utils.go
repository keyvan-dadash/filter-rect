package rect_c

import (
	"github.com/sod-lol/filter-rect/core/models/rect_model"
)

type RectangleResponse struct {
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Time   string `json:"time"`
}

func ConvertRectangleToRectangleResponse(rect *rect_model.Rectangle) RectangleResponse {
	return RectangleResponse{
		X:      rect.X,
		Y:      rect.Y,
		Width:  rect.Width,
		Height: rect.Height,
		Time:   rect.CreatedAt.Format("RFC1123Z"),
	}
}
