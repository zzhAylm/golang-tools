package testing

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMonster_Save(t *testing.T) {

	monster := Monster{
		Age:   20,
		Name:  "zzh",
		Skill: "skill",
	}
	monster.Save()
}
func TestMonster_Download(t *testing.T) {

	monster := Monster{}
	monster.Download()
	marshal, err := json.Marshal(monster)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
