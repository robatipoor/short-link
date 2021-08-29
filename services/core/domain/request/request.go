package request

type Redirect struct {
	ShortUrl string
}

type CreateNewUrl struct {
	OriginalUrl string `json:"original_url"`
	ExpireTime  int    `json:"expire_time"`
}

type DeleteUrl struct {
	ShortUrl string
}
