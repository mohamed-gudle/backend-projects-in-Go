package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Activity struct {
	Activity_type string `json:"type"`
	Repo          struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func NewActivity(activity_type string) *Activity {
	return &Activity{
		Activity_type: activity_type,
	}
}

func GetActivities(username string) ([]Activity, error) {
	response, err := http.Get("https://api.github.com/users/" + username + "/events")

	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}

	var activities []Activity

	print(data)

	err = json.Unmarshal(data, &activities)

	if err != nil {
		print(err.Error())
		fmt.Errorf(err.Error())
		return nil, err
	}

	fmt.Printf("%+v", activities)
	return nil, nil
}
