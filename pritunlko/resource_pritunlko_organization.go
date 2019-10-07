package pritunlko

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePritunlkoOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourcePritunlkoOrganizationCreate,
		Read:   resourcePritunlkoOrganizationRead,
		Update: resourcePritunlkoOrganizationUpdate,
		Delete: resourcePritunlkoOrganizationDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePritunlkoOrganizationCreate(d *schema.ResourceData, m interface{}) error {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationGet(prvdr, sch)
	if err != nil {
		return err
	}

	if data != nil {
		sch.Id = data.Id

		data, err = organizationPut(prvdr, sch)
		if err != nil {
			return err
		}
	}

	if data == nil {
		data, err = organizationPost(prvdr, sch)
		if err != nil {
			return err
		}
	}

	d.SetId(data.Id)

	return resourcePritunlkoOrganizationRead(d, m)
}

func resourcePritunlkoOrganizationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePritunlkoOrganizationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePritunlkoOrganizationRead(d, m)
}

func resourcePritunlkoOrganizationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
