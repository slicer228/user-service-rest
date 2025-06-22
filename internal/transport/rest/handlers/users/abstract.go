package users

import "github.com/gin-gonic/gin"

type IUsersHandler interface {
	getUsersHandler(c *gin.Context)
	deleteUsersHandler(c *gin.Context)
	patchUsersHandler(c *gin.Context)
	addUserHandler(c *gin.Context)
	Load()
}
