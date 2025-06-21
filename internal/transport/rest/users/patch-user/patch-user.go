package patch_user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log/slog"
	"net/http"
	um "user-service/internal/service/user-manager"
	"user-service/internal/storage"
)

type RequestPatchUser struct {
	storage.UserData
	UserIds []int `json:"user_ids"`
}

func NewPatchUserHandler(userManager *um.UserManager, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		log = log.With("request_id", c.GetHeader("X-Request-ID"))
		log.Info("New PATCH users request")
		var body RequestPatchUser

		if err = c.Bind(&body); err != nil {
			log.Error("Error binding userdata", err.Error())
			return
		}

		var userData storage.UserData
		copier.Copy(&userData, &body)

		err = userManager.PatchUsers(log, &userData, body.UserIds)

		if err != nil {
			log.Error("Error patching user", err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
