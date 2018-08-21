package freenas

import (
	"context"
	"log"
	"strconv"

	"github.com/fishman/go-freenas/freenas"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFreenasNfsShare() *schema.Resource {
	return &schema.Resource{
		Create: resourceFreenasNfsShareCreate,
		Read:   resourceFreenasNfsShareRead,
		Update: resourceFreenasNfsShareUpdate,
		Delete: resourceFreenasNfsShareDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A Comment describing the purpose of the share.",
				Optional:    true,
			},
			"paths": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The paths this share is linking.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
			},
		},
	}
}

func resourceFreenasNfsShareCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client

	pathItems := d.Get("paths").([]interface{})
	var paths []string
	for _, item := range pathItems {
		paths = append(paths, item.(string))
	}
	log.Printf("[DEBUG] Creating nfs share: %v\n", paths)

	share := freenas.NfsShare{
		Comment: d.Get("comment").(string),
		Paths:   paths,
	}
	newShare, _, err := client.NfsShares.Create(context.TODO(), share)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(newShare.ID, 10))

	return resourceFreenasNfsShareRead(d, meta)
}

func resourceFreenasNfsShareRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client

	share, err := getFreenasNfsShare(d, client)
	if err != nil {
		log.Printf("[WARN] NfsShare (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	d.Set("comment", share.Comment)
	d.Set("paths", share.Paths)
	return nil
}

func resourceFreenasNfsShareUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client
	share, err := getFreenasNfsShare(d, client)
	if err != nil {
		d.SetId("")
		return nil
	}

	pathItems := d.Get("paths").([]interface{})
	var paths []string
	for _, item := range pathItems {
		paths = append(paths, item.(string))
	}

	editedShare := freenas.NfsShare{
		Comment: d.Get("comment").(string),
		Paths:   paths,
	}

	share, _, err = client.NfsShares.Edit(context.TODO(), share.ID, editedShare)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(share.ID, 10))
	return resourceFreenasNfsShareRead(d, meta)
}

func resourceFreenasNfsShareDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client

	id, err := strconv.ParseInt(d.Id(), 10, 32)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	_, err = client.NfsShares.Delete(context.TODO(), id)
	return err
}

func getFreenasNfsShare(d *schema.ResourceData, freenas *freenas.Client) (*freenas.NfsShare, error) {
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return nil, unconvertibleIdErr(d.Id(), err)
	}

	nfsShare, _, err := freenas.NfsShares.Get(context.TODO(), id)
	return nfsShare, err
}
