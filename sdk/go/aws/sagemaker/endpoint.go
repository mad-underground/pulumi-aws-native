// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Endpoint
//
// Deprecated: Endpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Endpoint struct {
	pulumi.CustomResourceState

	DeploymentConfig                 EndpointDeploymentConfigPtrOutput  `pulumi:"deploymentConfig"`
	EndpointConfigName               pulumi.StringOutput                `pulumi:"endpointConfigName"`
	EndpointName                     pulumi.StringPtrOutput             `pulumi:"endpointName"`
	ExcludeRetainedVariantProperties EndpointVariantPropertyArrayOutput `pulumi:"excludeRetainedVariantProperties"`
	RetainAllVariantProperties       pulumi.BoolPtrOutput               `pulumi:"retainAllVariantProperties"`
	RetainDeploymentConfig           pulumi.BoolPtrOutput               `pulumi:"retainDeploymentConfig"`
	Tags                             EndpointTagArrayOutput             `pulumi:"tags"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointConfigName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointConfigName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("aws-native:sagemaker:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("aws-native:sagemaker:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	DeploymentConfig                 *EndpointDeploymentConfig `pulumi:"deploymentConfig"`
	EndpointConfigName               string                    `pulumi:"endpointConfigName"`
	EndpointName                     *string                   `pulumi:"endpointName"`
	ExcludeRetainedVariantProperties []EndpointVariantProperty `pulumi:"excludeRetainedVariantProperties"`
	RetainAllVariantProperties       *bool                     `pulumi:"retainAllVariantProperties"`
	RetainDeploymentConfig           *bool                     `pulumi:"retainDeploymentConfig"`
	Tags                             []EndpointTag             `pulumi:"tags"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	DeploymentConfig                 EndpointDeploymentConfigPtrInput
	EndpointConfigName               pulumi.StringInput
	EndpointName                     pulumi.StringPtrInput
	ExcludeRetainedVariantProperties EndpointVariantPropertyArrayInput
	RetainAllVariantProperties       pulumi.BoolPtrInput
	RetainDeploymentConfig           pulumi.BoolPtrInput
	Tags                             EndpointTagArrayInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) DeploymentConfig() EndpointDeploymentConfigPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointDeploymentConfigPtrOutput { return v.DeploymentConfig }).(EndpointDeploymentConfigPtrOutput)
}

func (o EndpointOutput) EndpointConfigName() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointConfigName }).(pulumi.StringOutput)
}

func (o EndpointOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) ExcludeRetainedVariantProperties() EndpointVariantPropertyArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointVariantPropertyArrayOutput { return v.ExcludeRetainedVariantProperties }).(EndpointVariantPropertyArrayOutput)
}

func (o EndpointOutput) RetainAllVariantProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.RetainAllVariantProperties }).(pulumi.BoolPtrOutput)
}

func (o EndpointOutput) RetainDeploymentConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.RetainDeploymentConfig }).(pulumi.BoolPtrOutput)
}

func (o EndpointOutput) Tags() EndpointTagArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointTagArrayOutput { return v.Tags }).(EndpointTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterOutputType(EndpointOutput{})
}
