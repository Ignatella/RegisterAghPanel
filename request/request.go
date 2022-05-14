package request

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Reserve(client *http.Client, targetUrl string) error {
	resp, err := client.PostForm(targetUrl, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(
			"[HTTP] Reserve with status code: " + strconv.Itoa(resp.StatusCode) + ". Status: " + resp.Status,
		)
	}

	return nil
}

func GetBookingPage(client *http.Client, pageUrl string) (string, error) {
	resp, err := client.Get(pageUrl)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(
			"[HTTP] GetBookingPage with status code: " + strconv.Itoa(resp.StatusCode) + ". Status: " + resp.Status,
		)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	html := string(body)

	return html, nil
}

func Login(client *http.Client, loginUrl string, username string, password string) error {
	resp, err := client.PostForm(loginUrl, url.Values{
		"_username": {username},
		"_password": {password},
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(
			"[HTTP] Login with status code: " + strconv.Itoa(resp.StatusCode) + ". Status: " + resp.Status,
		)
	}

	return nil
}
