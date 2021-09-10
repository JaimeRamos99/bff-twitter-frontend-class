package tweets

import (
	"encoding/json"
	"io/ioutil"
)

type Tweet struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Time_ago        string `json:"time_ago"`
	Content         string `json:"content"`
	Include_hashtag bool   `json:"include_hashtag"`
	Include_image   bool   `json:"include_image"`
	Is_verified     bool   `json:"is_verified"`
	Content_image   string `json:"content_image"`
	Profile_picture string `json:"profile_picture"`
}

type Tweets struct {
	Tweets []Tweet `json:"data"`
}

func GetTweets() *Tweets {
	file, _ := ioutil.ReadFile("mock-data/tweets.json")

	var data *Tweets

	_ = json.Unmarshal([]byte(file), &data)
	return data
}
