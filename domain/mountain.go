package domain

type Mountain struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	ReadingJP     string    `json:"reading_jp"`
	ReadingEN     string    `json:"reading_en"`
	Height        int       `json:"height"`
	Location      [2]string `json:"location"`
	MountainRange string    `json:"mountain_range"`
	Image         string    `json:"image"`
}

type Mountains []Mountain
