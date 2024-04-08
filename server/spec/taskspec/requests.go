package taskspec

type CreateTaskRequest struct {
	ID          int    `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
	DateCreated string `json:"dateCreated"`
	Reporter    string `json:"reporter"`
}

type GetTasksRequest struct {
	ID                    int    `json:"id"`
	DueDate               string `json:"dueDate"`
	DueDateComparison     string `json:"dueDateComparison"`
	Priority              string `json:"priority"`
	PriorityComparison    string `json:"priorityComparison"`
	Status                string `json:"status"`
	Assignee              string `json:"assignee"`
	DateCreated           string `json:"dateCreated"`
	DateCreatedComparison string `json:"dateCreatedComparison"`
	Reporter              string `json:"reporter"`
}

type Comment struct {
	ID          int    `json:"id" db:"id"`
	Content     string `json:"content" db:"content"`
	TicketID    int    `json:"ticketid" db:"ticketid"`
	AuthorName  string `json:"author_name" db:"author_name"`
	DateCreated string `json:"date_created" db:"date_created"`
}

type Notification struct {
	ID                int    `json:"id" db:"id"`
	Type              string `json:"type" db:"type"`
	SubjectCustomerID string `json:"subject_customer_id" db:"subject_customer_id"`
	SubjectTaskID     int    `json:"subject_task_id" db:"subject_task_id"`
	CustomerRelation  string `json:"customer_relation" db:"customer_relation"`
	TSCreated         string `json:"ts_created" db:"ts_created"`
}
