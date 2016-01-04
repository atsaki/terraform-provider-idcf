package idcf

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IDCF_API_KEY", nil),
				Description: "The API key for compute and DNS operations",
			},

			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IDCF_SECRET_KEY", nil),
				Description: "The secret key for compute and DNS operations",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			// "idcf_dns_zone":   resourceIDCFDNSZone(),
			"idcf_dns_record": resourceIDCFDNSRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey:    d.Get("api_key").(string),
		SecretKey: d.Get("secret_key").(string),
	}

	return config.Client()
}
