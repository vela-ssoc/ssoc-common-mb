package model

type ExecuteStat struct {
	Goos       string   `bson:"goos,omitempty"       json:"goos,omitzero"`
	Goarch     string   `bson:"goarch,omitempty"     json:"goarch,omitzero"`
	Version    string   `bson:"version"              json:"version"`    // 版本数字
	Unstable   bool     `bson:"unstable"             json:"unstable"`   // 内测版本，主要用于 agent 节点。
	Customized string   `bson:"customized"           json:"customized"` // 定制版
	PID        int      `bson:"pid,omitempty"        json:"pid,omitzero"`
	Args       []string `bson:"args,omitempty"       json:"args,omitzero"`
	Hostname   string   `bson:"hostname,omitempty"   json:"hostname,omitzero"`
	Workdir    string   `bson:"workdir,omitempty"    json:"workdir,omitzero"`
	Executable string   `bson:"executable,omitempty" json:"executable,omitzero"`
}
