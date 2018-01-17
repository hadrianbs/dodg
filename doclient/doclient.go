package doclient

import (
	"github.com/digitalocean/godo"
	"github.com/digitalocean/godo/context"
	"github.com/hadrianbs/dodg/config"
	"golang.org/x/oauth2"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func InitClient(patConfig config.PATConfig) (*godo.Client, error) {
	tokenSource := &TokenSource{
		AccessToken: patConfig.TokenSecret,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	doClient := godo.NewClient(oauthClient)
	return doClient, nil
}
