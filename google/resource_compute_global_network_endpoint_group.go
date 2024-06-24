// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeGlobalNetworkEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalNetworkEndpointGroupCreate,
		Read:   resourceComputeGlobalNetworkEndpointGroupRead,
		Delete: resourceComputeGlobalNetworkEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalNetworkEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCEName,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network_endpoint_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"INTERNET_IP_PORT", "INTERNET_FQDN_PORT"}),
				Description:  `Type of network endpoints in this network endpoint group. Possible values: ["INTERNET_IP_PORT", "INTERNET_FQDN_PORT"]`,
			},
			"default_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The default port used if the port number is not specified in the
network endpoint.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeGlobalNetworkEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeGlobalNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeGlobalNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeGlobalNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !isEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	defaultPortProp, err := expandComputeGlobalNetworkEndpointGroupDefaultPort(d.Get("default_port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_port"); !isEmptyValue(reflect.ValueOf(defaultPortProp)) && (ok || !reflect.DeepEqual(v, defaultPortProp)) {
		obj["defaultPort"] = defaultPortProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networkEndpointGroups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalNetworkEndpointGroup: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalNetworkEndpointGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalNetworkEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/networkEndpointGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating GlobalNetworkEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalNetworkEndpointGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating GlobalNetworkEndpointGroup %q: %#v", d.Id(), res)

	return resourceComputeGlobalNetworkEndpointGroupRead(d, meta)
}

func resourceComputeGlobalNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalNetworkEndpointGroup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalNetworkEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenComputeGlobalNetworkEndpointGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenComputeGlobalNetworkEndpointGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network_endpoint_type", flattenComputeGlobalNetworkEndpointGroupNetworkEndpointType(res["networkEndpointType"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("default_port", flattenComputeGlobalNetworkEndpointGroupDefaultPort(res["defaultPort"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpointGroup: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalNetworkEndpointGroup: %s", err)
	}

	return nil
}

func resourceComputeGlobalNetworkEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalNetworkEndpointGroup: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalNetworkEndpointGroup %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalNetworkEndpointGroup")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting GlobalNetworkEndpointGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalNetworkEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalNetworkEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/networkEndpointGroups/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/networkEndpointGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalNetworkEndpointGroupName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalNetworkEndpointGroupDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalNetworkEndpointGroupNetworkEndpointType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalNetworkEndpointGroupDefaultPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandComputeGlobalNetworkEndpointGroupName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalNetworkEndpointGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalNetworkEndpointGroupNetworkEndpointType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalNetworkEndpointGroupDefaultPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
