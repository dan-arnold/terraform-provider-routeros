package routeros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceRoutingPimSmInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceRoutingPimSmInterfaceRead,
		Schema: map[string]*schema.Schema{
			MetaResourcePath: PropResourcePath("/routing/pimsm/interface"),
			MetaId:           PropId(Id),
			KeyFilter: PropFilterRw,
      "address": {
        Type:     schema.TypeString,
        Computed: true,
      },
      "designated-router": {
        Type:     schema.TypeBool,
        Computed: true,
      },
      "dr": {
        Type:     schema.TypeBool,
        Computed: true,
      },
      "dynamic": {
        Type:     schema.TypeBool,
        Computed: true,
      },
      "instance": {
        Type:     schema.TypeString,
        Computed: true,
        Description: "Name of the PIM instance this interface template belongs to.",
      },
      "join-tracking": {
        Type:     schema.TypeBool,
        Computed: true,
      },
		},
	}
}

func datasourceRoutingPimSmInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := DatasourceRoutingPimSmInterface().Schema
	path := s[MetaResourcePath].Default.(string)

	res, err := ReadItemsFiltered(buildReadFilter(d.Get(KeyFilter).(map[string]interface{})), path, m.(Client))
	if err != nil {
		return diag.FromErr(err)
	}

	return MikrotikResourceDataToTerraformDatasource(res, "interfaces", s, d)
}
