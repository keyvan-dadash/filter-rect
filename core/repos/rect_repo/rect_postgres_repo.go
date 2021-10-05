package rect_repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/sod-lol/filter-rect/core/models/rect_model"
	"gorm.io/gorm"
)

type RectanglePostgresDB struct {
	db *gorm.DB
}

/*
type RectangleRepo interface {

	//Query api's
	GetRectByID(ctx context.Context, Id string) rect_model.Rectangle
	GetAllRect(ctx context.Context) []rect_model.Rectangle

	//modify api's
	Migrate(ctx context.Context)
	SaveRectangle(ctx context.Context, rect *rect_model.Rectangle)
	UpdateRectangle(ctx context.Context, rect *rect_model.Rectangle) error
	DeleteRectangleByID(ctx context.Context, Id string)
}
*/

func (rpd *RectanglePostgresDB) GetRectByID(ctx context.Context, Id string) *rect_model.Rectangle {

	rect := new(rect_model.Rectangle)

	rpd.db.First(&rect, "Id = ?", Id)

	return rect
}

func (rpd *RectanglePostgresDB) GetAllRect(ctx context.Context) []rect_model.Rectangle {

	var rects []rect_model.Rectangle

	rpd.db.Find(&rects)

	return rects
}

func (rpd *RectanglePostgresDB) Migrate(ctx context.Context) {
	rpd.db.AutoMigrate(&rect_model.Rectangle{})
}

func (rpd *RectanglePostgresDB) SaveRectangle(ctx context.Context, rect *rect_model.Rectangle) error {

	result := rpd.db.Create(rect)

	if result.Error != nil {
		logrus.Errorf("[ERROR] Save rectangle with key %v faild with err %v\n", rect.Id, result.Error)
	}

	return result.Error

}

func (rpd *RectanglePostgresDB) UpdateRectangle(ctx context.Context, rect *rect_model.Rectangle) error {

	result := rpd.db.Save(rect)

	if result.Error != nil {
		logrus.Errorf("[ERROR] Update rectangle with key %v faild with err %v\n", rect.Id, result.Error)
	}

	return result.Error

}

func (rpd *RectanglePostgresDB) DeleteRectangleByID(ctx context.Context, Id string) error {

	result := rpd.db.Delete(rect_model.Rectangle{}, Id)

	if result.Error != nil {
		logrus.Errorf("[ERROR] delete rectangle with key %v faild with err %v\n", Id, result.Error)
	}

	return result.Error

}
