package pritunlko

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/dropbox/godropbox/errors"
	"github.com/vi7/terraform-provider-pritunlko/pritunlko/internal/errortypes"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: 2 * time.Minute,
}

type Request struct {
	Method string
	Path   string
	Query  map[string]string
	Json   interface{}
}

func (r *Request) Do(pritunlClient *PritunlClient, respVal interface{}) (*http.Response, error) {

	var resp *http.Response

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: pritunlClient.PritunlInsecure,
		},
	}
	client.Transport = tr

	url := "https://" + pritunlClient.PritunlHost + r.Path

	authTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	authNonce := bson.NewObjectId().Hex()
	authString := strings.Join([]string{
		pritunlClient.PritunlToken,
		authTimestamp,
		authNonce,
		r.Method,
		r.Path,
	}, "&")

	hashFunc := hmac.New(sha256.New, []byte(pritunlClient.PritunlSecret))
	hashFunc.Write([]byte(authString))
	rawSignature := hashFunc.Sum(nil)
	authSig := base64.StdEncoding.EncodeToString(rawSignature)

	var body io.Reader
	if r.Json != nil {
		data, e := json.Marshal(r.Json)
		if e != nil {
			err := errortypes.RequestError{
				errors.Wrap(e, "request: Json marshal error"),
			}
			return resp, err
		}

		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(r.Method, url, body)
	if err != nil {
		err = &errortypes.RequestError{
			errors.Wrap(err, "request: Failed to create request"),
		}
		return resp, err
	}

	if r.Query != nil {
		query := req.URL.Query()

		for key, val := range r.Query {
			query.Add(key, val)
		}

		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("Auth-Token", pritunlClient.PritunlToken)
	req.Header.Set("Auth-Timestamp", authTimestamp)
	req.Header.Set("Auth-Nonce", authNonce)
	req.Header.Set("Auth-Signature", authSig)

	if r.Json != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// send http request to the server
	log.Printf("request: Sending request to the server: %v\n", req)
	resp, err = client.Do(req)
	if err != nil {
		err = &errortypes.RequestError{
			errors.Wrap(err, "request: Request error:\n"),
		}
		return nil, err
	}

	respBody, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		readErr = &errortypes.ReadError{
			errors.Wrap(readErr, "request: Response body read error"),
		}
		return nil, readErr
	}
	defer resp.Body.Close()

	log.Printf("request: Server replied with status code: %d\n", resp.StatusCode)
	log.Printf("request: Server replied with response body: %s\n", respBody)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = &errortypes.RequestError{
			errors.Wrapf(err, "request: Bad response status %d\nServer replied: %s",
				resp.StatusCode, respBody),
		}
		return nil, err
	}

	if respVal != nil {
		if err := json.Unmarshal(respBody, respVal); err != nil {
			err = &errortypes.ParseError{
				errors.Wrapf(err, "request: Failed to parse response body. Server replied: %s\n", respBody),
			}
			return nil, err
		}
	}

	return resp, err
}
