package pritunlko

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"pritunl_host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Pritunl server address or hostname",
			},
			"pritunl_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Pritunl server API token",
			},
			"pritunl_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Pritunl server API secret",
			},
			"pritunl_insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Skip server TLS certificate validation",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pritunlko_organization": resourcePritunlkoOrganization(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	pritunlClient := NewClient(d)

	return pritunlClient, nil
}
