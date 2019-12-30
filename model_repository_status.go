/*
 * Titan API
 *
 * API used by the Titan CLI
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package titan-client
// RepositoryStatus struct for RepositoryStatus
type RepositoryStatus struct {
	// The latest commit ID for the repository
	LastCommit string `json:"lastCommit,omitempty"`
	// The source commit for the current state (last checkout or commit)
	SourceCommit string `json:"sourceCommit,omitempty"`
}
