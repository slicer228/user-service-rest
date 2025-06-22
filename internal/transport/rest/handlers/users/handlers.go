package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log/slog"
	"net/http"
	um "user-service/internal/service/user-manager"
	"user-service/internal/storage"
)

type UsersRouter struct {
	userManager *um.UserManager
	log         *slog.Logger
	rootRouter  *gin.RouterGroup
}

func (ur *UsersRouter) Load() {
	userRouter := ur.rootRouter.Group("users")
	{
		userRouter.POST("", ur.addUserHandler)
		userRouter.DELETE("", ur.deleteUsersHandler)
		userRouter.PATCH("", ur.patchUsersHandler)
		userRouter.GET("", ur.getUsersHandler)
	}

}

// @Summary      Добавление пользователя
// @Tags         Users
// @Description  Добавляет нового пользователя в систему
// @ID           add-user
// @Accept       json
// @Produce      json
// @Param        user  body      RequestAddUser  true  "Данные пользователя"
// @Success      200   {object}  map[string]int  "user_id"
// @Failure      400,404,500     {string}        string "ошибка"
// @Router       /users [post]
func (ur *UsersRouter) addUserHandler(c *gin.Context) {
	var err error
	log := ur.log.With("request_id", c.GetHeader("X-Request-ID"), "http-method", "POST")
	log.Info("New POST users request")
	var body RequestAddUser

	if err = c.Bind(&body); err != nil {
		log.Error("Error binding userdata", err.Error())
		return
	}

	var userData um.PrimaryUserData
	copier.Copy(&userData, &body)

	userId, err := ur.userManager.AddUser(log, &userData)

	if err != nil {
		log.Error("Error adding user", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(200, gin.H{"user_id": userId})
}

// @Summary      Получение списка пользователей
// @Tags         Users
// @Description  Возвращает пользователей по фильтру
// @ID           get-users
// @Accept       json
// @Produce      json
// @Param        id           query []int    false  "ID пользователя"
// @Param        name         query []string false  "Имя"
// @Param        surname      query []string false  "Фамилия"
// @Param        patronymic   query []string false  "Отчество"
// @Param        gender       query []string false  "Пол"
// @Param        age          query []int    false  "Возраст"
// @Param        nationality  query []string false  "Национальность"
// @Param        page         query int      false  "Номер страницы"
// @Param        itemsPerPage query int      false  "Элементов на страницу"
// @Success      200  {array}   storage.UserData
// @Failure      400,404,500  {string} string "ошибка"
// @Router       /users [get]
func (ur *UsersRouter) getUsersHandler(c *gin.Context) {
	var err error
	log := ur.log.With("request_id", c.GetHeader("X-Request-ID"), "http-method", "GET")
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

	users := ur.userManager.GetUsers(log, &filter, &pag)

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, users)
}

// @Summary      Удаление пользователей
// @Tags         Users
// @Description  Удаляет пользователей по списку ID
// @ID           delete-users
// @Accept       json
// @Produce      json
// @Param        user_ids  body      RequestDeleteUser  true  "ID пользователей"
// @Success      204
// @Failure      400,404,500  {string} string "ошибка"
// @Router       /users [delete]
func (ur *UsersRouter) deleteUsersHandler(c *gin.Context) {
	var err error
	log := ur.log.With("request_id", c.GetHeader("X-Request-ID"), "http-method", "DELETE")
	log.Info("New DELETE user request")
	var body RequestDeleteUser

	if err = c.BindJSON(&body); err != nil {
		log.Error("Bind json error", "error", err)
		return
	}

	err = ur.userManager.DeleteUsers(log, body.UserIds)

	if err != nil {
		log.Error("Delete users error", "error", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary      Обновление данных пользователей
// @Tags         Users
// @Description  Обновляет данные пользователей по списку ID
// @ID           patch-users
// @Accept       json
// @Produce      json
// @Param        user_data  body      RequestPatchUser  true  "Обновляемые поля и ID пользователей"
// @Success      200
// @Failure      400,404,500  {string} string "ошибка"
// @Router       /users [patch]
func (ur *UsersRouter) patchUsersHandler(c *gin.Context) {
	var err error
	log := ur.log.With("request_id", c.GetHeader("X-Request-ID"), "http-method", "PATCH")
	log.Info("New PATCH users request")
	var body RequestPatchUser

	if err = c.Bind(&body); err != nil {
		log.Error("Error binding userdata", err.Error())
		return
	}

	var userData storage.UserData
	copier.Copy(&userData, &body)

	err = ur.userManager.PatchUsers(log, &userData, body.UserIds)

	if err != nil {
		log.Error("Error patching user", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func NewUsersRouter(log *slog.Logger, userManager *um.UserManager, rootRouter *gin.RouterGroup) *UsersRouter {
	return &UsersRouter{
		userManager: userManager,
		log:         log,
		rootRouter:  rootRouter,
	}
}
