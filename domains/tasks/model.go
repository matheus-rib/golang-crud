package tasks

type Task struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type createTaskInput struct {
	Name string `json:"name" binding:"required"`
}

type updateTaskNameInput struct {
	Name string `json:"name" binding:"required"`
}
