package rect_c

import (
	"time"

	"github.com/sod-lol/filter-rect/core/models/rect_model"
)

//RectangleResponse
//struct for return reponse to client
type RectangleResponse struct {
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Time   string `json:"time"`
}

//ConvertRectangleToRectangleResponse
//convert give Rectangle model to Rectangle Reponse struct
func ConvertRectangleToRectangleResponse(rect *rect_model.Rectangle) RectangleResponse {
	return RectangleResponse{
		X:      rect.X,
		Y:      rect.Y,
		Width:  rect.Width,
		Height: rect.Height,
		Time:   rect.CreatedAt.Format(time.RFC1123Z),
	}
}

//ValidateRectangleRequest
//is to validate rectangle since zero value can accepted we use pointer instead of number
//so we should check their value with nil to ensure value no present in request
func ValidateRectangleRequest(recReq rectangleReq) bool {
	return (recReq.X != nil) && (recReq.Y != nil) && (recReq.Width != nil) && (recReq.Height != nil)
}
