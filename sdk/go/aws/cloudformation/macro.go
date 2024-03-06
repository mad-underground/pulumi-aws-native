// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFormation::Macro
//
// Deprecated: Macro is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Macro struct {
	pulumi.CustomResourceState

	AwsId        pulumi.StringOutput    `pulumi:"awsId"`
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	FunctionName pulumi.StringOutput    `pulumi:"functionName"`
	LogGroupName pulumi.StringPtrOutput `pulumi:"logGroupName"`
	LogRoleArn   pulumi.StringPtrOutput `pulumi:"logRoleArn"`
	Name         pulumi.StringOutput    `pulumi:"name"`
}

// NewMacro registers a new resource with the given unique name, arguments, and options.
func NewMacro(ctx *pulumi.Context,
	name string, args *MacroArgs, opts ...pulumi.ResourceOption) (*Macro, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Macro
	err := ctx.RegisterResource("aws-native:cloudformation:Macro", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMacro gets an existing Macro resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMacro(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MacroState, opts ...pulumi.ResourceOption) (*Macro, error) {
	var resource Macro
	err := ctx.ReadResource("aws-native:cloudformation:Macro", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Macro resources.
type macroState struct {
}

type MacroState struct {
}

func (MacroState) ElementType() reflect.Type {
	return reflect.TypeOf((*macroState)(nil)).Elem()
}

type macroArgs struct {
	Description  *string `pulumi:"description"`
	FunctionName string  `pulumi:"functionName"`
	LogGroupName *string `pulumi:"logGroupName"`
	LogRoleArn   *string `pulumi:"logRoleArn"`
	Name         *string `pulumi:"name"`
}

// The set of arguments for constructing a Macro resource.
type MacroArgs struct {
	Description  pulumi.StringPtrInput
	FunctionName pulumi.StringInput
	LogGroupName pulumi.StringPtrInput
	LogRoleArn   pulumi.StringPtrInput
	Name         pulumi.StringPtrInput
}

func (MacroArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*macroArgs)(nil)).Elem()
}

type MacroInput interface {
	pulumi.Input

	ToMacroOutput() MacroOutput
	ToMacroOutputWithContext(ctx context.Context) MacroOutput
}

func (*Macro) ElementType() reflect.Type {
	return reflect.TypeOf((**Macro)(nil)).Elem()
}

func (i *Macro) ToMacroOutput() MacroOutput {
	return i.ToMacroOutputWithContext(context.Background())
}

func (i *Macro) ToMacroOutputWithContext(ctx context.Context) MacroOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MacroOutput)
}

type MacroOutput struct{ *pulumi.OutputState }

func (MacroOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Macro)(nil)).Elem()
}

func (o MacroOutput) ToMacroOutput() MacroOutput {
	return o
}

func (o MacroOutput) ToMacroOutputWithContext(ctx context.Context) MacroOutput {
	return o
}

func (o MacroOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Macro) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o MacroOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Macro) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MacroOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Macro) pulumi.StringOutput { return v.FunctionName }).(pulumi.StringOutput)
}

func (o MacroOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Macro) pulumi.StringPtrOutput { return v.LogGroupName }).(pulumi.StringPtrOutput)
}

func (o MacroOutput) LogRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Macro) pulumi.StringPtrOutput { return v.LogRoleArn }).(pulumi.StringPtrOutput)
}

func (o MacroOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Macro) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MacroInput)(nil)).Elem(), &Macro{})
	pulumi.RegisterOutputType(MacroOutput{})
}
