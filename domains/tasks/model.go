package tasks

type Task struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type CreateTaskInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskNameInput struct {
	Name string `json:"name" binding:"required"`
}
