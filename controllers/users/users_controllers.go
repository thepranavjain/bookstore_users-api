package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thepranavjain/bookstore_users-api/domain/users"
	"github.com/thepranavjain/bookstore_users-api/services"
	"github.com/thepranavjain/bookstore_users-api/utils/errors"
)

func getUserId(userId string) (int64, *errors.RestErr) {
	userIdInt, userErr := strconv.ParseInt(userId, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userIdInt, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	isPartial := c.Request.Method == "PATCH"
	user.Id = userId
	result, updateErr := services.UpdateUser(user, isPartial)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	deleteErr := services.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": "user deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
