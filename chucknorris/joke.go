package chucknorris

import (
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	fmt.Println(string(b))
	date, err := time.Parse(`"2006-01-02 15:04:05.000000"`, string(b))
	if err != nil {
		return err
	}

	ct.Time = date
	return nil
}

type Joke struct {
	Categories []string   `json:"categories"`
	CreatedAt  CustomTime `json:"created_at"`
	IconUrl    string     `json:"icon_url"`
	Id         string     `json:"id"`
	UpdatedAt  CustomTime `jsob:"updated_at"`
	Url        string     `json:"url"`
	Value      string     `json:"value"`
}
