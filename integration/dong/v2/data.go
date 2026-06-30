package dong

type requestBody struct {
	Prop requestProp `json:"prop"`
	Data requestData `json:"data"`
}

type requestProp struct {
	JobNumbers []string `json:"job_numbers"`
	RuleID     string   `json:"rule_id"`
	OriginName string   `json:"origin_name"`
}

type requestData struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type requestDirectBody struct {
	UserIDs  string `json:"userIds,omitempty"`  // 消息接收用户，多个 以,隔开（groupIds 为 空时必填）
	GroupIDs string `json:"groupIds,omitempty"` // 消息接收群，多个以,隔开（userIds 为空时必填）
	Title    string `json:"title"`              // 标题（必填）长度：300 个字符（150 个中文）
	Detail   string `json:"detail"`             // 卡片消息详细（必填）长度限制：2000 个字符（中文：1000 个）支持 html 标签
}

type requestTunnelBody struct {
	UserIDs  []string `json:"user_ids"`
	GroupIDs []string `json:"group_ids"`
	Detail   string   `json:"detail"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
}

type ResponseEntity struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	IssueID string `json:"issueId"`
	Data    any    `json:"data"`
}

func (r *ResponseEntity) Error() string {
	return r.Msg + ": " + r.IssueID + "（咚咚服务）"
}

func (r *ResponseEntity) checkError() error {
	if r.Code == "200" {
		return nil
	}

	return r
}

type BroadcastResponse struct {
	FailUsers []string `json:"failUsers"` // 失败用户
}
