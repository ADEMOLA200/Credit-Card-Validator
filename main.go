package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CreditCardRequest struct {
	CardNumber string `json:"card_number"`
	ExpMonth   string `json:"exp_month"`
	ExpYear    string `json:"exp_year"`
}

func main() {
	fmt.Println("running")

	// http.HandleFunc("/", handler)
	http.HandleFunc("/pages", ok)

	port := os.Getenv("PORT")

	fmt.Println(port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func ok(w http.ResponseWriter, r *http.Request) {
	// Define your response data
	responseData := map[string]string{"status": "ok"}

	// Marshal the response data into JSON format
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	w.Write(responseJSON)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodOptions {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}

// 	// Set CORS headers for POST requests
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.Header().Set("Access-Control-Allow-Credentials", "true")

// 	// Handle actual POST requests
// 	if r.Method == http.MethodPost {
// 		var creditCardReq CreditCardRequest
// 		// Decode the JSON data from the request body directly using json.Unmarshal
// 		if err := json.NewDecoder(r.Body).Decode(&creditCardReq); err != nil {
// 			// If there's an error decoding JSON, respond with a 400 Bad Request status
// 			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// 			return
// 		}

// 		valid, err := validateCreditCard(creditCardReq.CardNumber, creditCardReq.ExpMonth, creditCardReq.ExpYear)
// 		if err != nil {
// 			http.Error(w, "Error validating credit card", http.StatusInternalServerError)
// 			return
// 		}

// 		response := map[string]interface{}{
// 			"valid":   valid,
// 			"message": getMessage(valid),
// 		}
// 		writeJSONResponse(w, http.StatusOK, response)
// 		return
// 	}

// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// }

// // New helper function to get the message based on card validity
// func getMessage(valid bool) string {
// 	if valid {
// 		return "Card is valid"
// 	}
// 	return "Card is not valid"
// }

// // Luhn algorithm-based credit card number validation
// func isValidCreditCardNumber(cardNumber string) bool {
// 	sum := 0
// 	shouldDouble := false

// 	// Iterate through each digit in reverse order
// 	for i := len(cardNumber) - 1; i >= 0; i-- {
// 		digit := int(cardNumber[i] - '0')

// 		if shouldDouble {
// 			digit *= 2
// 			if digit > 9 {
// 				digit -= 9
// 			}
// 		}

// 		sum += digit
// 		shouldDouble = !shouldDouble
// 	}

// 	// The card number is valid if the sum is a multiple of 10
// 	return sum%10 == 0
// }

// func validateCreditCard(cardNumber, expMonth, expYear string) (bool, error) {
// 	// Check if cardNumber is empty
// 	if cardNumber == "" {
// 		return false, errors.New("card_number is required")
// 	}

// 	// Check if expMonth is empty or not a valid integer
// 	month, err := strconv.Atoi(expMonth)
// 	if err != nil || month < 1 || month > 12 {
// 		return false, errors.New("invalid exp_month")
// 	}

// 	// Check if expYear is empty or not a valid integer
// 	year, err := strconv.Atoi(expYear)
// 	if err != nil || year < 2000 || year > 2100 {
// 		return false, errors.New("invalid exp_year")
// 	}

// 	// Check if card is valid using Luhn algorithm
// 	if !isValidCreditCardNumber(cardNumber) {
// 		return false, nil
// 	}

// 	// If there's no error, the card is valid
// 	return true, nil
// }

// // Helper function to write JSON response
// func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
// 	}
// }
