package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sod-lol/filter-rect/core/repos/rect_repo"
	"github.com/sod-lol/filter-rect/routers"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var doOnce sync.Once

var router *gin.Engine

var rectRepo rect_repo.RectanglePostgresDB

func setUpRouters() {
	gin.DefaultWriter = ioutil.Discard

	router = gin.New()

	root := context.Background()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("PS_HOST"),
		os.Getenv("PS_USERNAME"),
		os.Getenv("PS_PASSWORD"),
		os.Getenv("PS_DBNAME"),
		os.Getenv("PS_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database")
	}

	rectRepo = rect_repo.RectanglePostgresDB{}
	rectRepo.SetDB(root, db)
	rectRepo.Migrate(root)

	ctxWithRepo := rect_repo.SetRectangleRepoInContext(root, &rectRepo)

	routers.InitRoutes(ctxWithRepo, &router.RouterGroup)
}

func GetRouterAndRepo() (*gin.Engine, *rect_repo.RectanglePostgresDB) {
	doOnce.Do(func() {
		setUpRouters()
	})

	return router, &rectRepo
}

func TestPostRectangle(t *testing.T) {
	assert := assert.New(t)
	routerSetup, repo := GetRouterAndRepo()

	w := httptest.NewRecorder()

	jsonBody := `
	{
		"main": {"x": 0, "y": 0, "width": 10, "height": 20},
		"inputs": [
			{"x": 2, "y": 18, "width": 5, "height": 4},
			{"x": 12, "y": 18, "width": 5, "height": 4},
			{"x": -1, "y": -1, "width": 5, "height": 4}
		]
	}
	
	
	{
		"main": {"x": 3, "y": 2, "width": 5, "height": 10},
		"inputs": [
			{"x": 4, "y": 10, "width": 1, "height": 1},
			{"x": 9, "y": 10, "width": 5, "height": 4}
		]
	}`

	req, _ := http.NewRequest("POST", "/", strings.NewReader(string(jsonBody)))
	routerSetup.ServeHTTP(w, req)

	assert.Equal(http.StatusCreated, w.Code)

	//clean up
	ctx := context.Background()
	repo.DeleteAllRectangle(ctx)
}

type Response []struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Time   string `json:"time"`
}

func TestGetAfterPostRectangle(t *testing.T) {
	assert := assert.New(t)
	routerSetup, repo := GetRouterAndRepo()

	w := httptest.NewRecorder()

	jsonBody := `
	{
		"main": {"x": 0, "y": 0, "width": 10, "height": 20},
		"inputs": [
			{"x": 2, "y": 18, "width": 5, "height": 4},
			{"x": 12, "y": 18, "width": 5, "height": 4}
		]
	}`

	req, _ := http.NewRequest("POST", "/", strings.NewReader(string(jsonBody)))
	routerSetup.ServeHTTP(w, req)

	assert.Equal(http.StatusCreated, w.Code)

	//Get
	w = httptest.NewRecorder()

	completeURL := string("/")
	req, _ = http.NewRequest("GET", completeURL, nil)
	routerSetup.ServeHTTP(w, req)

	respInfo := w.Result()
	body, _ := io.ReadAll(respInfo.Body)

	fmt.Println(string(body))

	var wholeResp Response
	err := json.Unmarshal(body, &wholeResp)
	assert.Equal(nil, err)

	fmt.Println(wholeResp)

	validRect := wholeResp[0]

	assert.Equal(validRect.X, 2)
	assert.Equal(validRect.Y, 18)
	assert.Equal(validRect.Width, 5)
	assert.Equal(validRect.Height, 4)

	//clean up
	ctx := context.Background()
	repo.DeleteAllRectangle(ctx)
}
