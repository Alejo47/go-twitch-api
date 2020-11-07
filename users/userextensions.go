package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUserExtensions(bearer_token, client_id) (UserExtensionsResponse, error) {
	var extensionsResponse UserExtensionsResponse
	client := http.Client{}
	uri := fmt.Sprintf("https://api.twitch.tv/helix/users/extensions/list")
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Bearer "+bearer_token)
	req.Header.Add("Client-Id", client_id)

	if err != nil {
		return extensionsResponse, err
	}

	res, err := client.Do(req)

	if err != nil {
		return extensionsResponse, err
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return extensionsResponse, err
	}
	err = json.Unmarshal(out, &extensionsResponse)
	if err != nil {
		return extensionsResponse, err
	}
	return extensionsResponse, err
}
