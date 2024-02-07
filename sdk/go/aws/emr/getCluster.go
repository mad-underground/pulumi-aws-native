// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EMR::Cluster
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("aws-native:emr:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	Id string `pulumi:"id"`
}

type LookupClusterResult struct {
	AutoTerminationPolicy *ClusterAutoTerminationPolicy  `pulumi:"autoTerminationPolicy"`
	Id                    *string                        `pulumi:"id"`
	Instances             *ClusterJobFlowInstancesConfig `pulumi:"instances"`
	ManagedScalingPolicy  *ClusterManagedScalingPolicy   `pulumi:"managedScalingPolicy"`
	MasterPublicDns       *string                        `pulumi:"masterPublicDns"`
	StepConcurrencyLevel  *int                           `pulumi:"stepConcurrencyLevel"`
	Tags                  []ClusterTag                   `pulumi:"tags"`
	VisibleToAllUsers     *bool                          `pulumi:"visibleToAllUsers"`
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

func (o LookupClusterResultOutput) AutoTerminationPolicy() ClusterAutoTerminationPolicyPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterAutoTerminationPolicy { return v.AutoTerminationPolicy }).(ClusterAutoTerminationPolicyPtrOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Instances() ClusterJobFlowInstancesConfigPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterJobFlowInstancesConfig { return v.Instances }).(ClusterJobFlowInstancesConfigPtrOutput)
}

func (o LookupClusterResultOutput) ManagedScalingPolicy() ClusterManagedScalingPolicyPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterManagedScalingPolicy { return v.ManagedScalingPolicy }).(ClusterManagedScalingPolicyPtrOutput)
}

func (o LookupClusterResultOutput) MasterPublicDns() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.MasterPublicDns }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) StepConcurrencyLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *int { return v.StepConcurrencyLevel }).(pulumi.IntPtrOutput)
}

func (o LookupClusterResultOutput) Tags() ClusterTagArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClusterTag { return v.Tags }).(ClusterTagArrayOutput)
}

func (o LookupClusterResultOutput) VisibleToAllUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *bool { return v.VisibleToAllUsers }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
