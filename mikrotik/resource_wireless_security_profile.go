package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

type wirelessSecurityProfile struct {
	client *client.Mikrotik
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &wirelessSecurityProfile{}
	_ resource.ResourceWithConfigure   = &wirelessSecurityProfile{}
	_ resource.ResourceWithImportState = &wirelessSecurityProfile{}
)

// NewWirelessSecurityProfileResource is a helper function to simplify the provider implementation.
func NewWirelessSecurityProfileResource() resource.Resource {
	return &wirelessSecurityProfile{}
}

func (r *wirelessSecurityProfile) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Mikrotik)
}

// Metadata returns the resource type name.
func (r *wirelessSecurityProfile) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_security_profile"
}

// Schema defines the schema for the resource.
func (s *wirelessSecurityProfile) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Creates a MikroTik Wireless Security Profile.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Unique identifier for this resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the security profile.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Security mode. Supported values: none, static-keys-optional, static-keys-required, dynamic-keys.",
			},
			"authentication_types": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: tftypes.StringType,
				Description: "List of authentication types. Supported values: wpa-psk, wpa2-psk, wpa-eap, wpa2-eap.",
			},
			"wpa2_pre_shared_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "WPA2 pre-shared key. Used when authentication includes wpa2-psk.",
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *wirelessSecurityProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var terraformModel wirelessSecurityProfileModel
	var mikrotikModel client.WirelessSecurityProfile
	GenericCreateResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

// Read refreshes the Terraform state with the latest data.
func (r *wirelessSecurityProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var terraformModel wirelessSecurityProfileModel
	var mikrotikModel client.WirelessSecurityProfile
	GenericReadResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *wirelessSecurityProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var terraformModel wirelessSecurityProfileModel
	var mikrotikModel client.WirelessSecurityProfile
	GenericUpdateResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *wirelessSecurityProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var terraformModel wirelessSecurityProfileModel
	var mikrotikModel client.WirelessSecurityProfile
	GenericDeleteResource(&terraformModel, &mikrotikModel, r.client)(ctx, req, resp)
}

func (r *wirelessSecurityProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type wirelessSecurityProfileModel struct {
	Id                  tftypes.String `tfsdk:"id"`
	Name                tftypes.String `tfsdk:"name"`
	Mode                tftypes.String `tfsdk:"mode"`
	AuthenticationTypes tftypes.Set    `tfsdk:"authentication_types"`
	WPA2PreSharedKey    tftypes.String `tfsdk:"wpa2_pre_shared_key"`
}
