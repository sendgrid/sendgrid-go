package jwt

import "fmt"

type TaskRouterGrant struct {
	WorkspaceSid string `json:"workspace_sid"`
	WorkerSid    string `json:"worker_sid"`
	Role         string `json:"role"`
}

func (taskRouterGrant *TaskRouterGrant) Key() string {
	return "task_router"
}

func (taskRouterGrant *TaskRouterGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})

	if taskRouterGrant.WorkspaceSid != "" {
		grant["workspace_sid"] = taskRouterGrant.WorkspaceSid
	}
	if taskRouterGrant.WorkerSid != "" {
		grant["worker_sid"] = taskRouterGrant.WorkerSid
	}
	if taskRouterGrant.Role != "" {
		grant["role"] = taskRouterGrant.Role
	}

	return grant
}

func (taskRouterGrant *TaskRouterGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "TaskRouterGrant", taskRouterGrant.ToPayload())
}
