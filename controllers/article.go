package controllers

import (
	"strconv"
	"strings"

	"lke-app/forms"
	"lke-app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// ArticleController ...
type ArticleController struct{}

var articleModel = new(models.ArticleModel)
var articleForm = new(forms.ArticleForm)

// Create creates a new article
// @Summary Create article
// @Description Create a new article
// @Tags Articles
// @Accept json
// @Produce json
// @Security Bearer
// @Param createForm body forms.CreateArticleForm true "Article data"
// @Success 200 {object} map[string]interface{} "Article created"
// @Failure 406 {object} map[string]interface{} "Validation failed"
// @Router /v1/article [post]
func (ctrl ArticleController) Create(c *gin.Context) {
	userID := getUserID(c)

	var form forms.CreateArticleForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := articleModel.Create(userID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created", "id": id})
}

// All retrieves all articles for the user
// @Summary Get all articles
// @Description Get all articles for the authenticated user
// @Tags Articles
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{} "Articles list"
// @Failure 406 {object} map[string]interface{} "Could not get articles"
// @Router /v1/articles [get]
func (ctrl ArticleController) All(c *gin.Context) {
	userID := getUserID(c)

	format := c.Param("format")

	results, err := articleModel.All(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get articles"})
		return
	}

	switch strings.ToLower(format) {
	case "json":
		c.JSON(http.StatusOK, gin.H{"results": results})
	case "xml":
		c.XML(http.StatusOK, gin.H{"results": results})
	case "yaml":
		c.YAML(http.StatusOK, gin.H{"results": results})
	default:
		c.JSON(http.StatusOK, gin.H{"results": results})
	}

}

// One retrieves a specific article by ID
// @Summary Get article by ID
// @Description Get a specific article by ID for the authenticated user
// @Tags Articles
// @Produce json
// @Security Bearer
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]interface{} "Article data"
// @Failure 404 {object} map[string]interface{} "Article not found or invalid parameter"
// @Router /v1/article/{id} [get]
func (ctrl ArticleController) One(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")
	format := c.Param("format")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	data, err := articleModel.One(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found"})
		return
	}

	switch strings.ToLower(format) {
	case "json":
		c.JSON(http.StatusOK, gin.H{"data": data})
	case "xml":
		c.XML(http.StatusOK, gin.H{"data": data})
	case "yaml":
		c.YAML(http.StatusOK, gin.H{"data": data})
	default:
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

// Update modifies an existing article
// @Summary Update article
// @Description Update an existing article by ID
// @Tags Articles
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Article ID"
// @Param updateForm body forms.CreateArticleForm true "Updated article data"
// @Success 200 {object} map[string]interface{} "Article updated"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Validation failed or could not update"
// @Router /v1/article/{id} [put]
func (ctrl ArticleController) Update(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.CreateArticleForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := articleForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = articleModel.Update(userID, getID, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
}

// Delete removes an article by ID
// @Summary Delete article
// @Description Delete an article by ID
// @Tags Articles
// @Produce json
// @Security Bearer
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]interface{} "Article deleted"
// @Failure 404 {object} map[string]interface{} "Invalid parameter"
// @Failure 406 {object} map[string]interface{} "Could not delete"
// @Router /v1/article/{id} [delete]
func (ctrl ArticleController) Delete(c *gin.Context) {
	userID := getUserID(c)

	id := c.Param("id")

	getID, err := strconv.ParseInt(id, 10, 64)
	if getID == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err = articleModel.Delete(userID, getID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})

}
