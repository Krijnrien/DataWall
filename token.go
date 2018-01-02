package main

import "golang.org/x/oauth2"

/**
 * Currently not used within the project. No unit tests either.
 */
func (TokenSourceParam *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: TokenSourceParam.AccessToken,
	}
	return token, nil
}
