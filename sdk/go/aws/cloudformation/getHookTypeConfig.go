// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the configuration data for a registered hook in CloudFormation Registry.
func LookupHookTypeConfig(ctx *pulumi.Context, args *LookupHookTypeConfigArgs, opts ...pulumi.InvokeOption) (*LookupHookTypeConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupHookTypeConfigResult
	err := ctx.Invoke("aws-native:cloudformation:getHookTypeConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHookTypeConfigArgs struct {
	// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
	ConfigurationArn string `pulumi:"configurationArn"`
}

type LookupHookTypeConfigResult struct {
	// The configuration data for the extension, in this account and region.
	Configuration *string `pulumi:"configuration"`
	// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
	ConfigurationArn *string `pulumi:"configurationArn"`
	// The Amazon Resource Name (ARN) of the type without version number.
	TypeArn *string `pulumi:"typeArn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName *string `pulumi:"typeName"`
}

func LookupHookTypeConfigOutput(ctx *pulumi.Context, args LookupHookTypeConfigOutputArgs, opts ...pulumi.InvokeOption) LookupHookTypeConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHookTypeConfigResult, error) {
			args := v.(LookupHookTypeConfigArgs)
			r, err := LookupHookTypeConfig(ctx, &args, opts...)
			var s LookupHookTypeConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHookTypeConfigResultOutput)
}

type LookupHookTypeConfigOutputArgs struct {
	// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
	ConfigurationArn pulumi.StringInput `pulumi:"configurationArn"`
}

func (LookupHookTypeConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHookTypeConfigArgs)(nil)).Elem()
}

type LookupHookTypeConfigResultOutput struct{ *pulumi.OutputState }

func (LookupHookTypeConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHookTypeConfigResult)(nil)).Elem()
}

func (o LookupHookTypeConfigResultOutput) ToLookupHookTypeConfigResultOutput() LookupHookTypeConfigResultOutput {
	return o
}

func (o LookupHookTypeConfigResultOutput) ToLookupHookTypeConfigResultOutputWithContext(ctx context.Context) LookupHookTypeConfigResultOutput {
	return o
}

// The configuration data for the extension, in this account and region.
func (o LookupHookTypeConfigResultOutput) Configuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookTypeConfigResult) *string { return v.Configuration }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) for the configuration data, in this account and region.
func (o LookupHookTypeConfigResultOutput) ConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookTypeConfigResult) *string { return v.ConfigurationArn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the type without version number.
func (o LookupHookTypeConfigResultOutput) TypeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookTypeConfigResult) *string { return v.TypeArn }).(pulumi.StringPtrOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o LookupHookTypeConfigResultOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookTypeConfigResult) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHookTypeConfigResultOutput{})
}
