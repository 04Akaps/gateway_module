package config

type Producer struct {
	URL       string  `yaml:"url"`
	ClientID  string  `yaml:"client_id"`
	Acks      string  `yaml:"acks"`
	Topic     string  `yaml:"topic"`
	BatchTime float64 `yaml:"batch_time"`
}
