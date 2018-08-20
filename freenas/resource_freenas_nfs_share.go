package freenas

import "github.com/hashicorp/terraform/helper/schema"

func resourceFreenasNfsShare() *schema.Resource {
	return &schema.Resource{
		Create: resourceFreenasNfsShareCreate,
		Read:   resourceFreenasNfsShareRead,
		Update: resourceFreenasNfsShareUpdate,
		Delete: resourceFreenasNfsShareDelete,
		Importer: &schema.ResourceImporter{
			State: resourceFreenasNfsShareImport,
		},

		SchemaVersion: 1,
		MigrateState:  resourceFreenasNfsShareMigrateState,
		Schema: map[string]*schema.Schema{
			"Comment": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A Comment describing the purpose of the share.",
				Required:    true,
			},
			"Paths": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The paths this share is linking.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceFreenasNfsShareCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFreenasNfsShareRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFreenasNfsShareUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFreenasNfsShareDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFreenasNfsShareImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, nil
}
