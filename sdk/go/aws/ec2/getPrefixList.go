// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema of AWS::EC2::PrefixList Type
func LookupPrefixList(ctx *pulumi.Context, args *LookupPrefixListArgs, opts ...pulumi.InvokeOption) (*LookupPrefixListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrefixListResult
	err := ctx.Invoke("aws-native:ec2:getPrefixList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrefixListArgs struct {
	// Id of Prefix List.
	PrefixListId string `pulumi:"prefixListId"`
}

type LookupPrefixListResult struct {
	// Ip Version of Prefix List.
	AddressFamily *PrefixListAddressFamily `pulumi:"addressFamily"`
	// The Amazon Resource Name (ARN) of the Prefix List.
	Arn *string `pulumi:"arn"`
	// Entries of Prefix List.
	Entries []PrefixListEntry `pulumi:"entries"`
	// Max Entries of Prefix List.
	MaxEntries *int `pulumi:"maxEntries"`
	// Owner Id of Prefix List.
	OwnerId *string `pulumi:"ownerId"`
	// Id of Prefix List.
	PrefixListId *string `pulumi:"prefixListId"`
	// Name of Prefix List.
	PrefixListName *string `pulumi:"prefixListName"`
	// Tags for Prefix List
	Tags []PrefixListTag `pulumi:"tags"`
	// Version of Prefix List.
	Version *int `pulumi:"version"`
}

func LookupPrefixListOutput(ctx *pulumi.Context, args LookupPrefixListOutputArgs, opts ...pulumi.InvokeOption) LookupPrefixListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrefixListResult, error) {
			args := v.(LookupPrefixListArgs)
			r, err := LookupPrefixList(ctx, &args, opts...)
			var s LookupPrefixListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrefixListResultOutput)
}

type LookupPrefixListOutputArgs struct {
	// Id of Prefix List.
	PrefixListId pulumi.StringInput `pulumi:"prefixListId"`
}

func (LookupPrefixListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrefixListArgs)(nil)).Elem()
}

type LookupPrefixListResultOutput struct{ *pulumi.OutputState }

func (LookupPrefixListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrefixListResult)(nil)).Elem()
}

func (o LookupPrefixListResultOutput) ToLookupPrefixListResultOutput() LookupPrefixListResultOutput {
	return o
}

func (o LookupPrefixListResultOutput) ToLookupPrefixListResultOutputWithContext(ctx context.Context) LookupPrefixListResultOutput {
	return o
}

// Ip Version of Prefix List.
func (o LookupPrefixListResultOutput) AddressFamily() PrefixListAddressFamilyPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *PrefixListAddressFamily { return v.AddressFamily }).(PrefixListAddressFamilyPtrOutput)
}

// The Amazon Resource Name (ARN) of the Prefix List.
func (o LookupPrefixListResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Entries of Prefix List.
func (o LookupPrefixListResultOutput) Entries() PrefixListEntryArrayOutput {
	return o.ApplyT(func(v LookupPrefixListResult) []PrefixListEntry { return v.Entries }).(PrefixListEntryArrayOutput)
}

// Max Entries of Prefix List.
func (o LookupPrefixListResultOutput) MaxEntries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *int { return v.MaxEntries }).(pulumi.IntPtrOutput)
}

// Owner Id of Prefix List.
func (o LookupPrefixListResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// Id of Prefix List.
func (o LookupPrefixListResultOutput) PrefixListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *string { return v.PrefixListId }).(pulumi.StringPtrOutput)
}

// Name of Prefix List.
func (o LookupPrefixListResultOutput) PrefixListName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *string { return v.PrefixListName }).(pulumi.StringPtrOutput)
}

// Tags for Prefix List
func (o LookupPrefixListResultOutput) Tags() PrefixListTagArrayOutput {
	return o.ApplyT(func(v LookupPrefixListResult) []PrefixListTag { return v.Tags }).(PrefixListTagArrayOutput)
}

// Version of Prefix List.
func (o LookupPrefixListResultOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPrefixListResult) *int { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrefixListResultOutput{})
}
