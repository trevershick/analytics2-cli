package info
// generated from http://mervine.net/json2struct

type CollectionNameAndSize struct {
	Name             string  `json:"name"`
	TotalStorageSize float64 `json:"totalStorageSize"`
}
type WorkspaceInfo struct {
	ApiHalted   interface{} `json:"apiHalted"`
	Collections []CollectionNameAndSize `json:"collections"`
	Database             string      `json:"database"`
	EarliestRevisionDate string      `json:"earliestRevisionDate"`
	EtlDate              interface{} `json:"etlDate"`
	Halted               bool        `json:"halted"`
	LastRebuild          struct {
		Data struct {
			EarliestRevisionDate          string  `json:"earliestRevisionDate"`
			ElapsedOperationTimeInSeconds float64 `json:"elapsedOperationTimeInSeconds"`
			LastDataRefreshTimestamp      string  `json:"lastDataRefreshTimestamp"`
			RevisionsRemainingInQueue     float64 `json:"revisionsRemainingInQueue"`
		} `json:"data"`
		OperationType  string  `json:"operationType"`
		SpecifiedDate  string  `json:"specifiedDate"`
		SubscriptionId float64 `json:"subscriptionId"`
		Timestamp      string  `json:"timestamp"`
		WorkspaceOid   float64 `json:"workspaceOid"`
	} `json:"lastRebuild"`
	Metadata []struct {
		ID struct {
			Class      string  `json:"class"`
			Inc        float64 `json:"inc"`
			Machine    float64 `json:"machine"`
			New        bool    `json:"new"`
			Time       float64 `json:"time"`
			TimeSecond float64 `json:"timeSecond"`
		} `json:"_id"`
		Name   string `json:"name"`
		Status struct {
			LastRevisionDate string `json:"lastRevisionDate"`
		} `json:"status"`
	} `json:"metadata"`
	RevisionsInQueue      float64 `json:"revisionsInQueue"`
	Subscription          float64 `json:"subscription"`
	SubscriptionName      string  `json:"subscriptionName"`
	TotalStorageSize      float64 `json:"totalStorageSize"`
	Workspace             float64 `json:"workspace"`
	WorkspaceCreationDate string  `json:"workspaceCreationDate"`
}

type HaltedWorkspaceInfo struct {
	Data struct {
		Reason string `json:"reason"`
	} `json:"data"`
	HealthCheckShouldFail bool    `json:"healthCheckShouldFail"`
	SubscriptionId        float64 `json:"subscriptionId"`
	Timestamp             string  `json:"timestamp"`
	WorkspaceOid          float64 `json:"workspaceOid"`
}
