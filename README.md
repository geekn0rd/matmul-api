# matmul-api
An HTTP server in Go accepts two matrices via a POST request and sends the matmul of them as the response. 
The code uses goroutines and channels to perform matrix multiplication in parallel, utilizing concurrent execution to potentially improve performance. 
# Usage 
First, you should go build and then run the server, 
after that, with the curl command, you can make POST requests and get the appropriate responses.
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "matrixA": {
    "rows": 2,
    "cols": 3,
    "data": [
      [1, 2, 3],
      [4, 5, 6]
    ]
  },
  "matrixB": {
    "rows": 3,
    "cols": 2,
    "data": [
      [7, 8],
      [9, 10],
      [11, 12]
    ]
  }
}' http://localhost:1379/matmul
```
we can expect to receive the following response:
```json
{
  "rows": 2,
  "cols": 2,
  "data": [
    [58, 64],
    [139, 154]
  ]
}

```
