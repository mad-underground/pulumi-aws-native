// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::Alias
func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAliasResult
	err := ctx.Invoke("aws-native:lambda:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAliasArgs struct {
	Id string `pulumi:"id"`
}

type LookupAliasResult struct {
	Description                  *string                                   `pulumi:"description"`
	FunctionVersion              *string                                   `pulumi:"functionVersion"`
	Id                           *string                                   `pulumi:"id"`
	ProvisionedConcurrencyConfig *AliasProvisionedConcurrencyConfiguration `pulumi:"provisionedConcurrencyConfig"`
	RoutingConfig                *AliasRoutingConfiguration                `pulumi:"routingConfig"`
}

func LookupAliasOutput(ctx *pulumi.Context, args LookupAliasOutputArgs, opts ...pulumi.InvokeOption) LookupAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAliasResult, error) {
			args := v.(LookupAliasArgs)
			r, err := LookupAlias(ctx, &args, opts...)
			var s LookupAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAliasResultOutput)
}

type LookupAliasOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAliasArgs)(nil)).Elem()
}

type LookupAliasResultOutput struct{ *pulumi.OutputState }

func (LookupAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAliasResult)(nil)).Elem()
}

func (o LookupAliasResultOutput) ToLookupAliasResultOutput() LookupAliasResultOutput {
	return o
}

func (o LookupAliasResultOutput) ToLookupAliasResultOutputWithContext(ctx context.Context) LookupAliasResultOutput {
	return o
}

func (o LookupAliasResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAliasResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAliasResultOutput) FunctionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAliasResult) *string { return v.FunctionVersion }).(pulumi.StringPtrOutput)
}

func (o LookupAliasResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAliasResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAliasResultOutput) ProvisionedConcurrencyConfig() AliasProvisionedConcurrencyConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAliasResult) *AliasProvisionedConcurrencyConfiguration {
		return v.ProvisionedConcurrencyConfig
	}).(AliasProvisionedConcurrencyConfigurationPtrOutput)
}

func (o LookupAliasResultOutput) RoutingConfig() AliasRoutingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAliasResult) *AliasRoutingConfiguration { return v.RoutingConfig }).(AliasRoutingConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAliasResultOutput{})
}
