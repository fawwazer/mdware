package task

// CreateTaskResponse represents the response returned by the CreateTask endpoint.
type CreateTaskResponse struct {
	Message string `json:"message"`
	Task    Task   `json:"task"`
}

// UpdateTaskResponse represents the response returned by the UpdateTask endpoint.
type UpdateTaskResponse struct {
	Message string `json:"message"`
	Task    Task   `json:"task"`
}

// Define any other response structures you may need for other endpoints...

// Task represents the task data structure.
type Task struct {
	ID       uint   `json:"id"`
	Taskname string `json:"taskname"`
	Pelaku   string `json:"pelaku"`
	// Add any other fields you need...
}
