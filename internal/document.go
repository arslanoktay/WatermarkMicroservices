package internal

type Document struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Topic     string `json:"topic"`
	Watermark string `json:"watermark,omitempty"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"` // If value is empty, just return everything but sorted with the key
}

type Status string

const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "InProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
