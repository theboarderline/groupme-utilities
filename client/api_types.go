package groupme

type APIResponse struct {
	Meta struct {
		Code   int      `json:"code"`
		Errors []string `json:"errors"`
	} `json:"meta"`
	Response struct {
		Count    int       `json:"count"`
		Messages []Message `json:"messages"`
	} `json:"response"`
}

type Request struct {
	BotID      string `json:"bot_id"`
	Text       string `json:"text"`
	PictureURL string `json:"picture_url"`
}

type ImageServiceResponse struct {
	Payload ImageServicePayload `json:"payload"`
}
type ImageServicePayload struct {
	URL        string `json:"BotMessagePostURL"`
	PictureURL string `json:"picture_url"`
}

type MessageResponse struct {
	Message  string
	ImageURL string
	Error    error
}
