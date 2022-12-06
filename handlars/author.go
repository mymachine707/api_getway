package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/blogpost"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatAuthor godoc
//	@Summary		Creat Author
//	@Description	Creat a new author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			author			body		models.CreateAuthorModul	true	"Author body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/author [post]
func (h *handler) CreatAuthor(c *gin.Context) {

	var body models.CreateAuthorModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// validation should be here
	// create new author

	author, err := h.grpcClient.Author.CreateAuthor(c.Request.Context(), &blogpost.CreateAuthorRequest{
		Firstname:  body.Firstname,
		Lastname:   body.Lastname,
		Middlename: body.Middlename,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Get author by id service ichida tekshirilgan

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatAuthor",
		Data:    author,
	})
}

// GetAuthorByID godoc
//	@Summary		GetAuthorByID
//	@Description	get an author by id
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Author id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.PackedAuthorModel}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/author/{id} [get]
func (h *handler) GetAuthorByID(c *gin.Context) {

	idStr := c.Param("id")

	// validation

	author, err := h.grpcClient.Author.GetAuthorById(c.Request.Context(), &blogpost.GetAuthorByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthorList godoc
//	@Summary		List authors
//	@Description	GetAuthorList
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Author}
//	@Router			/v2/author/ [get]
func (h *handler) GetAuthorList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "100")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	authorList, err := h.grpcClient.Author.GetAuthorList(c.Request.Context(), &blogpost.GetAuthorListRequest{
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
		Message: "GetList OK",
		Data:    authorList,
	})
}

// AuthorUpdate godoc
//	@Summary		My work !!! -- Update Author
//	@Description	Update Author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			author			body		models.UpdateAuthorModul	true	"Author body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/author/ [put]
func (h *handler) AuthorUpdate(c *gin.Context) {
	var body models.UpdateAuthorModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	author, err := h.grpcClient.Author.UpdateAuthor(c.Request.Context(), &blogpost.UpdateAuthorRequest{
		Id:         body.ID,
		Firstname:  body.Firstname,
		Lastname:   body.Lastname,
		Middlename: body.Middlename,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Get author by id service ichida tekshirilgan

	c.JSON(http.StatusOK, gin.H{
		"message": "Author Update",
		"data":    author,
	})

}

// DeleteAuthor godoc
//	@Summary		My work!!! -- Delete Author
//	@Description	get element by id and delete this author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Author id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v2/author/{id} [delete]
func (h *handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	// Get author by id service ichida tekshirilgan

	author, err := h.grpcClient.Author.DeleteAuthor(c.Request.Context(), &blogpost.DeleteAuthorRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author Deleted",
		"data":    author,
	})

}
