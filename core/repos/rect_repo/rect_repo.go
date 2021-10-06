package rect_repo

import (
	"context"

	"github.com/sod-lol/filter-rect/core/models/rect_model"
)

//RectangleRepo should be implemented for all databases.
type RectangleRepo interface {

	//Query api's
	GetRectByID(ctx context.Context, Id string) *rect_model.Rectangle
	GetAllRect(ctx context.Context) []rect_model.Rectangle

	//modify api's
	Migrate(ctx context.Context)
	SaveRectangle(ctx context.Context, rect *rect_model.Rectangle) error
	UpdateRectangle(ctx context.Context, rect *rect_model.Rectangle) error
	DeleteRectangleByID(ctx context.Context, Id string) error
	DeleteAllRectangle(tx context.Context)
}
