package freenas

import (
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

func NewConfig(d *schema.ResourceData) (*Config, error) {
	// Handle backcompat support for vcenter_server; once that is removed,
	// vsphere_server can just become a Required field that is referenced inline
	// in Config below.
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
