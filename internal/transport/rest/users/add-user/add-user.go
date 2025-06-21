package add_user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log/slog"
	"net/http"
	um "user-service/internal/service/user-manager"
)

type RequestAddUser struct {
	um.PrimaryUserData
}

func NewAddUserHandler(userManager *um.UserManager, log *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		var err error
		log = log.With("request_id", c.GetHeader("X-Request-ID"))
		log.Info("New POST users request")
		var body RequestAddUser

		if err = c.Bind(&body); err != nil {
			log.Error("Error binding userdata", err.Error())
			return
		}

		var userData um.PrimaryUserData
		copier.Copy(&userData, &body)

		userId, err := userManager.AddUser(log, &userData)

		if err != nil {
			log.Error("Error adding user", err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("content-type", "application/json")
		c.JSON(200, gin.H{"user_id": userId})
	}
}
