// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::StoredQuery
func LookupStoredQuery(ctx *pulumi.Context, args *LookupStoredQueryArgs, opts ...pulumi.InvokeOption) (*LookupStoredQueryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStoredQueryResult
	err := ctx.Invoke("aws-native:configuration:getStoredQuery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStoredQueryArgs struct {
	QueryName string `pulumi:"queryName"`
}

type LookupStoredQueryResult struct {
	QueryArn         *string `pulumi:"queryArn"`
	QueryDescription *string `pulumi:"queryDescription"`
	QueryExpression  *string `pulumi:"queryExpression"`
	QueryId          *string `pulumi:"queryId"`
	// The tags for the stored query.
	Tags []StoredQueryTag `pulumi:"tags"`
}

func LookupStoredQueryOutput(ctx *pulumi.Context, args LookupStoredQueryOutputArgs, opts ...pulumi.InvokeOption) LookupStoredQueryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStoredQueryResult, error) {
			args := v.(LookupStoredQueryArgs)
			r, err := LookupStoredQuery(ctx, &args, opts...)
			var s LookupStoredQueryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStoredQueryResultOutput)
}

type LookupStoredQueryOutputArgs struct {
	QueryName pulumi.StringInput `pulumi:"queryName"`
}

func (LookupStoredQueryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStoredQueryArgs)(nil)).Elem()
}

type LookupStoredQueryResultOutput struct{ *pulumi.OutputState }

func (LookupStoredQueryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStoredQueryResult)(nil)).Elem()
}

func (o LookupStoredQueryResultOutput) ToLookupStoredQueryResultOutput() LookupStoredQueryResultOutput {
	return o
}

func (o LookupStoredQueryResultOutput) ToLookupStoredQueryResultOutputWithContext(ctx context.Context) LookupStoredQueryResultOutput {
	return o
}

func (o LookupStoredQueryResultOutput) QueryArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryArn }).(pulumi.StringPtrOutput)
}

func (o LookupStoredQueryResultOutput) QueryDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryDescription }).(pulumi.StringPtrOutput)
}

func (o LookupStoredQueryResultOutput) QueryExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryExpression }).(pulumi.StringPtrOutput)
}

func (o LookupStoredQueryResultOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

// The tags for the stored query.
func (o LookupStoredQueryResultOutput) Tags() StoredQueryTagArrayOutput {
	return o.ApplyT(func(v LookupStoredQueryResult) []StoredQueryTag { return v.Tags }).(StoredQueryTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStoredQueryResultOutput{})
}
