package freenas

import (
  "fmt"
  "log"

  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
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
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREENAS_HOST", nil),
			},
		},

		ConfigureFunc: providerConfigure,
	}
}


func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host:     data.Get("host").(string),
		User:     data.Get("user").(string),
		Password: data.Get("password").(string),
	}

	apis, err := config.APIs()
	if err != nil {
		return nil, fmt.Errorf("Error creating APIs: %s", err)
	}

	log.Println("[INFO] Initializing FreeNAS client")

	err = apis.AuthAPI.Login(config.User, config.Password)
	if err != nil {
		return nil, fmt.Errorf("Error logging in user %s: %s", config.User, err)
	}

	return apis, nil
}
