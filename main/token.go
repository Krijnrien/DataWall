package main

import "golang.org/x/oauth2"

func (TokenSourceParam *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: TokenSourceParam.AccessToken,
	}
	return token, nil
}
