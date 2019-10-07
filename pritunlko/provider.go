package pritunlko

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"pritunlko_organization": resourcePritunlkoOrganization(),
		},
	}
}
