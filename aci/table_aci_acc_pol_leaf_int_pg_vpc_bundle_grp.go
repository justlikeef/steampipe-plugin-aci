package aci

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"steampipe-plugin-aci/client"
	"steampipe-plugin-aci/container"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type AccPolLeafIntPolVPCBundleGrp struct {
	Id          string
	Name        string
	Annotation  string
	Description string
	LagT        string
	NameAlias   string
	//	RelationInfraRtAccBaseGrp         string
	RelationInfraRsStormctrlIfPol     string
	RelationInfraRsLldpIfPol          string
	RelationInfraRsMacsecIfPol        string
	RelationInfraRsQosDppIfPol        string
	RelationInfraRsHIfPol             string
	RelationInfraRsMcpIfPol           string
	RelationInfraRsL2PortSecurityPol  string
	RelationInfraRsCoppIfPol          string
	RelationInfraRsLacpPol            string
	RelationInfraRsCdpIfPol           string
	RelationInfraRsQosPfcIfPol        string
	RelationInfraRsQosSdIfPol         string
	RelationInfraRsMonIfInfraPol      string
	RelationInfraRsFcIfPol            string
	RelationInfraRsQosIngressDppIfPol string
	RelationInfraRsQosEgressDppIfPol  string
	RelationInfraRsL2IfPol            string
	RelationInfraRsStpIfPol           string
	RelationInfraRsAttEntP            string
	RelationInfraRsL2InstPol          string
}

