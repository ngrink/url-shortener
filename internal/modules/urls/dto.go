package urls

type CreateUrlDto struct {
	OriginalUrl string `json:"original_url"`
	CustomKey   string `json:"custom_key"`
}
