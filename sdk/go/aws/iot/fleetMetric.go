// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An aggregated metric of certain devices in your fleet
type FleetMetric struct {
	pulumi.CustomResourceState

	// The aggregation field to perform aggregation and metric emission
	AggregationField pulumi.StringPtrOutput              `pulumi:"aggregationField"`
	AggregationType  FleetMetricAggregationTypePtrOutput `pulumi:"aggregationType"`
	// The creation date of a fleet metric
	CreationDate pulumi.Float64Output `pulumi:"creationDate"`
	// The description of a fleet metric
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The index name of a fleet metric
	IndexName pulumi.StringPtrOutput `pulumi:"indexName"`
	// The last modified date of a fleet metric
	LastModifiedDate pulumi.Float64Output `pulumi:"lastModifiedDate"`
	// The Amazon Resource Number (ARN) of a fleet metric metric
	MetricArn pulumi.StringOutput `pulumi:"metricArn"`
	// The name of the fleet metric
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// The period of metric emission in seconds
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The Fleet Indexing query used by a fleet metric
	QueryString pulumi.StringPtrOutput `pulumi:"queryString"`
	// The version of a Fleet Indexing query used by a fleet metric
	QueryVersion pulumi.StringPtrOutput `pulumi:"queryVersion"`
	// An array of key-value pairs to apply to this resource
	Tags FleetMetricTagArrayOutput `pulumi:"tags"`
	// The unit of data points emitted by a fleet metric
	Unit pulumi.StringPtrOutput `pulumi:"unit"`
	// The version of a fleet metric
	Version pulumi.Float64Output `pulumi:"version"`
}

// NewFleetMetric registers a new resource with the given unique name, arguments, and options.
func NewFleetMetric(ctx *pulumi.Context,
	name string, args *FleetMetricArgs, opts ...pulumi.ResourceOption) (*FleetMetric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	var resource FleetMetric
	err := ctx.RegisterResource("aws-native:iot:FleetMetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleetMetric gets an existing FleetMetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleetMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetMetricState, opts ...pulumi.ResourceOption) (*FleetMetric, error) {
	var resource FleetMetric
	err := ctx.ReadResource("aws-native:iot:FleetMetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FleetMetric resources.
type fleetMetricState struct {
}

type FleetMetricState struct {
}

func (FleetMetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetMetricState)(nil)).Elem()
}

type fleetMetricArgs struct {
	// The aggregation field to perform aggregation and metric emission
	AggregationField *string                     `pulumi:"aggregationField"`
	AggregationType  *FleetMetricAggregationType `pulumi:"aggregationType"`
	// The description of a fleet metric
	Description *string `pulumi:"description"`
	// The index name of a fleet metric
	IndexName *string `pulumi:"indexName"`
	// The name of the fleet metric
	MetricName string `pulumi:"metricName"`
	// The period of metric emission in seconds
	Period *int `pulumi:"period"`
	// The Fleet Indexing query used by a fleet metric
	QueryString *string `pulumi:"queryString"`
	// The version of a Fleet Indexing query used by a fleet metric
	QueryVersion *string `pulumi:"queryVersion"`
	// An array of key-value pairs to apply to this resource
	Tags []FleetMetricTag `pulumi:"tags"`
	// The unit of data points emitted by a fleet metric
	Unit *string `pulumi:"unit"`
}

// The set of arguments for constructing a FleetMetric resource.
type FleetMetricArgs struct {
	// The aggregation field to perform aggregation and metric emission
	AggregationField pulumi.StringPtrInput
	AggregationType  FleetMetricAggregationTypePtrInput
	// The description of a fleet metric
	Description pulumi.StringPtrInput
	// The index name of a fleet metric
	IndexName pulumi.StringPtrInput
	// The name of the fleet metric
	MetricName pulumi.StringInput
	// The period of metric emission in seconds
	Period pulumi.IntPtrInput
	// The Fleet Indexing query used by a fleet metric
	QueryString pulumi.StringPtrInput
	// The version of a Fleet Indexing query used by a fleet metric
	QueryVersion pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource
	Tags FleetMetricTagArrayInput
	// The unit of data points emitted by a fleet metric
	Unit pulumi.StringPtrInput
}

func (FleetMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetMetricArgs)(nil)).Elem()
}

