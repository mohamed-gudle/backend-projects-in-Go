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
	return tasks, err
}

func AddTask(description string) error {

	t:= *NewTask(6,description,"in-progress")

	tasks,err := GetTasks()

	if err != nil {
		err:=WriteTasks(&[]Task{t})

		if err!=nil{
			fmt.Errorf(err.Error())
			return err
		}

	}

	tasks = append(tasks,t)

	err=WriteTasks(&tasks)
	return err
}

func DeleteTask(id int) error {
	tasks,err:=GetTasks()

	if err!=nil {
		fmt.Errorf(err.Error())
		return err
	}

	for i,t:=range tasks {
		if t.ID == id {
			tasks = append(tasks[:i],tasks[i+1:]...)
			break
		}
	}

	err=WriteTasks(&tasks)
	return err
}

func UpdateTask(id int, status string) error {
	tasks,err:=GetTasks()

	if err!=nil {
		fmt.Errorf(err.Error())
		return err
	}

	for i,t:=range tasks {
		if t.ID == id {
			fmt.Println("updating")
			tasks[i].Status = status
			break
		}
	}
	fmt.Printf("%+v",tasks)

	err=WriteTasks(&tasks)
	return err
}

func ListTasks(filter string) error {
	tasks,err:=GetTasks()

	if err!=nil{
		return err
	}

	filteredTasks := make([]Task,len(tasks))
	switch filter {
	case "done":
		for _,t:=range tasks {
			if t.Status == "done"{
				filteredTasks = append(filteredTasks,t )
				fmt.Printf(`%+v`,t)
			}
			
		}
		return nil
	case "not-done":
		for _,t:=range tasks {
			if t.Status == "not-done"{
				filteredTasks = append(filteredTasks,t )
				fmt.Printf(`%+v`,t)
			}
		}
		return nil
	case "in-progress":
		for _,t:=range tasks {
			if t.Status == "in-progress"{
				filteredTasks = append(filteredTasks,t )
				fmt.Printf(`%+v`,t)
			}
		}
		return nil
	default:
		fmt.Printf(`%+v`,tasks)
		return nil
	}
}
