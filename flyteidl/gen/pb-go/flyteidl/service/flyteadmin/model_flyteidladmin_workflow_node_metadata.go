/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type FlyteidladminWorkflowNodeMetadata struct {
	// The identifier for a workflow execution launched by a node.
	ExecutionId *CoreWorkflowExecutionIdentifier `json:"executionId,omitempty"`
}
