package structs

type Anime struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description_anime"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
}
