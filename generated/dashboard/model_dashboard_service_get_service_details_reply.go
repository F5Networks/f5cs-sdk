/*
 * dashboard-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dashboard
// DashboardServiceGetServiceDetailsReply struct for DashboardServiceGetServiceDetailsReply
type DashboardServiceGetServiceDetailsReply struct {
	Type DashboardServiceServiceType `json:"type,omitempty"`
	Health DashboardServiceServiceHealth `json:"health,omitempty"`
	Name string `json:"name,omitempty"`
	Stats DashboardServiceServiceMeteringStats `json:"stats,omitempty"`
}