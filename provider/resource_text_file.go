package provider

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceTextFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceTextFileCreate,
		Read:   resourceTextFileRead,
		Update: resourceTextFileUpdate,
		Delete: resourceTextFileDelete,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTextFileCreate(d *schema.ResourceData, m interface{}) error {
	var (
		path = d.Get("path").(string)
		text = d.Get("text").(string)
	)

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	if _, err := f.Write([]byte(text)); err != nil {
		return errors.Wrap(err, "writing file")
	}
	if err := f.Close(); err != nil {
		return errors.Wrap(err, "closing file")
	}

	d.SetId(path)

	return errors.Wrap(resourceTextFileRead(d, m), "reading created file")
}

func resourceTextFileRead(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	d.Set("text", string(contents))

	return nil
}

func resourceTextFileUpdate(d *schema.ResourceData, m interface{}) error {
	var (
		path = d.Get("path").(string)
		text = d.Get("text").(string)
	)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	if _, err := f.Write([]byte(text)); err != nil {
		return errors.Wrap(err, "writing file")
	}
	if err := f.Close(); err != nil {
		return errors.Wrap(err, "closing file")
	}

	return resourceTextFileRead(d, m)
}

func resourceTextFileDelete(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	if err := os.Remove(path); err != nil {
		return err
	}

	d.SetId("")

	return nil
}
