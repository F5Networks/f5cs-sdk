/*
 * account-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package account
// V1Cascade Only max may change, the rest must stay the same.
type V1Cascade string

// List of v1Cascade
const (
	UPWARD V1Cascade = "CASCADE_UPWARD"
	DOWNWARD V1Cascade = "CASCADE_DOWNWARD"
	ALL V1Cascade = "CASCADE_ALL"
	NONE V1Cascade = "CASCADE_NONE"
	MAX V1Cascade = "CASCADE_MAX"
)
