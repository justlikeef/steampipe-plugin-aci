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

type AccPortBlock struct {
	Id                           string
	AccessPortSelectorDn         string
	Description                  string
	Name                         string
	Annotation                   string
	FromCard                     string
	FromPort                     string
	NameAlias                    string
	ToCard                       string
	ToPort                       string
	RelationInfrarsAccBndlSubgrp string
}

func tableACIAccPortBlock() *plugin.Table {
	return &plugin.Table{
		Name:        "aci_access_port_block",
		Description: "ACI Access Port Block",
		List: &plugin.ListConfig{
			Hydrate: listLeafAccPortBlock,
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
				Name:        "access_port_selector_dn",
				Description: "(Required) Distinguished name of parent Access Port Selector or Spine Access Port Selector object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "(Optional) Specifies a description of the policy definition.",
				Type:        proto.ColumnType_STRING,
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
				Name:        "from_card",
				Description: "(Optional) The beginning (from-range) of the card range block for the leaf access port block. Allowed value range is 1-100. Default value is \"1\".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "from_port",
				Description: "(Optional) The beginning (from-range) of the port range block for the leaf access port block. Allowed value range is 1-127. Default value is \"1\".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name_alias",
				Description: "(Optional) Name alias for object leaf access bundle policy group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "to_card",
				Description: "(Optional) The beginning (to-range) of the card range block for the leaf access port block. Allowed value range is 1-100. Default value is \"1\".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "to_port",
				Description: "(Optional) The beginning (to-range) of the port range block for the leaf access port block. Allowed value range is 1-127. Default value is \"1\".",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "relation_infra_rs_acc_bndl_subgrp",
				Description: "(Optional) Relation to class infraAccBndlSubgrp. Cardinality - N_TO_ONE. Type - String.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

func AccPortBlockFromContainerList(cont *container.Container, ctx context.Context, d *plugin.QueryData) error {

	ObjContainerAttribs := cont.S("attributes")
	ObjContainerChildren, err := cont.S("children").Children()

	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Error getting objects: %v", err))
	}
	log.Printf("[TRACE] ObjContainerChildren: %s ", ObjContainerChildren)

	for i := 0; i < len(ObjContainerChildren); i++ {
		ObjChildrenMap, err := ObjContainerChildren[i].ChildrenMap()
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("Error getting object attributes: %v", err))
		}

		log.Printf("[TRACE] ObjChildrenMap: %s ", ObjChildrenMap)

		for reltype, relattribs := range ObjChildrenMap {
			log.Printf("[TRACE] reltype: %s", reltype)

			if reltype == "infraHPortS" {
				infraHPortSAttribs := relattribs.S("attributes")
				infraHPortsChildren, err := relattribs.S("children").Children()

				log.Printf("[TRACE] infaHPortsChildren: %s ", infraHPortsChildren)

				if err != nil {
					return fmt.Errorf(fmt.Sprintf("Error getting object attributes: %v", err))
				}

				for j := 0; j < len(infraHPortsChildren); j++ {

					infraHPortsChildrenMap, err := infraHPortsChildren[j].ChildrenMap()
					if err != nil {
						return fmt.Errorf(fmt.Sprintf("Error getting object attributes: %v", err))
					}

					for relchildtype, relchildattribs := range infraHPortsChildrenMap {

						if relchildtype == "infraPortBlk" {
							recordobj := AccPortBlock{}
							recordobj.Id = client.G(ObjContainerAttribs, "dn") + "/" + client.G(infraHPortSAttribs, "rn") + "/" + client.G(relchildattribs.S("attributes"), "rn")
							recordobj.AccessPortSelectorDn = client.G(ObjContainerAttribs, "dn") + "/" + client.G(infraHPortSAttribs, "rn")
							recordobj.Description = client.G(relchildattribs.S("attributes"), "descr")
							recordobj.Name = client.G(relchildattribs.S("attributes"), "name")
							recordobj.Annotation = client.G(relchildattribs.S("attributes"), "annotation")
							recordobj.FromCard = client.G(relchildattribs.S("attributes"), "fromCard")
							recordobj.FromPort = client.G(relchildattribs.S("attributes"), "fromPort")
							recordobj.NameAlias = client.G(relchildattribs.S("attributes"), "nameAlias")
							recordobj.ToCard = client.G(relchildattribs.S("attributes"), "toCard")
							recordobj.ToPort = client.G(relchildattribs.S("attributes"), "toPort")
							//recordobj.RelationInfrarsAccBndlSubgrp = client.G(ObjContainer, "dn")

							log.Printf("[TRACE]: Recordobj: %s ", recordobj)
							d.StreamListItem(ctx, recordobj)
						}
					}
				}
			}
		}
	}

	return nil
}

//// LIST FUNCTION
func listLeafAccPortBlock(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	aciclient, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to APIC: %v", err)
	}

	dnUrl := "/api/node/mo/uni/infra.json?query-target=subtree&target-subtree-class=infraAccPortP&query-target-filter=not(wcard(infraAccPortP.dn,\"__ui_\"))&query-target=children&target-subtree-class=infraAccPortP&rsp-subtree=full"
	leafintprofcont, err := aciclient.ServiceManager.GetViaURL(dnUrl)

	if err != nil {
		return nil, fmt.Errorf("Error getting Port Blocks: %v\nURL: %v", err, dnUrl)
	}

	length, _ := strconv.Atoi(client.G(leafintprofcont, "totalCount"))
	for i := 0; i < length; i++ {
		err = AccPortBlockFromContainerList(leafintprofcont.S("imdata").Index(i).S("infraAccPortP"), ctx, d)
	}

	return nil, nil
}
