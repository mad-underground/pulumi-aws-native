// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::NetworkInterfacePermission
func LookupNetworkInterfacePermission(ctx *pulumi.Context, args *LookupNetworkInterfacePermissionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfacePermissionResult, error) {
	var rv LookupNetworkInterfacePermissionResult
	err := ctx.Invoke("aws-native:ec2:getNetworkInterfacePermission", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfacePermissionArgs struct {
	Id string `pulumi:"id"`
}

type LookupNetworkInterfacePermissionResult struct {
	Id *string `pulumi:"id"`
}

func LookupNetworkInterfacePermissionOutput(ctx *pulumi.Context, args LookupNetworkInterfacePermissionOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfacePermissionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfacePermissionResult, error) {
			args := v.(LookupNetworkInterfacePermissionArgs)
			r, err := LookupNetworkInterfacePermission(ctx, &args, opts...)
			return *r, err
		}).(LookupNetworkInterfacePermissionResultOutput)
}

type LookupNetworkInterfacePermissionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupNetworkInterfacePermissionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfacePermissionArgs)(nil)).Elem()
}

type LookupNetworkInterfacePermissionResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfacePermissionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfacePermissionResult)(nil)).Elem()
}

func (o LookupNetworkInterfacePermissionResultOutput) ToLookupNetworkInterfacePermissionResultOutput() LookupNetworkInterfacePermissionResultOutput {
	return o
}

func (o LookupNetworkInterfacePermissionResultOutput) ToLookupNetworkInterfacePermissionResultOutputWithContext(ctx context.Context) LookupNetworkInterfacePermissionResultOutput {
	return o
}

func (o LookupNetworkInterfacePermissionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfacePermissionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfacePermissionResultOutput{})
}
