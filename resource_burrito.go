package main

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBurrito() *schema.Resource {
	return &schema.Resource{
		Create: resourceBurritoCreate,
		Read:   resourceBurritoRead,
		Update: resourceBurritoUpdate,
		Delete: resourceBurritoDelete,

		Schema: map[string]*schema.Schema{
			"ingredients": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func resourceBurritoCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(fmt.Sprintf("%x", rand.Int()))
	return resourceBurritoRead(d, m)
}

func resourceBurritoRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBurritoUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceBurritoRead(d, m)
}

func resourceBurritoDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
