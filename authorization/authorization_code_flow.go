package authorization

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAccessToken(client_id, client_secret, redirect_uri, code string) (AccessTokenResponse, error) {
	var tokenOut AccessTokenResponse
	var err error

	client := http.Client{}
	uri := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code&redirect_uri=%s",
		client_id,
		client_secret,
		code,
		redirect_uri,
	)

	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		return tokenOut, err
	}

	res, err := client.Do(req)
	if err != nil {
		return tokenOut, err
	} else if res.StatusCode < 200 || res.StatusCode >= 400 {
		err = errors.New(fmt.Sprintf("Invalid status code returned: %d", res.StatusCode))
		return tokenOut, err
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return tokenOut, err
	}

	err = json.Unmarshal(out, &tokenOut)
	if err != nil {
		return tokenOut, err
	}

	return tokenOut, nil
}

func GetAccessTokenFromWebRequest(r *http.Request, client_id, client_secret, redirect_uri string) (AccessTokenResponse, error) {
	var response AccessTokenResponse
	var err error

	if r.URL.Query().Get("error") != "" && r.URL.Query().Get("error_description") != "" {
		err = errors.New(fmt.Sprintf(`{"error": "%s, "error_description": "%s"}`,
			r.URL.Query().Get("error"),
			r.URL.Query().Get("error_description"),
		))
	}

	code := r.URL.Query().Get("code")
	if code != "" {
		response, err = GetAccessToken(client_id, client_secret, redirect_uri, code)
	}

	return response, err
}
