package messages

//Command container. Will be used to marshal and unmarshal messages on the event bus
type Command struct {
	Hosts []string `json:"hosts"`
	Path string `json:"path"`
	Options []string `json:"options"`
	CMD string `json: "cmd"`
}