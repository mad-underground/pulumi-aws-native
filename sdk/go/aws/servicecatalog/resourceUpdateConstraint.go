// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::ResourceUpdateConstraint
//
// Deprecated: ResourceUpdateConstraint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ResourceUpdateConstraint struct {
	pulumi.CustomResourceState

	AcceptLanguage                pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	Description                   pulumi.StringPtrOutput `pulumi:"description"`
	PortfolioId                   pulumi.StringOutput    `pulumi:"portfolioId"`
	ProductId                     pulumi.StringOutput    `pulumi:"productId"`
	TagUpdateOnProvisionedProduct pulumi.StringOutput    `pulumi:"tagUpdateOnProvisionedProduct"`
}

// NewResourceUpdateConstraint registers a new resource with the given unique name, arguments, and options.
func NewResourceUpdateConstraint(ctx *pulumi.Context,
	name string, args *ResourceUpdateConstraintArgs, opts ...pulumi.ResourceOption) (*ResourceUpdateConstraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.TagUpdateOnProvisionedProduct == nil {
		return nil, errors.New("invalid value for required argument 'TagUpdateOnProvisionedProduct'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceUpdateConstraint
	err := ctx.RegisterResource("aws-native:servicecatalog:ResourceUpdateConstraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceUpdateConstraint gets an existing ResourceUpdateConstraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceUpdateConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceUpdateConstraintState, opts ...pulumi.ResourceOption) (*ResourceUpdateConstraint, error) {
	var resource ResourceUpdateConstraint
	err := ctx.ReadResource("aws-native:servicecatalog:ResourceUpdateConstraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceUpdateConstraint resources.
type resourceUpdateConstraintState struct {
}

type ResourceUpdateConstraintState struct {
}

func (ResourceUpdateConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceUpdateConstraintState)(nil)).Elem()
}

type resourceUpdateConstraintArgs struct {
	AcceptLanguage                *string `pulumi:"acceptLanguage"`
	Description                   *string `pulumi:"description"`
	PortfolioId                   string  `pulumi:"portfolioId"`
	ProductId                     string  `pulumi:"productId"`
	TagUpdateOnProvisionedProduct string  `pulumi:"tagUpdateOnProvisionedProduct"`
}

// The set of arguments for constructing a ResourceUpdateConstraint resource.
type ResourceUpdateConstraintArgs struct {
	AcceptLanguage                pulumi.StringPtrInput
	Description                   pulumi.StringPtrInput
	PortfolioId                   pulumi.StringInput
	ProductId                     pulumi.StringInput
	TagUpdateOnProvisionedProduct pulumi.StringInput
}

func (ResourceUpdateConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceUpdateConstraintArgs)(nil)).Elem()
}

type ResourceUpdateConstraintInput interface {
	pulumi.Input

	ToResourceUpdateConstraintOutput() ResourceUpdateConstraintOutput
	ToResourceUpdateConstraintOutputWithContext(ctx context.Context) ResourceUpdateConstraintOutput
}

func (*ResourceUpdateConstraint) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceUpdateConstraint)(nil)).Elem()
}

func (i *ResourceUpdateConstraint) ToResourceUpdateConstraintOutput() ResourceUpdateConstraintOutput {
	return i.ToResourceUpdateConstraintOutputWithContext(context.Background())
}

func (i *ResourceUpdateConstraint) ToResourceUpdateConstraintOutputWithContext(ctx context.Context) ResourceUpdateConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceUpdateConstraintOutput)
}

type ResourceUpdateConstraintOutput struct{ *pulumi.OutputState }

func (ResourceUpdateConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceUpdateConstraint)(nil)).Elem()
}

func (o ResourceUpdateConstraintOutput) ToResourceUpdateConstraintOutput() ResourceUpdateConstraintOutput {
	return o
}

func (o ResourceUpdateConstraintOutput) ToResourceUpdateConstraintOutputWithContext(ctx context.Context) ResourceUpdateConstraintOutput {
	return o
}

func (o ResourceUpdateConstraintOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceUpdateConstraint) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o ResourceUpdateConstraintOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceUpdateConstraint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceUpdateConstraintOutput) PortfolioId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceUpdateConstraint) pulumi.StringOutput { return v.PortfolioId }).(pulumi.StringOutput)
}

func (o ResourceUpdateConstraintOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceUpdateConstraint) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o ResourceUpdateConstraintOutput) TagUpdateOnProvisionedProduct() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceUpdateConstraint) pulumi.StringOutput { return v.TagUpdateOnProvisionedProduct }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceUpdateConstraintInput)(nil)).Elem(), &ResourceUpdateConstraint{})
	pulumi.RegisterOutputType(ResourceUpdateConstraintOutput{})
}
