package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()

	log.Println("JSON API server running on port: ", s.listenAddr)

	mux.HandleFunc("GET /api/classify-number", classifyNumberHandler)

	http.ListenAndServe(s.listenAddr, mux)
}

func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("number")
	fmt.Println(param)

	number, err := strconv.Atoi(param)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Number: param, Error: true})
		return
	}

	fmt.Println(number)

	response, err := classifyNumber(number)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An unknown error occurred"))
		return
	}

	err = WriteJSON(w, http.StatusOK, response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

}

func getFunFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d?json", n)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var numbersApiResponse NumbersAPIResponse

	err = json.Unmarshal(body, &numbersApiResponse)

	if err != nil {
		return "", err
	}

	return numbersApiResponse.Text, nil
}

func classifyNumber(n int) (*ClassifyNumberResponse, error) {
	var result ClassifyNumberResponse

	funFact, err := getFunFact(n)

	if err != nil {
		return nil, err
	}

	isPrime := isPrime(n)

	isPerfect := isPerfect(n)

	digitSum := digitSum(n)

	isOdd := isOdd(n)

	isEven := isEven(n)

	isArmstrong := isArmstrong(n)

	if isArmstrong {
		result.Properties = append(result.Properties, ARMSTRONG)
	}

	if isOdd {
		result.Properties = append(result.Properties, ODD)
	}

	if isEven {
		result.Properties = append(result.Properties, EVEN)
	}

	result.FunFact = funFact

	result.IsPerfect = isPerfect

	result.IsPrime = isPrime

	result.DigitSum = digitSum

	result.Number = n

	return &result, nil

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
