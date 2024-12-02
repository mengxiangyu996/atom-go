package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// 项目配置
	App struct {
		Name       string `yaml:"name"`       // 项目名称
		Domain     string `yaml:"domain"`     // 项目域名
		SSL        bool   `yaml:"ssl"`        // 是否启用SSL
		PublicPath string `yaml:"publicPath"` // 资源路径
		UploadPath string `yaml:"uploadPath"` // 文件上传路径

		// 服务器配置
		Server struct {
			Port int    `yaml:"port"` // 服务端口
			Mode string `yaml:"mode"` // 运行模式
		} `yaml:"server"`
	} `yaml:"app"`

	// Mysql配置
	Mysql struct {
		Host         string `yaml:"host"`         // 数据库地址
		Port         int    `yaml:"port"`         // 数据库端口
		Database     string `yaml:"database"`     // 数据库名称
		Username     string `yaml:"username"`     // 数据库用户名
		Password     string `yaml:"password"`     // 数据库密码
		Charset      string `yaml:"charset"`      // 数据库编码
		MaxIdleConns int    `yaml:"maxIdleConns"` // 连接池最大连接数
		MaxOpenConns int    `yaml:"maxOpenConns"` // 连接池最大打开连接数
	} `yaml:"mysql"`

	// Redis配置
	Redis struct {
		Host     string `yaml:"host"`     // Redis地址
		Port     int    `yaml:"port"`     // Redis端口
		Database int    `yaml:"database"` // Redis数据库索引
		Password string `yaml:"password"` // Redis密码
	} `yaml:"redis"`

	// 授权配置
	Token struct {
		Header     string `yaml:"header"`     // 令牌自定义标识
		Secret     string `yaml:"secret"`     // 令牌密钥
		ExpireTime int    `yaml:"expireTime"` // 令牌有效期
	} `yaml:"token"`
}

var Data *Config

func InitConfig() {

	file, err := os.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &Data)
	if err != nil {
		panic(err)
	}
}
