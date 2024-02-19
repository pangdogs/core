package core

var With _Option

type _Option struct {
	EntityCreator _EntityCreatorOption // 实体构建器的选项
	Runtime       _RuntimeOption       // 运行时的选项
	Service       _ServiceOption       // 服务的选项
}
