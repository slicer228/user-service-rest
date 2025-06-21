package users

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	um "user-service/internal/service/user-manager"
	"user-service/internal/transport/rest/users/add-user"
	"user-service/internal/transport/rest/users/delete-user"
	"user-service/internal/transport/rest/users/get-users"
	"user-service/internal/transport/rest/users/patch-user"
)

func LoadUsersRouter(userManager *um.UserManager, log *slog.Logger, rootRouter *gin.RouterGroup) {
	userRouter := rootRouter.Group("users")
	{
		userRouter.POST("", add_user.NewAddUserHandler(userManager, log.With("http-method", "POST")))
		userRouter.DELETE("", delete_user.NewDeleteUserHandler(userManager, log.With("http-method", "DELETE")))
		userRouter.PATCH("", patch_user.NewPatchUserHandler(userManager, log.With("http-method", "PATCH")))
		userRouter.GET("", get_users.NewGetUserHandler(userManager, log.With("http-method", "GET")))
	}

}
