// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.
type FlowLog struct {
	pulumi.CustomResourceState

	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole pulumi.StringPtrOutput `pulumi:"deliverCrossAccountRole"`
	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	DeliverLogsPermissionArn pulumi.StringPtrOutput                `pulumi:"deliverLogsPermissionArn"`
	DestinationOptions       DestinationOptionsPropertiesPtrOutput `pulumi:"destinationOptions"`
	// Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
	LogDestination pulumi.StringPtrOutput `pulumi:"logDestination"`
	// Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
	LogDestinationType FlowLogLogDestinationTypePtrOutput `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringPtrOutput `pulumi:"logFormat"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	LogGroupName pulumi.StringPtrOutput `pulumi:"logGroupName"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
	MaxAggregationInterval pulumi.IntPtrOutput `pulumi:"maxAggregationInterval"`
	// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
	ResourceType FlowLogResourceTypeOutput `pulumi:"resourceType"`
	// The tags to apply to the flow logs.
	Tags FlowLogTagArrayOutput `pulumi:"tags"`
	// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType FlowLogTrafficTypePtrOutput `pulumi:"trafficType"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"deliverCrossAccountRole",
		"deliverLogsPermissionArn",
		"destinationOptions",
		"logDestination",
		"logDestinationType",
		"logFormat",
		"logGroupName",
		"maxAggregationInterval",
		"resourceId",
		"resourceType",
		"trafficType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlowLog
	err := ctx.RegisterResource("aws-native:ec2:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("aws-native:ec2:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
}

type FlowLogState struct {
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole *string `pulumi:"deliverCrossAccountRole"`
	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	DeliverLogsPermissionArn *string                       `pulumi:"deliverLogsPermissionArn"`
	DestinationOptions       *DestinationOptionsProperties `pulumi:"destinationOptions"`
	// Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
	LogDestination *string `pulumi:"logDestination"`
	// Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
	LogDestinationType *FlowLogLogDestinationType `pulumi:"logDestinationType"`
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat *string `pulumi:"logFormat"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	LogGroupName *string `pulumi:"logGroupName"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
	MaxAggregationInterval *int `pulumi:"maxAggregationInterval"`
	// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
	ResourceId string `pulumi:"resourceId"`
	// The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
	ResourceType FlowLogResourceType `pulumi:"resourceType"`
	// The tags to apply to the flow logs.
	Tags []FlowLogTag `pulumi:"tags"`
	// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType *FlowLogTrafficType `pulumi:"trafficType"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole pulumi.StringPtrInput
	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	DeliverLogsPermissionArn pulumi.StringPtrInput
	DestinationOptions       DestinationOptionsPropertiesPtrInput
	// Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
	LogDestination pulumi.StringPtrInput
	// Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
	LogDestinationType FlowLogLogDestinationTypePtrInput
	// The fields to include in the flow log record, in the order in which they should appear.
	LogFormat pulumi.StringPtrInput
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
	LogGroupName pulumi.StringPtrInput
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
	MaxAggregationInterval pulumi.IntPtrInput
	// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
	ResourceId pulumi.StringInput
	// The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
	ResourceType FlowLogResourceTypeInput
	// The tags to apply to the flow logs.
	Tags FlowLogTagArrayInput
	// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType FlowLogTrafficTypePtrInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}

type FlowLogInput interface {
	pulumi.Input

	ToFlowLogOutput() FlowLogOutput
	ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput
}

func (*FlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (i *FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i *FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

type FlowLogOutput struct{ *pulumi.OutputState }

func (FlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLog)(nil)).Elem()
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
func (o FlowLogOutput) DeliverCrossAccountRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.DeliverCrossAccountRole }).(pulumi.StringPtrOutput)
}

// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
func (o FlowLogOutput) DeliverLogsPermissionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.DeliverLogsPermissionArn }).(pulumi.StringPtrOutput)
}

func (o FlowLogOutput) DestinationOptions() DestinationOptionsPropertiesPtrOutput {
	return o.ApplyT(func(v *FlowLog) DestinationOptionsPropertiesPtrOutput { return v.DestinationOptions }).(DestinationOptionsPropertiesPtrOutput)
}

// Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
func (o FlowLogOutput) LogDestination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.LogDestination }).(pulumi.StringPtrOutput)
}

// Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
func (o FlowLogOutput) LogDestinationType() FlowLogLogDestinationTypePtrOutput {
	return o.ApplyT(func(v *FlowLog) FlowLogLogDestinationTypePtrOutput { return v.LogDestinationType }).(FlowLogLogDestinationTypePtrOutput)
}

// The fields to include in the flow log record, in the order in which they should appear.
func (o FlowLogOutput) LogFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.LogFormat }).(pulumi.StringPtrOutput)
}

// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
func (o FlowLogOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringPtrOutput { return v.LogGroupName }).(pulumi.StringPtrOutput)
}

// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
func (o FlowLogOutput) MaxAggregationInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.IntPtrOutput { return v.MaxAggregationInterval }).(pulumi.IntPtrOutput)
}

// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
func (o FlowLogOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowLog) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
func (o FlowLogOutput) ResourceType() FlowLogResourceTypeOutput {
	return o.ApplyT(func(v *FlowLog) FlowLogResourceTypeOutput { return v.ResourceType }).(FlowLogResourceTypeOutput)
}

// The tags to apply to the flow logs.
func (o FlowLogOutput) Tags() FlowLogTagArrayOutput {
	return o.ApplyT(func(v *FlowLog) FlowLogTagArrayOutput { return v.Tags }).(FlowLogTagArrayOutput)
}

// The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
func (o FlowLogOutput) TrafficType() FlowLogTrafficTypePtrOutput {
	return o.ApplyT(func(v *FlowLog) FlowLogTrafficTypePtrOutput { return v.TrafficType }).(FlowLogTrafficTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowLogInput)(nil)).Elem(), &FlowLog{})
	pulumi.RegisterOutputType(FlowLogOutput{})
}
