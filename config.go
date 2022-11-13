package swagfs

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/a-kataev/swagfs/files"
	"github.com/a-kataev/swagfs/tmpl"
)

type Layout string

const (
	BaseLayout       = "BaseLayout"
	StandaloneLayout = "StandaloneLayout"
)

type URL struct {
	URL  string `json:"url"`
	Name string `json:"name,omitempty"`
}

type Config struct {
	Urls   []URL  `json:"urls"`
	Layout Layout `json:"layout"`
}

var _ tmpl.File = (*Config)(nil)

func defaultConfig() *Config {
	return &Config{
		Urls: []URL{
			{
				URL:  "/api/v1/swagger.yaml",
				Name: "v1",
			},
		},
		Layout: BaseLayout,
	}
}

func NewConfig() *Config {
	return defaultConfig()
}

func (c *Config) Filename() string {
	return files.ConfigName
}

func (c *Config) Template() string {
	return string(files.ConfigData())
}

func (c *Config) TagFunc(w io.Writer, tag string) (int, error) {
	switch tag {
	case "layout":
		return w.Write([]byte(c.Layout))
	case "urls":
		if c.Layout == BaseLayout {
			return w.Write([]byte("null"))
		}

		b, err := json.Marshal(c.Urls)
		if err != nil {
			return 0, fmt.Errorf("urls: %v", err)
		}

		return w.Write(b)
	case "url":
		if c.Layout == BaseLayout {
			if len(c.Urls) > 0 {
				return w.Write([]byte(c.Urls[0].URL))
			}
		}

		return 0, nil
	}

	return 0, nil
}

func (c *Config) AddURL(url, name string) *Config {
	c.Urls = append(c.Urls, URL{
		URL:  url,
		Name: name,
	})

	return c
}

func (c *Config) SetLayout(layout Layout) *Config {
	c.Layout = layout

	return c
}
