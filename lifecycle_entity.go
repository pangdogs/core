package core

// LifecycleEntityAwake 实体的生命周期进入唤醒（awake）时的回调，实体实现此接口即可使用
type LifecycleEntityAwake interface {
	Awake()
}

// LifecycleEntityStart 实体的生命周期进入开始（start）时的回调，实体实现此接口即可使用
type LifecycleEntityStart interface {
	Start()
}

// LifecycleEntityUpdate 如果开启运行时的帧更新特性，那么实体状态为活跃（living）时，将会收到这个帧更新（update）回调，实体实现此接口即可使用
type LifecycleEntityUpdate = eventUpdate

// LifecycleEntityLateUpdate 如果开启运行时的帧更新特性，那么实体状态为活跃（living）时，将会收到这个帧迟滞更新（late update）回调，实体实现此接口即可使用
type LifecycleEntityLateUpdate = eventLateUpdate

// LifecycleEntityShut 实体的生命周期进入结束（shut）时的回调，实体实现此接口即可使用
type LifecycleEntityShut interface {
	Shut()
}

// LifecycleEntityDispose 实体的生命周期进入死亡（death）时的回调，实体实现此接口即可使用
type LifecycleEntityDispose interface {
	Dispose()
}
