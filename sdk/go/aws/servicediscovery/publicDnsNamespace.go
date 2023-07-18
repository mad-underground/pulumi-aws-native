// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace
//
// Deprecated: PublicDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type PublicDnsNamespace struct {
	pulumi.CustomResourceState

	Arn          pulumi.StringOutput                   `pulumi:"arn"`
	Description  pulumi.StringPtrOutput                `pulumi:"description"`
	HostedZoneId pulumi.StringOutput                   `pulumi:"hostedZoneId"`
	Name         pulumi.StringOutput                   `pulumi:"name"`
	Properties   PublicDnsNamespacePropertiesPtrOutput `pulumi:"properties"`
	Tags         PublicDnsNamespaceTagArrayOutput      `pulumi:"tags"`
}

// NewPublicDnsNamespace registers a new resource with the given unique name, arguments, and options.
func NewPublicDnsNamespace(ctx *pulumi.Context,
	name string, args *PublicDnsNamespaceArgs, opts ...pulumi.ResourceOption) (*PublicDnsNamespace, error) {
	if args == nil {
		args = &PublicDnsNamespaceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicDnsNamespace
	err := ctx.RegisterResource("aws-native:servicediscovery:PublicDnsNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicDnsNamespace gets an existing PublicDnsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicDnsNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicDnsNamespaceState, opts ...pulumi.ResourceOption) (*PublicDnsNamespace, error) {
	var resource PublicDnsNamespace
	err := ctx.ReadResource("aws-native:servicediscovery:PublicDnsNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicDnsNamespace resources.
type publicDnsNamespaceState struct {
}

type PublicDnsNamespaceState struct {
}

func (PublicDnsNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDnsNamespaceState)(nil)).Elem()
}

type publicDnsNamespaceArgs struct {
	Description *string                       `pulumi:"description"`
	Name        *string                       `pulumi:"name"`
	Properties  *PublicDnsNamespaceProperties `pulumi:"properties"`
	Tags        []PublicDnsNamespaceTag       `pulumi:"tags"`
}

// The set of arguments for constructing a PublicDnsNamespace resource.
type PublicDnsNamespaceArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Properties  PublicDnsNamespacePropertiesPtrInput
	Tags        PublicDnsNamespaceTagArrayInput
}

func (PublicDnsNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDnsNamespaceArgs)(nil)).Elem()
}

type PublicDnsNamespaceInput interface {
	pulumi.Input

	ToPublicDnsNamespaceOutput() PublicDnsNamespaceOutput
	ToPublicDnsNamespaceOutputWithContext(ctx context.Context) PublicDnsNamespaceOutput
}

func (*PublicDnsNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDnsNamespace)(nil)).Elem()
}

func (i *PublicDnsNamespace) ToPublicDnsNamespaceOutput() PublicDnsNamespaceOutput {
	return i.ToPublicDnsNamespaceOutputWithContext(context.Background())
}

func (i *PublicDnsNamespace) ToPublicDnsNamespaceOutputWithContext(ctx context.Context) PublicDnsNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDnsNamespaceOutput)
}

type PublicDnsNamespaceOutput struct{ *pulumi.OutputState }

func (PublicDnsNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDnsNamespace)(nil)).Elem()
}

func (o PublicDnsNamespaceOutput) ToPublicDnsNamespaceOutput() PublicDnsNamespaceOutput {
	return o
}

func (o PublicDnsNamespaceOutput) ToPublicDnsNamespaceOutputWithContext(ctx context.Context) PublicDnsNamespaceOutput {
	return o
}

func (o PublicDnsNamespaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o PublicDnsNamespaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PublicDnsNamespaceOutput) HostedZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) pulumi.StringOutput { return v.HostedZoneId }).(pulumi.StringOutput)
}

func (o PublicDnsNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PublicDnsNamespaceOutput) Properties() PublicDnsNamespacePropertiesPtrOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) PublicDnsNamespacePropertiesPtrOutput { return v.Properties }).(PublicDnsNamespacePropertiesPtrOutput)
}

func (o PublicDnsNamespaceOutput) Tags() PublicDnsNamespaceTagArrayOutput {
	return o.ApplyT(func(v *PublicDnsNamespace) PublicDnsNamespaceTagArrayOutput { return v.Tags }).(PublicDnsNamespaceTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicDnsNamespaceInput)(nil)).Elem(), &PublicDnsNamespace{})
	pulumi.RegisterOutputType(PublicDnsNamespaceOutput{})
}
