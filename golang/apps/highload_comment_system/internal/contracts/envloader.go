package contracts

type EnvLoader interface {
	MustLoad(cfg interface{})
}
