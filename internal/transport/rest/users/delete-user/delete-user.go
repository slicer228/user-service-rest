package delete_user

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	um "user-service/internal/service/user-manager"
)

type RequestDeleteUser struct {
	UserIds []int `json:"user_id"`
}

func NewDeleteUserHandler(userManager *um.UserManager, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		log = log.With("request_id", c.GetHeader("X-Request-ID"))
		log.Info("New DELETE user request")
		var body RequestDeleteUser

		if err = c.BindJSON(&body); err != nil {
			log.Error("Bind json error", "error", err)
			return
		}

		err = userManager.DeleteUsers(log, body.UserIds)

		if err != nil {
			log.Error("Delete users error", "error", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
