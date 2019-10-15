package pritunlko

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vi7/terraform-provider-pritunlko/pritunlko/internal/errortypes"
)

type organization struct {
	Id   string
	Name string
}

type organizationPostData struct {
	Name string `json:"name"`
}

type organizationPutData struct {
	Name string `json:"name"`
}

type organizationData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/*******************
	RESOURCE FUNCTIONS
********************/

func resourcePritunlkoOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourcePritunlkoOrganizationCreate,
		Read:   resourcePritunlkoOrganizationRead,
		Update: resourcePritunlkoOrganizationUpdate,
		Delete: resourcePritunlkoOrganizationDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Organization name",
				Required:    true,
			},
		},
	}
}

func resourcePritunlkoOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	pritunlClient := meta.(*PritunlClient)
	org := loadOrganization(d)

	data, err := organizationGet(pritunlClient, org)

	if err != nil {
		return err
	}

	if data != nil {
		/* pritunl itself doesn't give a shit about non-unique org names
			but we are handling that on our side */
		err = &errortypes.RequestError{
			errors.Newf("resource_organization: Organization \"%s\" already exists on the server and doesn't seem to be managed by this terraform config", data.Name),
		}
		return err
	}

	if data == nil {
		if err = organizationPost(pritunlClient, org); err != nil {
			return err
		}
	}

	return resourcePritunlkoOrganizationRead(d, meta)
}

func resourcePritunlkoOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	pritunlClient := meta.(*PritunlClient)
	org := loadOrganization(d)

	data, err := organizationGet(pritunlClient, org)
	if err != nil {
		return err
	}

	if data == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", data.Name)
	d.SetId(data.Id)

	return nil
}

func resourcePritunlkoOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	pritunlClient := meta.(*PritunlClient)
	org := loadOrganization(d)

	err := organizationPut(pritunlClient, org)
	if err != nil {
		return err
	}

	return resourcePritunlkoOrganizationRead(d, meta)
}

func resourcePritunlkoOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	pritunlClient := meta.(*PritunlClient)
	org := loadOrganization(d)

	err := organizationDel(pritunlClient, org)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

/*****************
	HELPER FUNCTIONS
******************/

func loadOrganization(d *schema.ResourceData) *organization {
	org := organization{
		Id:   d.Id(),
		Name: d.Get("name").(string),
	}

	return &org
}

func organizationGet(pritunlClient *PritunlClient, org *organization) (*organizationData, error) {

	req := Request{
		Method: "GET",
		Path:   "/organization",
	}

	orgs := &[]organizationData{}

	_, err := req.Do(pritunlClient, orgs)
	if err != nil {
		err = &errortypes.RequestError{
			errors.Wrapf(err, "resource_organization: Error getting the org: %s", org.Name),
		}
		return nil, err
	}

	for _, data := range *orgs {
		if data.Name == org.Name {
			return &data, nil
		}
	}
	return nil, nil
}

func organizationPut(pritunlClient *PritunlClient, org *organization) error {

	req := Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/organization/%s", org.Id),
		Json: &organizationPutData{
			Name: org.Name,
		},
	}

	_, err := req.Do(pritunlClient, nil)
	if err != nil {
		return err
	}

	return nil
}

func organizationPost(pritunlClient *PritunlClient, org *organization) error {

	req := Request{
		Method: "POST",
		Path:   "/organization",
		Json: &organizationPostData{
			Name: org.Name,
		},
	}

	_, err := req.Do(pritunlClient, nil)
	if err != nil {
		return err
	}

	return nil
}

func organizationDel(pritunlClient *PritunlClient, org *organization) error {

	req := Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/organization/%s", org.Id),
	}

	_, err := req.Do(pritunlClient, nil)
	if err != nil {
		return err
	}

	return nil
}
