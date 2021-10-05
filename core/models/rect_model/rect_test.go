package rect_model_test

import (
	"testing"

	"github.com/sod-lol/filter-rect/core/models/rect_model"
	"github.com/stretchr/testify/assert"
)

func TestCreatingRect(t *testing.T) {
	assert := assert.New(t)

	rectObj := rect_model.CreateRectangle(4, 5, 10, 15)

	assert.Equal(rectObj.X, 4)
	assert.Equal(rectObj.Y, 5)
	assert.Equal(rectObj.Width, 10)
	assert.Equal(rectObj.Height, 15)
}
