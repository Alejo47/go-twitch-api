package authorization

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RefreshAccessToken(client_id, client_secret, redirect_uri, refresh_token string) (RefreshAccessTokenResponse, error) {
	var tokenOut RefreshAccessTokenResponse
	var err error

	client := http.Client{}
	uri := fmt.Sprintf("https://id.twitch.tv/oauth2/token?grant_type=refresh_token&refresh_token=%s&client_id=%s&client_secret=%s",
		refresh_token,
		client_id,
		client_secret,
	)

	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		return tokenOut, err
	}

	res, err := client.Do(req)
	if err != nil {
		return tokenOut, err
	}

	out, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		err = errors.New(fmt.Sprintf("Invalid status code returned: %d", res.StatusCode))
		return tokenOut, err
	}
	if err != nil {
		return tokenOut, err
	}

	err = json.Unmarshal(out, &tokenOut)
	if err != nil {
		return tokenOut, err
	}

	return tokenOut, nil
}
