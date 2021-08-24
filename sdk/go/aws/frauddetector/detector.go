// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package frauddetector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html
type Detector struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
	AssociatedModels DetectorModelArrayOutput `pulumi:"associatedModels"`
	CreatedTime      pulumi.StringOutput      `pulumi:"createdTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
	DetectorId        pulumi.StringOutput `pulumi:"detectorId"`
	DetectorVersionId pulumi.StringOutput `pulumi:"detectorVersionId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
	DetectorVersionStatus pulumi.StringPtrOutput `pulumi:"detectorVersionStatus"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
	EventType                DetectorEventTypeOutput `pulumi:"eventType"`
	EventTypeArn             pulumi.StringOutput     `pulumi:"eventTypeArn"`
	EventTypeCreatedTime     pulumi.StringOutput     `pulumi:"eventTypeCreatedTime"`
	EventTypeLastUpdatedTime pulumi.StringOutput     `pulumi:"eventTypeLastUpdatedTime"`
	LastUpdatedTime          pulumi.StringOutput     `pulumi:"lastUpdatedTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
	RuleExecutionMode pulumi.StringPtrOutput `pulumi:"ruleExecutionMode"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
	Rules DetectorRuleArrayOutput `pulumi:"rules"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewDetector registers a new resource with the given unique name, arguments, and options.
func NewDetector(ctx *pulumi.Context,
	name string, args *DetectorArgs, opts ...pulumi.ResourceOption) (*Detector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.EventType == nil {
		return nil, errors.New("invalid value for required argument 'EventType'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource Detector
	err := ctx.RegisterResource("aws-native:frauddetector:Detector", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:frauddetector:Detector", name, id, state, &resource, opts...)
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
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
	AssociatedModels []DetectorModel `pulumi:"associatedModels"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
	DetectorId string `pulumi:"detectorId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
	DetectorVersionStatus *string `pulumi:"detectorVersionStatus"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
	EventType DetectorEventType `pulumi:"eventType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
	RuleExecutionMode *string `pulumi:"ruleExecutionMode"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
	Rules []DetectorRule `pulumi:"rules"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Detector resource.
type DetectorArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-associatedmodels
	AssociatedModels DetectorModelArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorid
	DetectorId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-detectorversionstatus
	DetectorVersionStatus pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-eventtype
	EventType DetectorEventTypeInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-ruleexecutionmode
	RuleExecutionMode pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-rules
	Rules DetectorRuleArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-detector.html#cfn-frauddetector-detector-tags
	Tags aws.TagArrayInput
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
	return reflect.TypeOf((*Detector)(nil))
}

func (i *Detector) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i *Detector) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil))
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DetectorOutput{})
}