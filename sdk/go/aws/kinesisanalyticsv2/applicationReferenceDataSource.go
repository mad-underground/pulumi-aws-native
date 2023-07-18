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

// Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource
//
// Deprecated: ApplicationReferenceDataSource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApplicationReferenceDataSource struct {
	pulumi.CustomResourceState

	ApplicationName     pulumi.StringOutput                                     `pulumi:"applicationName"`
	ReferenceDataSource ApplicationReferenceDataSourceReferenceDataSourceOutput `pulumi:"referenceDataSource"`
}

// NewApplicationReferenceDataSource registers a new resource with the given unique name, arguments, and options.
func NewApplicationReferenceDataSource(ctx *pulumi.Context,
	name string, args *ApplicationReferenceDataSourceArgs, opts ...pulumi.ResourceOption) (*ApplicationReferenceDataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.ReferenceDataSource == nil {
		return nil, errors.New("invalid value for required argument 'ReferenceDataSource'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationReferenceDataSource
	err := ctx.RegisterResource("aws-native:kinesisanalyticsv2:ApplicationReferenceDataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationReferenceDataSource gets an existing ApplicationReferenceDataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationReferenceDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationReferenceDataSourceState, opts ...pulumi.ResourceOption) (*ApplicationReferenceDataSource, error) {
	var resource ApplicationReferenceDataSource
	err := ctx.ReadResource("aws-native:kinesisanalyticsv2:ApplicationReferenceDataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationReferenceDataSource resources.
type applicationReferenceDataSourceState struct {
}

type ApplicationReferenceDataSourceState struct {
}

func (ApplicationReferenceDataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationReferenceDataSourceState)(nil)).Elem()
}

type applicationReferenceDataSourceArgs struct {
	ApplicationName     string                                            `pulumi:"applicationName"`
	ReferenceDataSource ApplicationReferenceDataSourceReferenceDataSource `pulumi:"referenceDataSource"`
}

// The set of arguments for constructing a ApplicationReferenceDataSource resource.
type ApplicationReferenceDataSourceArgs struct {
	ApplicationName     pulumi.StringInput
	ReferenceDataSource ApplicationReferenceDataSourceReferenceDataSourceInput
}

func (ApplicationReferenceDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationReferenceDataSourceArgs)(nil)).Elem()
}

type ApplicationReferenceDataSourceInput interface {
	pulumi.Input

	ToApplicationReferenceDataSourceOutput() ApplicationReferenceDataSourceOutput
	ToApplicationReferenceDataSourceOutputWithContext(ctx context.Context) ApplicationReferenceDataSourceOutput
}

func (*ApplicationReferenceDataSource) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationReferenceDataSource)(nil)).Elem()
}

func (i *ApplicationReferenceDataSource) ToApplicationReferenceDataSourceOutput() ApplicationReferenceDataSourceOutput {
	return i.ToApplicationReferenceDataSourceOutputWithContext(context.Background())
}

func (i *ApplicationReferenceDataSource) ToApplicationReferenceDataSourceOutputWithContext(ctx context.Context) ApplicationReferenceDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationReferenceDataSourceOutput)
}

type ApplicationReferenceDataSourceOutput struct{ *pulumi.OutputState }

func (ApplicationReferenceDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationReferenceDataSource)(nil)).Elem()
}

func (o ApplicationReferenceDataSourceOutput) ToApplicationReferenceDataSourceOutput() ApplicationReferenceDataSourceOutput {
	return o
}

func (o ApplicationReferenceDataSourceOutput) ToApplicationReferenceDataSourceOutputWithContext(ctx context.Context) ApplicationReferenceDataSourceOutput {
	return o
}

func (o ApplicationReferenceDataSourceOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationReferenceDataSource) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o ApplicationReferenceDataSourceOutput) ReferenceDataSource() ApplicationReferenceDataSourceReferenceDataSourceOutput {
	return o.ApplyT(func(v *ApplicationReferenceDataSource) ApplicationReferenceDataSourceReferenceDataSourceOutput {
		return v.ReferenceDataSource
	}).(ApplicationReferenceDataSourceReferenceDataSourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationReferenceDataSourceInput)(nil)).Elem(), &ApplicationReferenceDataSource{})
	pulumi.RegisterOutputType(ApplicationReferenceDataSourceOutput{})
}
