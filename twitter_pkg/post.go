package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

type Tweet struct {}

func (Tweet) Name() string {
	return "tweet"
}

func (Tweet) Version() string {
	return "1.0"
}

func (Tweet) Execute() {
	input := TweetInput{}
	step.BindInputs(&input)
	output := TweetOutput{}

	credentials := Credentials{
		AccessToken: input.AccessToken,
		AccessTokenSecret: input.AccessTokenSecret,
		ConsumerKey: "I5UseC18M5siiqXWmlPAsLhj0",
		ConsumerSecret: "Q3Lb57NS9Zhf1KjbzY0vzEq6CgwNhtQhoNd8aCriiDeWA0ZqC3",
	}

	client, err := GetUserClient(&credentials)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	tweet, resp, err := client.Statuses.Update(
		input.Text, nil)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)

	output.Success = true
	step.SetOutput(output)
}

type Credentials struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func GetUserClient(credentials *Credentials) (*twitter.Client, error){
	config := oauth1.NewConfig(credentials.ConsumerKey,
		credentials.ConsumerSecret)
	token := oauth1.NewToken(credentials.AccessToken, credentials.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's Account:\n%+v\n", user)
	return client, nil
}

type TweetInput struct {
	Text string
	AccessToken string
	AccessTokenSecret string
}

type TweetOutput struct {
	Success bool
}
