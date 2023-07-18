// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBSubnetGroup
func LookupDBSubnetGroup(ctx *pulumi.Context, args *LookupDBSubnetGroupArgs, opts ...pulumi.InvokeOption) (*LookupDBSubnetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDBSubnetGroupResult
	err := ctx.Invoke("aws-native:neptune:getDBSubnetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBSubnetGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupDBSubnetGroupResult struct {
	DBSubnetGroupDescription *string            `pulumi:"dBSubnetGroupDescription"`
	Id                       *string            `pulumi:"id"`
	SubnetIds                []string           `pulumi:"subnetIds"`
	Tags                     []DBSubnetGroupTag `pulumi:"tags"`
}

func LookupDBSubnetGroupOutput(ctx *pulumi.Context, args LookupDBSubnetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDBSubnetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBSubnetGroupResult, error) {
			args := v.(LookupDBSubnetGroupArgs)
			r, err := LookupDBSubnetGroup(ctx, &args, opts...)
			var s LookupDBSubnetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBSubnetGroupResultOutput)
}

type LookupDBSubnetGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDBSubnetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBSubnetGroupArgs)(nil)).Elem()
}

type LookupDBSubnetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDBSubnetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBSubnetGroupResult)(nil)).Elem()
}

func (o LookupDBSubnetGroupResultOutput) ToLookupDBSubnetGroupResultOutput() LookupDBSubnetGroupResultOutput {
	return o
}

func (o LookupDBSubnetGroupResultOutput) ToLookupDBSubnetGroupResultOutputWithContext(ctx context.Context) LookupDBSubnetGroupResultOutput {
	return o
}

func (o LookupDBSubnetGroupResultOutput) DBSubnetGroupDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSubnetGroupResult) *string { return v.DBSubnetGroupDescription }).(pulumi.StringPtrOutput)
}

func (o LookupDBSubnetGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSubnetGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDBSubnetGroupResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDBSubnetGroupResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o LookupDBSubnetGroupResultOutput) Tags() DBSubnetGroupTagArrayOutput {
	return o.ApplyT(func(v LookupDBSubnetGroupResult) []DBSubnetGroupTag { return v.Tags }).(DBSubnetGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBSubnetGroupResultOutput{})
}
