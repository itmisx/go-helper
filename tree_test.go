package helper

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTree(*testing.T) {
	list := []struct {
		ID   int    `json:"id"`
		PID  int    `json:"pid"`
		Name string `json:"name"`
	}{{ID: 1, PID: 0, Name: "smally1"}, {ID: 2, PID: 1, Name: "small2"}}
	l, _ := NewTree("ID", "PID").ListToTree(0, list)
	js, _ := json.Marshal(l[0])
	fmt.Println(string(js))
}
