package api

import "embed"

// Mode 开发者模式
type Mode uint8

// IsDev 是否是开发者模式。
func (m Mode) IsDev() bool {
	return m == DevMode
}

const (
	// DevMode 开发模式
	DevMode Mode = iota

	// ProdMode 生产模式
	ProdMode
)

type (
	StaticFSConfig struct {
		// 静态资源是否开启
		staticOn bool

		Mode Mode

		// dev mode: 使用路径加载静态文件
		StaticFilesPath string

		// prod mode: 使用 embed 加载静态文件
		embedFS embed.FS
	}

	StaticFSOption func(*StaticFSConfig)
)

func NewStaticFSConfig(opts ...StaticFSOption) *StaticFSConfig {
	defConfig := StaticFSConfig{}

	for _, opt := range opts {
		opt(&defConfig)
	}

	return &defConfig
}

// WithDevMode 使用开发者模式。
// 开发者模式下，静态文件使用路径加载。
func WithDevMode(staticFilesPath string) StaticFSOption {

	return func(c *StaticFSConfig) {
		c.staticOn = true
		c.Mode = DevMode
		c.StaticFilesPath = staticFilesPath
	}
}

func WithProdMode(embedFS embed.FS) StaticFSOption {
	return func(c *StaticFSConfig) {
		c.staticOn = true
		c.Mode = ProdMode
	}
}
