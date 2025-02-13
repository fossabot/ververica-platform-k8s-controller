/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

type DeploymentTemplateSpec struct {
	Artifact *Artifact `json:"artifact"`
	Parallelism int32 `json:"parallelism,omitempty"`
	NumberOfTaskManagers int32 `json:"numberOfTaskManagers,omitempty"`
	Resources map[string]ResourceSpec `json:"resources,omitempty"`
	FlinkConfiguration map[string]string `json:"flinkConfiguration,omitempty"`
	Logging *Logging `json:"logging,omitempty"`
	Kubernetes *KubernetesOptions `json:"kubernetes,omitempty"`
}
