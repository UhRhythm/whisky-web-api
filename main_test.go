package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    
    router.GET("/whiskies", getWhiskies)
    router.GET("/whiskies/:id", getWhiskiyByID)
    router.POST("/whiskies", createWhisky)
    router.DELETE("/whiskies/:id", deleteWhisky)
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "OK",
            "message": "Whisky API is running",
            "whisky_count": len(whiskies),
        })
    })
    
    return router
}

func TestGetWhiskies(t *testing.T) {
    router := setupRouter()
    
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/whiskies", nil)
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "デュワーズ")
}

func TestGetWhiskyByID(t *testing.T) {
    router := setupRouter()
    
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/whiskies/1", nil)
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "デュワーズ")
}

func TestCreateWhisky(t *testing.T) {
    router := setupRouter()
    
    newWhisky := Whisky{
        Name:   "山崎",
        Region: "日本",
        Type:   "シングルモルト",
        ABV:    43.0,
        Price:  15000,
    }
    
    jsonValue, _ := json.Marshal(newWhisky)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/whiskies", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 201, w.Code)
    assert.Contains(t, w.Body.String(), "山崎")
}

func TestHealthCheck(t *testing.T) {
    router := setupRouter()
    
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/health", nil)
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "OK")
}
func TestDeleteWhisky(t *testing.T) {
    router := setupRouter()
    
    // まず削除対象が存在することを確認
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/whiskies/3", nil)
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "削除しました")
}

func TestDeleteWhiskyNotFound(t *testing.T) {
    router := setupRouter();

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("DELETE", "/whiskies/999", nil)

    router.ServeHTTP(w, req)

    assert.Equal(t, 404, w.Code)

    assert.Contains(t, w.Body.String(), "見つかりません")
}