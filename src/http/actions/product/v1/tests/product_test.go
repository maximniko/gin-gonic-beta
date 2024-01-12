package tests

import (
	"awesomeProject/src/http/actions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductRoute(t *testing.T) {
	e := gin.Default()
	actions.InitRoutes(e)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/product/1", nil)
	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestProductRouteError(t *testing.T) {
	e := gin.Default()
	actions.InitRoutes(e)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/product/0", nil)
	e.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestProductRouteError2(t *testing.T) {
	e := gin.Default()
	actions.InitRoutes(e)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/product/lol", nil)
	e.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
