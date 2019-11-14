package pritunlko

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type PritunlClient struct {
	PritunlHost     string
	PritunlToken    string
	PritunlSecret   string
	PritunlInsecure bool
}

func NewClient(d *schema.ResourceData) *PritunlClient {
	c := PritunlClient{
		PritunlHost:     d.Get("pritunl_host").(string),
		PritunlToken:    d.Get("pritunl_token").(string),
		PritunlSecret:   d.Get("pritunl_secret").(string),
		PritunlInsecure: d.Get("pritunl_insecure").(bool),
	}

	return &c
}
