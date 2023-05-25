// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::NetworkManager::SiteToSiteVpnAttachment Resource Type definition.
type SiteToSiteVpnAttachment struct {
	pulumi.CustomResourceState

	// The ID of the attachment.
	AttachmentId pulumi.StringOutput `pulumi:"attachmentId"`
	// The policy rule number associated with the attachment.
	AttachmentPolicyRuleNumber pulumi.IntOutput `pulumi:"attachmentPolicyRuleNumber"`
	// The type of attachment.
	AttachmentType pulumi.StringOutput `pulumi:"attachmentType"`
	// The ARN of a core network for the VPC attachment.
	CoreNetworkArn pulumi.StringOutput `pulumi:"coreNetworkArn"`
	// The ID of a core network where you're creating a site-to-site VPN attachment.
	CoreNetworkId pulumi.StringOutput `pulumi:"coreNetworkId"`
	// Creation time of the attachment.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The Region where the edge is located.
	EdgeLocation pulumi.StringOutput `pulumi:"edgeLocation"`
	// Owner account of the attachment.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The attachment to move from one segment to another.
	ProposedSegmentChange SiteToSiteVpnAttachmentProposedSegmentChangePtrOutput `pulumi:"proposedSegmentChange"`
	// The ARN of the Resource.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// The name of the segment that attachment is in.
	SegmentName pulumi.StringOutput `pulumi:"segmentName"`
	// The state of the attachment.
	State pulumi.StringOutput `pulumi:"state"`
	// Tags for the attachment.
	Tags SiteToSiteVpnAttachmentTagArrayOutput `pulumi:"tags"`
	// Last update time of the attachment.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The ARN of the site-to-site VPN attachment.
	VpnConnectionArn pulumi.StringOutput `pulumi:"vpnConnectionArn"`
}

// NewSiteToSiteVpnAttachment registers a new resource with the given unique name, arguments, and options.
func NewSiteToSiteVpnAttachment(ctx *pulumi.Context,
	name string, args *SiteToSiteVpnAttachmentArgs, opts ...pulumi.ResourceOption) (*SiteToSiteVpnAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CoreNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'CoreNetworkId'")
	}
	if args.VpnConnectionArn == nil {
		return nil, errors.New("invalid value for required argument 'VpnConnectionArn'")
	}
	var resource SiteToSiteVpnAttachment
	err := ctx.RegisterResource("aws-native:networkmanager:SiteToSiteVpnAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteToSiteVpnAttachment gets an existing SiteToSiteVpnAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteToSiteVpnAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteToSiteVpnAttachmentState, opts ...pulumi.ResourceOption) (*SiteToSiteVpnAttachment, error) {
	var resource SiteToSiteVpnAttachment
	err := ctx.ReadResource("aws-native:networkmanager:SiteToSiteVpnAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteToSiteVpnAttachment resources.
type siteToSiteVpnAttachmentState struct {
}

type SiteToSiteVpnAttachmentState struct {
}

func (SiteToSiteVpnAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteToSiteVpnAttachmentState)(nil)).Elem()
}

type siteToSiteVpnAttachmentArgs struct {
	// The ID of a core network where you're creating a site-to-site VPN attachment.
	CoreNetworkId string `pulumi:"coreNetworkId"`
	// The attachment to move from one segment to another.
	ProposedSegmentChange *SiteToSiteVpnAttachmentProposedSegmentChange `pulumi:"proposedSegmentChange"`
	// Tags for the attachment.
	Tags []SiteToSiteVpnAttachmentTag `pulumi:"tags"`
	// The ARN of the site-to-site VPN attachment.
	VpnConnectionArn string `pulumi:"vpnConnectionArn"`
}

// The set of arguments for constructing a SiteToSiteVpnAttachment resource.
type SiteToSiteVpnAttachmentArgs struct {
	// The ID of a core network where you're creating a site-to-site VPN attachment.
	CoreNetworkId pulumi.StringInput
	// The attachment to move from one segment to another.
	ProposedSegmentChange SiteToSiteVpnAttachmentProposedSegmentChangePtrInput
	// Tags for the attachment.
	Tags SiteToSiteVpnAttachmentTagArrayInput
	// The ARN of the site-to-site VPN attachment.
	VpnConnectionArn pulumi.StringInput
}

