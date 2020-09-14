package client

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

type Client struct {
	client *zk.Conn
	host   string
	port   int
}

type ZNode struct {
	Path string
	Data []byte
	Zxid int64
}

func NewClient(host string, port int) (*Client, error) {
	c := new(Client)
	var err error
	c.client, _, err = zk.Connect([]string{fmt.Sprintf("%v:%v", host, port)}, time.Second*10)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot create zookeeper client. err=%s", err))
	}
	return c, err
}

func (c *Client) CreateZNode(path, data string) error {
	fmt.Println(fmt.Sprintf("[INFO] Creating znode %v", path))
	_, err := c.client.Create(path, []byte(data), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Println(fmt.Sprintf("[INFO] ZNode %v created with data %v", path, data))
	}
	return err
}

func (c *Client) ReadZNode(path string) (znode ZNode, err error) {
	fmt.Println(fmt.Sprintf("[INFO] Reading znode %v", path))
	data, stat, err := c.client.Get(path)
	if err != nil {
		fmt.Println(fmt.Sprintf("[ERROR] Error reading znode %s", err))
	}
	znode.Path = path
	znode.Zxid = stat.Czxid
	znode.Data = data
	return
}

func (c *Client) DeleteZnode(path string) error {
	fmt.Println(fmt.Sprintf("[INFO] Deleting znode %v", path))
	// get version first
	stat, err := c.GetZnodeStats(path)
	if err != nil {
		if err.Error() == "zk: node does not exists" {
			return nil
		}
		return err
	}
	return c.client.Delete(path, stat.Version)
}

func (c *Client) GetZnodeStats(path string) (*zk.Stat, error) {
	// get version first
	_, stat, err := c.client.Get(path)
	return stat, err
}

func (c *Client) ZnodeExists(path string) (bool,error) {
	fmt.Println(fmt.Sprintf("[INFO] Checking if znode %v exists", path))
	exists, _, err := c.client.Exists(path)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

func (c *Client) UpdateZnode(path string, data string) error {
	fmt.Println(fmt.Sprintf("[INFO] Updating znode %v", path))
	stat, err := c.GetZnodeStats(path)
	_, err = c.client.Set(path, []byte(data), stat.Version)
	if err != nil {
		return err
	}
	return nil
}