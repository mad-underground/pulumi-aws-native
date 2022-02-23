// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::OpsWorks::Volume
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("aws-native:opsworks:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	Id string `pulumi:"id"`
}

type LookupVolumeResult struct {
	Id         *string `pulumi:"id"`
	MountPoint *string `pulumi:"mountPoint"`
	Name       *string `pulumi:"name"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			return *r, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) MountPoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.MountPoint }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}