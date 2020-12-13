package apperrors

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type user struct {
	FirstName       string `form:"firstname" json:"firstname,omitempty" binding:"required,max=12,min=3"`
	LastName        string `form:"lastname" json:"lastname,omitempty" binding:"required,max=10,min=3"`
	Email           string `form:"email" json:"email,omitempty" binding:"required,email"`
	Password        string `form:"password" json:"password" binding:"required,max=10,min=6"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required,max=10,min=6"`
}

func TestNewValidatorError(t *testing.T) {
	t.Run("Nil error generated", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		testUser := `{"firstname":"Test123","lastname":"Test12345","email":"Test@123.com","password":"test12345","confirmPassword":"test12345"}`
		r := gin.Default()

		r.POST("/testregister", func(c *gin.Context) {
			var u user
			if err := c.ShouldBindJSON(&u); err != nil {
				bindErr := NewValidatorError(err)
				c.JSON(400, bindErr)
				return
			}
			c.JSON(200, gin.H{
				"message": "All Ok!!",
			})
			return

		})

		req, err := http.NewRequest("POST", "/testregister", bytes.NewBufferString(testUser))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

	})

	t.Run("error generated", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		testUser := `{"firstname":"","lastname":"T","email":"Test@123","password":"test12345","confirmPassword":"test12345"}`
		r := gin.Default()

		r.POST("/testregister", func(c *gin.Context) {
			var u user
			if err := c.ShouldBindJSON(&u); err != nil {
				bindErr := NewValidatorError(err)
				c.JSON(400, bindErr)
				return
			}
			c.JSON(200, gin.H{
				"message": "All Ok!!",
			})
			return

		})

		req, err := http.NewRequest("POST", "/testregister", bytes.NewBufferString(testUser))
		req.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)

	})

}
