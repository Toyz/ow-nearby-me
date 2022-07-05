package schemes

import (
	"encoding/json"
)

type (
	Player struct {
		Version string `json:"version"`
		Build   string `json:"build"`

		Account   string `json:"account"`
		AccountId string `json:"account_id"`

		SecondaryAccount   string `json:"secondary_account"`
		SecondaryAccountId string `json:"secondary_account_id"`

		Avatar    int64  `json:"avatar"`
		AvatarSID string `json:"avatar_sid"`

		Level int `json:"level"`

		PlayerLevelFrame    int64  `json:"pframe"`
		PlayerLevelFrameSID string `json:"pframe_sid"`

		EndorsementLevel int           `json:"elevel"`
		Endorsement      []Endorsement `json:"endors"`
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

func (p Player) FindEndorsement(id int64) Endorsement {
	for _, e := range p.Endorsement {
		if e.ID == id {
			return e
		}
	}

	return Endorsement{}
}
