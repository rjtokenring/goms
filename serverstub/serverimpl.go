package serverstub

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rjtokenring/goms/dbaccess"
	"net/http"
	"strconv"
	"strings"
)

type GoMsServerImpl struct {
}

func (impl *GoMsServerImpl) DeleteUser(ctx echo.Context, id int64) error {
	var stringId = strconv.FormatInt(id, 10)
	log.Info("Deleting user " + stringId)
	dbaccess.DeleteUserByID(id)
	return ctx.NoContent(http.StatusNoContent)
}

func (impl *GoMsServerImpl) FindUserById(ctx echo.Context, id int64) error {
	var stringId = strconv.FormatInt(id, 10)
	log.Info("Getting user " + stringId)

	byID, nm, surnm := dbaccess.GetUserByID(id)

	var iUser = BaseUser{nm, surnm, nil}
	var user = User{iUser, byID}
	if user.Id == 0 {
		log.Warn("User not found - bad request")
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, user)
}

// (GET /users)
func (impl *GoMsServerImpl) FindUsers(ctx echo.Context, params FindUsersParams) error {
	if params.Names != nil {
		log.Info("Finding users: " + strings.Join(*params.Names, ";"))
	} else {
		log.Warn("Finding users for no filter - bad request")
		return ctx.NoContent(http.StatusBadRequest)
	}

	var iUserA = BaseUser{"UserA", "SurnameA", nil}
	var userA = User{iUserA, 1}

	var iUserB = BaseUser{"UserB", "SurnameB", nil}
	var userB = User{iUserB, 1}

	var retList = [2]User{userA, userB}

	return ctx.JSON(http.StatusOK, retList)
}

// (POST /users)
func (impl *GoMsServerImpl) AddUser(ctx echo.Context) error {

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(ctx.Request().Body); err != nil {
		log.Error("Bad request: " + err.Error())
		return ctx.NoContent(http.StatusBadRequest)
	}
	newStr := buf.String()
	log.Info("Adding user: " + newStr)

	var user User
	if err := json.Unmarshal(buf.Bytes(), &user); err != nil {
		log.Error("Bad request: " + err.Error())
		return ctx.NoContent(http.StatusBadRequest)
	}

	dbaccess.AddUser(user.Id, user.Name, user.Surname)

	return ctx.JSON(http.StatusOK, user)
}
