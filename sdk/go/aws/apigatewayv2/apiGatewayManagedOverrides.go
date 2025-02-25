// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::ApiGatewayManagedOverrides
//
// Deprecated: ApiGatewayManagedOverrides is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApiGatewayManagedOverrides struct {
	pulumi.CustomResourceState

	ApiId       pulumi.StringOutput                                     `pulumi:"apiId"`
	AwsId       pulumi.StringOutput                                     `pulumi:"awsId"`
	Integration ApiGatewayManagedOverridesIntegrationOverridesPtrOutput `pulumi:"integration"`
	Route       ApiGatewayManagedOverridesRouteOverridesPtrOutput       `pulumi:"route"`
	Stage       ApiGatewayManagedOverridesStageOverridesPtrOutput       `pulumi:"stage"`
}

// NewApiGatewayManagedOverrides registers a new resource with the given unique name, arguments, and options.
func NewApiGatewayManagedOverrides(ctx *pulumi.Context,
	name string, args *ApiGatewayManagedOverridesArgs, opts ...pulumi.ResourceOption) (*ApiGatewayManagedOverrides, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiGatewayManagedOverrides
	err := ctx.RegisterResource("aws-native:apigatewayv2:ApiGatewayManagedOverrides", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiGatewayManagedOverrides gets an existing ApiGatewayManagedOverrides resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiGatewayManagedOverrides(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiGatewayManagedOverridesState, opts ...pulumi.ResourceOption) (*ApiGatewayManagedOverrides, error) {
	var resource ApiGatewayManagedOverrides
	err := ctx.ReadResource("aws-native:apigatewayv2:ApiGatewayManagedOverrides", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiGatewayManagedOverrides resources.
type apiGatewayManagedOverridesState struct {
}

type ApiGatewayManagedOverridesState struct {
}

func (ApiGatewayManagedOverridesState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayManagedOverridesState)(nil)).Elem()
}

type apiGatewayManagedOverridesArgs struct {
	ApiId       string                                          `pulumi:"apiId"`
	Integration *ApiGatewayManagedOverridesIntegrationOverrides `pulumi:"integration"`
	Route       *ApiGatewayManagedOverridesRouteOverrides       `pulumi:"route"`
	Stage       *ApiGatewayManagedOverridesStageOverrides       `pulumi:"stage"`
}

// The set of arguments for constructing a ApiGatewayManagedOverrides resource.
type ApiGatewayManagedOverridesArgs struct {
	ApiId       pulumi.StringInput
	Integration ApiGatewayManagedOverridesIntegrationOverridesPtrInput
	Route       ApiGatewayManagedOverridesRouteOverridesPtrInput
	Stage       ApiGatewayManagedOverridesStageOverridesPtrInput
}

func (ApiGatewayManagedOverridesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayManagedOverridesArgs)(nil)).Elem()
}

type ApiGatewayManagedOverridesInput interface {
	pulumi.Input

	ToApiGatewayManagedOverridesOutput() ApiGatewayManagedOverridesOutput
	ToApiGatewayManagedOverridesOutputWithContext(ctx context.Context) ApiGatewayManagedOverridesOutput
}

func (*ApiGatewayManagedOverrides) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayManagedOverrides)(nil)).Elem()
}

func (i *ApiGatewayManagedOverrides) ToApiGatewayManagedOverridesOutput() ApiGatewayManagedOverridesOutput {
	return i.ToApiGatewayManagedOverridesOutputWithContext(context.Background())
}

func (i *ApiGatewayManagedOverrides) ToApiGatewayManagedOverridesOutputWithContext(ctx context.Context) ApiGatewayManagedOverridesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayManagedOverridesOutput)
}

type ApiGatewayManagedOverridesOutput struct{ *pulumi.OutputState }

func (ApiGatewayManagedOverridesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayManagedOverrides)(nil)).Elem()
}

func (o ApiGatewayManagedOverridesOutput) ToApiGatewayManagedOverridesOutput() ApiGatewayManagedOverridesOutput {
	return o
}

func (o ApiGatewayManagedOverridesOutput) ToApiGatewayManagedOverridesOutputWithContext(ctx context.Context) ApiGatewayManagedOverridesOutput {
	return o
}

func (o ApiGatewayManagedOverridesOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayManagedOverrides) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

func (o ApiGatewayManagedOverridesOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayManagedOverrides) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o ApiGatewayManagedOverridesOutput) Integration() ApiGatewayManagedOverridesIntegrationOverridesPtrOutput {
	return o.ApplyT(func(v *ApiGatewayManagedOverrides) ApiGatewayManagedOverridesIntegrationOverridesPtrOutput {
		return v.Integration
	}).(ApiGatewayManagedOverridesIntegrationOverridesPtrOutput)
}

func (o ApiGatewayManagedOverridesOutput) Route() ApiGatewayManagedOverridesRouteOverridesPtrOutput {
	return o.ApplyT(func(v *ApiGatewayManagedOverrides) ApiGatewayManagedOverridesRouteOverridesPtrOutput { return v.Route }).(ApiGatewayManagedOverridesRouteOverridesPtrOutput)
}

func (o ApiGatewayManagedOverridesOutput) Stage() ApiGatewayManagedOverridesStageOverridesPtrOutput {
	return o.ApplyT(func(v *ApiGatewayManagedOverrides) ApiGatewayManagedOverridesStageOverridesPtrOutput { return v.Stage }).(ApiGatewayManagedOverridesStageOverridesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiGatewayManagedOverridesInput)(nil)).Elem(), &ApiGatewayManagedOverrides{})
	pulumi.RegisterOutputType(ApiGatewayManagedOverridesOutput{})
}
