/*
 * Titan API
 *
 * API used by the Titan CLI
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package titan-client
// ApiError struct for ApiError
type ApiError struct {
	// Unique code for the error
	Code string `json:"code,omitempty"`
	// Human readable description of the error
	Message string `json:"message"`
	// Additional details, such as server-side stack trace
	Details string `json:"details,omitempty"`
}
