// Package gitio is a client for http://git.io URL shortener.
package gitio

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const gitioAPI = "http://git.io"

// Shorten returns a short version of an URL, or an error otherwise.
// Please note that it's not guaranteed the code will be accepted by git.io,
// the random one may be used instead.
func Shorten(longURL, code string) (shortURL *url.URL, err error) {
	if len(longURL) == 0 {
		return nil, errors.New("no URL provided")
	}

	form := make(url.Values)
	form.Add("url", longURL)
	form.Add("code", code)
	resp, err := http.PostForm(gitioAPI, form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusCreated:
		return resp.Location()
	case 422:
		return nil, errors.New("only GitGub/Gist links are accepted")
	default:
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
}

// CheckTaken checks if the provided custom code has already been taken on git.io.
func CheckTaken(code string) (bool, error) {
	if len(code) == 0 {
		return false, errors.New("no code provided")
	}
	resp, err := http.Get(fmt.Sprintf("%s/%s", gitioAPI, code))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusNotFound:
		return false, nil
	case http.StatusFound:
		return true, nil
	default:
		// probably it's taken
		return true, nil
	}
}
