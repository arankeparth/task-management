package taskspec

type GetTaskResponse []*Task

type Task struct {
	ID          int    `json:"id" db:"id"`
	Summary     string `json:"summary" db:"summary"`
	Description string `json:"description" db:"description"`
	DueDate     string `json:"dueDate" db:"dueDate"`
	Priority    string `json:"priority" db:"priority"`
	Status      string `json:"status" db:"status"`
	Assignee    string `json:"assignee" db:"assignee"`
	DateCreated string `json:"dateCreated" db:"dateCreated"`
	Reporter    string `json:"reporter" db:"reporter"`
}
