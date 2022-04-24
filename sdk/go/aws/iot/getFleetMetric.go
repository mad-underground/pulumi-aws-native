// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An aggregated metric of certain devices in your fleet
func LookupFleetMetric(ctx *pulumi.Context, args *LookupFleetMetricArgs, opts ...pulumi.InvokeOption) (*LookupFleetMetricResult, error) {
	var rv LookupFleetMetricResult
	err := ctx.Invoke("aws-native:iot:getFleetMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetMetricArgs struct {
	// The name of the fleet metric
	MetricName string `pulumi:"metricName"`
}

type LookupFleetMetricResult struct {
	// The aggregation field to perform aggregation and metric emission
	AggregationField *string                     `pulumi:"aggregationField"`
	AggregationType  *FleetMetricAggregationType `pulumi:"aggregationType"`
	// The creation date of a fleet metric
	CreationDate *float64 `pulumi:"creationDate"`
	// The description of a fleet metric
	Description *string `pulumi:"description"`
	// The index name of a fleet metric
	IndexName *string `pulumi:"indexName"`
	// The last modified date of a fleet metric
	LastModifiedDate *float64 `pulumi:"lastModifiedDate"`
	// The Amazon Resource Number (ARN) of a fleet metric metric
	MetricArn *string `pulumi:"metricArn"`
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
	// The version of a fleet metric
	Version *float64 `pulumi:"version"`
}

func LookupFleetMetricOutput(ctx *pulumi.Context, args LookupFleetMetricOutputArgs, opts ...pulumi.InvokeOption) LookupFleetMetricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetMetricResult, error) {
			args := v.(LookupFleetMetricArgs)
			r, err := LookupFleetMetric(ctx, &args, opts...)
			var s LookupFleetMetricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetMetricResultOutput)
}

type LookupFleetMetricOutputArgs struct {
	// The name of the fleet metric
	MetricName pulumi.StringInput `pulumi:"metricName"`
}

func (LookupFleetMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetMetricArgs)(nil)).Elem()
}

type LookupFleetMetricResultOutput struct{ *pulumi.OutputState }

func (LookupFleetMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetMetricResult)(nil)).Elem()
}

func (o LookupFleetMetricResultOutput) ToLookupFleetMetricResultOutput() LookupFleetMetricResultOutput {
	return o
}

func (o LookupFleetMetricResultOutput) ToLookupFleetMetricResultOutputWithContext(ctx context.Context) LookupFleetMetricResultOutput {
	return o
}

// The aggregation field to perform aggregation and metric emission
func (o LookupFleetMetricResultOutput) AggregationField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.AggregationField }).(pulumi.StringPtrOutput)
}

func (o LookupFleetMetricResultOutput) AggregationType() FleetMetricAggregationTypePtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *FleetMetricAggregationType { return v.AggregationType }).(FleetMetricAggregationTypePtrOutput)
}

// The creation date of a fleet metric
func (o LookupFleetMetricResultOutput) CreationDate() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *float64 { return v.CreationDate }).(pulumi.Float64PtrOutput)
}

// The description of a fleet metric
func (o LookupFleetMetricResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The index name of a fleet metric
func (o LookupFleetMetricResultOutput) IndexName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.IndexName }).(pulumi.StringPtrOutput)
}

// The last modified date of a fleet metric
func (o LookupFleetMetricResultOutput) LastModifiedDate() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *float64 { return v.LastModifiedDate }).(pulumi.Float64PtrOutput)
}

// The Amazon Resource Number (ARN) of a fleet metric metric
func (o LookupFleetMetricResultOutput) MetricArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.MetricArn }).(pulumi.StringPtrOutput)
}

// The period of metric emission in seconds
func (o LookupFleetMetricResultOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *int { return v.Period }).(pulumi.IntPtrOutput)
}

// The Fleet Indexing query used by a fleet metric
func (o LookupFleetMetricResultOutput) QueryString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.QueryString }).(pulumi.StringPtrOutput)
}

// The version of a Fleet Indexing query used by a fleet metric
func (o LookupFleetMetricResultOutput) QueryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.QueryVersion }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource
func (o LookupFleetMetricResultOutput) Tags() FleetMetricTagArrayOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) []FleetMetricTag { return v.Tags }).(FleetMetricTagArrayOutput)
}

// The unit of data points emitted by a fleet metric
func (o LookupFleetMetricResultOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

// The version of a fleet metric
func (o LookupFleetMetricResultOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFleetMetricResult) *float64 { return v.Version }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetMetricResultOutput{})
}
