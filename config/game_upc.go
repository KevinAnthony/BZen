package config

type GameCPUConfig struct {
	APIKey     string `env:"GAME_UPC_API_CONFIG" envDefault:"test_test_test_test_test"`
	URL        string `env:"GAME_UPC_URL" envDefault:"api.gameupc.com"`
	PathPrefix string `env:"GAME_UPC_PREFIX" envDefault:"test"`
}
