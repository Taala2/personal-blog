package internal

import (
	"encoding/json"
	"os"
	"time"
)

type Article struct {
	Name    string    `json:"name"`
	Date    time.Time `json:"time"`
	Content string    `json:"content"`
}

func SaveArticle(art Article) error {
	file, err := os.Create(art.Name + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(&art)
}

func UpdateArticle(filename string) {
	file, err := os.OpenFile(filename,)
}
