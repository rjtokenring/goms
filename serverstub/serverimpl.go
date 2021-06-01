package serverstub

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strconv"
	"strings"
)

type GoMsServerImpl struct {

}

func (impl *GoMsServerImpl) DeleteUser(ctx echo.Context, id int64) error  {
	var stringId = strconv.FormatInt(id, 10)
	log.Info("Deleting user " + stringId)
	return ctx.NoContent(204)
}

func (impl *GoMsServerImpl) FindUserById(ctx echo.Context, id int64) error  {
	var stringId = strconv.FormatInt(id, 10)
	log.Info("Getting user " + stringId)

	var iUser = BaseUser{"John", "Tom", nil}
	var user = User{iUser, id}
	return ctx.JSON(200, user)
}

// (GET /users)
func (impl *GoMsServerImpl) FindUsers(ctx echo.Context, params FindUsersParams) error {
	if params.Names != nil {
		log.Info("Finding users: " + strings.Join(*params.Names,";"))
	} else {
		log.Warn("Finding users for no filter - bad request")
		return ctx.NoContent(400)
	}

	var iUserA = BaseUser{"UserA", "SurnameA", nil}
	var userA = User{iUserA, 1}

	var iUserB = BaseUser{"UserB", "SurnameB", nil}
	var userB = User{iUserB, 1}

	var retList = [2]User{userA, userB}

	return ctx.JSON(200, retList)
}

// (POST /users)
func (impl *GoMsServerImpl) AddUser(ctx echo.Context) error {
	log.Info("Adding user ")

	var iUser = BaseUser{"Added", "User", nil}
	var user = User{iUser, -1}
	return ctx.JSON(200, user)
}

