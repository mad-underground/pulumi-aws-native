// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Detective::Graph
func LookupGraph(ctx *pulumi.Context, args *LookupGraphArgs, opts ...pulumi.InvokeOption) (*LookupGraphResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGraphResult
	err := ctx.Invoke("aws-native:detective:getGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphArgs struct {
	// The Detective graph ARN
	Arn string `pulumi:"arn"`
}

type LookupGraphResult struct {
	// The Detective graph ARN
	Arn *string `pulumi:"arn"`
	// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
	AutoEnableMembers *bool      `pulumi:"autoEnableMembers"`
	Tags              []GraphTag `pulumi:"tags"`
}

func LookupGraphOutput(ctx *pulumi.Context, args LookupGraphOutputArgs, opts ...pulumi.InvokeOption) LookupGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphResult, error) {
			args := v.(LookupGraphArgs)
			r, err := LookupGraph(ctx, &args, opts...)
			var s LookupGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphResultOutput)
}

type LookupGraphOutputArgs struct {
	// The Detective graph ARN
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphArgs)(nil)).Elem()
}

type LookupGraphResultOutput struct{ *pulumi.OutputState }

func (LookupGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphResult)(nil)).Elem()
}

func (o LookupGraphResultOutput) ToLookupGraphResultOutput() LookupGraphResultOutput {
	return o
}

func (o LookupGraphResultOutput) ToLookupGraphResultOutputWithContext(ctx context.Context) LookupGraphResultOutput {
	return o
}

// The Detective graph ARN
func (o LookupGraphResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.
func (o LookupGraphResultOutput) AutoEnableMembers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGraphResult) *bool { return v.AutoEnableMembers }).(pulumi.BoolPtrOutput)
}

func (o LookupGraphResultOutput) Tags() GraphTagArrayOutput {
	return o.ApplyT(func(v LookupGraphResult) []GraphTag { return v.Tags }).(GraphTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphResultOutput{})
}
