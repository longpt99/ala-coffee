package utils

// func HandlerResponse(w http.ResponseWriter, data interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(data)
// }

// func HandlerError(w http.ResponseWriter, err error) {
// 	fmt.Printf("Error: %v\n", err)
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusInternalServerError)
// 	json.NewEncoder(w).Encode(struct {
// 		Message string `json:"message"`
// 	}{Message: err.Error()})
// }
