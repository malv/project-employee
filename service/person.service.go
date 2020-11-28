package service

import (
	"encoding/json"
	"log"
)

func GetPersons() (data []map[string]interface{}, e error) {
	respBody, err := GetRequest("http://camskoleksi.com:8090/api/person")
	m := make([]map[string]interface{}, 0, 0)
	err = json.Unmarshal(respBody, &m)
	return m, err
}

func GetPersonById(id string) (data map[string]interface{}, e error) {
	log.Print(id)
	respBody, err := GetRequest("http://camskoleksi.com:8090/api/person/" + id)
	m := make(map[string]interface{})
	err = json.Unmarshal(respBody, &m)
	log.Print(string(respBody))
	return m, err
}

func GetPersonByIdWithToken(id string, token string) (data map[string]interface{}, e error) {
	log.Print(id)
	respBody, err := GetRequestWithProto("http://camskoleksi.com:8090/api/person/"+id, token)
	m := make(map[string]interface{})
	err = json.Unmarshal(respBody, &m)
	log.Print(string(respBody))
	return m, err
}
