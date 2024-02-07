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

// Resource Type definition for AWS::ServiceCatalog::LaunchTemplateConstraint
//
// Deprecated: LaunchTemplateConstraint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type LaunchTemplateConstraint struct {
	pulumi.CustomResourceState

	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	PortfolioId    pulumi.StringOutput    `pulumi:"portfolioId"`
	ProductId      pulumi.StringOutput    `pulumi:"productId"`
	Rules          pulumi.StringOutput    `pulumi:"rules"`
}

// NewLaunchTemplateConstraint registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplateConstraint(ctx *pulumi.Context,
	name string, args *LaunchTemplateConstraintArgs, opts ...pulumi.ResourceOption) (*LaunchTemplateConstraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"portfolioId",
		"productId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LaunchTemplateConstraint
	err := ctx.RegisterResource("aws-native:servicecatalog:LaunchTemplateConstraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplateConstraint gets an existing LaunchTemplateConstraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplateConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateConstraintState, opts ...pulumi.ResourceOption) (*LaunchTemplateConstraint, error) {
	var resource LaunchTemplateConstraint
	err := ctx.ReadResource("aws-native:servicecatalog:LaunchTemplateConstraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplateConstraint resources.
type launchTemplateConstraintState struct {
}

type LaunchTemplateConstraintState struct {
}

func (LaunchTemplateConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateConstraintState)(nil)).Elem()
}

type launchTemplateConstraintArgs struct {
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	Description    *string `pulumi:"description"`
	PortfolioId    string  `pulumi:"portfolioId"`
	ProductId      string  `pulumi:"productId"`
	Rules          string  `pulumi:"rules"`
}

// The set of arguments for constructing a LaunchTemplateConstraint resource.
type LaunchTemplateConstraintArgs struct {
	AcceptLanguage pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	PortfolioId    pulumi.StringInput
	ProductId      pulumi.StringInput
	Rules          pulumi.StringInput
}

func (LaunchTemplateConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateConstraintArgs)(nil)).Elem()
}

type LaunchTemplateConstraintInput interface {
	pulumi.Input

	ToLaunchTemplateConstraintOutput() LaunchTemplateConstraintOutput
	ToLaunchTemplateConstraintOutputWithContext(ctx context.Context) LaunchTemplateConstraintOutput
}

func (*LaunchTemplateConstraint) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplateConstraint)(nil)).Elem()
}

func (i *LaunchTemplateConstraint) ToLaunchTemplateConstraintOutput() LaunchTemplateConstraintOutput {
	return i.ToLaunchTemplateConstraintOutputWithContext(context.Background())
}

func (i *LaunchTemplateConstraint) ToLaunchTemplateConstraintOutputWithContext(ctx context.Context) LaunchTemplateConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateConstraintOutput)
}

type LaunchTemplateConstraintOutput struct{ *pulumi.OutputState }

func (LaunchTemplateConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplateConstraint)(nil)).Elem()
}

func (o LaunchTemplateConstraintOutput) ToLaunchTemplateConstraintOutput() LaunchTemplateConstraintOutput {
	return o
}

func (o LaunchTemplateConstraintOutput) ToLaunchTemplateConstraintOutputWithContext(ctx context.Context) LaunchTemplateConstraintOutput {
	return o
}

func (o LaunchTemplateConstraintOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplateConstraint) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o LaunchTemplateConstraintOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LaunchTemplateConstraint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LaunchTemplateConstraintOutput) PortfolioId() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplateConstraint) pulumi.StringOutput { return v.PortfolioId }).(pulumi.StringOutput)
}

func (o LaunchTemplateConstraintOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplateConstraint) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

func (o LaunchTemplateConstraintOutput) Rules() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplateConstraint) pulumi.StringOutput { return v.Rules }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateConstraintInput)(nil)).Elem(), &LaunchTemplateConstraint{})
	pulumi.RegisterOutputType(LaunchTemplateConstraintOutput{})
}
