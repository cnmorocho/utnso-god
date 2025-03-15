package dtos

type CreateProcessRequestDTO struct {
	FilePath       string `json:"file_path"`
	ProcessSize    int    `json:"process_size"`
	ThreadPriority int    `json:"thread_priority"`
}
