package freenas

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	var p *schema.Provider
	// The actual provider
	p = &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREENAS_USER", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREENAS_PASSWORD", nil),
				Sensitive:   true,
			},
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREENAS_SERVER", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"freenas_nfs_share": resourceFreenasNfsShare(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := Config{
			FreenasServer: d.Get("server").(string),
			User:          d.Get("user").(string),
			Password:      d.Get("password").(string),
		}

		log.Println("[DEBUG] Initializing FreeNAS client")
		meta, err := config.Client()
		if err != nil {
			return nil, err
		}

		meta.(*FreenasClient).StopContext = p.StopContext()

		return meta, nil
	}
}
