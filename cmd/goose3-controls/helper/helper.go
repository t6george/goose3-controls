package helper

// JSON for the command
type Command struct {
    Time int64      `json:"time"`
    Type string     `json:"type"`
    Id uint8        `json:"id"`
    Data []float32  `json:"data"`
}
