package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
    "afi" "",
    "bsm-forward-back": "",
    "crp-advertise-contained": "",
    "name": "",
    "rp-hash-mask-length": "",
    "rp-static-override": "",
    "ssm-range": "",
    "switch-to-spt": "",
    "switch-to-spt-bytes": "",
    "switch-to-spt-interval": "",
    "vrf": ""
  }
*/

// ResourceRoutingPimSmInstance https://help.mikrotik.com/docs/display/ROS/PIM-SM
func ResourceRoutingPimSmInstance() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/routing/pimsm/instance"),
		MetaId:           PropId(Id),

		KeyComment:  PropCommentRw,
		KeyDisabled: PropDisabledRw,
		"afi": {
			Type:        schema.TypeString,
			Optional:    true,
      Default:     "ipv4",
			Description: "Specifies address family for PIM.",
      ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
		},
		"bsm-forward-back": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Currently not implemented.",
		},
		"crp-advertise-contained": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Currently not implemented.",
		},
		KeyInactive: PropInactiveRo,
		KeyName: PropNameForceNewRw,
		"rp-hash-mask-length": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The hash mask allows changing how many groups to map to one of the matching RPs.",
			ValidateFunc: validation.IntBetween(0, 4294967295),
		},
		"rp-static-override": {
			Type:     schema.TypeBool,
			Optional: true,
      Default:  false,
			Description: "Changes the selection priority for static RP. When disabled, the bootstrap RP set has a higher priority. When enabled, static RP has a higher priority.",
		},
		"ssm-range": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Currently not implemented.",
      ValidateFunc: validation.StringInSlice([]string{"IPv4", "IPv6"}, false),
		},
		"switch-to-spt ": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether to switch to Shortest Path Tree (SPT) if multicast data bandwidth threshold is reached. The router will not proceed from protocol phase one (register encapsulation) to native multicast traffic flow if this option is disabled. It is recommended to enable this option.",
		},
		"switch-to-spt-bytes": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Multicast data bandwidth threshold. Switching to Shortest Path Tree (SPT) happens if this threshold is reached in the specified time interval. If a value of 0 is configured, switching will happen immediately.",
			ValidateFunc: validation.IntBetween(0, 4294967295),
		},
    "switch-to-spt-interval": {
      Type:        schema.TypeString,
      Optional:    true,
      Description: "Time interval in which to account for multicast data bandwidth, used in conjunction with switch-to-spt-bytes to determine if the switching threshold is reached.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
    },
		KeyVrf: PropVrfRw,
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