type FleetMetricInput interface {
	pulumi.Input

	ToFleetMetricOutput() FleetMetricOutput
	ToFleetMetricOutputWithContext(ctx context.Context) FleetMetricOutput
}

func (*FleetMetric) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetMetric)(nil)).Elem()
}

func (i *FleetMetric) ToFleetMetricOutput() FleetMetricOutput {
	return i.ToFleetMetricOutputWithContext(context.Background())
}

func (i *FleetMetric) ToFleetMetricOutputWithContext(ctx context.Context) FleetMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetMetricOutput)
}

type FleetMetricOutput struct{ *pulumi.OutputState }

func (FleetMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetMetric)(nil)).Elem()
}

func (o FleetMetricOutput) ToFleetMetricOutput() FleetMetricOutput {
	return o
}

func (o FleetMetricOutput) ToFleetMetricOutputWithContext(ctx context.Context) FleetMetricOutput {
	return o
}

// The aggregation field to perform aggregation and metric emission
func (o FleetMetricOutput) AggregationField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringPtrOutput { return v.AggregationField }).(pulumi.StringPtrOutput)
}

func (o FleetMetricOutput) AggregationType() FleetMetricAggregationTypePtrOutput {
	return o.ApplyT(func(v *FleetMetric) FleetMetricAggregationTypePtrOutput { return v.AggregationType }).(FleetMetricAggregationTypePtrOutput)
}

// The creation date of a fleet metric
func (o FleetMetricOutput) CreationDate() pulumi.Float64Output {
	return o.ApplyT(func(v *FleetMetric) pulumi.Float64Output { return v.CreationDate }).(pulumi.Float64Output)
}

// The description of a fleet metric
func (o FleetMetricOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The index name of a fleet metric
func (o FleetMetricOutput) IndexName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringPtrOutput { return v.IndexName }).(pulumi.StringPtrOutput)
}

// The last modified date of a fleet metric
func (o FleetMetricOutput) LastModifiedDate() pulumi.Float64Output {
	return o.ApplyT(func(v *FleetMetric) pulumi.Float64Output { return v.LastModifiedDate }).(pulumi.Float64Output)
}

// The Amazon Resource Number (ARN) of a fleet metric metric
func (o FleetMetricOutput) MetricArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringOutput { return v.MetricArn }).(pulumi.StringOutput)
}

// The name of the fleet metric
func (o FleetMetricOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

// The period of metric emission in seconds
func (o FleetMetricOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The Fleet Indexing query used by a fleet metric
func (o FleetMetricOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringPtrOutput { return v.QueryString }).(pulumi.StringPtrOutput)
}

// The version of a Fleet Indexing query used by a fleet metric
func (o FleetMetricOutput) QueryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringPtrOutput { return v.QueryVersion }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource
func (o FleetMetricOutput) Tags() FleetMetricTagArrayOutput {
	return o.ApplyT(func(v *FleetMetric) FleetMetricTagArrayOutput { return v.Tags }).(FleetMetricTagArrayOutput)
}

// The unit of data points emitted by a fleet metric
func (o FleetMetricOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetMetric) pulumi.StringPtrOutput { return v.Unit }).(pulumi.StringPtrOutput)
}

// The version of a fleet metric
func (o FleetMetricOutput) Version() pulumi.Float64Output {
	return o.ApplyT(func(v *FleetMetric) pulumi.Float64Output { return v.Version }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FleetMetricInput)(nil)).Elem(), &FleetMetric{})
	pulumi.RegisterOutputType(FleetMetricOutput{})
}
