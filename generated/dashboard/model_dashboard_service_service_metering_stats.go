/*
 * dashboard-service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dashboard
// DashboardServiceServiceMeteringStats struct for DashboardServiceServiceMeteringStats
type DashboardServiceServiceMeteringStats struct {
	Granularity DashboardServiceStatTimeGranularity `json:"granularity,omitempty"`
	StatType string `json:"stat_type,omitempty"`
	Duration string `json:"duration,omitempty"`
	Stats []DashboardServiceServiceMeteringStat `json:"stats,omitempty"`
}
