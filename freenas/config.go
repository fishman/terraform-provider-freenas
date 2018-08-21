package freenas

import (
	"context"
	"fmt"

	"github.com/fishman/go-freenas/freenas"
	"github.com/hashicorp/terraform/helper/schema"
)

type FreenasClient struct {
	client      *freenas.Client
	StopContext context.Context
}

type Config struct {
	Debug         bool
	User          string
	Password      string
	FreenasServer string
	Insecure      bool
}

func (c *Config) Client() (interface{}, error) {
	var client FreenasClient
	ctx := context.Background()

	client.client = freenas.NewClient(c.FreenasServer, c.User, c.Password)

	_, _, err := client.client.Users.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error logging in user %s: %s", c.User, err)
	}

	return &client, nil

}

func NewConfig(d *schema.ResourceData) (*Config, error) {
	server := d.Get("freenas_server").(string)

	if server == "" {
		return nil, fmt.Errorf("freenas_server must be provided")
	}

	c := &Config{
		User:          d.Get("user").(string),
		Password:      d.Get("password").(string),
		Insecure:      d.Get("allow_unverified_ssl").(bool),
		FreenasServer: server,
	}

	return c, nil
}
