package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Matrix struct {
	Rows int     `json:"rows"`
	Cols int     `json:"cols"`
	Data [][]int `json:"data"`
}

func multiplyMatrices(matrixA, matrixB Matrix) Matrix {
	if matrixA.Cols != matrixB.Rows {
		// Dimension mismatch
		return Matrix{}
	}

	resultRows := matrixA.Rows
	resultCols := matrixB.Cols
	resultData := make([][]int, resultRows)
	for i := 0; i < resultRows; i++ {
		resultData[i] = make([]int, resultCols)
		for j := 0; j < resultCols; j++ {
			sum := 0
			for k := 0; k < matrixA.Cols; k++ {
				sum += matrixA.Data[i][k] * matrixB.Data[k][j]
			}
			resultData[i][j] = sum
		}
	}

	return Matrix{
		Rows: resultRows,
		Cols: resultCols,
		Data: resultData,
	}
}

func main() {
	fmt.Println("abbas")
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON("Hello")
	})

	if err := app.Listen("0.0.0.0:1379"); err != nil {
		fmt.Printf("we have error on listening %s\n", err)
	}
}
