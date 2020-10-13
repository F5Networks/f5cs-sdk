/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
import (
	"time"
)
// V1TermsAcceptance Terms of service acceptance.
type V1TermsAcceptance struct {
	// User ID for terms of service acceptance.
	UserId string `json:"user_id,omitempty"`
	// Accept terms of service. Must be set it to 'true'.
	Accepted bool `json:"accepted,omitempty"`
	Name string `json:"name,omitempty"`
	// Date and time when terms have been accepted in ISO8601 format.
	AcceptTime time.Time `json:"accept_time,omitempty"`
}