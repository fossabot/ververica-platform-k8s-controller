/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

type VolumeAndMount struct {
	Name string `json:"name,omitempty"`
	Volume *JsonNode `json:"volume,omitempty"`
	VolumeMount *JsonNode `json:"volumeMount,omitempty"`
}
