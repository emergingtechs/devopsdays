package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type InputSchedule map[string][]InputScheduleObject

type InputScheduleObject struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	TalkTopic string `json:"talk_topic"`
	Speaker   string `json:"speaker"`
}

type SessionData struct {
	ID          int    `yaml:"id"`
	Title       string `yaml:"title"`
	Service     string `yaml:"service"`
	SubType     string `yaml:"subtype"`
	Description string `yaml:"description"`
	Speakers    []int  `yaml:"speakers"`
}

type ScheduleData struct {
	StartTime string `yaml:"startTime"`
	EndTime   string `yaml:"endTime"`
	SessionID []int  `yaml:"sessionIds"`
}

type SpeakerData struct {
	ID           int          `yaml:"id"`
	Name         string       `yaml:"name"`
	Surname      string       `yaml:"surname"`
	Company      string       `yaml:"company"`
	Title        string       `yaml:"title"`
	Bio          string       `yaml:"bio"`
	ThumbnailURL string       `yaml:"thumbnailUrl"`
	Rockstar     string       `yaml:"rockstar"`
	Social       []SocialData `yaml:"social"`
}

type SocialData struct {
	Name string `yaml:"name"`
	Link string `link:"link"`
}

func main() {
	var input InputSchedule

	jsonFile, err := os.Open("schedule_2022.json")
	if err != nil {
		log.Fatalf("cannot open file: %v", err)
	}

	contents, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("cannot read file contents: %v", err)
	}

	err = json.Unmarshal(contents, &input)
	if err != nil {
		log.Fatalf("cannot unmarshal file contents: %v", err)
	}

	speakerIndex := 0
	speakers := []SpeakerData{}

	for day, details := range input {
		schedule := []ScheduleData{}
		session := []SessionData{}

		for index, s := range details {
			id := getID(day, index)

			schedule = append(schedule, ScheduleData{
				StartTime: s.StartTime,
				EndTime:   s.EndTime,
				SessionID: []int{id},
			})

			speakerData := getSpeakers(s.Speaker, speakerIndex)
			speakers = append(speakers, speakerData...)
			speakerIndex = speakerIndex + len(speakerData)

			speakerIDs := []int{}
			for _, sData := range speakerData {
				speakerIDs = append(speakerIDs, sData.ID)
			}

			session = append(session, SessionData{
				ID:          id,
				Title:       s.TalkTopic,
				Service:     "",
				Description: "",
				Speakers:    speakerIDs,
			})

		}

		scheduleString, err := yaml.Marshal(schedule)
		if err != nil {
			log.Fatalf("cannot marshal schedule: %v", err)
		}

		sessionString, err := yaml.Marshal(session)
		if err != nil {
			log.Fatalf("cannot marshal schedule: %v", err)
		}

		fmt.Printf("Generating for: %s\n", day)

		filename := fmt.Sprintf("schedule_%s.yaml", day)
		ioutil.WriteFile(filename, scheduleString, 0777)

		filename = fmt.Sprintf("session_%s.yaml", day)
		ioutil.WriteFile(filename, sessionString, 0777)

	}
	speakerString, err := yaml.Marshal(speakers)
	if err != nil {
		log.Fatalf("cannot marshal schedule: %v", err)
	}

	filename := "speaker.yaml"
	ioutil.WriteFile(filename, speakerString, 0777)
}

func getID(day string, index int) int {
	if day != "day_1" {
		return index + 100
	}
	return index
}

func getSpeakers(speakers string, currentIndex int) []SpeakerData {
	speakerData := []SpeakerData{}
	if speakers == "" {
		return speakerData
	}

	for _, s := range strings.Split(speakers, "|") {
		name := strings.Split(s, " ")
		firstname := ""
		lastname := ""
		if len(name) == 2 {
			firstname = name[0]
			lastname = name[1]
		}
		if len(name) == 1 {
			firstname = name[0]
		}
		speakerData = append(speakerData, SpeakerData{
			ID:           currentIndex + 1,
			Name:         firstname,
			Surname:      lastname,
			Company:      "",
			Title:        "",
			Bio:          "",
			ThumbnailURL: "",
			Rockstar:     "",
			Social:       []SocialData{},
		})
		currentIndex = currentIndex + 1
	}

	return speakerData
}
