package tests

import (
	"github.com/lyft/flyteadmin/pkg/manager/mocks"
	"github.com/lyft/flyteadmin/pkg/rpc/adminservice"
	mockScope "github.com/lyft/flytestdlib/promutils"
)

type NewMockAdminServerInput struct {
	executionManager               *mocks.MockExecutionManager
	launchPlanManager              *mocks.MockLaunchPlanManager
	nodeExecutionManager           *mocks.MockNodeExecutionManager
	projectManager                 *mocks.MockProjectManager
	projectDomainAttributesManager *mocks.MockProjectDomainAttributesManager
	taskManager                    *mocks.MockTaskManager
	workflowManager                *mocks.MockWorkflowManager
	taskExecutionManager           *mocks.MockTaskExecutionManager
}

func NewMockAdminServer(input NewMockAdminServerInput) *adminservice.AdminService {
	var testScope = mockScope.NewTestScope()
	return &adminservice.AdminService{
		ExecutionManager:               input.executionManager,
		LaunchPlanManager:              input.launchPlanManager,
		NodeExecutionManager:           input.nodeExecutionManager,
		TaskManager:                    input.taskManager,
		ProjectManager:                 input.projectManager,
		ProjectDomainAttributesManager: input.projectDomainAttributesManager,
		WorkflowManager:                input.workflowManager,
		TaskExecutionManager:           input.taskExecutionManager,
		Metrics:                        adminservice.InitMetrics(testScope),
	}
}
