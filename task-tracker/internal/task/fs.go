package task

import (
	"encoding/json"
	"fmt"
	"os"
)



func ReadTasks(tasks *[]Task) error {
	path,err := os.Getwd()
	if err!=nil {
		fmt.Errorf(err.Error())
	}
	path = path+ "\\tasks.json"
	data,err := os.ReadFile(path)
	if os.IsNotExist(err){
		fmt.Println("Error",err)
		return err
	}

	err=json.Unmarshal(data,tasks)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func WriteTasks(tasks *[]Task) error {
	file,err:=os.Create("tasks.json")
	defer file.Close()

	if err!=nil {
		fmt.Errorf(err.Error())
		return err
	}

	data,err:=json.Marshal(tasks)

	if err!=nil {
		fmt.Errorf(err.Error())
		return err
	}

	file.Write(data)

	return nil

}