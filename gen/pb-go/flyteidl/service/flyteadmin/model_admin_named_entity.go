/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminNamedEntity struct {
	ResourceType *CoreResourceType `json:"resource_type,omitempty"`
	Id *AdminNamedEntityIdentifier `json:"id,omitempty"`
	Metadata *AdminNamedEntityMetadata `json:"metadata,omitempty"`
}
