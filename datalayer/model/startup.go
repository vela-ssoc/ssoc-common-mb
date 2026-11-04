package model

type StartupV1 struct {
	Logger StartupV1Logger `json:"logger"`
}

type StartupV1Logger struct {
	Level    string `json:"level"    validate:"oneof=debug info warn error"`
	Filename string `json:"filename"`
	Console  bool   `json:"console"`
	Format   string `json:"format"   validate:"oneof=text json"`
	Caller   bool   `json:"caller"`
	Skip     int    `json:"skip"     validate:"gt=-100,lt=100"`
}
