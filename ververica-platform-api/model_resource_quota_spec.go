/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

type ResourceQuotaSpec struct {
	Limits *ResourceQuotaQuantity `json:"limits,omitempty"`
	ToleratedOveruse *ResourceQuotaQuantity `json:"toleratedOveruse,omitempty"`
	Type_ string `json:"type,omitempty"`
}
