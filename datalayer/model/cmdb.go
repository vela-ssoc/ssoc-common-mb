package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	CIDepartment  = "department"   // 部门
	CIServer      = "server"       // 资产
	CIVServer     = "vserver"      // 虚拟机
	CIFundVServer = "fund_vserver" // 基金虚拟机
	CIDocker      = "docker"       // Docker
	CIPublicCloud = "public_cloud" // 公有云
)

// CMDB 服务器资产。
//
// 物理机：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=9e7fb369-3695-406d-85ae-aa51579f94aa
// 虚拟机：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=20ab7bdd-beeb-4031-bb46-ec40986592e3
// 容器：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=3ce158ee-31f8-4755-b7ad-0fb5c1a24ffb
// 公有云：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=44bd2034-a9b1-4c74-9b25-0c147783f4e4
//
// 不同类型（ci_type）的唯一标识：
//   - server:       sn (SN)
//   - vserver:      uuid (UUID)
//   - docker:       instance_id (InstanceID)
//   - public_cloud: id (PublicCloudID)
type CMDB struct {
	ID     bson.ObjectID `json:"id"           bson:"_id,omitempty" jsonschema:"-"`
	CIID   int64         `json:"ci_id,string" bson:"ci_id"         jsonschema:"CMDB 资产 ID"`
	CIType string        `json:"ci_type"      bson:"ci_type"       jsonschema:"资产类型"`

	AccountID               string       `json:"account_id,omitzero"                 bson:"account_id"   jsonschema:"account_id"`
	AppDuty                 SimpleDuties `json:"app_duty,omitzero"                   bson:"app_duty"`
	AppName                 string       `json:"app_name,omitzero"                   bson:"app_name"     jsonschema:"应用名称"`
	Appid                   string       `json:"appid,omitzero"                      bson:"appid"`
	AssetID                 string       `json:"asset_id,omitzero"                   bson:"asset_id"     jsonschema:"财务系统资产编号"`
	AssetStatus             string       `json:"asset_status,omitzero"               bson:"asset_status" jsonschema:"资产状态 未使用/使用中/已报废"`
	AutoRenew               string       `json:"auto_renew,omitzero"                 bson:"auto_renew"`
	BaoleijiIdentity        string       `json:"baoleiji_identity,omitzero"          bson:"baoleiji_identity" jsonschema:"堡垒机可登录账号"`
	BeetleService           string       `json:"beetle_service,omitzero"             bson:"beetle_service"`
	BillingCycle            string       `json:"billing_cycle,omitzero"              bson:"billing_cycle"`
	BillingType             string       `json:"billing_type,omitzero"               bson:"billing_type"`
	Business                string       `json:"business,omitzero"                   bson:"business"          jsonschema:"业务分类"`
	BusinessEnv             string       `json:"business_env,omitzero"               bson:"business_env"`
	BusinessStatus          string       `json:"business_status,omitzero"            bson:"business_status"`
	CPU                     string       `json:"cpu,omitzero"                        bson:"cpu"              jsonschema:"CPU 型号"`
	CPUCount                int          `json:"cpu_count,omitzero"                  bson:"cpu_count"        jsonschema:"CPU 核数"`
	CmcIP                   string       `json:"cmc_ip,omitzero"                     bson:"cmc_ip"`
	CncIP                   string       `json:"cnc_ip,omitzero"                     bson:"cnc_ip"`
	Comment                 string       `json:"comment,omitzero"                    bson:"comment"         jsonschema:"作用"`
	CostDepCasID            string       `json:"cost_dep_cas_id,omitzero"            bson:"cost_dep_cas_id"`
	CreatedTime             string       `json:"created_time,omitzero"               bson:"created_time"`
	CtcIP                   string       `json:"ctc_ip,omitzero"                     bson:"ctc_ip"          jsonschema:"电信 IP"`
	Deleted                 string       `json:"deleted,omitzero"                    bson:"deleted"`
	Department              string       `json:"department,omitzero"                 bson:"department"`
	Description             string       `json:"description,omitzero"                bson:"description"`
	DockerCPUCount          string       `json:"docker_cpu_count,omitzero"           bson:"docker_cpu_count"`
	Env                     string       `json:"env,omitzero"                        bson:"env"`
	ExternalID              string       `json:"external_id,omitzero"                bson:"external_id"`
	FloatIP                 []string     `json:"float_ip,omitzero"                   bson:"float_ip"`
	Harddisk                string       `json:"harddisk,omitzero"                   bson:"harddisk"`
	HostBuyDate             string       `json:"host_buy_date,omitzero"              bson:"host_buy_date"`
	HostIP                  []string     `json:"host_ip,omitzero"                    bson:"host_ip"         jsonschema:"宿主机 IP"`
	HostType                string       `json:"host_type,omitzero"                  bson:"host_type"`
	Hostname                string       `json:"hostname,omitzero"                   bson:"hostname"        jsonschema:"主机名"`
	HyperThreading          string       `json:"hyper_threading,omitzero"            bson:"hyper_threading" jsonschema:"超线程 enable/disable"`
	IDC                     string       `json:"idc,omitzero"                        bson:"idc"             jsonschema:"IDC"`
	IPv6                    []string     `json:"ipv6,omitzero"                       bson:"ipv6"            jsonschema:"内网 IPv6"`
	Image                   string       `json:"image,omitzero"                      bson:"image"`
	ImageName               string       `json:"image_name,omitzero"                 bson:"image_name"`
	ImageVersion            string       `json:"image_version,omitzero"              bson:"image_version"`
	ImportedAt              string       `json:"imported_at,omitzero"                bson:"imported_at"`
	InScalingGroup          string       `json:"in_scaling_group,omitzero"           bson:"in_scaling_group"`
	InstanceID              string       `json:"instance_id,omitzero"                bson:"instance_id"`
	InstanceType            string       `json:"instance_type,omitzero"              bson:"instance_type"`
	InternetMaxBandwidthOut int          `json:"internet_max_bandwidth_out,omitzero" bson:"internet_max_bandwidth_out" jsonschema:""`
	K8SCluster              string       `json:"k8s_cluster,omitzero"                bson:"k8s_cluster"`
	KernelVersion           string       `json:"kernel_version,omitzero"             bson:"kernel_version"  jsonschema:"内核版本"`
	LogicCPUCount           int          `json:"logic_cpu_count,omitzero"            bson:"logic_cpu_count"`
	MinionNotCheck          string       `json:"minion_not_check,omitzero"           bson:"minion_not_check"`
	Name                    string       `json:"name"                                bson:"name"`
	Namespace               string       `json:"namespace,omitzero"                  bson:"namespace"`
	NetOpen                 string       `json:"net_open,omitzero"                   bson:"net_open"`
	OSArch                  string       `json:"os_arch,omitzero"                    bson:"os_arch"`
	OSName                  string       `json:"os_name,omitzero"                    bson:"os_name"`
	OSType                  string       `json:"os_type,omitzero"                    bson:"os_type"`
	OSVersion               string       `json:"os_version,omitzero"                 bson:"os_version"`
	OpDuty                  SimpleDuties `json:"op_duty,omitzero"                    bson:"op_duty"`
	Payer                   string       `json:"payer,omitzero"                      bson:"payer"`
	PowerStates             string       `json:"power_states,omitzero"               bson:"power_states"`
	PrivateCloudIP          string       `json:"private_cloud_ip,omitzero"           bson:"private_cloud_ip"`
	PrivateCloudType        string       `json:"private_cloud_type,omitzero"         bson:"private_cloud_type"`
	PrivateIP               []string     `json:"private_ip,omitzero"                 bson:"private_ip"         jsonschema:"内网 IP"`
	PublicCloudID           string       `json:"public_cloud_id,omitzero"            bson:"public_cloud_id"`         // public_cloud 云管 ID
	PublicCloudIDC          string       `json:"public_cloud_idc,omitzero"           bson:"public_cloud_idc"`        // public_cloud 公有云厂商
	PublicCloudCreatedAt    string       `json:"public_cloud_created_at,omitzero"    bson:"public_cloud_created_at"` // public_cloud 云管创建时间 YYYY-mm-DD HH:MM:SS
	PublicCloudExpiredAt    string       `json:"public_cloud_expired_at,omitzero"    bson:"public_cloud_expired_at"` // 过期时间
	PublicCloudUpdatedAt    string       `json:"public_cloud_updated_at,omitzero"    bson:"public_cloud_updated_at"` // public_cloud 云管更新时间 YYYY-mm-DD HH:MM:SS
	RAID                    string       `json:"raid,omitzero"                       bson:"raid"`
	RAM                     string       `json:"ram,omitzero"                        bson:"ram"`
	RAMSize                 string       `json:"ram_size,omitzero"                   bson:"ram_size"`
	Rack                    string       `json:"rack,omitzero"                       bson:"rack"`
	RdDuty                  SimpleDuties `json:"rd_duty,omitzero"                    bson:"rd_duty"`
	Region                  string       `json:"region,omitzero"                     bson:"region"`
	RegionExtID             string       `json:"region_ext_id,omitzero"              bson:"region_ext_id"`
	ResourceRequests        string       `json:"resource_requests,omitzero"          bson:"resource_requests"`
	SN                      string       `json:"sn,omitzero"                         bson:"sn"                 jsonschema:"序列化"`
	SecurityInfo            string       `json:"security_info,omitzero"              bson:"security_info"      jsonschema:"安全信息"`
	Server                  string       `json:"server,omitzero"                     bson:"server"             jsonschema:"宿主机序列号"`
	ServerRoom              string       `json:"server_room,omitzero"                bson:"server_room"        jsonschema:"机房"`
	ShutdownBehavior        string       `json:"shutdown_behavior,omitzero"          bson:"shutdown_behavior"`
	ShutdownMode            string       `json:"shutdown_mode,omitzero"              bson:"shutdown_mode"`
	Status                  string       `json:"status,omitzero"                     bson:"status"              jsonschema:"业务状态 待用/在线"`
	SysDuty                 SimpleDuties `json:"sys_duty,omitzero"                   bson:"sys_duty"`
	Tags                    []string     `json:"tags,omitzero"                       bson:"tags"`
	Throughput              int          `json:"throughput,omitzero"                 bson:"throughput"`
	TradeType               string       `json:"trade_type,omitzero"                 bson:"trade_type"`
	UUID                    string       `json:"uuid,omitzero"                       bson:"uuid"`
	UpdateTime              string       `json:"update_time,omitzero"                bson:"update_time"`
	Use                     string       `json:"use,omitzero,omitzero"               bson:"use"`
	VCPUCount               int          `json:"vcpu_count,omitzero"                 bson:"vcpu_count"`
	VMemSize                int          `json:"vmem_size,omitzero"                  bson:"vmem_size"`
	VServerType             string       `json:"vserver_type,omitzero"               bson:"vserver_type"`
	ZabbixNotCheck          string       `json:"zabbix_not_check,omitzero"           bson:"zabbix_not_check"`
	Zone                    string       `json:"zone,omitzero"                       bson:"zone"`

	// ci_type = department 特有的字段
	DepPath             string       `json:"dep_path,omitzero"             bson:"dep_path"`
	Owner               SimpleDuties `json:"owner,omitzero"                bson:"owner"`
	BusinessType        string       `json:"business_type,omitzero"        bson:"business_type"`
	ManagementType      string       `json:"management_type,omitzero"      bson:"management_type"`
	ProjectID           string       `json:"project_id,omitzero"           bson:"project_id"`
	UserDepartment      []string     `json:"user_department,omitzero"      bson:"user_department"`
	User                SimpleDuties `json:"user,omitzero"                 bson:"user"`
	WorkflowApprover    SimpleDuties `json:"workflow_approver,omitzero"    bson:"workflow_approver"`
	TemporaryDepartment string       `json:"temporary_department,omitzero" bson:"temporary_department"`

	// private_dns_record
	DomainName  string       `json:"domain_name,omitzero"  bson:"domain_name"`
	Duty        SimpleDuties `json:"duty,omitzero"         bson:"duty"`
	RecordType  string       `json:"record_type,omitzero"  bson:"record_type"`
	RecordValue string       `json:"record_value,omitzero" bson:"record_value"`
	TTL         int          `json:"ttl,omitzero"          bson:"ttl"`

	// public_dns_record
	ISP              string `json:"isp,omitzero"                bson:"isp"`
	PublicDomainName string `json:"public_domain_name,omitzero" bson:"public_domain_name"`
	DeviceName       string `json:"device_name,omitzero"        bson:"device_name"`

	// vip
	AccessControl string              `json:"access_control,omitzero" bson:"access_control"`
	IdleTimeout   string              `json:"idle_timeout,omitzero"   bson:"idle_timeout"`
	IPType        string              `json:"ip_type,omitzero"        bson:"ip_type"`
	LoadBalance   string              `json:"load_balance,omitzero"   bson:"load_balance"`
	NAT           int                 `json:"nat,omitzero"            bson:"nat"`
	Protocol      string              `json:"protocol,omitzero"       bson:"protocol"`
	RServer       []CMDBVIPRealServer `json:"rserver,omitzero"        bson:"rserver"`
	Source        string              `json:"source,omitzero"         bson:"source"`
	VIP           string              `json:"vip,omitzero"            bson:"vip"`
	VIPUniqueID   string              `json:"vip_unique_id,omitzero"  bson:"vip_unique_id"`
	VPort         int                 `json:"vport,omitzero"          bson:"vport"`
	VSName        string              `json:"vs_name,omitzero"        bson:"vs_name"`
	GetClientIP   string              `json:"get_client_ip,omitzero"  bson:"get_client_ip"`
	Certificate   string              `json:"certificate,omitzero"    bson:"certificate"`

	SSOCAgent *SSOCAgentSimple `json:"ssoc_agent,omitempty" bson:"ssoc_agent,omitempty"`

	// Retirement bool      `json:"retirement" bson:"retirement"           jsonschema:"是否已清退"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty" jsonschema:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty" jsonschema:"更新时间"`

	// 下面的字段定义和接口文档描述不一致，或者是同样的字段会返回多种类型。
	// 这些字段本身对于我们的系统也没有什么价值，就忽略掉了。
	// AppLoadType  []string `json:"app_load_type,omitzero" bson:"app_load_type"`
	// SecurityRisk int      `json:"security_risk,omitzero" bson:"security_risk"`
}

