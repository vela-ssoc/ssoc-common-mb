package accord

const (
	PathPrefix       = "/api/v1"
	PathTaskLoad     = "/task/load"     // 加载指定配置
	PathTaskSync     = "/task/sync"     // 同步配置
	PathTaskTable    = "/task/table"    // 同步配置
	PathThirdDiff    = "/third/diff"    // 三方文件变更
	PathElasticReset = "/elastic/reset" // 三方文件变更
	PathEmcReset     = "/emc/reset"     // 三方文件变更
	PathStoreReset   = "/store/reset"   // 三方文件变更
	PathCmdbReset    = "/cmdb/reset"    // 三方文件变更
	FPTaskLoad       = PathPrefix + PathTaskLoad
	FPTaskSync       = PathPrefix + PathTaskSync
	FPTaskTable      = PathPrefix + PathTaskTable
	FPThirdDiff      = PathPrefix + PathThirdDiff
	FPElasticReset   = PathPrefix + PathElasticReset
	FPEmcReset       = PathPrefix + PathEmcReset
	FPStoreReset     = PathPrefix + PathStoreReset
	FPCmdbReset      = PathPrefix + PathCmdbReset
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

type StoreRestRequest struct {
	ID string `json:"id"`
}

type ThirdDiff struct {
	Name  string `json:"name"`  // 三方文件名
	Event string `json:"event"` // 事件类型：delete-文件删除 update-文件更新
}

const (
	ThirdUpdate = "update"
	ThirdDelete = "delete"
)

type TaskTable struct {
	TaskID int64 `json:"task_id"`
}
