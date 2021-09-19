package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

type timesStruct struct {
	MaxTimes struct {
		Fri int `json:"Fri"`
		Mon int `json:"Mon"`
		Sat int `json:"Sat"`
		Sun int `json:"Sun"`
		Thu int `json:"Thu"`
		Tue int `json:"Tue"`
		Wed int `json:"Wed"`
	} `json:"maxTimes"`
	TimeLeft struct {
		Left int    `json:"left"`
		Date string `json:"date"`
	} `json:"timeLeft"`
}

//get timelimit of today
func (times timesStruct) getLimit(now time.Time) int {
	switch now.Format("Mon") {
	case "Mon":
		return times.MaxTimes.Mon
	case "Tue":
		return times.MaxTimes.Tue
	case "Wed":
		return times.MaxTimes.Wed
	case "Thu":
		return times.MaxTimes.Thu
	case "Fri":
		return times.MaxTimes.Fri
	case "Sat":
		return times.MaxTimes.Sat
	case "Sun":
		return times.MaxTimes.Sun
	}
	return 4
}

func main() {
	//get json data
	data, err := os.ReadFile("./time.json")
	if err != nil {
		//if file doesnt exist
		data = []byte(`{"maxTimes":{"Fri":120,"Mon":120,"Sat":120,"Sun":120,"Thu":120,"Tue":120,"Wed":120},"timeLeft":{"left":0,"date":"` + time.Now().Format("01-02-2006") + `"}}`)
		err = os.WriteFile("./time.json", data, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
	var times timesStruct

	err = json.Unmarshal([]byte(data), &times)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(times.TimeLeft.Left)

	startTime := time.Now()
	now := time.Now()
	var timeLeft int

	//time for logout
	endTime := startTime.Add(time.Minute * time.Duration(times.TimeLeft.Left))

	for range time.Tick(time.Minute * 1) {
		//current time
		now = time.Now()
		//if next day started, start timer from limit
		if times.TimeLeft.Date != now.Format("01-02-2006") {
			startTime = time.Now()
			times.TimeLeft.Date = now.Format("01-02-2006")
			times.TimeLeft.Left = times.getLimit(now)
			endTime = startTime.Add(time.Minute * time.Duration(times.TimeLeft.Left))
		}

		//time left
		timeLeft = int(endTime.Sub(now).Minutes())
		times.TimeLeft.Left = timeLeft

		//notify when 10 and 2 minutes left
		if timeLeft == 9 || timeLeft == 1 {
			err := beeep.Notify("TimeLimiter", strconv.Itoa(timeLeft+1)+" minutes left", "")
			if err != nil {
				panic(err)
			}
		}

		//write to json
		res, err := json.Marshal(times)
		if err != nil {
			fmt.Println(err)
		}
		err = os.WriteFile("./time.json", res, 0644)
		if err != nil {
			fmt.Println(err)
		}

		//logout after time runs out
		if now.After(endTime) {
			if err := exec.Command("cmd", "/C", "shutdown", "/l").Run(); err != nil {
				fmt.Println("Failed to initiate logout:", err)
			}
		}
	}
}
