package groupme

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

const (
	RequestLimit = "100"
)

type Client struct {
	BotID       string
	GroupID     string
	AccessToken string
	Messages    []Message
	httpClient  *resty.Client
}

func NewClient(botID, groupID, accessToken string) *Client {
	return &Client{
		BotID:       botID,
		GroupID:     groupID,
		AccessToken: accessToken,
		httpClient:  resty.New(),
	}
}

func (c Client) GetTopMemeBetweenDates(startDate, endDate time.Time) (message Message, err error) {
	messages, err := c.GetMemesInWindow(&startDate, &endDate)
	if err != nil {
		log.Err(err).Msg("could not get Memes from day")
		return Message{}, err
	}

	memes := GetTopNMessages(messages, 1)
	if err != nil {
		log.Err(err).Msg("could not get top meme from day")
		return Message{}, err
	}

	if len(memes) == 0 {
		return Message{}, nil
	}

	return memes[0], nil
}

func (c Client) GetMemesInWindow(startDate, endDate *time.Time) (messages []Message, err error) {

	if startDate == nil || endDate == nil {
		start := time.Now().AddDate(0, -1, 0)
		startDate = &start
		end := time.Now()
		endDate = &end
	}

	var beforeID string

	for {
		url := fmt.Sprintf("https://api.groupme.com/v3/groups/%s/messages", c.GroupID)
		response, err := c.httpClient.R().
			SetQueryParams(map[string]string{
				"token":     c.AccessToken,
				"before_id": beforeID,
				"limit":     RequestLimit,
			}).
			Get(url)

		if err != nil {
			log.Error().Err(err).Msg("could not get messages from groupme api")
			return nil, err
		}

		if response.IsError() {
			err = errors.New("groupme api returned error response")
			log.Error().Err(err).Msgf("status: %d response: %s", response.StatusCode(), string(response.Body()))
			return nil, err
		}

		if response.StatusCode() == http.StatusOK {
			log.Info().Msg("successfully retrieved messages from groupme api")
			var res APIResponse
			if err = json.Unmarshal(response.Body(), &res); err != nil {
				log.Err(err).Msg("could not unmarshal response from groupme api")
				return nil, err
			}

			if len(res.Response.Messages) == 0 || res.Response.Messages[0].SentBefore(*startDate) {
				break
			}

			beforeID = res.Response.Messages[len(res.Response.Messages)-1].ID

			messages = append(messages, FilterMemesByTimespan(res.Response.Messages, *startDate, *endDate)...)
		}

	}

	c.Messages = messages
	return messages, nil
}

func (c Client) SendMessage(text, pictureURL string) error {

	messageRequest := Request{
		BotID:      c.BotID,
		Text:       text,
		PictureURL: pictureURL,
	}

	log.Debug().Msg("sending message to groupme api")

	response, err := c.httpClient.R().
		SetQueryParams(map[string]string{
			"token": c.AccessToken,
		}).
		SetBody(messageRequest).
		Post(BotMessagePostURL)

	if err != nil {
		log.Err(err).Msg("could not send message to groupme api")
		return err
	}

	if response.IsError() {
		err = errors.New(string(response.Body()))
		log.Err(err).Msgf("status: %d", response.StatusCode())
		return err
	}

	return nil
}

func (c Client) UploadPicture(file io.Reader) (string, error) {

	response, err := c.httpClient.R().
		SetHeader(AccessTokenHeaderKey, c.AccessToken).
		SetHeader(ContentTypeHeaderKey, ImageJPEGContentType).
		SetBody(file).
		Post(ImageServiceURL)

	if err != nil {
		log.Err(err).Msg("could not upload picture to groupme api")
		return "", err
	}

	if response.IsError() {
		err = errors.New(string(response.Body()))
		log.Err(err).Msgf("status: %d", response.StatusCode())
		return "", err
	}

	log.Debug().Msg(string(response.Body()))

	var imageResponse ImageServiceResponse
	if err = json.Unmarshal(response.Body(), &imageResponse); err != nil {
		log.Err(err).Msg("could not unmarshal response from groupme api")
		return "", err
	}

	return imageResponse.Payload.URL, nil
}
