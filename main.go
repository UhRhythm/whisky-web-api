package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ウィスキー情報の構造体
type Whisky struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Region string  `json:"region"`
	Type   string  `json:"type"`
	ABV    float64 `json:"abv"`
	Price  int     `json:"price"`
}

var whiskies = []Whisky{
	{1, "デュワーズ", "スコットランド", "ブレンデッド", 40.0, 12000},
	{2, "駒ヶ岳", "日本", "シングルモルト", 52.0, 9460},
	{3, "ラガブーリン", "スコットランド", "シングルモルト", 43.0, 9680},
}

// ウイスキー一覧取得API
func getWhiskies(c *gin.Context) {
	c.JSON(http.StatusOK,whiskies)
}

// 単体取得
func getWhiskiyByID(c *gin.Context) {
	// URLパラメータからIDを取得
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"向こうなIDです￥"})
		return
	}

	// ウイスキーを検索
	for _, whisky := range whiskies {
		if whisky.ID == id {
			c.JSON(http.StatusOK, whisky)
			return
		}
	}

	// 見つからなkった場合
	c.JSON(http.StatusNotFound, gin.H{"error": "ウイスキーが見つかりません"})
}

// 削除
func deleteWhisky(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	// ウイスキーを検索して削除
	for i, whisky := range whiskies {
		if whisky.ID == id {
			whiskies = append(whiskies[:i], whiskies[i+1:]...)
			c.JSON(http.StatusOK,gin.H{
				"message": "削除しました",
				"deleted_whisky" : whisky.Name,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "ウイスキーが見つかりません"})
}

// 新しいウイスキーを追加
func createWhisky(c *gin.Context) {
	var newWhisky Whisky

	if err := c.ShouldBindJSON(&newWhisky); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//新しいIDを生成
	newWhisky.ID = len(whiskies) + 1
	whiskies = append(whiskies, newWhisky)

	// レスポンス返却
	c.JSON(http.StatusCreated, newWhisky)
}

func main(){ 
	router := gin.Default();
	router.GET("/whiskies", getWhiskies)
	router.GET("/whiskies/:id", getWhiskiyByID)
	// main関数に追加
	router.POST("/whiskies", createWhisky)
	router.DELETE("/whiskies/:id", deleteWhisky)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"status": "OK",
			"message": "Whisky API is running",
			"whisky_count": len(whiskies),
		})
	})
	// サーバー起動
	router.Run(":8000")
}