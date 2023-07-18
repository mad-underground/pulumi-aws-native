// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DocDB::DBInstance
func LookupDBInstance(ctx *pulumi.Context, args *LookupDBInstanceArgs, opts ...pulumi.InvokeOption) (*LookupDBInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDBInstanceResult
	err := ctx.Invoke("aws-native:docdb:getDBInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBInstanceArgs struct {
	Id string `pulumi:"id"`
}

type LookupDBInstanceResult struct {
	AutoMinorVersionUpgrade    *bool           `pulumi:"autoMinorVersionUpgrade"`
	DBInstanceClass            *string         `pulumi:"dBInstanceClass"`
	EnablePerformanceInsights  *bool           `pulumi:"enablePerformanceInsights"`
	Endpoint                   *string         `pulumi:"endpoint"`
	Id                         *string         `pulumi:"id"`
	Port                       *string         `pulumi:"port"`
	PreferredMaintenanceWindow *string         `pulumi:"preferredMaintenanceWindow"`
	Tags                       []DBInstanceTag `pulumi:"tags"`
}

func LookupDBInstanceOutput(ctx *pulumi.Context, args LookupDBInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupDBInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBInstanceResult, error) {
			args := v.(LookupDBInstanceArgs)
			r, err := LookupDBInstance(ctx, &args, opts...)
			var s LookupDBInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBInstanceResultOutput)
}

type LookupDBInstanceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDBInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBInstanceArgs)(nil)).Elem()
}

type LookupDBInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupDBInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBInstanceResult)(nil)).Elem()
}

func (o LookupDBInstanceResultOutput) ToLookupDBInstanceResultOutput() LookupDBInstanceResultOutput {
	return o
}

func (o LookupDBInstanceResultOutput) ToLookupDBInstanceResultOutputWithContext(ctx context.Context) LookupDBInstanceResultOutput {
	return o
}

func (o LookupDBInstanceResultOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *bool { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupDBInstanceResultOutput) DBInstanceClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *string { return v.DBInstanceClass }).(pulumi.StringPtrOutput)
}

func (o LookupDBInstanceResultOutput) EnablePerformanceInsights() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *bool { return v.EnablePerformanceInsights }).(pulumi.BoolPtrOutput)
}

func (o LookupDBInstanceResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o LookupDBInstanceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDBInstanceResultOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o LookupDBInstanceResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o LookupDBInstanceResultOutput) Tags() DBInstanceTagArrayOutput {
	return o.ApplyT(func(v LookupDBInstanceResult) []DBInstanceTag { return v.Tags }).(DBInstanceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBInstanceResultOutput{})
}
