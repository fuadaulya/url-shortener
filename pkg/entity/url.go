package entity

type URL struct {
	ID        int    `json:"id"`
	URLShort  string `json:"url_short"`
	URLTarget string `json:"url_target"`
}
