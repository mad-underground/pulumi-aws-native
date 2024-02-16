// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Backup::BackupVault
func LookupBackupVault(ctx *pulumi.Context, args *LookupBackupVaultArgs, opts ...pulumi.InvokeOption) (*LookupBackupVaultResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupVaultResult
	err := ctx.Invoke("aws-native:backup:getBackupVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupVaultArgs struct {
	BackupVaultName string `pulumi:"backupVaultName"`
}

type LookupBackupVaultResult struct {
	AccessPolicy      interface{}                        `pulumi:"accessPolicy"`
	BackupVaultArn    *string                            `pulumi:"backupVaultArn"`
	BackupVaultTags   map[string]string                  `pulumi:"backupVaultTags"`
	LockConfiguration *BackupVaultLockConfigurationType  `pulumi:"lockConfiguration"`
	Notifications     *BackupVaultNotificationObjectType `pulumi:"notifications"`
}

func LookupBackupVaultOutput(ctx *pulumi.Context, args LookupBackupVaultOutputArgs, opts ...pulumi.InvokeOption) LookupBackupVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupVaultResult, error) {
			args := v.(LookupBackupVaultArgs)
			r, err := LookupBackupVault(ctx, &args, opts...)
			var s LookupBackupVaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupVaultResultOutput)
}

type LookupBackupVaultOutputArgs struct {
	BackupVaultName pulumi.StringInput `pulumi:"backupVaultName"`
}

func (LookupBackupVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultArgs)(nil)).Elem()
}

type LookupBackupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupBackupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultResult)(nil)).Elem()
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutput() LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutputWithContext(ctx context.Context) LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) AccessPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) interface{} { return v.AccessPolicy }).(pulumi.AnyOutput)
}

func (o LookupBackupVaultResultOutput) BackupVaultArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) *string { return v.BackupVaultArn }).(pulumi.StringPtrOutput)
}

func (o LookupBackupVaultResultOutput) BackupVaultTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) map[string]string { return v.BackupVaultTags }).(pulumi.StringMapOutput)
}

func (o LookupBackupVaultResultOutput) LockConfiguration() BackupVaultLockConfigurationTypePtrOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) *BackupVaultLockConfigurationType { return v.LockConfiguration }).(BackupVaultLockConfigurationTypePtrOutput)
}

func (o LookupBackupVaultResultOutput) Notifications() BackupVaultNotificationObjectTypePtrOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) *BackupVaultNotificationObjectType { return v.Notifications }).(BackupVaultNotificationObjectTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupVaultResultOutput{})
}
