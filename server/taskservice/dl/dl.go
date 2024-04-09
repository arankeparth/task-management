package dl

import (
	"errors"
	"fmt"
	"task-management/server/spec/taskspec"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// TaskDL provides methods to interact with the task data
type TaskDL struct {
	db *sqlx.DB
}

// NewTaskDL creates a new TaskDL instance
func NewTaskDL(db *sqlx.DB) *TaskDL {
	return &TaskDL{
		db: db,
	}
}

func parseCondition(cond, parameter, value string) string {
	if value == "" {
		return ""
	}
	switch cond {

	case "exact":
		cond = fmt.Sprintf("%s = '%s'", parameter, value)
	case "greater_than":
		cond = fmt.Sprintf("%s > '%s'", parameter, value)
	case "smaller_than":
		cond = fmt.Sprintf("%s < '%s'", parameter, value)
	case "smaller_than_eq":
		cond = fmt.Sprintf("%s <= '%s'", parameter, value)
	case "greater_than_eq":
		cond = fmt.Sprintf("%s >= '%s'", parameter, value)
	default:
		return ""
	}

	return cond
}

// CreateTask creates a new task in the database
func (t *TaskDL) CreateTask(task *taskspec.CreateTaskRequest) error {
	task.DateCreated = time.Now().Format("2006-01-02 15:04:05")
	query := `INSERT INTO tasks (summary, description, dueDate, priority, status, assignee, dateCreated, reporter) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := t.db.Exec(query, task.Summary, task.Description, task.DueDate, task.Priority, task.Status, task.Assignee, task.DateCreated, task.Reporter)
	if err != nil {
		return err
	}
	return nil
}

// GetTasks retrieves tasks from the database
func (t *TaskDL) GetTasks(req *taskspec.GetTasksRequest) (taskspec.GetTaskResponse, error) {

	dueDateCondition := parseCondition(req.DueDateComparison, "dueDate", req.DueDate)
	priorityCondition := parseCondition(req.PriorityComparison, "priority", req.Priority)
	dateCreatedCondition := parseCondition(req.DateCreatedComparison, "dateCreated", req.DateCreated)

	query := `SELECT id, summary, description, dueDate, priority, status, assignee, dateCreated 
			  FROM tasks 
			  WHERE 1=1`
	if dueDateCondition != "" {
		query += fmt.Sprintf(" AND %s", dueDateCondition)
	}
	if priorityCondition != "" {
		query += fmt.Sprintf(" AND %s", priorityCondition)
	}
	if dateCreatedCondition != "" {
		query += fmt.Sprintf(" AND %s", dateCreatedCondition)
	}
	if len(req.Assignee) > 0 {
		query += fmt.Sprintf(" AND assignee LIKE '%%%s%%'", req.Assignee)
	}
	if len(req.Status) > 0 {
		query += fmt.Sprintf(" AND status LIKE '%%%s%%'", req.Status)
	}
	if req.ID > 0 {
		query += fmt.Sprintf(" AND id = %d", req.ID)
	}
	if req.Reporter != "" {
		query += fmt.Sprintf(" AND reporter LIKE '%%%s%%'", req.Reporter)
	}

	query += " ORDER BY dateCreated DESC"

	rows, err := t.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks taskspec.GetTaskResponse

	for rows.Next() {
		var task taskspec.Task
		err := rows.StructScan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

// GetTaskByID retrieves a task from the database based on the task ID
func (t *TaskDL) GetTaskByID(taskID int) (*taskspec.Task, error) {
	query := `SELECT id, summary, description, dued_ate, priority, status, assignee, dateCreated 
			  FROM tasks 
			  WHERE id = ?`

	var task taskspec.Task
	err := t.db.Get(&task, query, taskID)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates an existing task in the database
func (t *TaskDL) UpdateTask(task *taskspec.Task) error {
	query := `UPDATE tasks 
			  SET summary = ?, description = ?, dueDate = ?, priority = ?, status = ?, assignee = ?, dateCreated = ?
			  WHERE id = ?`

	_, err := t.db.Exec(query, task.Summary, task.Description, task.DueDate, task.Priority, task.Status, task.Assignee, task.DateCreated, task.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask deletes a task from the database
func (t *TaskDL) DeleteTask(id int) error {
	// TODO: Implement the logic to delete a task from the database
	return errors.New("not implemented")
}

func (t *TaskDL) GetComments(taskID int) ([]taskspec.Comment, error) {

	query := `SELECT * FROM comments WHERE ticketid = ? order by date_created desc`
	rows, err := t.db.Queryx(query, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []taskspec.Comment

	for rows.Next() {
		var comment taskspec.Comment
		err := rows.Scan(&comment.ID, &comment.Content, &comment.TicketID, &comment.AuthorName, &comment.DateCreated)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

// CreateComment creates a new comment for a task in the database
func (t *TaskDL) CreateComment(comment *taskspec.Comment) error {

	query := `INSERT INTO comments (ticketid, author_name, content, date_created) 
			  VALUES (?, ?, ?, ?)`
	_, err := t.db.Exec(query, comment.TicketID, comment.AuthorName, comment.Content, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}

// CreateNotification creates a new notification in the database
func (t *TaskDL) CreateNotification(notification *taskspec.Notification) error {
	query := `INSERT INTO notifications (type, subject_customer_id, customer_relation, subject_task_id, ts_created) 
			  VALUES (?, ?, ?, ?)`
	_, err := t.db.Exec(query, notification.Type, notification.SubjectCustomerID, notification.CustomerRelation, notification.SubjectTaskID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}
