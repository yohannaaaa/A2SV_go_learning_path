package domain

type Role string

const (
    RoleAdmin Role = "admin"
    Roleuser  Role = "user"
)

type Status string

const (
    StatusNotStarted Status = "not_started"
    StatusInProgress Status = "in_progress"
    StatusCompleted  Status = "completed"
)

type User struct {
    ID       string
    Username    string
    Password string
    Role     Role
}

type Task struct {
    ID          string
    UserID      string
    Title       string
    Description string
    DueDate     string
    Status      string
}

type UserRepository interface {
	Create(user *User) error
	FetchByUsername(username string) (*User, error)
}

type TaskRepository interface {
	GetAll() ([]Task, error)
    GetByID(id string) (*Task, error)
    Create(task *Task) error
    Update(id string, task *Task) error
    Delete(id string) error
}