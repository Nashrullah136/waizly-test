package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-backend-1/dto"
)

type RequestHandlerInterface interface {
	Login(c *gin.Context)
}

func NewRequestHandler(authController ControllerInterface) RequestHandlerInterface {
	return requestHandler{authController: authController}
}

type requestHandler struct {
	authController ControllerInterface
}

func (h requestHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest("Invalid Username/Password"))
		return
	}
	response, err := h.authController.Login(request)
	if err != nil {
		c.JSON(response.Code, response)
		return
	}
	c.JSON(response.Code, response)
}
