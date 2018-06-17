package ipinfo

import "fmt"

// RateLimitedError is the error used when the API has reached a rate limit.
type RateLimitedError struct {
	Message string
}

// ErrorResponseError is the error used when the API returns an error response.
type ErrorResponseError struct {
	Response *ErrorResponse
}

// NoSuchCountryError is the error used to tell the user that the ISO2 of the response was incorrect.
type NoSuchCountryError struct {
	CountryCode string
}

// NewRateLimitedError returns a RateLimitedError but with the message constructed.
func NewRateLimitedError() *RateLimitedError {
	return &RateLimitedError{Message: "Rate limit reached."}
}

// NewErrorResponseError returns an ErrorResponseError with the ErrorResponse pre filled.
func NewErrorResponseError(e *ErrorResponse) *ErrorResponseError {
	return &ErrorResponseError{Response: e}
}

// NewNoSuchCountryError returns a NoSuchCountryError with the country code prefilled.
func NewNoSuchCountryError(countryCode string) *NoSuchCountryError {
	return &NoSuchCountryError{CountryCode: countryCode}
}

// Error Implements the Error interface.
func (e *RateLimitedError) Error() string {
	return e.Message
}

// Error Implements the Error interface.
func (e *ErrorResponseError) Error() string {
	return e.Response.Error
}

// Error Implements the Error interface.
func (e *NoSuchCountryError) Error() string {
	return fmt.Sprintf("Country code %s not found.", e.CountryCode)
}
