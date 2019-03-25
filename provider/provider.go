package provider

import "github.com/hashicorp/terraform/helper/schema"

func New() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"myfs_text_file": resourceTextFile(),
		},
	}
}
