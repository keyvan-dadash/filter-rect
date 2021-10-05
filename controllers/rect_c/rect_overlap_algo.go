package rect_c

import "github.com/sod-lol/filter-rect/core/models/rect_model"

type point struct {
	x int64
	y int64
}

func IsRectangleOverlap(rect1 *rect_model.Rectangle, rect2 *rect_model.Rectangle) bool {

	rect1_bottomright := point{
		x: rect1.X,
		y: rect1.Y,
	}

	rect1_topleft := point{
		x: rect1.X - rect1.Width,
		y: rect1.Y + rect1.Height,
	}

	rect2_bottomright := point{
		x: rect2.X,
		y: rect2.Y,
	}

	rect2_topleft := point{
		x: rect2.X - rect2.Width,
		y: rect2.Y + rect2.Height,
	}

	if rect1_topleft.x == rect1_bottomright.x ||
		rect1_topleft.y == rect1_bottomright.y ||
		rect2_topleft.x == rect2_bottomright.x ||
		rect2_topleft.y == rect2_bottomright.y {
		return false
	}

	if rect1_topleft.x >= rect2_bottomright.x ||
		rect2_topleft.x >= rect1_bottomright.x {
		return false
	}

	if rect1_bottomright.y >= rect2_topleft.y ||
		rect2_bottomright.y >= rect1_topleft.y {
		return false
	}

	return true
}
