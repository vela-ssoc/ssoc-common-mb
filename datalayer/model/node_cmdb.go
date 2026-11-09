package model

type NodeCMDB struct {
	CIType      string       `json:"ci_type"      bson:"ci_type"       jsonschema:"资产类型"`
	AppDuty     SimpleDuties `json:"app_duty,omitzero"                   bson:"app_duty"`
	AppName     string       `json:"app_name,omitzero"                   bson:"app_name"     jsonschema:"应用名称"`
	CPU         string       `json:"cpu,omitzero"                        bson:"cpu"              jsonschema:"CPU 型号"`
	CPUCount    int          `json:"cpu_count,omitzero"                  bson:"cpu_count"        jsonschema:"CPU 核数"`
	Department  string       `json:"department,omitzero"                 bson:"department"`
	Description string       `json:"description,omitzero"                bson:"description"`
	IDC         string       `json:"idc,omitzero"                        bson:"idc"             jsonschema:"IDC"`
	IPv6        []string     `json:"ipv6,omitzero"                       bson:"ipv6"            jsonschema:"内网 IPv6"`
	OpDuty      SimpleDuties `json:"op_duty,omitzero"                    bson:"op_duty"`
	PrivateIP   []string     `json:"private_ip,omitzero"                 bson:"private_ip"         jsonschema:"内网 IP"`
}
