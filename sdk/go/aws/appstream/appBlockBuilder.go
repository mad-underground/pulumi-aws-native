// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::AppBlockBuilder.
type AppBlockBuilder struct {
	pulumi.CustomResourceState

	AccessEndpoints             AppBlockBuilderAccessEndpointArrayOutput `pulumi:"accessEndpoints"`
	AppBlockArns                pulumi.StringArrayOutput                 `pulumi:"appBlockArns"`
	Arn                         pulumi.StringOutput                      `pulumi:"arn"`
	CreatedTime                 pulumi.StringOutput                      `pulumi:"createdTime"`
	Description                 pulumi.StringPtrOutput                   `pulumi:"description"`
	DisplayName                 pulumi.StringPtrOutput                   `pulumi:"displayName"`
	EnableDefaultInternetAccess pulumi.BoolPtrOutput                     `pulumi:"enableDefaultInternetAccess"`
	IamRoleArn                  pulumi.StringPtrOutput                   `pulumi:"iamRoleArn"`
	InstanceType                pulumi.StringOutput                      `pulumi:"instanceType"`
	Name                        pulumi.StringOutput                      `pulumi:"name"`
	Platform                    pulumi.StringOutput                      `pulumi:"platform"`
	Tags                        aws.TagArrayOutput                       `pulumi:"tags"`
	VpcConfig                   AppBlockBuilderVpcConfigOutput           `pulumi:"vpcConfig"`
}

// NewAppBlockBuilder registers a new resource with the given unique name, arguments, and options.
func NewAppBlockBuilder(ctx *pulumi.Context,
	name string, args *AppBlockBuilderArgs, opts ...pulumi.ResourceOption) (*AppBlockBuilder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.VpcConfig == nil {
		return nil, errors.New("invalid value for required argument 'VpcConfig'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppBlockBuilder
	err := ctx.RegisterResource("aws-native:appstream:AppBlockBuilder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppBlockBuilder gets an existing AppBlockBuilder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppBlockBuilder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppBlockBuilderState, opts ...pulumi.ResourceOption) (*AppBlockBuilder, error) {
	var resource AppBlockBuilder
	err := ctx.ReadResource("aws-native:appstream:AppBlockBuilder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppBlockBuilder resources.
type appBlockBuilderState struct {
}

type AppBlockBuilderState struct {
}

func (AppBlockBuilderState) ElementType() reflect.Type {
	return reflect.TypeOf((*appBlockBuilderState)(nil)).Elem()
}

type appBlockBuilderArgs struct {
	AccessEndpoints             []AppBlockBuilderAccessEndpoint `pulumi:"accessEndpoints"`
	AppBlockArns                []string                        `pulumi:"appBlockArns"`
	Description                 *string                         `pulumi:"description"`
	DisplayName                 *string                         `pulumi:"displayName"`
	EnableDefaultInternetAccess *bool                           `pulumi:"enableDefaultInternetAccess"`
	IamRoleArn                  *string                         `pulumi:"iamRoleArn"`
	InstanceType                string                          `pulumi:"instanceType"`
	Name                        *string                         `pulumi:"name"`
	Platform                    string                          `pulumi:"platform"`
	Tags                        []aws.Tag                       `pulumi:"tags"`
	VpcConfig                   AppBlockBuilderVpcConfig        `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a AppBlockBuilder resource.
type AppBlockBuilderArgs struct {
	AccessEndpoints             AppBlockBuilderAccessEndpointArrayInput
	AppBlockArns                pulumi.StringArrayInput
	Description                 pulumi.StringPtrInput
	DisplayName                 pulumi.StringPtrInput
	EnableDefaultInternetAccess pulumi.BoolPtrInput
	IamRoleArn                  pulumi.StringPtrInput
	InstanceType                pulumi.StringInput
	Name                        pulumi.StringPtrInput
	Platform                    pulumi.StringInput
	Tags                        aws.TagArrayInput
	VpcConfig                   AppBlockBuilderVpcConfigInput
}

func (AppBlockBuilderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appBlockBuilderArgs)(nil)).Elem()
}

type AppBlockBuilderInput interface {
	pulumi.Input

	ToAppBlockBuilderOutput() AppBlockBuilderOutput
	ToAppBlockBuilderOutputWithContext(ctx context.Context) AppBlockBuilderOutput
}

func (*AppBlockBuilder) ElementType() reflect.Type {
	return reflect.TypeOf((**AppBlockBuilder)(nil)).Elem()
}

func (i *AppBlockBuilder) ToAppBlockBuilderOutput() AppBlockBuilderOutput {
	return i.ToAppBlockBuilderOutputWithContext(context.Background())
}

func (i *AppBlockBuilder) ToAppBlockBuilderOutputWithContext(ctx context.Context) AppBlockBuilderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppBlockBuilderOutput)
}

type AppBlockBuilderOutput struct{ *pulumi.OutputState }

func (AppBlockBuilderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppBlockBuilder)(nil)).Elem()
}

func (o AppBlockBuilderOutput) ToAppBlockBuilderOutput() AppBlockBuilderOutput {
	return o
}

func (o AppBlockBuilderOutput) ToAppBlockBuilderOutputWithContext(ctx context.Context) AppBlockBuilderOutput {
	return o
}

func (o AppBlockBuilderOutput) AccessEndpoints() AppBlockBuilderAccessEndpointArrayOutput {
	return o.ApplyT(func(v *AppBlockBuilder) AppBlockBuilderAccessEndpointArrayOutput { return v.AccessEndpoints }).(AppBlockBuilderAccessEndpointArrayOutput)
}

func (o AppBlockBuilderOutput) AppBlockArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringArrayOutput { return v.AppBlockArns }).(pulumi.StringArrayOutput)
}

func (o AppBlockBuilderOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o AppBlockBuilderOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o AppBlockBuilderOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AppBlockBuilderOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AppBlockBuilderOutput) EnableDefaultInternetAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.BoolPtrOutput { return v.EnableDefaultInternetAccess }).(pulumi.BoolPtrOutput)
}

func (o AppBlockBuilderOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringPtrOutput { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

func (o AppBlockBuilderOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

func (o AppBlockBuilderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppBlockBuilderOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

func (o AppBlockBuilderOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *AppBlockBuilder) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o AppBlockBuilderOutput) VpcConfig() AppBlockBuilderVpcConfigOutput {
	return o.ApplyT(func(v *AppBlockBuilder) AppBlockBuilderVpcConfigOutput { return v.VpcConfig }).(AppBlockBuilderVpcConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppBlockBuilderInput)(nil)).Elem(), &AppBlockBuilder{})
	pulumi.RegisterOutputType(AppBlockBuilderOutput{})
}
