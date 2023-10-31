package groupme

import (
	"github.com/rs/zerolog/log"
	"time"
)

func FilterMemesByTimespan(messages []Message, begin, end time.Time) (filteredMessages []Message) {

	for _, message := range messages {

		if message.SentBefore(begin) {
			break
		}

		if message.IsMeme() && message.SentDuringTimespan(begin, end) {
			log.Debug().Msgf("Found a meme on given day: %s", time.Unix(message.CreatedAt, 0))
			filteredMessages = append(filteredMessages, message)
		}
	}

	return filteredMessages
}
