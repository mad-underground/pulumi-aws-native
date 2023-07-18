// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DAX::Cluster
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("aws-native:dax:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	Id string `pulumi:"id"`
}

type LookupClusterResult struct {
	Arn                         *string     `pulumi:"arn"`
	AvailabilityZones           []string    `pulumi:"availabilityZones"`
	ClusterDiscoveryEndpoint    *string     `pulumi:"clusterDiscoveryEndpoint"`
	ClusterDiscoveryEndpointURL *string     `pulumi:"clusterDiscoveryEndpointURL"`
	Description                 *string     `pulumi:"description"`
	Id                          *string     `pulumi:"id"`
	NotificationTopicARN        *string     `pulumi:"notificationTopicARN"`
	ParameterGroupName          *string     `pulumi:"parameterGroupName"`
	PreferredMaintenanceWindow  *string     `pulumi:"preferredMaintenanceWindow"`
	ReplicationFactor           *int        `pulumi:"replicationFactor"`
	SecurityGroupIds            []string    `pulumi:"securityGroupIds"`
	Tags                        interface{} `pulumi:"tags"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) ClusterDiscoveryEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterDiscoveryEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) ClusterDiscoveryEndpointURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ClusterDiscoveryEndpointURL }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) NotificationTopicARN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.NotificationTopicARN }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) ParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ParameterGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) ReplicationFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.ReplicationFactor }).(pulumi.IntPtrOutput)
}

func (o LookupClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupClusterResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
