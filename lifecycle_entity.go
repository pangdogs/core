package golaxy

// LifecycleEntityInit 实体的生命周期进入初始化（init）时的回调，实体实现此接口即可使用
type LifecycleEntityInit interface {
	Init()
}

// LifecycleEntityInited 实体的生命周期进入初始化完成（inited）时的回调，实体实现此接口即可使用
type LifecycleEntityInited interface {
	Inited()
}

// LifecycleEntityUpdate 如果开启运行时的帧更新特性，那么实体状态为活跃（living）时，将会收到这个帧更新（update）回调，实体实现此接口即可使用
type LifecycleEntityUpdate = eventUpdate

// LifecycleEntityLateUpdate 如果开启运行时的帧更新特性，那么实体状态为活跃（living）时，将会收到这个帧迟滞更新（late update）回调，实体实现此接口即可使用
type LifecycleEntityLateUpdate = eventLateUpdate

// LifecycleEntityShut 实体的生命周期进入结束（shut）时的回调，实体实现此接口即可使用
type LifecycleEntityShut interface {
	Shut()
}

// LifecycleEntityDestroy 实体的生命周期进入死亡（death）时的回调，实体实现此接口即可使用
type LifecycleEntityDestroy interface {
	Destroy()
}
