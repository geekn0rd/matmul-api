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

func multiplyMatricesHandler(c *fiber.Ctx) error {
	var matrices struct {
		MatrixA Matrix `json:"matrixA"`
		MatrixB Matrix `json:"matrixB"`
	}

	if err := c.BodyParser(&matrices); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	resultMatrix := multiplyMatrices(matrices.MatrixA, matrices.MatrixB)

	return c.JSON(resultMatrix)
}

func main() {
	fmt.Println("abbas")
	app := fiber.New()

	app.Post("/matmul", multiplyMatricesHandler)

	if err := app.Listen("0.0.0.0:1379"); err != nil {
		fmt.Printf("we have error on listening %s\n", err)
	}
}
