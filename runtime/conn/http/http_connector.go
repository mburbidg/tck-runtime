package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/opengql/tck/runtime/conn"
	"github.com/spyzhov/ajson"
	"io"
	"net/http"
)

type HTTPConnector struct {
	baseUrl string
	client  *http.Client
}

type authRequestBody struct {
	PrincipalID string `json:"principal_id"`
	Password    string `json:"password"`
}

type authResponseBody struct {
	AuthID string `json:"auth_id"`
}

type sessionResponse struct {
	SessionID string `json:"session_id"`
}

func New(baseUrl string) (conn *HTTPConnector) {
	return &HTTPConnector{
		baseUrl: baseUrl,
		client:  &http.Client{},
	}
}

func (c HTTPConnector) Authenticate(principalID, password string) (authID string, err error) {
	body, err := json.Marshal(&authRequestBody{
		PrincipalID: principalID,
		Password:    password,
	})
	if err != nil {
		return "", err
	}
	res, err := http.Post(c.baseUrl+"/auth/id", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http request failed: statusCode=%d", res.StatusCode)
	}
	authResponse := authResponseBody{}
	err = json.NewDecoder(res.Body).Decode(&authResponse)
	if err != nil {
		return "", err
	}
	return authResponse.AuthID, nil
}

func (c HTTPConnector) NewSession(authID string) (sessionID string, err error) {
	req, err := http.NewRequest("POST", c.baseUrl+"/session", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", authID)
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http request failed: statusCode=%d", res.StatusCode)
	}
	sessionRes := sessionResponse{}
	err = json.NewDecoder(res.Body).Decode(&sessionRes)
	if err != nil {
		return "", err
	}
	return sessionRes.SessionID, nil
}

func (c HTTPConnector) Request(authID, sessionID, program string) (outcome *conn.Outcome, err error) {
	req, err := http.NewRequest("POST", c.baseUrl+"/request", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", authID)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request failed: statusCode=%d", res.StatusCode)
	}
	switch v := res.Header.Get("Content-Type"); v {
	case "application/omitted+json":
		outcome, err = c.parseOmittedResponse(res.Body)
	case "application/value+json":
		outcome, err = c.parseValueResponse(res.Body)
	case "application/table+json":
		outcome, err = c.parseTableResponse(res.Body)
	default:
		return nil, fmt.Errorf("unexpected content-type=%s", v)
	}
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c HTTPConnector) parseOmittedResponse(body io.Reader) (outcome *conn.Outcome, err error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	root, err := ajson.Unmarshal(b)
	if err != nil {
		return nil, err
	}
	rootObj, err := root.GetObject()
	if err != nil {
		return nil, err
	}
	if status, ok := rootObj["gql_status"]; ok {
		statusStr, err := status.GetString()
		if err != nil {
			return nil, err
		}
		return &conn.Outcome{
			GQLStatus: statusStr,
			Result:    &conn.OmittedResult{},
		}, nil
	} else {
		return nil, fmt.Errorf("response did not contain the 'gql_status' property")
	}
}

func (c HTTPConnector) parseValueResponse(body io.Reader) (outcome *conn.Outcome, err error) {
	return nil, nil
}

func (c HTTPConnector) parseTableResponse(body io.Reader) (outcome *conn.Outcome, err error) {
	return nil, nil
}
