package accord

const (
	PathPrefix    = "/api/v1"
	PathTaskLoad  = "/task/load"  // 加载指定配置
	PathTaskSync  = "/task/sync"  // 同步配置
	PathTaskTable = "/task/table" // 同步配置
	PathThirdDiff = "/third/diff" // 三方文件变更
	FPTaskLoad    = PathPrefix + PathTaskLoad
	FPTaskSync    = PathPrefix + PathTaskSync
	FPTaskTable   = PathPrefix + PathTaskTable
	FPThirdDiff   = PathPrefix + PathThirdDiff
)

type TaskLoadRequest struct {
	MinionID    int64  `json:"minion_id"`
	SubstanceID int64  `json:"substance_id"`
	Inet        string `json:"inet"`
}

type TaskSyncRequest struct {
	MinionID int64  `json:"minion_id"`
	Inet     string `json:"inet"`
}

type ThirdDiff struct {
	Name  string `json:"name"`  // 三方文件名
	Event string `json:"event"` // 事件类型：delete-文件删除 update-文件更新
}

const (
	ThirdUpdate = "update"
	ThirdDelete = "delete"
)
