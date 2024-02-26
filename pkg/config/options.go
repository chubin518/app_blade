package config

type Option func(*appConfig)

func WithPath(path string) Option {
	return func(ac *appConfig) {
		ac.path = path
	}
}

func WithName(name string) Option {
	return func(ac *appConfig) {
		ac.name = name
	}
}

func WithType(tp string) Option {
	return func(ac *appConfig) {
		ac.suffix = tp
	}
}
