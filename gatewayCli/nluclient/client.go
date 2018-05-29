package nluclient

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type Client struct {
	baseSling *sling.Sling
}

var DefaultClient = New(http.DefaultClient)

func New(client *http.Client) *Client {
	return &Client{
		baseSling: sling.New().Client(client),
	}
}

func (nlu *Client) SetBase(url *string) *Client {
	nlu.baseSling.Base(*url)
	return nlu
}

func (nlu *Client) Query(authToken string, req *QueryRequest) (*http.Response, *QueryResponse, error) {
	sling := nlu.baseSling.New()
	respBody := &QueryResponse{}
	req.User.AccessTokens = []map[string]interface{}{
		{
			"value":         "dncjcmksdncfdjhfjmdkncvjdfvdmvdfjvf",
			"refresh_token": "bb7f01d9cabcfe1838cf4284662f88253649526506c56825771e13fcee9f93ec",
			"provider":      "iamplus",
			"expiry":        1616780470,
		},
	}
	sling = sling.Post("/api/v1/ai/query").BodyJSON(req)
	sling.Add("Accept", "application/vnd.iamplus+json;version=0").
		Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
	httpResp, err := sling.ReceiveSuccess(respBody)
	return httpResp, respBody, err
}
