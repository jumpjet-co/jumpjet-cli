package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetEnvVariables(apiKey string, appSlug string, env string) map[string]string {

	url := fmt.Sprintf("http://app.jumpjet.co/api/v1/app/%s/env/%s/var", appSlug, env)
	authorization := fmt.Sprintf("ApiKey %s", apiKey)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", authorization)

	res,err := client.Do(req)

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	// TODO: check status code
	//		fmt.Printf("%d %s", res.StatusCode, res.Status)
	defer res.Body.Close()

	var vars map[string]string

	err2 := json.NewDecoder(res.Body).Decode(&vars)

	if err2 != nil {
		println(err2.Error())
		os.Exit(1)
	}

	return vars
}