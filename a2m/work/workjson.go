package work

type Tasks struct {
	Page     float64 `json:"page"`
	PageSize float64 `json:"pageSize"`
	Tasks    []struct {
		Cancellable           bool        `json:"cancellable"`
		CancellationRequested bool        `json:"cancellationRequested"`
		Cancelled             bool        `json:"cancelled"`
		Completed             bool        `json:"completed"`
		CurrentMessage        interface{} `json:"currentMessage"`
		CurrentStep           float64     `json:"currentStep"`
		Done                  bool        `json:"done"`
		EndDate               string      `json:"endDate"`
		EtaInSeconds          float64     `json:"etaInSeconds"`
		Failed                bool        `json:"failed"`
		FailureReason         interface{} `json:"failureReason"`
		HighPriority          bool        `json:"highPriority"`
		ID                    string      `json:"id"`
		Parameters            struct{}    `json:"parameters"`
		PercentDone           float64     `json:"percentDone"`
		RequeueRequested      bool        `json:"requeueRequested"`
		StartDate             string      `json:"startDate"`
		Started               bool        `json:"started"`
		Starting              bool        `json:"starting"`
		Status                string      `json:"status"`
		TaskKey               string      `json:"taskKey"`
		TaskName              string      `json:"taskName"`
		TotalSteps            float64     `json:"totalSteps"`
		WorkspaceId           interface{} `json:"workspaceId"`
	} `json:"tasks"`
	Total float64 `json:"total"`
}

type Catalog struct {
	Available []struct {
		Description         string        `json:"description"`
		Name                string        `json:"name"`
		Parameters          []interface{} `json:"parameters"`
		RequiresWorkspaceId bool          `json:"requiresWorkspaceId"`
	} `json:"available"`
}
