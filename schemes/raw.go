package schemes

type RawAccountStruct struct {
	Version          string `json:"version"`
	Build            string `json:"build"`
	Account          string `json:"account"`
	SecondaryAccount string `json:"secondary_account"`
	Avatar           int64  `json:"avatar"`
	Level            int    `json:"level"`
	Pframe           int64  `json:"pframe"`
	Platform         int    `json:"platform"`
	Elevel           int    `json:"elevel"`
	Endors           []struct {
		ID    int64 `json:"id"`
		Count int   `json:"count"`
	} `json:"endors"`
}
