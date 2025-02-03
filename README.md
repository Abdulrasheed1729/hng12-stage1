## HNG 12 Stage 1 Task

Submission for [HNG](https://hng.tech) 12 internship stage 1 task.
This is a simple backend API built with the [Go](https://go.dev) programming language.

### API Info

A number clssification API that uses gives some interesting properties of a number and also returns a fun fact about the number using the [Numbers API](http://numbersapi.com/).

### Programming language
The [Go](https://go.dev) programming language

Backlink: https://hng.tech/hire/golang-developers

### API Endpoints
There is only one endpoint in the API:

- `"/api/classify-number?number=371"`: which has the following sample JSON response
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```
On any bad request error we return the following format
```json
{
    "number": "alphabet",
    "error": true
}
```

### Usage

```sh
curl -X GET https://hng12-stage0-f7pg.onrender.com/api/classify-number?=number=371
```

### Local Setup

To run this project locally following the following steps:
- Install the [Go](https://go.dev) programming language on your machine.
- Clone the repository

```sh
git clone https://github.com/Abdulrasheed1729/hng12-stage1

cd hng12-stage0/
```

- Build the project
```
go build .
```

- Run the project
```
./hng12-stage1
```

- Make a Sample request
```
curl -X GET http://localhost:8080
```
