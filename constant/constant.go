package constant

const (
	ProjectID = "cascade-masters"
	TopicID   = "dev-example-topic"
	SubName   = "your-subscription-name"
)

type Message struct {
	Data string `json:"data"`
}
