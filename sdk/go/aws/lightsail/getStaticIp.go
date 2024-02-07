// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lightsail::StaticIp
func LookupStaticIp(ctx *pulumi.Context, args *LookupStaticIpArgs, opts ...pulumi.InvokeOption) (*LookupStaticIpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticIpResult
	err := ctx.Invoke("aws-native:lightsail:getStaticIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticIpArgs struct {
	// The name of the static IP address.
	StaticIpName string `pulumi:"staticIpName"`
}

type LookupStaticIpResult struct {
	// The instance where the static IP is attached.
	AttachedTo *string `pulumi:"attachedTo"`
	// The static IP address.
	IpAddress *string `pulumi:"ipAddress"`
	// A Boolean value indicating whether the static IP is attached.
	IsAttached  *bool   `pulumi:"isAttached"`
	StaticIpArn *string `pulumi:"staticIpArn"`
}

func LookupStaticIpOutput(ctx *pulumi.Context, args LookupStaticIpOutputArgs, opts ...pulumi.InvokeOption) LookupStaticIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticIpResult, error) {
			args := v.(LookupStaticIpArgs)
			r, err := LookupStaticIp(ctx, &args, opts...)
			var s LookupStaticIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticIpResultOutput)
}

type LookupStaticIpOutputArgs struct {
	// The name of the static IP address.
	StaticIpName pulumi.StringInput `pulumi:"staticIpName"`
}

func (LookupStaticIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticIpArgs)(nil)).Elem()
}

type LookupStaticIpResultOutput struct{ *pulumi.OutputState }

func (LookupStaticIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticIpResult)(nil)).Elem()
}

func (o LookupStaticIpResultOutput) ToLookupStaticIpResultOutput() LookupStaticIpResultOutput {
	return o
}

func (o LookupStaticIpResultOutput) ToLookupStaticIpResultOutputWithContext(ctx context.Context) LookupStaticIpResultOutput {
	return o
}

// The instance where the static IP is attached.
func (o LookupStaticIpResultOutput) AttachedTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticIpResult) *string { return v.AttachedTo }).(pulumi.StringPtrOutput)
}

// The static IP address.
func (o LookupStaticIpResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticIpResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// A Boolean value indicating whether the static IP is attached.
func (o LookupStaticIpResultOutput) IsAttached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStaticIpResult) *bool { return v.IsAttached }).(pulumi.BoolPtrOutput)
}

func (o LookupStaticIpResultOutput) StaticIpArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticIpResult) *string { return v.StaticIpArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticIpResultOutput{})
}
