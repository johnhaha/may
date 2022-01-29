package may

import (
	"time"

	"github.com/johnhaha/hakit/hadata"
)

type IndexRes struct {
	ErrRes
	Uid        string    `json:"uid"`
	PrimaryKey string    `json:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ModifyIndexRes struct {
	ErrRes
	Uid        int64     `json:"uid"`
	IndexUid   string    `json:"indexUid"`
	Status     string    `json:"status"`
	Type       string    `json:"type"`
	EnqueuedAt time.Time `json:"enqueuedAt"`
}

type ErrRes struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
	ErrorType string `json:"errorType"`
	ErrorLink string `json:"errorLink"`
}

func (e ErrRes) Error() string {
	return e.Message
}

type Doc map[string]interface{}

func (doc *Doc) Decode(data interface{}) error {
	return hadata.MapToStruct(doc, data)
}

func (doc Doc) Error() string {
	if _, ok := doc["errorCode"]; ok {
		return doc["message"].(string)
	}
	return ""
}

type SearchRes struct {
	ErrRes
	Hits             []Doc  `json:"hits"`
	Offset           int64  `json:"offset"`
	Limit            int64  `json:"limit"`
	NbHits           int64  `json:"nbHits"`
	ExhaustiveNbHits bool   `json:"exhaustiveNbHits"`
	ProcessingTimeMS int64  `json:"processingTimeMs"`
	Query            string `json:"query"`
}
