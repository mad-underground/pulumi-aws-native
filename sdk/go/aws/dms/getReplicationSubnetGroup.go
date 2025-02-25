// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DMS::ReplicationSubnetGroup
func LookupReplicationSubnetGroup(ctx *pulumi.Context, args *LookupReplicationSubnetGroupArgs, opts ...pulumi.InvokeOption) (*LookupReplicationSubnetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationSubnetGroupResult
	err := ctx.Invoke("aws-native:dms:getReplicationSubnetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationSubnetGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupReplicationSubnetGroupResult struct {
	Id                                *string   `pulumi:"id"`
	ReplicationSubnetGroupDescription *string   `pulumi:"replicationSubnetGroupDescription"`
	SubnetIds                         []string  `pulumi:"subnetIds"`
	Tags                              []aws.Tag `pulumi:"tags"`
}

func LookupReplicationSubnetGroupOutput(ctx *pulumi.Context, args LookupReplicationSubnetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationSubnetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationSubnetGroupResult, error) {
			args := v.(LookupReplicationSubnetGroupArgs)
			r, err := LookupReplicationSubnetGroup(ctx, &args, opts...)
			var s LookupReplicationSubnetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationSubnetGroupResultOutput)
}

type LookupReplicationSubnetGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupReplicationSubnetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationSubnetGroupArgs)(nil)).Elem()
}

type LookupReplicationSubnetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationSubnetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationSubnetGroupResult)(nil)).Elem()
}

func (o LookupReplicationSubnetGroupResultOutput) ToLookupReplicationSubnetGroupResultOutput() LookupReplicationSubnetGroupResultOutput {
	return o
}

func (o LookupReplicationSubnetGroupResultOutput) ToLookupReplicationSubnetGroupResultOutputWithContext(ctx context.Context) LookupReplicationSubnetGroupResultOutput {
	return o
}

func (o LookupReplicationSubnetGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) ReplicationSubnetGroupDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) *string { return v.ReplicationSubnetGroupDescription }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationSubnetGroupResultOutput{})
}