func tableACIAccPolLeafIntPolVPCBundleGrp() *plugin.Table {
	return &plugin.Table{
		Name:        "aci_access_pol_leaf_int_pg_vpc_bundle_grp",
		Description: "ACI Access Policy Leaf Interface Policy Groups VPC Bundle Groups",
		List: &plugin.ListConfig{
			Hydrate: listPortBundleGrps,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "id",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Distinguished Name of an object within the Cisco Application Policy Infrastructure Controller (APIC) GUI",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "name",
				Description: "The bundled ports group name. This name can be up to 64 alphanumeric characters. Note that you cannot change this name after the object has been saved.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "annotation",
				Description: "(Optional) Annotation for object Access Port Selector.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "(Optional) Specifies a description of the policy definition.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "lag_t",
				Description: "(Optional) The bundled ports group link aggregation type: port channel vs virtual port channel. Allowed values are \"not-aggregated\", \"node\" and \"link\". Default is \"link\".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name_alias",
				Description: "(Optional) Name alias for object leaf access bundle policy group.",
				Type:        proto.ColumnType_STRING,
			},
			//{
			//	Name:        "relation_infra_rt_acc_base_grp",
			//	Description: "A target relation to the access policy group providing port configuration. Cardinality - N_TO_ONE. Type - String.",
			//	Type:        proto.ColumnType_STRING,
			//},
			//			{
			//				Name:        "relation_infra_rs_span_v_src_grp",
			//				Description: "(Optional) Relation to class spanVSrcGrp. Cardinality - N_TO_M. Type - Set of String.",
			//				Type:        proto.ColumnType_STRING,
			//			},
			{
				Name:        "relation_infra_rs_stormctrl_if_pol",
				Description: "(Optional) Relation to class stormctrlIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_lldp_if_pol",
				Description: "(Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_macsec_if_pol",
				Description: "(Optional) Relation to class macsecIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_qos_dpp_if_pol",
				Description: "(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_h_if_pol",
				Description: "(Optional) Relation to class fabricHIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			//			{
			//				Name:        "relation_infra_rs_netflow_monitor_pol",
			//				Description: "(Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map.",
			//				Type:        proto.ColumnType_STRING,
			//			},
			//			{
			//				Name:        "flt_type",
			//				Description: "(Required) Netflow IP filter type. Allowed values: "ce", "ipv4", "ipv6".",
			//				Type:        proto.ColumnType_STRING,
			//			},
			//			{
			//				Name:        "target_dn",
			//				Description: "(Required) Distinguished name of target Netflow Monitor object.",
			//				Type:        proto.ColumnType_STRING,
			//			},
			//			{
			//				Name:        "relation_infra_rs_l2_port_auth_pol",
			//				Description: "(Optional) Relation to class l2PortAuthPol. Cardinality - N_TO_ONE. Type - String.",
			//				Type:        proto.ColumnType_STRING,
			//			},
			{
				Name:        "relation_infra_rs_mcp_if_pol",
				Description: "(Optional) Relation to class mcpIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_l2_port_security_pol",
				Description: "(Optional) Relation to class l2PortSecurityPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_copp_if_pol",
				Description: "(Optional) Relation to class coppIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			//			{
			//				Name:        "relation_infra_rs_span_v_dest_grp",
			//				Description: "(Optional) Relation to class spanVDestGrp. Cardinality - N_TO_M. Type - Set of String.",
			//				Type:        proto.ColumnType_STRING,
			//			},
			{
				Name:        "relation_infra_rs_lacp_pol",
				Description: "(Optional) Relation to class lacpLagPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_cdp_if_pol",
				Description: "(Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_qos_pfc_if_pol",
				Description: "(Optional) Relation to class qosPfcIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_qos_sd_if_pol",
				Description: "(Optional) Relation to class qosSdIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_mon_if_infra_pol",
				Description: "(Optional) Relation to class monInfraPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_fc_if_pol",
				Description: "(Optional) Relation to class fcIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_qos_ingress_dpp_if_pol",
				Description: "(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_qos_egress_dpp_if_pol",
				Description: "(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_l2_if_pol",
				Description: "(Optional) Relation to class l2IfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_stp_if_pol",
				Description: "(Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_att_ent_p",
				Description: "(Optional) Relation to class infraAttEntityP. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_l2_inst_pol",
				Description: "(Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

func AccPolLeafIntProfVPCBundleGrpFromContainerList(cont *container.Container, index int) (*AccPolLeafIntPolVPCBundleGrp, error) {

	PCVPCInterfacePolicyGroupCont := cont.S("imdata").Index(index).S("infraAccBndlGrp", "attributes")
	PCVPCInterfacePolicyGroupChildren, err := cont.S("imdata").Index(index).S("infraAccBndlGrp", "children").Children()

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error getting bundles: %v", err))
	}

	BPGobj := AccPolLeafIntPolVPCBundleGrp{}

	BPGobj.Id = client.G(PCVPCInterfacePolicyGroupCont, "dn")
	BPGobj.Name = client.G(PCVPCInterfacePolicyGroupCont, "name")
	BPGobj.Annotation = client.G(PCVPCInterfacePolicyGroupCont, "annotation")
	BPGobj.Description = client.G(PCVPCInterfacePolicyGroupCont, "descr")
	BPGobj.LagT = client.G(PCVPCInterfacePolicyGroupCont, "lagT")
	BPGobj.NameAlias = client.G(PCVPCInterfacePolicyGroupCont, "nameAlias")

	for i := 0; i < len(PCVPCInterfacePolicyGroupChildren); i++ {
		PCVPCInterfacePolicyGroupChildrenMap, err := PCVPCInterfacePolicyGroupChildren[i].ChildrenMap()

		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("Error getting bundles: %v", err))
		}

		for reltype, relattribs := range PCVPCInterfacePolicyGroupChildrenMap {
			switch reltype {
			//case "infraRtAccBaseGrp":
			//	BPGobj.RelationInfraRtAccBaseGrp = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsStormctrlIfPol":
				BPGobj.RelationInfraRsStormctrlIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsLldpIfPol":
				BPGobj.RelationInfraRsLldpIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsMacsecIfPol":
				BPGobj.RelationInfraRsMacsecIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsQosDppIfPol":
				BPGobj.RelationInfraRsQosDppIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsHIfPol":
				BPGobj.RelationInfraRsHIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsMcpIfPol":
				BPGobj.RelationInfraRsMcpIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsL2PortSecurityPol":
				BPGobj.RelationInfraRsL2PortSecurityPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsCoppIfPol":
				BPGobj.RelationInfraRsCoppIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsLacpPol":
				BPGobj.RelationInfraRsLacpPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsCdpIfPol":
				BPGobj.RelationInfraRsCdpIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsQosPfcIfPol":
				BPGobj.RelationInfraRsQosPfcIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsQosSdIfPol":
				BPGobj.RelationInfraRsQosSdIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsMonIfInfraPol":
				BPGobj.RelationInfraRsMonIfInfraPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsFcIfPol":
				BPGobj.RelationInfraRsFcIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsQosIngressDppIfPol":
				BPGobj.RelationInfraRsQosIngressDppIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsQosEgressDppIfPol":
				BPGobj.RelationInfraRsQosEgressDppIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsL2IfPol":
				BPGobj.RelationInfraRsL2IfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsStpIfPol":
				BPGobj.RelationInfraRsStpIfPol = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsAttEntP":
				BPGobj.RelationInfraRsAttEntP = client.G(relattribs.S("attributes"), "tDn")
			case "infraRsL2InstPol":
				BPGobj.RelationInfraRsL2InstPol = client.G(relattribs.S("attributes"), "tDn")
			}
		}
	}

	return &BPGobj, nil
}

//// LIST FUNCTION
func listPortBundleGrps(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	aciclient, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to APIC: %v", err)
	}

	dnUrl := "/api/node/mo/uni/infra/funcprof.json?query-target=subtree&target-subtree-class=infraAccBndlGrp&query-target-filter=and(not(wcard(infraAccBndlGrp.dn,\"__ui_\")),not(eq(infraAccBndlGrp.lagT,\"link\")))&rsp-subtree=children&order-by=infraAccBndlGrp.name|asc"
	bpgcont, err := aciclient.ServiceManager.GetViaURL(dnUrl)

	if err != nil {
		return nil, fmt.Errorf("Error getting bundles: %v\nURL: %v", err, dnUrl)
	}

	recordobj := &AccPolLeafIntPolVPCBundleGrp{}
	length, _ := strconv.Atoi(client.G(bpgcont, "totalCount"))
	for i := 0; i < length; i++ {
		recordobj, _ = AccPolLeafIntProfVPCBundleGrpFromContainerList(bpgcont, i)
		log.Printf("[TRACE]: Using Proxy Server: %s ", recordobj)
		d.StreamListItem(ctx, recordobj)
	}

	return nil, nil
}
