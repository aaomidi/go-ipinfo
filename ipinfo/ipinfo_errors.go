package ipinfo

// RateLimitError is the error used when the API has reached a rate limit.
type RateLimitedError struct {
	Message string
}

// ErrorResponseError is the error used when the API returns an error response.
type ErrorResponseError struct {
	Response *ErrorResponse
}

// NewRateLimitedError returns a RateLimitedError but with the message constructed.
func NewRateLimitedError() *RateLimitedError {
	return &RateLimitedError{Message: "Rate limit reached."}
}

// NewErrorResponseError returns an ErrorResponseError with the ErrorResponse pre filled.
func NewErrorResponseError(e *ErrorResponse) *ErrorResponseError {
	return &ErrorResponseError{Response: e}
}

// Error Implements the Error interface.
func (e *RateLimitedError) Error() string {
	return e.Message
}

// Error Implements the Error interface.
func (e *ErrorResponseError) Error() string {
	return e.Response.Error
}