func (CMDB) CollectionInfo() (string, string) {
	return "cmdb", "CMDB"
}

type CMDBVIPRealServer struct {
	Port       int    `json:"port,omitzero"        bson:"port"`
	Server     string `json:"server,omitzero"      bson:"server"`
	Status     string `json:"status,omitzero"      bson:"status"`
	UpdateTime string `json:"update_time,omitzero" bson:"update_time"`
}

type SimpleDuty struct {
	JobNumber string `json:"job_number" bson:"job_number" validate:"required"                                   jsonschema:"责任人工号"` // employee_id
	Name      string `json:"name"       bson:"name"       validate:"lte=10"                                     jsonschema:"责任人名字"` // nickname
	Company   string `json:"company"    bson:"company"    validate:"oneof=group security futures ttfund dcfund" jsonschema:"所属公司"`  // 子公司
}

type SimpleDuties []SimpleDuty

type SSOCAgentSimple struct {
	ID         int64  `bson:"id"          json:"id"`
	Status     int    `bson:"status"      json:"status"` // 0-离线 1-在线 2-已删除
	MachineID  string `bson:"machine_id"  json:"machine_id"`
	Version    string `bson:"version"     json:"version"`
	BrokerSlug string `bson:"broker_slug" json:"broker_slug"`
	BrokerName string `bson:"broker_name" json:"broker_name"`
}
