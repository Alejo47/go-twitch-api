package authorization

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ValidateToken(code string) (ValidateTokenResponse, error) {
	var validationOut ValidateTokenResponse

	client := http.Client{}
	uri := "https://id.twitch.tv/oauth2/validate"

	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Bearer "+code)
	if err != nil {
		return validationOut, err
	}

	res, err := client.Do(req)
	if err != nil {
		return validationOut, err
	}

	out, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return validationOut, errors.New(fmt.Sprintf("Invalid status code returned: %d", res.StatusCode))
	}
	if err != nil {
		return validationOut, err
	}

	err = json.Unmarshal(out, &validationOut)
	if err != nil {
		return validationOut, err
	}

	return validationOut, nil
}
