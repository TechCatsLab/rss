/*
 *  Revision History:
 *      Initial: 2018/08/02    Wang Huajian
 */

package client

import (
	"io"
	"net/http"
	"crypto/tls"
)

//Read a string url and returns a Channel struct, error
func Read(url string) (io.ReadCloser, error) {
	return ReadWithClient(url, http.DefaultClient)
}

//Read without certificate check
func InsecureRead(url string) (io.ReadCloser, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return ReadWithClient(url, client)
}

//ReadWithClient a string url and custom client that must match the Fetcher interface
//returns a Channel struct, error
func ReadWithClient(url string, client *http.Client) (io.ReadCloser, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, http.ErrNoLocation
	}

	return response.Body, nil
}
