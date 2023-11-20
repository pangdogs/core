package golaxy

// LifecycleComponentAwake 组件的生命周期进入唤醒（awake）时的回调，组件实现此接口即可使用
type LifecycleComponentAwake interface {
	Awake()
}

// LifecycleComponentStart 组件的生命周期进入开始（start）时的回调，组件实现此接口即可使用
type LifecycleComponentStart interface {
	Start()
}

// LifecycleComponentUpdate 如果开启运行时的帧更新特性，那么组件状态为活跃（living）时，将会收到这个帧更新（update）回调，组件实现此接口即可使用
type LifecycleComponentUpdate = eventUpdate

// LifecycleComponentLateUpdate 如果开启运行时的帧更新特性，那么组件状态为活跃（living）时，将会收到这个帧迟滞更新（late update）回调，组件实现此接口即可使用
type LifecycleComponentLateUpdate = eventLateUpdate

// LifecycleComponentShut 组件的生命周期进入结束（shut）时的回调，组件实现此接口即可使用
type LifecycleComponentShut interface {
	Shut()
}

// LifecycleComponentDispose 组件的生命周期进入死亡（death）时的回调，组件实现此接口即可使用
type LifecycleComponentDispose interface {
	Dispose()
}
