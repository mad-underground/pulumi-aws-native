// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaTailor::LiveSource Resource Type
type LiveSource struct {
	pulumi.CustomResourceState

	// <p>The ARN of the live source.</p>
	Arn pulumi.StringOutput `pulumi:"arn"`
	// <p>A list of HTTP package configuration parameters for this live source.</p>
	HttpPackageConfigurations LiveSourceHttpPackageConfigurationArrayOutput `pulumi:"httpPackageConfigurations"`
	LiveSourceName            pulumi.StringOutput                           `pulumi:"liveSourceName"`
	SourceLocationName        pulumi.StringOutput                           `pulumi:"sourceLocationName"`
	// The tags to assign to the live source.
	Tags LiveSourceTagArrayOutput `pulumi:"tags"`
}

// NewLiveSource registers a new resource with the given unique name, arguments, and options.
func NewLiveSource(ctx *pulumi.Context,
	name string, args *LiveSourceArgs, opts ...pulumi.ResourceOption) (*LiveSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HttpPackageConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'HttpPackageConfigurations'")
	}
	if args.SourceLocationName == nil {
		return nil, errors.New("invalid value for required argument 'SourceLocationName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"liveSourceName",
		"sourceLocationName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LiveSource
	err := ctx.RegisterResource("aws-native:mediatailor:LiveSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiveSource gets an existing LiveSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiveSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveSourceState, opts ...pulumi.ResourceOption) (*LiveSource, error) {
	var resource LiveSource
	err := ctx.ReadResource("aws-native:mediatailor:LiveSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiveSource resources.
type liveSourceState struct {
}

type LiveSourceState struct {
}

func (LiveSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*liveSourceState)(nil)).Elem()
}

type liveSourceArgs struct {
	// <p>A list of HTTP package configuration parameters for this live source.</p>
	HttpPackageConfigurations []LiveSourceHttpPackageConfiguration `pulumi:"httpPackageConfigurations"`
	LiveSourceName            *string                              `pulumi:"liveSourceName"`
	SourceLocationName        string                               `pulumi:"sourceLocationName"`
	// The tags to assign to the live source.
	Tags []LiveSourceTag `pulumi:"tags"`
}

// The set of arguments for constructing a LiveSource resource.
type LiveSourceArgs struct {
	// <p>A list of HTTP package configuration parameters for this live source.</p>
	HttpPackageConfigurations LiveSourceHttpPackageConfigurationArrayInput
	LiveSourceName            pulumi.StringPtrInput
	SourceLocationName        pulumi.StringInput
	// The tags to assign to the live source.
	Tags LiveSourceTagArrayInput
}

func (LiveSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveSourceArgs)(nil)).Elem()
}

type LiveSourceInput interface {
	pulumi.Input

	ToLiveSourceOutput() LiveSourceOutput
	ToLiveSourceOutputWithContext(ctx context.Context) LiveSourceOutput
}

func (*LiveSource) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveSource)(nil)).Elem()
}

func (i *LiveSource) ToLiveSourceOutput() LiveSourceOutput {
	return i.ToLiveSourceOutputWithContext(context.Background())
}

func (i *LiveSource) ToLiveSourceOutputWithContext(ctx context.Context) LiveSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveSourceOutput)
}

type LiveSourceOutput struct{ *pulumi.OutputState }

func (LiveSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveSource)(nil)).Elem()
}

func (o LiveSourceOutput) ToLiveSourceOutput() LiveSourceOutput {
	return o
}

func (o LiveSourceOutput) ToLiveSourceOutputWithContext(ctx context.Context) LiveSourceOutput {
	return o
}

// <p>The ARN of the live source.</p>
func (o LiveSourceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveSource) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// <p>A list of HTTP package configuration parameters for this live source.</p>
func (o LiveSourceOutput) HttpPackageConfigurations() LiveSourceHttpPackageConfigurationArrayOutput {
	return o.ApplyT(func(v *LiveSource) LiveSourceHttpPackageConfigurationArrayOutput { return v.HttpPackageConfigurations }).(LiveSourceHttpPackageConfigurationArrayOutput)
}

func (o LiveSourceOutput) LiveSourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveSource) pulumi.StringOutput { return v.LiveSourceName }).(pulumi.StringOutput)
}

func (o LiveSourceOutput) SourceLocationName() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveSource) pulumi.StringOutput { return v.SourceLocationName }).(pulumi.StringOutput)
}

// The tags to assign to the live source.
func (o LiveSourceOutput) Tags() LiveSourceTagArrayOutput {
	return o.ApplyT(func(v *LiveSource) LiveSourceTagArrayOutput { return v.Tags }).(LiveSourceTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LiveSourceInput)(nil)).Elem(), &LiveSource{})
	pulumi.RegisterOutputType(LiveSourceOutput{})
}
