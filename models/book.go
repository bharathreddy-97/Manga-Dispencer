package models

type Book struct {
	Name   string   `json:"name"`
	Author string   `json:"author"`
	Genres []string `json:"genres"`
}

func (b *Book) GetDataFromDict(data map[string]interface{}) {
	if name, ok := data["name"].(string); ok {
		b.Name = name
	}

	if author, ok := data["author"].(string); ok {
		b.Author = author
	}

	if genres, ok := data["genres"].([]interface{}); ok {
		for _, genre := range genres {
			if genreString, ok := genre.(string); ok {
				b.Genres = append(b.Genres, genreString)
			}
		}
	}
}
