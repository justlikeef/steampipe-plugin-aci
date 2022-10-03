/*
Package aws implements a steampipe plugin for aci.

This plugin provides data that Steampipe uses to present foreign
tables that represent Cisco ACI resources.
*/
package aci

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-aci"

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-aci",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"aci_access_pol_leaf_int_pg_vpc_bundle_grp": tableACIAccPolLeafIntPolVPCBundleGrp(),
			"aci_access_pol_leaf_int_prof":              tableACIAccPolLeafIntProf(),
			"aci_access_port_selector":                  tableACIAccPortSelector(),
			"aci_access_port_block":                     tableACIAccPortBlock(),
			"aci_access_sub_port_block":                 tableACIAccSubPortBlock(),
		},
	}
	return p
}
