package freenas

import (
	"context"
	"fmt"
	"log"
	"path"

	"github.com/fishman/go-freenas/freenas"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFreenasStorageDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceFreenasStorageDatasetCreate,
		Read:   resourceFreenasStorageDatasetRead,
		Update: resourceFreenasStorageDatasetUpdate,
		Delete: resourceFreenasStorageDatasetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"parent": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The parent of the Dataset",
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the Dataset",
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
			},
			"atime": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(Toggles),
			},
			"case_sensitivity": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(Toggles),
			},
			"comment": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A Comment describing the purpose of the dataset.",
				Optional:    true,
			},
			"compression": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CompressionLevels),
			},
			"dedup": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(Toggles),
			},
			"quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"readonly": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(Toggles),
			},
			"recordsize": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(RecordSizes),
			},
			"refquota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"refreservation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reservation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFreenasStorageDatasetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client

	name := d.Get("name").(string)
	parent := d.Get("parent").(string)

	dataset := freenas.Dataset{
		Name:    name,
		Comment: d.Get("comment").(string),
	}
	_, _, err := client.Datasets.Create(context.TODO(), parent, dataset)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", parent, name))

	return resourceFreenasStorageDatasetRead(d, meta)
}

func resourceFreenasStorageDatasetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client

	dataset, err := getFreenasStorageDataset(d, client)
	if err != nil {
		log.Printf("[WARN] Dataset (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	d.Set("comment", dataset.Comment)
	d.Set("name", path.Base(dataset.Name))
	d.Set("parent", path.Dir(dataset.Name))
	return nil
}

func resourceFreenasStorageDatasetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client
	_, err := getFreenasStorageDataset(d, client)
	if err != nil {
		d.SetId("")
		return nil
	}

	editedDataset := freenas.Dataset{
		Comment: d.Get("comment").(string),
	}

	_, _, err = client.Datasets.Edit(context.TODO(), d.Id(), editedDataset)
	if err != nil {
		return err
	}

	return resourceFreenasStorageDatasetRead(d, meta)
}

func resourceFreenasStorageDatasetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*FreenasClient).client

	_, err := client.Datasets.Delete(context.TODO(), d.Id())
	return err
}

func getFreenasStorageDataset(d *schema.ResourceData, freenas *freenas.Client) (*freenas.Dataset, error) {
	dataset, _, err := freenas.Datasets.Get(context.TODO(), d.Id())
	return dataset, err
}
