// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ce

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.
type AnomalyMonitor struct {
	pulumi.CustomResourceState

	// The date when the monitor was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The value for evaluated dimensions.
	DimensionalValueCount pulumi.IntOutput `pulumi:"dimensionalValueCount"`
	// The date when the monitor last evaluated for anomalies.
	LastEvaluatedDate pulumi.StringOutput `pulumi:"lastEvaluatedDate"`
	// The date when the monitor was last updated.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	MonitorArn      pulumi.StringOutput `pulumi:"monitorArn"`
	// The dimensions to evaluate
	MonitorDimension AnomalyMonitorMonitorDimensionPtrOutput `pulumi:"monitorDimension"`
	// The name of the monitor.
	MonitorName          pulumi.StringOutput             `pulumi:"monitorName"`
	MonitorSpecification pulumi.StringPtrOutput          `pulumi:"monitorSpecification"`
	MonitorType          AnomalyMonitorMonitorTypeOutput `pulumi:"monitorType"`
	// Tags to assign to monitor.
	ResourceTags AnomalyMonitorResourceTagArrayOutput `pulumi:"resourceTags"`
}

// NewAnomalyMonitor registers a new resource with the given unique name, arguments, and options.
func NewAnomalyMonitor(ctx *pulumi.Context,
	name string, args *AnomalyMonitorArgs, opts ...pulumi.ResourceOption) (*AnomalyMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.MonitorType == nil {
		return nil, errors.New("invalid value for required argument 'MonitorType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"monitorDimension",
		"monitorSpecification",
		"monitorType",
		"resourceTags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnomalyMonitor
	err := ctx.RegisterResource("aws-native:ce:AnomalyMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnomalyMonitor gets an existing AnomalyMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnomalyMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomalyMonitorState, opts ...pulumi.ResourceOption) (*AnomalyMonitor, error) {
	var resource AnomalyMonitor
	err := ctx.ReadResource("aws-native:ce:AnomalyMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnomalyMonitor resources.
type anomalyMonitorState struct {
}

type AnomalyMonitorState struct {
}

func (AnomalyMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalyMonitorState)(nil)).Elem()
}

type anomalyMonitorArgs struct {
	// The dimensions to evaluate
	MonitorDimension *AnomalyMonitorMonitorDimension `pulumi:"monitorDimension"`
	// The name of the monitor.
	MonitorName          string                    `pulumi:"monitorName"`
	MonitorSpecification *string                   `pulumi:"monitorSpecification"`
	MonitorType          AnomalyMonitorMonitorType `pulumi:"monitorType"`
	// Tags to assign to monitor.
	ResourceTags []AnomalyMonitorResourceTag `pulumi:"resourceTags"`
}

// The set of arguments for constructing a AnomalyMonitor resource.
type AnomalyMonitorArgs struct {
	// The dimensions to evaluate
	MonitorDimension AnomalyMonitorMonitorDimensionPtrInput
	// The name of the monitor.
	MonitorName          pulumi.StringInput
	MonitorSpecification pulumi.StringPtrInput
	MonitorType          AnomalyMonitorMonitorTypeInput
	// Tags to assign to monitor.
	ResourceTags AnomalyMonitorResourceTagArrayInput
}

func (AnomalyMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalyMonitorArgs)(nil)).Elem()
}

type AnomalyMonitorInput interface {
	pulumi.Input

	ToAnomalyMonitorOutput() AnomalyMonitorOutput
	ToAnomalyMonitorOutputWithContext(ctx context.Context) AnomalyMonitorOutput
}

func (*AnomalyMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalyMonitor)(nil)).Elem()
}

func (i *AnomalyMonitor) ToAnomalyMonitorOutput() AnomalyMonitorOutput {
	return i.ToAnomalyMonitorOutputWithContext(context.Background())
}

func (i *AnomalyMonitor) ToAnomalyMonitorOutputWithContext(ctx context.Context) AnomalyMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalyMonitorOutput)
}

type AnomalyMonitorOutput struct{ *pulumi.OutputState }

func (AnomalyMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalyMonitor)(nil)).Elem()
}

func (o AnomalyMonitorOutput) ToAnomalyMonitorOutput() AnomalyMonitorOutput {
	return o
}

func (o AnomalyMonitorOutput) ToAnomalyMonitorOutputWithContext(ctx context.Context) AnomalyMonitorOutput {
	return o
}

// The date when the monitor was created.
func (o AnomalyMonitorOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The value for evaluated dimensions.
func (o AnomalyMonitorOutput) DimensionalValueCount() pulumi.IntOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.IntOutput { return v.DimensionalValueCount }).(pulumi.IntOutput)
}

// The date when the monitor last evaluated for anomalies.
func (o AnomalyMonitorOutput) LastEvaluatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.StringOutput { return v.LastEvaluatedDate }).(pulumi.StringOutput)
}

// The date when the monitor was last updated.
func (o AnomalyMonitorOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

func (o AnomalyMonitorOutput) MonitorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.StringOutput { return v.MonitorArn }).(pulumi.StringOutput)
}

// The dimensions to evaluate
func (o AnomalyMonitorOutput) MonitorDimension() AnomalyMonitorMonitorDimensionPtrOutput {
	return o.ApplyT(func(v *AnomalyMonitor) AnomalyMonitorMonitorDimensionPtrOutput { return v.MonitorDimension }).(AnomalyMonitorMonitorDimensionPtrOutput)
}

// The name of the monitor.
func (o AnomalyMonitorOutput) MonitorName() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.StringOutput { return v.MonitorName }).(pulumi.StringOutput)
}

func (o AnomalyMonitorOutput) MonitorSpecification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalyMonitor) pulumi.StringPtrOutput { return v.MonitorSpecification }).(pulumi.StringPtrOutput)
}

func (o AnomalyMonitorOutput) MonitorType() AnomalyMonitorMonitorTypeOutput {
	return o.ApplyT(func(v *AnomalyMonitor) AnomalyMonitorMonitorTypeOutput { return v.MonitorType }).(AnomalyMonitorMonitorTypeOutput)
}

// Tags to assign to monitor.
func (o AnomalyMonitorOutput) ResourceTags() AnomalyMonitorResourceTagArrayOutput {
	return o.ApplyT(func(v *AnomalyMonitor) AnomalyMonitorResourceTagArrayOutput { return v.ResourceTags }).(AnomalyMonitorResourceTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnomalyMonitorInput)(nil)).Elem(), &AnomalyMonitor{})
	pulumi.RegisterOutputType(AnomalyMonitorOutput{})
}
