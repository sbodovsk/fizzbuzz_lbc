package api

// FizzBuzzRes reprensents json struct returned by fizzbuzz controller
type FizzBuzzRes struct {
	Result []string `json:"result,omitempty"`
	Error  string   `json:"error,omitempty"`
}
