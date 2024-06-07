package pkg

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) middleware(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}

type MatrixRequest struct {
	Matrix [][]int `json:"matrix"`
}

func (h *Handler) floyd(c *gin.Context) {
	var inputMatr MatrixRequest
	if err := c.ShouldBindJSON(&inputMatr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	floyd := h.services.floyd(parseMatr(inputMatr.Matrix))

	c.JSON(http.StatusOK, gin.H{"result": floyd})
}
func parseMatr(matr [][]int) [][]int {
	for i := 0; i < len(matr); i++ {
		for j := 0; j < len(matr[i]); j++ {
			if matr[i][j] == 0 && i != j {
				matr[i][j] = math.MaxInt32
			}
		}
	}
	return matr
}
