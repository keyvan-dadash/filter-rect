package rect_model_test

import (
	"testing"

	"github.com/sod-lol/filter-rect/core/models/rect_model"
	"github.com/stretchr/testify/assert"
)

func TestCreatingRect(t *testing.T) {
	assert := assert.New(t)

	rectObj := rect_model.CreateRectangle(4, 5, 10, 15)

	assert.Equal(rectObj.X, int64(4))
	assert.Equal(rectObj.Y, int64(5))
	assert.Equal(rectObj.Width, int64(10))
	assert.Equal(rectObj.Height, int64(15))
}
