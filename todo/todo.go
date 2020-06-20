package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sureeratend/finalexam/databases"
)

func getCustomersHandler(c *gin.Context) {
	rows, err := databases.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	cuss := []Customer{}
	for rows.Next() {
		cus := Customer{}
		rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		cuss = append(cuss, cus)
	}
	c.JSON(http.StatusOK, cuss)
}
func getCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")

	row, err := databases.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	cus := &Customer{}

	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cus)
}

func deleteCustomerByIDHandler(c *gin.Context) {

	id := c.Param("id")
	err := databases.DeleteCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	var m = Message{Message: "customer deleted"}

	c.JSON(http.StatusOK, m)
}
func updateCustomerByIDHandler(c *gin.Context) {

	id := c.Param("id")

	row, err := databases.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cus := &Customer{}

	err = row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	err = c.ShouldBindJSON(cus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	databases.UpdateCustomerByID(id, cus.Name, cus.Email, cus.Status)

	c.JSON(http.StatusOK, cus)
}

func createCustomersHandler(c *gin.Context) {
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := databases.Conn().QueryRow("INSERT INTO customers (name,email,status) values ($1, $2,$3)  RETURNING id", cus.Name, cus.Email, cus.Status)

	err := row.Scan(&cus.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, cus)
}

func authMiddleware(c *gin.Context) {

	token := c.GetHeader("Authorization")
	fmt.Println("start #middleware token", token)
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have the right!!"})
		c.Abort()
		return
	}

	c.Next()

	fmt.Println("end #middleware")

}

// SetupRouter  xxxx
func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("")

	apiV1.Use(authMiddleware)

	apiV1.POST("/customers", createCustomersHandler)
	apiV1.GET("/customers/:id", getCustomerByIDHandler)
	apiV1.GET("/customers", getCustomersHandler)
	apiV1.DELETE("/customers/:id", deleteCustomerByIDHandler)
	apiV1.PUT("/customers/:id", updateCustomerByIDHandler)

	return r
}
