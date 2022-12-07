package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/blogpost"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
//
//	@param	Authorization	header	string	false	"Authorization"
func (h *handler) AuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		hasAccessTokenResponse, err := h.grpcClient.Users.HasAccess(c.Request.Context(), &blogpost.TokenRequest{
			Token: token,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		if !hasAccessTokenResponse.HasAccess {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if userType != "*" {
			if hasAccessTokenResponse.User.UserType != userType {
				c.JSON(http.StatusUnauthorized, "permission denied")
				c.Abort()
			}
		}

		c.Set("auth_username", hasAccessTokenResponse.User.Username)
		c.Set("auth_userID", hasAccessTokenResponse.User.Id)

		c.Next()
		//
	}
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			users	body		models.LoginModul	true	"Login body"
//	@Success		201		{object}	models.JSONResult{data=models.TokenResponse}
//	@Failure		400		{object}	models.JSONErrorResponse
//	@Router			/v2/login [post]
func (h *handler) Login(c *gin.Context) {

	var body models.LoginModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	tokenResponse, err := h.grpcClient.Users.Login(c.Request.Context(), &blogpost.LoginUserRequest{
		Username: body.Username,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Login",
		Data:    tokenResponse,
	})
}

// CreateUser godoc
//
//	@Summary		Creat User
//	@Description	Creat a new User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			User	body		models.CreateUserModul	true	"User body"
//	@Success		201		{object}	models.JSONResult{data=models.User}
//	@Failure		400		{object}	models.JSONErrorResponse
//	@Router			/v2/user [post]
func (h *handler) CreateUser(c *gin.Context) {

	var body models.CreateUserModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	user, err := h.grpcClient.Users.CreateUser(c.Request.Context(), &blogpost.CreateUserRequest{
		Username: body.Username,
		Password: body.Password,
		UserType: body.User_type,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreateUser",
		Data:    user,
	})
}
