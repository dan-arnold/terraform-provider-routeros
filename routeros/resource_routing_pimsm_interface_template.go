package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
  {
  }
*/

// ResourceRoutingPimSmInterfaceTemplate https://help.mikrotik.com/docs/display/ROS/PIM-SM
func ResourceRoutingPimSmInterfaceTemplate() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath:   PropResourcePath("/routing/pimsm/interface-template"),
		MetaId:             PropId(Id),
		MetaSetUnsetFields: PropSetUnsetFields("passive"),

		KeyComment: PropCommentRw,
		KeyDisabled: PropDisabledRw,
		KeyInactive: PropInactiveRo,
		"hello-delay": {
			Type:        schema.TypeString,
			Required:    false,
      Default:     "5s",
			Description: "Randomized interval for the initial Hello message on interface startup or detecting new neighbor.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"hello-period": {
			Type:         schema.TypeString,
			Required:     false,
      Default:      "30s",
			Description:  "Periodic interval for Hello messages.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"instance": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Name of the PIM instance this interface template belongs to.",
		},
		"interfaces": {
			Type:      schema.TypeString,
			Optional:  true,
      Default:   "all",
			Description: "List of interfaces that will participate in PIM.",
		},
		"join-prune-period ": {
			Type:         schema.TypeString,
			Required:     false,
      Default:      "1m",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"join-tracking-support": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
			Description: "Sets the value of a Tracking (T) bit in the LAN Prune Delay option in the Hello message. When enabled, a router advertises its willingness to disable Join suppression. it is possible for upstream routers to explicitly track the join membership of individual downstream routers if Join suppression is disabled. Unless all PIM routers on a link negotiate this capability, explicit tracking and the disabling of the Join suppression mechanism are not possible.",
		},
		"override-interval": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "2s500ms",
			Description:      "Sets the maximum time period over which to randomize when scheduling a delayed override Join message on a network that has join suppression enabled.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"priority": {
			Type:        schema.TypeInt,
			Optional:    true,
      Default:     1,
			Description: "The Designated Router (DR) priority. A single Designated Router is elected on each network. The priority is used only if all neighbors have advertised a priority option. Numerically largest priority is preferred. In case of a tie or if priority is not used - the numerically largest IP address is preferred.",
			ValidateFunc: validation.IntBetween(0, 4294967295),
		},
		"propagation-delay": {
			Type:             schema.TypeString,
			Optional:         true,
			Default:          "500ms",
			Description:      "Sets the value for a prune pending timer. It is used by upstream routers to figure out how long they should wait for a Join override message before pruning an interface that has join suppression enabled.",
			ValidateFunc:     ValidationTime,
			DiffSuppressFunc: TimeEquall,
		},
		"source-addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPAddress,
			},
		},
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
