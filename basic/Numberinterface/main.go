package main

import (
	"encoding/json"
	"os"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	dateString := "2022-07-28T02:29:56.396Z"
	date, _ := time.Parse("2006-01-02", dateString)
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", date},
	)
}

//func (u *MyUser) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		ID       int64  `json:"id"`
//		Name     string `json:"name"`
//		LastSeen int64  `json:"lastSeen"`
//	}{
//		ID:       u.ID,
//		Name:     u.Name,
//		LastSeen: u.LastSeen.Unix(),
//	})
//}
