// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::DomainNameApiAssociation
type DomainNameApiAssociation struct {
	pulumi.CustomResourceState

	ApiAssociationIdentifier pulumi.StringOutput `pulumi:"apiAssociationIdentifier"`
	ApiId                    pulumi.StringOutput `pulumi:"apiId"`
	DomainName               pulumi.StringOutput `pulumi:"domainName"`
}

// NewDomainNameApiAssociation registers a new resource with the given unique name, arguments, and options.
func NewDomainNameApiAssociation(ctx *pulumi.Context,
	name string, args *DomainNameApiAssociationArgs, opts ...pulumi.ResourceOption) (*DomainNameApiAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainNameApiAssociation
	err := ctx.RegisterResource("aws-native:appsync:DomainNameApiAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainNameApiAssociation gets an existing DomainNameApiAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainNameApiAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainNameApiAssociationState, opts ...pulumi.ResourceOption) (*DomainNameApiAssociation, error) {
	var resource DomainNameApiAssociation
	err := ctx.ReadResource("aws-native:appsync:DomainNameApiAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainNameApiAssociation resources.
type domainNameApiAssociationState struct {
}

type DomainNameApiAssociationState struct {
}

func (DomainNameApiAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameApiAssociationState)(nil)).Elem()
}

type domainNameApiAssociationArgs struct {
	ApiId      string `pulumi:"apiId"`
	DomainName string `pulumi:"domainName"`
}

// The set of arguments for constructing a DomainNameApiAssociation resource.
type DomainNameApiAssociationArgs struct {
	ApiId      pulumi.StringInput
	DomainName pulumi.StringInput
}

func (DomainNameApiAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameApiAssociationArgs)(nil)).Elem()
}

type DomainNameApiAssociationInput interface {
	pulumi.Input

	ToDomainNameApiAssociationOutput() DomainNameApiAssociationOutput
	ToDomainNameApiAssociationOutputWithContext(ctx context.Context) DomainNameApiAssociationOutput
}

func (*DomainNameApiAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainNameApiAssociation)(nil)).Elem()
}

func (i *DomainNameApiAssociation) ToDomainNameApiAssociationOutput() DomainNameApiAssociationOutput {
	return i.ToDomainNameApiAssociationOutputWithContext(context.Background())
}

func (i *DomainNameApiAssociation) ToDomainNameApiAssociationOutputWithContext(ctx context.Context) DomainNameApiAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameApiAssociationOutput)
}

type DomainNameApiAssociationOutput struct{ *pulumi.OutputState }

func (DomainNameApiAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainNameApiAssociation)(nil)).Elem()
}

func (o DomainNameApiAssociationOutput) ToDomainNameApiAssociationOutput() DomainNameApiAssociationOutput {
	return o
}

func (o DomainNameApiAssociationOutput) ToDomainNameApiAssociationOutputWithContext(ctx context.Context) DomainNameApiAssociationOutput {
	return o
}

func (o DomainNameApiAssociationOutput) ApiAssociationIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNameApiAssociation) pulumi.StringOutput { return v.ApiAssociationIdentifier }).(pulumi.StringOutput)
}

func (o DomainNameApiAssociationOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNameApiAssociation) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

func (o DomainNameApiAssociationOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainNameApiAssociation) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNameApiAssociationInput)(nil)).Elem(), &DomainNameApiAssociation{})
	pulumi.RegisterOutputType(DomainNameApiAssociationOutput{})
}
