package aci

import (
	"context"
	"fmt"
	"log"
	"strings"

	"steampipe-plugin-aci/client"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type AccPortSelector struct {
	Id                        string
	LeafInterfaceProfileDn    string
	Name                      string
	AccessPortSelectorType    string
	Annotation                string
	Description               string
	NameAlias                 string
	RelationInfraRsAccBaseGrp string
}

func tableACIAccPortSelector() *plugin.Table {
	return &plugin.Table{
		Name:        "aci_access_port_selector",
		Description: "ACI Access Port Selector",
		List: &plugin.ListConfig{
			Hydrate: listLeafAccPortSelector,
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
				Description: "(Required) Distinguished name of this access port selector.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "leaf_interface_profile_dn",
				Description: "(Required) Distinguished name of parent Leaf Interface Profile object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The bundled ports group name. This name can be up to 64 alphanumeric characters. Note that you cannot change this name after the object has been saved.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "access_port_selector_type",
				Description: "(Required) The host port selector type. Allowed values are \"ALL\" and \"range\". Default is \"ALL\".",
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
				Name:        "name_alias",
				Description: "(Optional) Name alias for object leaf access bundle policy group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_acc_base_grp",
				Description: "(Optional) Relation to class infraAccBaseGrp. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION
func listLeafAccPortSelector(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	aciclient, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to APIC: %v", err)
	}

	dnUrl := "/api/node/class/infraHPortS.json?query-target=self&target-subtree-class=infraRsAccBaseGrp&rsp-subtree=children"
	log.Printf("[TRACE] Getting HPortList from DN: %s ", dnUrl)
	hportresponse, err := aciclient.ServiceManager.GetViaURL(dnUrl)
	if err != nil {
		return nil, fmt.Errorf("Error getting Hport List: %v\nURL: %v", err, dnUrl)
	}
	log.Printf("[TRACE] hportresponse: %s ", hportresponse)

	hportobjlist, err := hportresponse.S("imdata").Children()
	if err != nil {
		return nil, fmt.Errorf("Error getting Hport List: %v", err)
	}
	log.Printf("[TRACE] hportobjlist: %s ", hportobjlist)

	for _, hportobj := range hportobjlist {
		ObjContainer := hportobj.S("infraHPortS")
		recordobj := &AccPortSelector{}
		recordobj.Id = client.G(ObjContainer.S("attributes"), "dn")
		dninfo := strings.Split(client.G(ObjContainer.S("attributes"), "dn"), "/")
		recordobj.LeafInterfaceProfileDn = dninfo[0] + "/" + dninfo[1] + "/" + dninfo[2]
		recordobj.Name = client.G(ObjContainer.S("attributes"), "name")
		recordobj.AccessPortSelectorType = client.G(ObjContainer.S("attributes"), "type")
		recordobj.Annotation = client.G(ObjContainer.S("attributes"), "annotation")
		recordobj.Description = client.G(ObjContainer.S("attributes"), "descr")
		recordobj.NameAlias = client.G(ObjContainer.S("attributes"), "nameAlias")
		infraRsAccBaseGrp := ObjContainer.S("children").S("infraRsAccBaseGrp").S("attributes").S("tDn").String()
		log.Printf("[TRACE] Basegroup: %v", infraRsAccBaseGrp)
		recordobj.RelationInfraRsAccBaseGrp = infraRsAccBaseGrp[2 : len(infraRsAccBaseGrp)-2]

		log.Printf("[TRACE] Recordobj: %s ", recordobj)
		d.StreamListItem(ctx, recordobj)
	}

	return nil, nil
}
