// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DMS::ReplicationInstance
func LookupReplicationInstance(ctx *pulumi.Context, args *LookupReplicationInstanceArgs, opts ...pulumi.InvokeOption) (*LookupReplicationInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationInstanceResult
	err := ctx.Invoke("aws-native:dms:getReplicationInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationInstanceArgs struct {
	Id string `pulumi:"id"`
}

type LookupReplicationInstanceResult struct {
	AllocatedStorage                      *int                     `pulumi:"allocatedStorage"`
	AllowMajorVersionUpgrade              *bool                    `pulumi:"allowMajorVersionUpgrade"`
	AutoMinorVersionUpgrade               *bool                    `pulumi:"autoMinorVersionUpgrade"`
	AvailabilityZone                      *string                  `pulumi:"availabilityZone"`
	EngineVersion                         *string                  `pulumi:"engineVersion"`
	Id                                    *string                  `pulumi:"id"`
	MultiAZ                               *bool                    `pulumi:"multiAZ"`
	PreferredMaintenanceWindow            *string                  `pulumi:"preferredMaintenanceWindow"`
	ReplicationInstanceClass              *string                  `pulumi:"replicationInstanceClass"`
	ReplicationInstanceIdentifier         *string                  `pulumi:"replicationInstanceIdentifier"`
	ReplicationInstancePrivateIpAddresses *string                  `pulumi:"replicationInstancePrivateIpAddresses"`
	ReplicationInstancePublicIpAddresses  *string                  `pulumi:"replicationInstancePublicIpAddresses"`
	Tags                                  []ReplicationInstanceTag `pulumi:"tags"`
	VpcSecurityGroupIds                   []string                 `pulumi:"vpcSecurityGroupIds"`
}

func LookupReplicationInstanceOutput(ctx *pulumi.Context, args LookupReplicationInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationInstanceResult, error) {
			args := v.(LookupReplicationInstanceArgs)
			r, err := LookupReplicationInstance(ctx, &args, opts...)
			var s LookupReplicationInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationInstanceResultOutput)
}

type LookupReplicationInstanceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupReplicationInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationInstanceArgs)(nil)).Elem()
}

type LookupReplicationInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationInstanceResult)(nil)).Elem()
}

func (o LookupReplicationInstanceResultOutput) ToLookupReplicationInstanceResultOutput() LookupReplicationInstanceResultOutput {
	return o
}

func (o LookupReplicationInstanceResultOutput) ToLookupReplicationInstanceResultOutputWithContext(ctx context.Context) LookupReplicationInstanceResultOutput {
	return o
}

func (o LookupReplicationInstanceResultOutput) AllocatedStorage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *int { return v.AllocatedStorage }).(pulumi.IntPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) AllowMajorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *bool { return v.AllowMajorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *bool { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) MultiAZ() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *bool { return v.MultiAZ }).(pulumi.BoolPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) ReplicationInstanceClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.ReplicationInstanceClass }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) ReplicationInstanceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.ReplicationInstanceIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) ReplicationInstancePrivateIpAddresses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.ReplicationInstancePrivateIpAddresses }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) ReplicationInstancePublicIpAddresses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) *string { return v.ReplicationInstancePublicIpAddresses }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationInstanceResultOutput) Tags() ReplicationInstanceTagArrayOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) []ReplicationInstanceTag { return v.Tags }).(ReplicationInstanceTagArrayOutput)
}

func (o LookupReplicationInstanceResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupReplicationInstanceResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationInstanceResultOutput{})
}
