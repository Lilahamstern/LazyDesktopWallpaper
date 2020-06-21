package unsplash

type Response struct {
	ID   string `json:"id"`
	Urls Url    `json:"urls"`
	User User   `json:"user"`
}

type Url struct {
	Raw string `json:"raw"`
}

type User struct {
	Name string `json:"name"`
}
