package config

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

/*注意：配置结构体中 mapstructure 标签需对应 config.ymal 中的配置名称， viper 会根据标签 value 值把 config.yaml 的数据赋予给结构体*/
