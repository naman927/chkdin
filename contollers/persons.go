package contollers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naman-dave/chkdin/modals"
)

// get person's detail based on id
func APIGetPersonDetails(c *gin.Context) {
	resp := map[string]interface{}{}
	ids := c.Param("id")

	id, err := ValidateID(ids)
	if err != nil {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	person, err := modals.NewPerson()
	if err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := person.GetPersonDetails(id); err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["data"] = person
	resp["error"] = nil
	resp["message"] = "Successfully retrived person's details"
	c.JSON(http.StatusOK, resp)
}

// create person
func APICreatePerson(c *gin.Context) {
	resp := map[string]interface{}{}

	person, err := modals.NewPerson()
	if err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := c.ShouldBindJSON(&person); err != nil {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = "unmarshal error"
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if person.FirstName == "" || person.Email == "" {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = "first_name as well as email is required"
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err := person.CreatePerson(); err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["data"] = nil
	resp["error"] = nil
	resp["message"] = "Successfully created a person"
	c.JSON(http.StatusOK, resp)
}

// delete person based on id
func APIDeletePerson(c *gin.Context) {
	resp := map[string]interface{}{}
	ids := c.Param("id")

	id, err := ValidateID(ids)
	if err != nil {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	person, err := modals.NewPerson()
	if err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := person.DeletePerson(id); err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["data"] = nil
	resp["error"] = nil
	resp["message"] = "Successfully deleted a person"
	c.JSON(http.StatusOK, resp)
}

// edit person based on id
func APIEditPerson(c *gin.Context) {
	resp := map[string]interface{}{}

	person, err := modals.NewPerson()
	if err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if err := c.ShouldBindJSON(person); err != nil {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = "unmarshal error"
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if person.ID <= 0 {
		resp["data"] = nil
		resp["error"] = "unprocessable entity"
		resp["message"] = "id should be passed with body to update person"
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err := person.UpdatePerson(); err != nil {
		resp["data"] = nil
		resp["error"] = "internal server error"
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["data"] = nil
	resp["error"] = nil
	resp["message"] = "Successfully updated a person"
	c.JSON(http.StatusOK, resp)
}

func ValidateID(ids string) (int, error) {
	if ids == "" {
		return 0, errors.New("id should be passed")
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		return 0, errors.New("id should be numaric")
	}

	return id, nil
}
