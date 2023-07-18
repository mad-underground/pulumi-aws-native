// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a studio component that connects a non-Nimble Studio resource in your account to your studio
func LookupStudioComponent(ctx *pulumi.Context, args *LookupStudioComponentArgs, opts ...pulumi.InvokeOption) (*LookupStudioComponentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStudioComponentResult
	err := ctx.Invoke("aws-native:nimblestudio:getStudioComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudioComponentArgs struct {
	StudioComponentId string `pulumi:"studioComponentId"`
	// <p>The studio ID. </p>
	StudioId string `pulumi:"studioId"`
}

type LookupStudioComponentResult struct {
	Configuration *StudioComponentConfiguration `pulumi:"configuration"`
	// <p>The description.</p>
	Description *string `pulumi:"description"`
	// <p>The EC2 security groups that control access to the studio component.</p>
	Ec2SecurityGroupIds []string `pulumi:"ec2SecurityGroupIds"`
	// <p>Initialization scripts for studio components.</p>
	InitializationScripts []StudioComponentInitializationScript `pulumi:"initializationScripts"`
	// <p>The name for the studio component.</p>
	Name           *string `pulumi:"name"`
	RuntimeRoleArn *string `pulumi:"runtimeRoleArn"`
	// <p>Parameters for the studio component scripts.</p>
	ScriptParameters            []StudioComponentScriptParameterKeyValue `pulumi:"scriptParameters"`
	SecureInitializationRoleArn *string                                  `pulumi:"secureInitializationRoleArn"`
	StudioComponentId           *string                                  `pulumi:"studioComponentId"`
	Type                        *StudioComponentType                     `pulumi:"type"`
}

func LookupStudioComponentOutput(ctx *pulumi.Context, args LookupStudioComponentOutputArgs, opts ...pulumi.InvokeOption) LookupStudioComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStudioComponentResult, error) {
			args := v.(LookupStudioComponentArgs)
			r, err := LookupStudioComponent(ctx, &args, opts...)
			var s LookupStudioComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStudioComponentResultOutput)
}

type LookupStudioComponentOutputArgs struct {
	StudioComponentId pulumi.StringInput `pulumi:"studioComponentId"`
	// <p>The studio ID. </p>
	StudioId pulumi.StringInput `pulumi:"studioId"`
}

func (LookupStudioComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioComponentArgs)(nil)).Elem()
}

type LookupStudioComponentResultOutput struct{ *pulumi.OutputState }

func (LookupStudioComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudioComponentResult)(nil)).Elem()
}

func (o LookupStudioComponentResultOutput) ToLookupStudioComponentResultOutput() LookupStudioComponentResultOutput {
	return o
}

func (o LookupStudioComponentResultOutput) ToLookupStudioComponentResultOutputWithContext(ctx context.Context) LookupStudioComponentResultOutput {
	return o
}

func (o LookupStudioComponentResultOutput) Configuration() StudioComponentConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *StudioComponentConfiguration { return v.Configuration }).(StudioComponentConfigurationPtrOutput)
}

// <p>The description.</p>
func (o LookupStudioComponentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// <p>The EC2 security groups that control access to the studio component.</p>
func (o LookupStudioComponentResultOutput) Ec2SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) []string { return v.Ec2SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// <p>Initialization scripts for studio components.</p>
func (o LookupStudioComponentResultOutput) InitializationScripts() StudioComponentInitializationScriptArrayOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) []StudioComponentInitializationScript {
		return v.InitializationScripts
	}).(StudioComponentInitializationScriptArrayOutput)
}

// <p>The name for the studio component.</p>
func (o LookupStudioComponentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupStudioComponentResultOutput) RuntimeRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.RuntimeRoleArn }).(pulumi.StringPtrOutput)
}

// <p>Parameters for the studio component scripts.</p>
func (o LookupStudioComponentResultOutput) ScriptParameters() StudioComponentScriptParameterKeyValueArrayOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) []StudioComponentScriptParameterKeyValue {
		return v.ScriptParameters
	}).(StudioComponentScriptParameterKeyValueArrayOutput)
}

func (o LookupStudioComponentResultOutput) SecureInitializationRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.SecureInitializationRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupStudioComponentResultOutput) StudioComponentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *string { return v.StudioComponentId }).(pulumi.StringPtrOutput)
}

func (o LookupStudioComponentResultOutput) Type() StudioComponentTypePtrOutput {
	return o.ApplyT(func(v LookupStudioComponentResult) *StudioComponentType { return v.Type }).(StudioComponentTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudioComponentResultOutput{})
}
