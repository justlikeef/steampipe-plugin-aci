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

type AccPolLeafIntProf struct {
	Id                      string
	Name                    string
	Annotation              string
	Description             string
	NameAlias               string
	RelationInfraRtAccPortP string
}

func tableACIAccPolLeafIntProf() *plugin.Table {
	return &plugin.Table{
		Name:        "aci_access_pol_leaf_int_prof",
		Description: "ACI Access Policy Leaf Interface Profile",
		List: &plugin.ListConfig{
			Hydrate: listLeafIntProf,
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
				Name:        "name_alias",
				Description: "(Optional) Name alias for object leaf access bundle policy group.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

func AccPolLeafIntProfFromContainerList(cont *container.Container, index int) (*AccPolLeafIntProf, error) {

	ObjContainer := cont.S("imdata").Index(index).S("infraAccPortP", "attributes")
	//ObjContainerChildren, err := cont.S("imdata").Index(index).S("infraAccBndlGrp", "children").Children()

	//if err != nil {
	//	return nil, fmt.Errorf(fmt.Sprintf("Error getting bundles: %v", err))
	//}

	returnObj := AccPolLeafIntProf{}

	returnObj.Id = client.G(ObjContainer, "dn")
	returnObj.Name = client.G(ObjContainer, "name")
	returnObj.Annotation = client.G(ObjContainer, "annotation")
	returnObj.Description = client.G(ObjContainer, "descr")
	returnObj.NameAlias = client.G(ObjContainer, "nameAlias")

	return &returnObj, nil
}

//// LIST FUNCTION
func listLeafIntProf(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	aciclient, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to APIC: %v", err)
	}

	dnUrl := "/api/node/class/infraAccPortP.json"
	pscont, err := aciclient.ServiceManager.GetViaURL(dnUrl)

	if err != nil {
		return nil, fmt.Errorf("Error getting Port Selectors: %v\nURL: %v", err, dnUrl)
	}

	recordobj := &AccPolLeafIntProf{}
	length, _ := strconv.Atoi(client.G(pscont, "totalCount"))
	for i := 0; i < length; i++ {
		recordobj, _ = AccPolLeafIntProfFromContainerList(pscont, i)
		log.Printf("[TRACE]: Recordobj: %s ", recordobj)
		d.StreamListItem(ctx, recordobj)
	}

	return nil, nil
}
