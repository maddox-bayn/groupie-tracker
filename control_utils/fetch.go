package control_utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/config"
	"net/http"
)

func Fetch(endpoint string, dest any) error {
	Urlresp, err := http.Get(config.Api_url + endpoint)
	if err != nil {
		return err
	}
	defer Urlresp.Body.Close()

	if Urlresp.StatusCode != http.StatusOK {
		return fmt.Errorf("Api return status code %d", Urlresp.StatusCode)
	}
	err = json.NewDecoder(Urlresp.Body).Decode(dest)
	if err != nil {
		return err
	}
	return nil
}
