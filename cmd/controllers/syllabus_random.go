package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"InfoLinkAPI/cmd/models"
)

// [メソッド名] [動詞で始まる簡潔な説明]。
//
// [引数についての詳細説明（必要な場合）]
// [戻り値についての詳細説明（必要な場合）]
// [その他特記事項があれば記述]
func (sc *SyllabusController) GetRandom(c *gin.Context) {
	var syllabus models.SyllabusBaseInfo
	// データベースからランダムなレコードを取得
	err := sc.DB.Order("RAND()").First(&syllabus).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データの取得に失敗しました"})
		return
	}
	// ランダムに取得したレコードを返す
	c.JSON(http.StatusOK, syllabus)
}
