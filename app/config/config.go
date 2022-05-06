package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type SiteConfig struct {
	Title     string   `yaml:"title"`
	Introduce string   `yaml:"introduce"`
	Limit     int      `yaml:"limit"`
	Theme     string   `yaml:"theme"`
	URL       string   `yaml:"url"`
	Comments  Comments `yaml:"comments"`
	Github    string   `yaml:"github"`
	Facebook  string   `yaml:"facebook"`
	Twitter   string   `yaml:"twitter"`
}

type Comments struct {
	Disqus DisqusComment `yaml:"disqus"`
	Giscus GiscusComment `yaml:"giscus"`
}

type DisqusComment struct {
	Enable bool   `yaml:"enable"`
	Site   string `yaml:"site"`
}

type GiscusComment struct {
	Enable           bool   `yaml:"enable"`
	Repo             string `yaml:"repo"`
	RepoID           string `yaml:"repo_id"`
	Category         string `yaml:"category"`
	CategoryID       string `yaml:"category_id"`
	Mapping          string `yaml:"mapping"`
	ReactionsEnabled bool   `yaml:"reactions_enabled"`
	Theme            string `yaml:"theme"`
}

type ServeConfig struct {
	Addr string `yaml:"addr"`
}

type PublishConfig struct {
	Cmd string `yaml:"cmd"`
}

type Config struct {
	Site       SiteConfig    `yaml:"site"`
	Serve      ServeConfig   `yaml:"serve"`
	Publish    PublishConfig `yaml:"publish"`
	RootPath   string
	ThemePath  string
	SourcePath string
	PublicPath string
}

func (c *Config) parseConfig() error {
	data, err := ioutil.ReadFile(filepath.Join(c.RootPath, "config.yml"))
	if err != nil {
		return fmt.Errorf("Read config: %v", err)
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		return fmt.Errorf("Unmarshal config: %v", err)
	}

	return nil
}

func New(c *cli.Context) (*Config, error) {
	var config = &Config{
		RootPath: ".",
	}
	if len(c.Args()) > 0 {
		config.RootPath = c.Args()[0]
	}
	err := config.parseConfig()
	if err != nil {
		return nil, fmt.Errorf("parse config: %v", err)
	}
	config.ThemePath = filepath.Join(config.RootPath, config.Site.Theme)
	config.SourcePath = filepath.Join(config.RootPath, "source")
	config.PublicPath = filepath.Join(config.RootPath, "public")
	return config, nil
}
