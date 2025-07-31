package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type CreditCardRequest struct {
	CardNumber string `json:"card_number"`
	ExpMonth   string `json:"exp_month"`
	ExpYear    string `json:"exp_year"`
}

func main() {
	fmt.Println("running")

	http.HandleFunc("/", handler)
	http.HandleFunc("/pages", ok)

	port := os.Getenv("PORT")

	fmt.Println("running on port: ", port)

	err := http.ListenAndServe(":"+port, allowCORS(http.DefaultServeMux))
	if err != nil {
		fmt.Println(err)
	}
}

func allowCORS(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func ok(w http.ResponseWriter, r *http.Request) {
	responseData := map[string]string{"status": "ok"}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(responseJSON)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var creditCardReq CreditCardRequest
		if err := json.NewDecoder(r.Body).Decode(&creditCardReq); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		valid, err := validateCreditCard(creditCardReq.CardNumber, creditCardReq.ExpMonth, creditCardReq.ExpYear)
		if err != nil {
			http.Error(w, "Error validating credit card", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"valid":   valid,
			"message": getMessage(valid),
		}
		writeJSONResponse(w, http.StatusOK, response)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func getMessage(valid bool) string {
	if valid {
		return "Card is valid"
	}
	return "Card is not valid"
}

func isValidCreditCardNumber(cardNumber string) bool {
	sum := 0
	shouldDouble := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}

func validateCreditCard(cardNumber, expMonth, expYear string) (bool, error) {
	if cardNumber == "" {
		return false, errors.New("card_number is required")
	}

	month, err := strconv.Atoi(expMonth)
	if err != nil || month < 1 || month > 12 {
		return false, errors.New("invalid exp_month")
	}

	year, err := strconv.Atoi(expYear)
	if err != nil || year < 2000 || year > 2100 {
		return false, errors.New("invalid exp_year")
	}

	if !isValidCreditCardNumber(cardNumber) {
		return false, nil
	}

	return true, nil
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
