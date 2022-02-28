package raspberry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/goodaye/fakeeyes/protos/command"
	"github.com/goodaye/fakeeyes/protos/response"
	"google.golang.org/protobuf/proto"
)

var APIPrefix = ""

type Client struct {
	address string
	url     *url.URL
}

func NewClient(address string) (*Client, error) {

	r := Client{
		address: address,
	}

	url, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	r.url = url

	return &r, nil
}

//
func (c *Client) httpproxy(api string, req interface{}, resp interface{}, header http.Header) error {
	var err error

	relurl := path.Join(APIPrefix, api)
	u, err := url.Parse(relurl)
	if err != nil {
		return err
	}
	queryURL := c.url.ResolveReference(u).String()
	var reader io.Reader
	var reqstr string
	if req == nil {
		reader = strings.NewReader("")
	} else if bydata, ok := req.([]byte); ok {
		reader = bytes.NewReader(bydata)
	} else {

		reqbody, err := json.Marshal(req)
		if err != nil {
			return err
		}
		reqstr = string(reqbody)
		reader = strings.NewReader(reqstr)
	}
	httpreq, err := http.NewRequest(http.MethodPost, queryURL, reader)
	if err != nil {
		return err
	}
	// sb := c.sign(reqstr)
	// httpreq.Header.Add("Timestamp", fmt.Sprintf("%d", sb.Timestamp))
	// httpreq.Header.Add("Signature", sb.Sign)
	// httpreq.Header.Add("Accesskey", sb.Accesskey)
	if header != nil {
		httpreq.Header = header
	}
	httpclt := http.Client{}
	httpresp, err := httpclt.Do(httpreq)
	if err != nil {
		return err
	}
	defer httpresp.Body.Close()
	body, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return err
	}
	if httpresp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Code: %d , HTTP Response: %s ", httpresp.StatusCode, string(body))
	}
	rm := response.ReturnMessage{}

	err = json.Unmarshal(body, &rm)
	if err != nil {
		return err
	}
	if !rm.Success {
		err = fmt.Errorf("ErrorCode: %s  ErrorMessage: %s", rm.ErrorCode, rm.ErrorMessage)
		return err
	}
	if resp == nil {
		return nil
	}
	err = json.Unmarshal(body, resp)
	return err
}
func (c *Client) HealthCheck() error {
	api := "/api/healthcheck"
	err := c.httpproxy(api, nil, nil, nil)
	return err
}
func (c *Client) Motion(op *command.DeviceOperation) (err error) {

	api := "/api/motion"
	req, err := proto.Marshal(op)
	if err != nil {
		return
	}
	fmt.Println(req)
	err = c.httpproxy(api, req, nil, nil)
	return
}
