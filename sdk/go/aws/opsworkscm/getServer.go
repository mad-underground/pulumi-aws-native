// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworkscm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::OpsWorksCM::Server
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("aws-native:opsworkscm:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	ServerName string `pulumi:"serverName"`
}

type LookupServerResult struct {
	Arn                        *string     `pulumi:"arn"`
	BackupRetentionCount       *int        `pulumi:"backupRetentionCount"`
	DisableAutomatedBackup     *bool       `pulumi:"disableAutomatedBackup"`
	Endpoint                   *string     `pulumi:"endpoint"`
	Id                         *string     `pulumi:"id"`
	PreferredBackupWindow      *string     `pulumi:"preferredBackupWindow"`
	PreferredMaintenanceWindow *string     `pulumi:"preferredMaintenanceWindow"`
	Tags                       []ServerTag `pulumi:"tags"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerResult, error) {
			args := v.(LookupServerArgs)
			r, err := LookupServer(ctx, &args, opts...)
			return *r, err
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) BackupRetentionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *int { return v.BackupRetentionCount }).(pulumi.IntPtrOutput)
}

func (o LookupServerResultOutput) DisableAutomatedBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *bool { return v.DisableAutomatedBackup }).(pulumi.BoolPtrOutput)
}

func (o LookupServerResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) PreferredBackupWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PreferredBackupWindow }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o LookupServerResultOutput) Tags() ServerTagArrayOutput {
	return o.ApplyT(func(v LookupServerResult) []ServerTag { return v.Tags }).(ServerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
