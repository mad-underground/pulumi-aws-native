// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::NetworkManager::GlobalNetwork type specifies a global network of the user's account
func LookupGlobalNetwork(ctx *pulumi.Context, args *LookupGlobalNetworkArgs, opts ...pulumi.InvokeOption) (*LookupGlobalNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGlobalNetworkResult
	err := ctx.Invoke("aws-native:networkmanager:getGlobalNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalNetworkArgs struct {
	// The ID of the global network.
	Id string `pulumi:"id"`
}

type LookupGlobalNetworkResult struct {
	// The Amazon Resource Name (ARN) of the global network.
	Arn *string `pulumi:"arn"`
	// The date and time that the global network was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the global network.
	Description *string `pulumi:"description"`
	// The ID of the global network.
	Id *string `pulumi:"id"`
	// The state of the global network.
	State *string `pulumi:"state"`
	// The tags for the global network.
	Tags []GlobalNetworkTag `pulumi:"tags"`
}

func LookupGlobalNetworkOutput(ctx *pulumi.Context, args LookupGlobalNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalNetworkResult, error) {
			args := v.(LookupGlobalNetworkArgs)
			r, err := LookupGlobalNetwork(ctx, &args, opts...)
			var s LookupGlobalNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalNetworkResultOutput)
}

type LookupGlobalNetworkOutputArgs struct {
	// The ID of the global network.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGlobalNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalNetworkArgs)(nil)).Elem()
}

type LookupGlobalNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalNetworkResult)(nil)).Elem()
}

func (o LookupGlobalNetworkResultOutput) ToLookupGlobalNetworkResultOutput() LookupGlobalNetworkResultOutput {
	return o
}

func (o LookupGlobalNetworkResultOutput) ToLookupGlobalNetworkResultOutputWithContext(ctx context.Context) LookupGlobalNetworkResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the global network.
func (o LookupGlobalNetworkResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalNetworkResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The date and time that the global network was created.
func (o LookupGlobalNetworkResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalNetworkResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the global network.
func (o LookupGlobalNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the global network.
func (o LookupGlobalNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The state of the global network.
func (o LookupGlobalNetworkResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalNetworkResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The tags for the global network.
func (o LookupGlobalNetworkResultOutput) Tags() GlobalNetworkTagArrayOutput {
	return o.ApplyT(func(v LookupGlobalNetworkResult) []GlobalNetworkTag { return v.Tags }).(GlobalNetworkTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalNetworkResultOutput{})
}
