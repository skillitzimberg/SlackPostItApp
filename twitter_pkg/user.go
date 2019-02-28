package main

import "github.com/dghubble/go-twitter/twitter"

type TwitterUser struct {}

func (TwitterUser) Name() string {
	panic("implement me")
}

func (TwitterUser) Version() string {
	panic("implement me")
}

func (TwitterUser) Execute() {
	panic("implement me")
}


type TwitterUserInput struct {
	Username string
	AccessToken string
	AccessTokenSecret string
}

type TwitterUserOutput struct {
	UserInformation twitter.User
}
