package controllers

import (
	"net/http"

	"github.com/koficypher/simple_server/models"
	"github.com/koficypher/simple_server/util"
	"gorm.io/gorm"
)

//ArticleHandler used to bind dependencies into the article service struct
type ArticleHandler struct {
	DB *gorm.DB
}

//NewArticleController - returns an instance of the article controller with its dependencies
func NewArticleController(db *gorm.DB) *ArticleHandler {
	return &ArticleHandler{
		DB: db,
	}
}

//GetAllArticles - fetches all articles from table
func (a *ArticleHandler) GetAllArticles(res http.ResponseWriter, req *http.Request) {
	var articles []models.Article
	if result := a.DB.Find(&articles); result.Error != nil {
		util.ERROR(res, http.StatusInternalServerError, "Could not fetch all articles from database ....")
		return
	}

	util.JSON(res, http.StatusOK, articles)
}
