package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/blogpost"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatArticle godoc
//
//	@Summary		Creat Article
//	@Description	Creat a new article
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article			body		models.CreateArticleModul	true	"Article body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Article}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/article [post]
func (h *handler) CreatArticle(c *gin.Context) {

	var body models.CreateArticleModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	article, err := h.grpcClient.Article.CreateArticle(c.Request.Context(), &blogpost.CreateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		AuthorId: body.AuthorID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatArticle",
		Data:    article,
	})
}

// GetArticleByID godoc
//
//	@Summary		GetArticleByID
//	@Description	get an article by id
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Article id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.PackedArticleModel}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/article/{id} [get]
func (h *handler) GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.grpcClient.Article.GetArticleById(c.Request.Context(), &blogpost.GetArticleByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "GetArticleByID",
		Data:    article,
	})
}

// GetArticleList godoc
//
//	@Summary		List articles
//	@Description	GetArticleList
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"			default()
//	@Param			limit			query		int		false	"100"		default()
//	@Param			search			query		string	false	"search"	default()
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Article}
//	@Router			/v2/article/ [get]
func (h *handler) GetArticleList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")

	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClient.Article.GetArticleList(c.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Data:    articleList,
		Message: "GetList OK",
	})
}

// SearchArticleByMyUsername godoc
//
//	@Summary		List articles
//	@Description	SearchArticleByMyUsername
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"		default()
//	@Param			limit			query		int		false	"100"	default()
//	@Param			search			query		string	false	"s"		default()
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Article}
//	@Router			/v2/my-articles/ [get]
func (h *handler) SearchArticleByMyUsername(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")

	usernameRaw, ok := c.Get("auth-username")
	if !ok {
		c.JSON(http.StatusBadRequest, "Something is wrong")
	}

	username, ok := usernameRaw.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, "Something is wrong")
	}

	searchStr := username
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClient.Article.GetArticleList(c.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Data:    articleList,
		Message: "GetList OK",
	})
}

// ArticleUpdate godoc
//
//	@Summary		Update Article
//	@Description	Update Article
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article			body		models.UpdateArticleModul	true	"Article body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Article}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/article/ [put]
func (h *handler) ArticleUpdate(c *gin.Context) {
	var body models.UpdateArticleModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	article, err := h.grpcClient.Article.UpdateArticle(c.Request.Context(), &blogpost.UpdateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		Id: body.ID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// get article by id service ichida tekshirib ketilgan

	c.JSON(http.StatusOK, gin.H{
		"message": "Article Update",
		"data":    article,
	})

}

// DeleteArticle godoc
//
//	@Summary		Delete Article
//	@Description	get element by id and delete this article
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Article id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Article}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/article/{id} [delete]
func (h *handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	// get article by id service ichida tekshirib ketilgan

	article, err := h.grpcClient.Article.DeleteArticle(c.Request.Context(), &blogpost.DeleteArticleRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article Deleted",
		"data":    article,
	})
}
