package pubsub

// Subscriber 事件接收者，通常对应一条 SSE 连接。
//
// 约定：
//   - 必须用指针实现：分发器以 Subscriber 作为 map 的 key，
//     动态类型不可比较（含 map/slice 的结构体）会在 AddSubscriber 时 panic。
//   - 分发器保证对同一 Subscriber 串行调用 OnEvent，实现里可以安全写 ResponseWriter。
//   - OnEvent 必须非阻塞（入队后立即返回），否则会拖慢 Publish 的调用方。
//   - OnEvent 不得调用 DelSubscriber，否则会与 Publish 的读锁互等。
//   - 返回 error 表示该订阅者已无法继续接收，分发器会将其注销，之后不再调用。
type Subscriber interface {
	OnEvent(event string, data any) error
}

// Publisher 发布者（发布规则集）：只负责发，不感知订阅者。
//
// 契约：
//   - 绝不阻塞：订阅者慢就丢，不能拖慢调用方（比如 broker 接入流程）。
//   - 不返回错误：没有存储、没有重试、没有补偿，调用方也无从处理。
type Publisher interface {
	Publish(event string, data any)
}

// Dispatcher 事件分发器：发布与订阅的唯一交汇点。
//
// 只做转发，不做权限、脱敏、级别过滤等任何策略——这些属于订阅侧的
// 守卫与投影，否则发布侧与订阅侧的规则会互相污染。
type Dispatcher interface {
	Publisher

	// AddSubscriber 注册订阅者。重复注册视为无操作。
	AddSubscriber(sub Subscriber)

	// DelSubscriber 注销订阅者，幂等。
	// 只从分发器摘除，不关闭订阅者：关闭由订阅者的持有方负责。
	DelSubscriber(sub Subscriber)

	// Count 当前订阅者数量，用于观测与连接数熔断。
	Count() int
}
