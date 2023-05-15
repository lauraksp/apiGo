package pkg

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"id"`
}

var tasks []Task

func AddTask(task Task) {
	task.ID = len(tasks) + 1
	tasks = append(tasks, task)
}

func GetTasks() []Task {
	return tasks
}
