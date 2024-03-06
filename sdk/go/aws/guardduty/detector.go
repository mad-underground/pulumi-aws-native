// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::Detector
type Detector struct {
	pulumi.CustomResourceState

	AwsId                      pulumi.StringOutput                          `pulumi:"awsId"`
	DataSources                DetectorCfnDataSourceConfigurationsPtrOutput `pulumi:"dataSources"`
	Enable                     pulumi.BoolOutput                            `pulumi:"enable"`
	Features                   DetectorCfnFeatureConfigurationArrayOutput   `pulumi:"features"`
	FindingPublishingFrequency pulumi.StringPtrOutput                       `pulumi:"findingPublishingFrequency"`
	Tags                       aws.TagArrayOutput                           `pulumi:"tags"`
}

// NewDetector registers a new resource with the given unique name, arguments, and options.
func NewDetector(ctx *pulumi.Context,
	name string, args *DetectorArgs, opts ...pulumi.ResourceOption) (*Detector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enable == nil {
		return nil, errors.New("invalid value for required argument 'Enable'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Detector
	err := ctx.RegisterResource("aws-native:guardduty:Detector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDetector gets an existing Detector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDetector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DetectorState, opts ...pulumi.ResourceOption) (*Detector, error) {
	var resource Detector
	err := ctx.ReadResource("aws-native:guardduty:Detector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Detector resources.
type detectorState struct {
}

type DetectorState struct {
}

func (DetectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*detectorState)(nil)).Elem()
}

type detectorArgs struct {
	DataSources                *DetectorCfnDataSourceConfigurations `pulumi:"dataSources"`
	Enable                     bool                                 `pulumi:"enable"`
	Features                   []DetectorCfnFeatureConfiguration    `pulumi:"features"`
	FindingPublishingFrequency *string                              `pulumi:"findingPublishingFrequency"`
	Tags                       []aws.Tag                            `pulumi:"tags"`
}

// The set of arguments for constructing a Detector resource.
type DetectorArgs struct {
	DataSources                DetectorCfnDataSourceConfigurationsPtrInput
	Enable                     pulumi.BoolInput
	Features                   DetectorCfnFeatureConfigurationArrayInput
	FindingPublishingFrequency pulumi.StringPtrInput
	Tags                       aws.TagArrayInput
}

func (DetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*detectorArgs)(nil)).Elem()
}

type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(ctx context.Context) DetectorOutput
}

func (*Detector) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (i *Detector) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i *Detector) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

func (o DetectorOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o DetectorOutput) DataSources() DetectorCfnDataSourceConfigurationsPtrOutput {
	return o.ApplyT(func(v *Detector) DetectorCfnDataSourceConfigurationsPtrOutput { return v.DataSources }).(DetectorCfnDataSourceConfigurationsPtrOutput)
}

func (o DetectorOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Detector) pulumi.BoolOutput { return v.Enable }).(pulumi.BoolOutput)
}

func (o DetectorOutput) Features() DetectorCfnFeatureConfigurationArrayOutput {
	return o.ApplyT(func(v *Detector) DetectorCfnFeatureConfigurationArrayOutput { return v.Features }).(DetectorCfnFeatureConfigurationArrayOutput)
}

func (o DetectorOutput) FindingPublishingFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) pulumi.StringPtrOutput { return v.FindingPublishingFrequency }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Detector) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorInput)(nil)).Elem(), &Detector{})
	pulumi.RegisterOutputType(DetectorOutput{})
}
