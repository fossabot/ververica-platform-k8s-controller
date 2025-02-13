/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

type JsonNode struct {
	ValueNode bool `json:"valueNode,omitempty"`
	ContainerNode bool `json:"containerNode,omitempty"`
	MissingNode bool `json:"missingNode,omitempty"`
	Object bool `json:"object,omitempty"`
	NodeType string `json:"nodeType,omitempty"`
	Pojo bool `json:"pojo,omitempty"`
	Number bool `json:"number,omitempty"`
	IntegralNumber bool `json:"integralNumber,omitempty"`
	FloatingPointNumber bool `json:"floatingPointNumber,omitempty"`
	Short bool `json:"short,omitempty"`
	Int_ bool `json:"int,omitempty"`
	Long bool `json:"long,omitempty"`
	Float bool `json:"float,omitempty"`
	Double bool `json:"double,omitempty"`
	BigDecimal bool `json:"bigDecimal,omitempty"`
	BigInteger bool `json:"bigInteger,omitempty"`
	Textual bool `json:"textual,omitempty"`
	Boolean bool `json:"boolean,omitempty"`
	Binary bool `json:"binary,omitempty"`
	Array bool `json:"array,omitempty"`
	Null bool `json:"null,omitempty"`
}
