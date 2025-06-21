package get_users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log/slog"
	"net/http"
	um "user-service/internal/service/user-manager"
	"user-service/internal/storage"
)

type RequestGetUsers struct {
	storage.UserFilter
	storage.Paginate
}

func NewGetUserHandler(userManager *um.UserManager, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		log = log.With("request_id", c.GetHeader("X-Request-ID"))
		log.Info("New GET users request")
		var body RequestGetUsers

		if err = c.BindQuery(&body); err != nil {
			log.Error("Error binding query params", err.Error())
			return
		}

		var filter storage.UserFilter
		var pag storage.Paginate

		copier.Copy(&filter, &body)
		copier.Copy(&pag, &body)

		users := userManager.GetUsers(log, &filter, &pag)

		c.Header("content-type", "application/json")
		c.JSON(http.StatusOK, users)
	}
}
