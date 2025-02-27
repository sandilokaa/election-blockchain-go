package domain

type Header struct {
	PrevHash string `json:"prev_hash"`
	Time     int64  `json:"time"`
}
