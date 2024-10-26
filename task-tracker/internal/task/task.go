package task

import "fmt"

const (
	STATUS_TODO        = "Todo"
	STATUS_DONE        = "Done"
	STATUS_IN_PROGRESS = "In-Progress"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTask(id int, description, status string) *Task {
	return &Task{
		Description: description,
		Status:      status,
	}
}

func (t *Task) SetStatus(status string) {
	t.Status = status
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	err:=ReadTasks(&tasks)
	fmt.Printf("%+v/n",tasks)
	return tasks, err
}

func AddTask(t Task) error {
	tasks,err := GetTasks()

	if err != nil {
		err:=WriteTasks(&[]Task{t})

		if err!=nil{
			fmt.Errorf(err.Error())
			return err
		}

	}

	tasks = append(tasks, t)

	err=WriteTasks(&tasks)
	return err
}
