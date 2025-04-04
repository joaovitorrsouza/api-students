package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joaovitorrsouza/api-students/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Handler
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}
	student, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Status not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) createStudent(c echo.Context) error {

	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}
	return c.String(http.StatusOK, "Create student")
}

func (api *API) updateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to update student")
	}

	receivedStudent := db.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Status not found")
	}

	student := updateStudentInfo(receivedStudent, updatingStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")

	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStud)

}

func updateStudentInfo(receivedStudent, student db.Student) db.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}

	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}

	if receivedStudent.CPF != "" {
		student.CPF = receivedStudent.CPF
	}

	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}
	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}
	return student
}
