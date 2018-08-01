//go:generate protoc --go_out=plugins=grpc:. service.proto

package osiw

import "encoding/json"

type Attachment struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	AuthorName string `json:"author_name"`
}

type Payload struct {
	Attachments []*Attachment `json:"attachments"`
}

func (p *PostRequest) GetPayload() ([]byte, error) {
	payload := &Payload{
		Attachments: []*Attachment{
			&Attachment{
				Title:      p.GetTitle(),
				Text:       p.GetText(),
				AuthorName: p.GetHostname(),
			},
		},
	}

	return json.Marshal(payload)
}
