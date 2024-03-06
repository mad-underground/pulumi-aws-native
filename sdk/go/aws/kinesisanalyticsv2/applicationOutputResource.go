// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisanalyticsv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationOutput
//
// Deprecated: ApplicationOutputResource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApplicationOutputResource struct {
	pulumi.CustomResourceState

	ApplicationName pulumi.StringOutput                       `pulumi:"applicationName"`
	AwsId           pulumi.StringOutput                       `pulumi:"awsId"`
	Output          ApplicationOutputResourceOutputTypeOutput `pulumi:"output"`
}

// NewApplicationOutputResource registers a new resource with the given unique name, arguments, and options.
func NewApplicationOutputResource(ctx *pulumi.Context,
	name string, args *ApplicationOutputResourceArgs, opts ...pulumi.ResourceOption) (*ApplicationOutputResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.Output == nil {
		return nil, errors.New("invalid value for required argument 'Output'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationOutputResource
	err := ctx.RegisterResource("aws-native:kinesisanalyticsv2:ApplicationOutputResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationOutputResource gets an existing ApplicationOutputResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationOutputResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationOutputResourceState, opts ...pulumi.ResourceOption) (*ApplicationOutputResource, error) {
	var resource ApplicationOutputResource
	err := ctx.ReadResource("aws-native:kinesisanalyticsv2:ApplicationOutputResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationOutputResource resources.
type applicationOutputResourceState struct {
}

type ApplicationOutputResourceState struct {
}

func (ApplicationOutputResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOutputResourceState)(nil)).Elem()
}

type applicationOutputResourceArgs struct {
	ApplicationName string                              `pulumi:"applicationName"`
	Output          ApplicationOutputResourceOutputType `pulumi:"output"`
}

// The set of arguments for constructing a ApplicationOutputResource resource.
type ApplicationOutputResourceArgs struct {
	ApplicationName pulumi.StringInput
	Output          ApplicationOutputResourceOutputTypeInput
}

func (ApplicationOutputResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOutputResourceArgs)(nil)).Elem()
}

type ApplicationOutputResourceInput interface {
	pulumi.Input

	ToApplicationOutputResourceOutput() ApplicationOutputResourceOutput
	ToApplicationOutputResourceOutputWithContext(ctx context.Context) ApplicationOutputResourceOutput
}

func (*ApplicationOutputResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOutputResource)(nil)).Elem()
}

func (i *ApplicationOutputResource) ToApplicationOutputResourceOutput() ApplicationOutputResourceOutput {
	return i.ToApplicationOutputResourceOutputWithContext(context.Background())
}

func (i *ApplicationOutputResource) ToApplicationOutputResourceOutputWithContext(ctx context.Context) ApplicationOutputResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutputResourceOutput)
}

type ApplicationOutputResourceOutput struct{ *pulumi.OutputState }

func (ApplicationOutputResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationOutputResource)(nil)).Elem()
}

func (o ApplicationOutputResourceOutput) ToApplicationOutputResourceOutput() ApplicationOutputResourceOutput {
	return o
}

func (o ApplicationOutputResourceOutput) ToApplicationOutputResourceOutputWithContext(ctx context.Context) ApplicationOutputResourceOutput {
	return o
}

func (o ApplicationOutputResourceOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOutputResource) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o ApplicationOutputResourceOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationOutputResource) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o ApplicationOutputResourceOutput) Output() ApplicationOutputResourceOutputTypeOutput {
	return o.ApplyT(func(v *ApplicationOutputResource) ApplicationOutputResourceOutputTypeOutput { return v.Output }).(ApplicationOutputResourceOutputTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationOutputResourceInput)(nil)).Elem(), &ApplicationOutputResource{})
	pulumi.RegisterOutputType(ApplicationOutputResourceOutput{})
}