func (SiteToSiteVpnAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteToSiteVpnAttachmentArgs)(nil)).Elem()
}

type SiteToSiteVpnAttachmentInput interface {
	pulumi.Input

	ToSiteToSiteVpnAttachmentOutput() SiteToSiteVpnAttachmentOutput
	ToSiteToSiteVpnAttachmentOutputWithContext(ctx context.Context) SiteToSiteVpnAttachmentOutput
}

func (*SiteToSiteVpnAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteToSiteVpnAttachment)(nil)).Elem()
}

func (i *SiteToSiteVpnAttachment) ToSiteToSiteVpnAttachmentOutput() SiteToSiteVpnAttachmentOutput {
	return i.ToSiteToSiteVpnAttachmentOutputWithContext(context.Background())
}

func (i *SiteToSiteVpnAttachment) ToSiteToSiteVpnAttachmentOutputWithContext(ctx context.Context) SiteToSiteVpnAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteToSiteVpnAttachmentOutput)
}

type SiteToSiteVpnAttachmentOutput struct{ *pulumi.OutputState }

func (SiteToSiteVpnAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteToSiteVpnAttachment)(nil)).Elem()
}

func (o SiteToSiteVpnAttachmentOutput) ToSiteToSiteVpnAttachmentOutput() SiteToSiteVpnAttachmentOutput {
	return o
}

func (o SiteToSiteVpnAttachmentOutput) ToSiteToSiteVpnAttachmentOutputWithContext(ctx context.Context) SiteToSiteVpnAttachmentOutput {
	return o
}

// The ID of the attachment.
func (o SiteToSiteVpnAttachmentOutput) AttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.AttachmentId }).(pulumi.StringOutput)
}

// The policy rule number associated with the attachment.
func (o SiteToSiteVpnAttachmentOutput) AttachmentPolicyRuleNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.IntOutput { return v.AttachmentPolicyRuleNumber }).(pulumi.IntOutput)
}

// The type of attachment.
func (o SiteToSiteVpnAttachmentOutput) AttachmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.AttachmentType }).(pulumi.StringOutput)
}

// The ARN of a core network for the VPC attachment.
func (o SiteToSiteVpnAttachmentOutput) CoreNetworkArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.CoreNetworkArn }).(pulumi.StringOutput)
}

// The ID of a core network where you're creating a site-to-site VPN attachment.
func (o SiteToSiteVpnAttachmentOutput) CoreNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.CoreNetworkId }).(pulumi.StringOutput)
}

// Creation time of the attachment.
func (o SiteToSiteVpnAttachmentOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The Region where the edge is located.
func (o SiteToSiteVpnAttachmentOutput) EdgeLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.EdgeLocation }).(pulumi.StringOutput)
}

// Owner account of the attachment.
func (o SiteToSiteVpnAttachmentOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// The attachment to move from one segment to another.
func (o SiteToSiteVpnAttachmentOutput) ProposedSegmentChange() SiteToSiteVpnAttachmentProposedSegmentChangePtrOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) SiteToSiteVpnAttachmentProposedSegmentChangePtrOutput {
		return v.ProposedSegmentChange
	}).(SiteToSiteVpnAttachmentProposedSegmentChangePtrOutput)
}

// The ARN of the Resource.
func (o SiteToSiteVpnAttachmentOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

// The name of the segment that attachment is in.
func (o SiteToSiteVpnAttachmentOutput) SegmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.SegmentName }).(pulumi.StringOutput)
}

// The state of the attachment.
func (o SiteToSiteVpnAttachmentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Tags for the attachment.
func (o SiteToSiteVpnAttachmentOutput) Tags() SiteToSiteVpnAttachmentTagArrayOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) SiteToSiteVpnAttachmentTagArrayOutput { return v.Tags }).(SiteToSiteVpnAttachmentTagArrayOutput)
}

// Last update time of the attachment.
func (o SiteToSiteVpnAttachmentOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The ARN of the site-to-site VPN attachment.
func (o SiteToSiteVpnAttachmentOutput) VpnConnectionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteToSiteVpnAttachment) pulumi.StringOutput { return v.VpnConnectionArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteToSiteVpnAttachmentInput)(nil)).Elem(), &SiteToSiteVpnAttachment{})
	pulumi.RegisterOutputType(SiteToSiteVpnAttachmentOutput{})
}
