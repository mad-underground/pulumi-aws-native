// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WorkSpaces::ConnectionAlias
func LookupConnectionAlias(ctx *pulumi.Context, args *LookupConnectionAliasArgs, opts ...pulumi.InvokeOption) (*LookupConnectionAliasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionAliasResult
	err := ctx.Invoke("aws-native:workspaces:getConnectionAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionAliasArgs struct {
	AliasId string `pulumi:"aliasId"`
}

type LookupConnectionAliasResult struct {
	AliasId              *string                      `pulumi:"aliasId"`
	Associations         []ConnectionAliasAssociation `pulumi:"associations"`
	ConnectionAliasState *ConnectionAliasStateEnum    `pulumi:"connectionAliasState"`
}

func LookupConnectionAliasOutput(ctx *pulumi.Context, args LookupConnectionAliasOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionAliasResult, error) {
			args := v.(LookupConnectionAliasArgs)
			r, err := LookupConnectionAlias(ctx, &args, opts...)
			var s LookupConnectionAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionAliasResultOutput)
}

type LookupConnectionAliasOutputArgs struct {
	AliasId pulumi.StringInput `pulumi:"aliasId"`
}

func (LookupConnectionAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionAliasArgs)(nil)).Elem()
}

type LookupConnectionAliasResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionAliasResult)(nil)).Elem()
}

func (o LookupConnectionAliasResultOutput) ToLookupConnectionAliasResultOutput() LookupConnectionAliasResultOutput {
	return o
}

func (o LookupConnectionAliasResultOutput) ToLookupConnectionAliasResultOutputWithContext(ctx context.Context) LookupConnectionAliasResultOutput {
	return o
}

func (o LookupConnectionAliasResultOutput) AliasId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionAliasResult) *string { return v.AliasId }).(pulumi.StringPtrOutput)
}

func (o LookupConnectionAliasResultOutput) Associations() ConnectionAliasAssociationArrayOutput {
	return o.ApplyT(func(v LookupConnectionAliasResult) []ConnectionAliasAssociation { return v.Associations }).(ConnectionAliasAssociationArrayOutput)
}

func (o LookupConnectionAliasResultOutput) ConnectionAliasState() ConnectionAliasStateEnumPtrOutput {
	return o.ApplyT(func(v LookupConnectionAliasResult) *ConnectionAliasStateEnum { return v.ConnectionAliasState }).(ConnectionAliasStateEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionAliasResultOutput{})
}
