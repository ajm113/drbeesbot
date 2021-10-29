package main

import (
	"os"
	"strings"

	"github.com/ajm113/drbeesbot/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dgraph-io/ristretto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type (
	handler struct {
		TwitterClient *twitter.Client
		Cache         *ristretto.Cache
		Config        *config.Config
	}
)

func setupLogging(config *config.Config) {
	switch strings.ToLower(config.Logging.LogLevel) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "warm":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "silent":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if config.Logging.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func newHandler(config *config.Config) (h *handler, err error) {
	h = &handler{
		Config: config,
	}

	twitterLog := log.With().
		Str("key", config.Twitter.Key).
		Str("token", config.Twitter.Token).Logger()

	twitterLog.Info().Msg("logging into Twitter")

	authConfig := oauth1.NewConfig(config.Twitter.Key, config.Twitter.Secret)
	authToken := oauth1.NewToken(config.Twitter.Token, config.Twitter.TokenSecret)

	httpClient := authConfig.Client(oauth1.NoContext, authToken)
	h.TwitterClient = twitter.NewClient(httpClient)

	user, _, err := h.TwitterClient.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	})

	if err != nil {
		twitterLog.Error().Err(err).Msg("failed verifying credentials")
	}

	log.Info().Str("user", user.Name).Str("id", user.IDStr).Msg("succesfully logged into Twitter")

	return
}

func run(context *cli.Context) (err error) {
	c, err := config.Load(context.Value("config").(string))

	if err != nil {
		return err
	}

	setupLogging(c)
	_, err = newHandler(c)

	return
}
