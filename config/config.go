package config

import "github.com/jinzhu/configor"

type (
	Twitter struct {
		Key         string
		Secret      string
		Token       string
		TokenSecret string
	}

	Interactions struct {
		DryRun             bool
		RespondsToComments bool
		RespondsToTweets   bool
	}

	Logging struct {
		Pretty   bool `default:"true"`
		LogLevel string
	}

	Config struct {
		Twitter      *Twitter
		Interactions *Interactions
		Logging      *Logging
	}
)

func Load(config string) (c *Config, err error) {
	c = &Config{}

	con := configor.New(&configor.Config{
		ENVPrefix: "DRBEES",
	})

	if config == "" {
		err = con.Load(c)
		return
	}

	err = con.Load(c, config)

	return
}
