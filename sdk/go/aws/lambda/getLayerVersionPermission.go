// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::LayerVersionPermission
func LookupLayerVersionPermission(ctx *pulumi.Context, args *LookupLayerVersionPermissionArgs, opts ...pulumi.InvokeOption) (*LookupLayerVersionPermissionResult, error) {
	var rv LookupLayerVersionPermissionResult
	err := ctx.Invoke("aws-native:lambda:getLayerVersionPermission", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLayerVersionPermissionArgs struct {
	Id string `pulumi:"id"`
}

type LookupLayerVersionPermissionResult struct {
	Id *string `pulumi:"id"`
}

func LookupLayerVersionPermissionOutput(ctx *pulumi.Context, args LookupLayerVersionPermissionOutputArgs, opts ...pulumi.InvokeOption) LookupLayerVersionPermissionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLayerVersionPermissionResult, error) {
			args := v.(LookupLayerVersionPermissionArgs)
			r, err := LookupLayerVersionPermission(ctx, &args, opts...)
			return *r, err
		}).(LookupLayerVersionPermissionResultOutput)
}

type LookupLayerVersionPermissionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLayerVersionPermissionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLayerVersionPermissionArgs)(nil)).Elem()
}

type LookupLayerVersionPermissionResultOutput struct{ *pulumi.OutputState }

func (LookupLayerVersionPermissionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLayerVersionPermissionResult)(nil)).Elem()
}

func (o LookupLayerVersionPermissionResultOutput) ToLookupLayerVersionPermissionResultOutput() LookupLayerVersionPermissionResultOutput {
	return o
}

func (o LookupLayerVersionPermissionResultOutput) ToLookupLayerVersionPermissionResultOutputWithContext(ctx context.Context) LookupLayerVersionPermissionResultOutput {
	return o
}

func (o LookupLayerVersionPermissionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLayerVersionPermissionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLayerVersionPermissionResultOutput{})
}
