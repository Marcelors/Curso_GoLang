package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/marcelors/curso/usuarios/models"
)

//Home is page initial the application
func Home(c echo.Context) error {
	var users []models.User

	if err := models.UserModel.Find().All(&users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error fetching data",
		})
	}

	data := map[string]interface{}{
		"title": "List of Users",
		"users": users,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

// Add is page of "add" users on application
func Add(c echo.Context) error {

	data := map[string]interface{}{
		"title": "Add user",
	}

	return c.Render(http.StatusOK, "add.html", data)
}

// Edit is page of "edit" users on application
func Edit(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	var user models.User

	result := models.UserModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Could not find user",
		})
	}

	if err := result.One(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Could not find user",
		})
	}

	data := map[string]interface{}{
		"title": "Edit User",
		"user":  user,
	}

	return c.Render(http.StatusOK, "update.html", data)
}

// Save is method for save a user
func Save(c echo.Context) error {

	name := c.FormValue("name")
	email := c.FormValue("email")

	var user models.User
	user.Name = name
	user.Email = email

	if name != "" && email != "" {
		_, err := models.UserModel.Insert(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Can not add user to database",
			})
		}

		c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "The fields are required",
	})
}

// Delete is method for delete a user
func Delete(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	result := models.UserModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Nao foi possivel encontrar o usuario",
		})
	}
	if err := result.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao deletar usuario!",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "Usuario Deletado com sucesso!",
	})
}

// Update is method for update a user
func Update(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	name := c.FormValue("nome")
	email := c.FormValue("email")

	var user = models.User{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	result := models.UserModel.Find("id=?", userID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Nao foi possivel encontrar o usuario",
		})
	}

	if err := result.Update(user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Erro ao tentar atualizar o regritro!",
		})
	}

	return c.JSON(http.StatusAccepted, user)
}
