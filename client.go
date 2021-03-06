package dumbno

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type FilterRequest struct {
	Src   string `json:"src"`
	Dst   string `json:"dst"`
	Sport int    `json:"sport"`
	Dport int    `json:"dport"`
	Proto string `json:"proto"`
}

type Client struct {
	conn net.Conn
	//The total timeout for the request/response operation
	Timeout time.Duration
}

func NewClient(endpoint string) (*Client, error) {
	var c Client
	conn, err := net.Dial("udp", endpoint)
	c.conn = conn
	c.Timeout = 2 * time.Second
	return &c, err
}

func (c *Client) AddACL(req FilterRequest) error {
	clone := req
	if clone.Proto == "" {
		clone.Proto = "ip"
	}
	clone.Proto = strings.ToLower(clone.Proto)
	bytes, err := json.Marshal(clone)
	if err != nil {
		return errors.Wrap(err, "AddACL: error building json request")
	}
	c.conn.SetDeadline(time.Now().Add(c.Timeout))
	_, err = c.conn.Write(bytes)
	if err != nil {
		return errors.Wrap(err, "AddACL: error sending json request")
	}

	resp := make([]byte, 128)
	n, err := c.conn.Read(resp)
	resp = resp[:n]
	//log.Printf("Got %d bytes in response: %s", n, string(resp))
	if err != nil {
		return errors.Wrap(err, "AddACL: no response from api")
	}
	if string(resp) != "ok" {
		return fmt.Errorf("AddACL: Expected 'ok', got: %q", resp)
	}
	return nil
}
