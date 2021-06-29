package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "https://mysibsau.ru/v2/"

func get_all_groups() {
	resp, err := http.Get(URL + "timetable/all_groups/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func get_timetable_by_group(group_id string) {
	resp, err := http.Get(URL + "timetable/group/" + group_id + "/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func get_group_id() (id string) {
	return "id"
}
