package schemes

import (
	"encoding/json"
)

type (
	Player struct {
		Version string `json:"version"`
		Build   string `json:"build"`

		Account   string `json:"account"`
		BNetAccount string `json:"bnet_account"`

		Level int `json:"level"`

		Avatar    int64  `json:"avatar"`
		PlayerLevelFrame int64 `json:"pframe"`
		NameCard int64 `json:"name_card"`
		Title int64 `json:"title"`

		Platform int `json:"platform"`

		EndorsementLevel int           `json:"elevel"`
	}

	Endorsement struct {
		ID    int64  `json:"id"`
		SID   string `json:"sid"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
)

var (
	EndorsementNames = map[string]string{
		"D80000000003944": "Shotcaller",
		"D80000000003945": "Good Teammate",
		"D80000000003946": "Sportsmanship",
	}
)

func (p Player) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p Player) MarshalIndent() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// im not worrying about this
//func (p Player) FindEndorsement(id int64) Endorsement {
//	for _, e := range p.Endorsement {
//		if e.ID == id {
//			return e
//		}
//	}
//
//	return Endorsement{}
//}
