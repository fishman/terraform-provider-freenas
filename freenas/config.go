package freenas

import (
	"context"
	"fmt"

	freenas "github.com/fishman/go-freenas"
	"github.com/hashicorp/terraform/helper/schema"
)

type FreenasClient struct {
	client *freenas.Client
}

type Config struct {
	Debug         bool
	User          string
	Password      string
	FreenasServer string
	InsecureFlag  bool
}

func (c *Config) Client() (*FreenasClient, error) {
	client := new(FreenasClient)
	client.client = freenas.NewClient(
		&freenas.Config{
			Address:  c.FreenasServer,
			User:     c.User,
			Password: c.Password,
		},
	)

	_, _, err := client.client.Users.List(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Error logging in user %s: %s", c.User, err)
	}

	return client, nil

}

func NewConfig(d *schema.ResourceData) (*Config, error) {
	server := d.Get("freenas_server").(string)

	if server == "" {
		return nil, fmt.Errorf("freenas_server must be provided")
	}

	c := &Config{
		User:          d.Get("user").(string),
		Password:      d.Get("password").(string),
		InsecureFlag:  d.Get("allow_unverified_ssl").(bool),
		FreenasServer: server,
	}

	return c, nil
}
