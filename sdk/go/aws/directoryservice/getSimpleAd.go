// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DirectoryService::SimpleAD
func LookupSimpleAd(ctx *pulumi.Context, args *LookupSimpleAdArgs, opts ...pulumi.InvokeOption) (*LookupSimpleAdResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSimpleAdResult
	err := ctx.Invoke("aws-native:directoryservice:getSimpleAd", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimpleAdArgs struct {
	// The unique identifier for a directory.
	DirectoryId string `pulumi:"directoryId"`
}

type LookupSimpleAdResult struct {
	// The alias for a directory.
	Alias *string `pulumi:"alias"`
	// The unique identifier for a directory.
	DirectoryId *string `pulumi:"directoryId"`
	// The IP addresses of the DNS servers for the directory, such as [ "172.31.3.154", "172.31.63.203" ].
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	// Whether to enable single sign-on for a Simple Active Directory in AWS.
	EnableSso *bool `pulumi:"enableSso"`
}

func LookupSimpleAdOutput(ctx *pulumi.Context, args LookupSimpleAdOutputArgs, opts ...pulumi.InvokeOption) LookupSimpleAdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimpleAdResult, error) {
			args := v.(LookupSimpleAdArgs)
			r, err := LookupSimpleAd(ctx, &args, opts...)
			var s LookupSimpleAdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimpleAdResultOutput)
}

type LookupSimpleAdOutputArgs struct {
	// The unique identifier for a directory.
	DirectoryId pulumi.StringInput `pulumi:"directoryId"`
}

func (LookupSimpleAdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimpleAdArgs)(nil)).Elem()
}

type LookupSimpleAdResultOutput struct{ *pulumi.OutputState }

func (LookupSimpleAdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimpleAdResult)(nil)).Elem()
}

func (o LookupSimpleAdResultOutput) ToLookupSimpleAdResultOutput() LookupSimpleAdResultOutput {
	return o
}

func (o LookupSimpleAdResultOutput) ToLookupSimpleAdResultOutputWithContext(ctx context.Context) LookupSimpleAdResultOutput {
	return o
}

// The alias for a directory.
func (o LookupSimpleAdResultOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimpleAdResult) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

// The unique identifier for a directory.
func (o LookupSimpleAdResultOutput) DirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimpleAdResult) *string { return v.DirectoryId }).(pulumi.StringPtrOutput)
}

// The IP addresses of the DNS servers for the directory, such as [ "172.31.3.154", "172.31.63.203" ].
func (o LookupSimpleAdResultOutput) DnsIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSimpleAdResult) []string { return v.DnsIpAddresses }).(pulumi.StringArrayOutput)
}

// Whether to enable single sign-on for a Simple Active Directory in AWS.
func (o LookupSimpleAdResultOutput) EnableSso() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSimpleAdResult) *bool { return v.EnableSso }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimpleAdResultOutput{})
}
