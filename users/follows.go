package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUserFollows(bearer_token, client_id, follow_request FollowRequest) (FollowsResponse, error) {
	var followResponse FollowsResponse
	client := http.Client{}
	uri := fmt.Sprintf("https://api.twitch.tv/helix/users/follows/?to_id=%s&from_id=%s&after=%s&first=%s",
		follow_request.ToID,
		follow_request.FromID,
		follow_request.After,
		follow_request.First,
	)

	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Bearer "+bearer_token)
	req.Header.Add("Client-Id", client_id)

	if err != nil {
		return followResponse, err
	}

	res, err := client.Do(req)

	if err != nil {
		return followResponse, err
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return followResponse, err
	}
	err = json.Unmarshal(out, &followResponse)
	if err != nil {
		return followResponse, err
	}
	return followResponse, err
}
