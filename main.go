package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var numWorkers int = 4 // Default number of workers
var numWorkersMutex sync.Mutex

type Matrix struct {
	Rows int     `json:"rows"`
	Cols int     `json:"cols"`
	Data [][]int `json:"data"`
}

func multiplyMatrices(matrixA, matrixB Matrix) Matrix {
	resultRows := matrixA.Rows
	resultCols := matrixB.Cols
	resultData := make([][]int, resultRows)

	// Use the numWorkers variable to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, numWorkers)

	// Create a mutex for synchronizing access to resultData
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(resultRows)

	for i := 0; i < resultRows; i++ {
		go func(row int) {
			defer wg.Done()

			// Acquire a semaphore
			semaphore <- struct{}{}

			// Perform matrix multiplication for the given row
			resultRow := make([]int, resultCols)
			for j := 0; j < resultCols; j++ {
				sum := 0
				for k := 0; k < matrixA.Cols; k++ {
					sum += matrixA.Data[row][k] * matrixB.Data[k][j]
				}
				resultRow[j] = sum
			}

			// Acquire the mutex to write to the resultData
			mutex.Lock()
			resultData[row] = resultRow
			mutex.Unlock()

			// Release the semaphore
			<-semaphore
		}(i)
	}

	wg.Wait()

	return Matrix{
		Rows: resultRows,
		Cols: resultCols,
		Data: resultData,
	}
}

func setNumberOfWorkersHandler(c *fiber.Ctx) error {
	numWorkersStr := c.Query("numWorkers")
	numOfWorkers, err := strconv.Atoi(numWorkersStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid number of workers"})
	}

	// Acquire the mutex to update the numWorkers variable
	numWorkersMutex.Lock()
	numWorkers = numOfWorkers
	numWorkersMutex.Unlock()

	return c.SendStatus(fiber.StatusOK)
}

func multiplyMatricesHandler(c *fiber.Ctx) error {
	var matrices struct {
		MatrixA Matrix `json:"matrixA"`
		MatrixB Matrix `json:"matrixB"`
	}

	if err := c.BodyParser(&matrices); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if matrices.MatrixA.Cols != matrices.MatrixB.Rows {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dimension mismatch"})
	}

	resultMatrix := multiplyMatrices(matrices.MatrixA, matrices.MatrixB)

	return c.JSON(resultMatrix)
}

func main() {
	fmt.Println("abbas")
	app := fiber.New()

	app.Post("/matmul", multiplyMatricesHandler)

	app.Get("/setNumberOfWorkers", setNumberOfWorkersHandler)

	if err := app.Listen("0.0.0.0:1379"); err != nil {
		fmt.Printf("we have error on listening %s\n", err)
	}
}
