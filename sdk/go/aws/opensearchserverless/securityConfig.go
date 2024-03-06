// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon OpenSearchServerless security config resource
type SecurityConfig struct {
	pulumi.CustomResourceState

	// The identifier of the security config
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// Security config description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name of the security config
	Name        pulumi.StringPtrOutput                   `pulumi:"name"`
	SamlOptions SecurityConfigSamlConfigOptionsPtrOutput `pulumi:"samlOptions"`
	Type        SecurityConfigTypePtrOutput              `pulumi:"type"`
}

// NewSecurityConfig registers a new resource with the given unique name, arguments, and options.
func NewSecurityConfig(ctx *pulumi.Context,
	name string, args *SecurityConfigArgs, opts ...pulumi.ResourceOption) (*SecurityConfig, error) {
	if args == nil {
		args = &SecurityConfigArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityConfig
	err := ctx.RegisterResource("aws-native:opensearchserverless:SecurityConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityConfig gets an existing SecurityConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConfigState, opts ...pulumi.ResourceOption) (*SecurityConfig, error) {
	var resource SecurityConfig
	err := ctx.ReadResource("aws-native:opensearchserverless:SecurityConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityConfig resources.
type securityConfigState struct {
}

type SecurityConfigState struct {
}

func (SecurityConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConfigState)(nil)).Elem()
}

type securityConfigArgs struct {
	// Security config description
	Description *string `pulumi:"description"`
	// The friendly name of the security config
	Name        *string                          `pulumi:"name"`
	SamlOptions *SecurityConfigSamlConfigOptions `pulumi:"samlOptions"`
	Type        *SecurityConfigType              `pulumi:"type"`
}

// The set of arguments for constructing a SecurityConfig resource.
type SecurityConfigArgs struct {
	// Security config description
	Description pulumi.StringPtrInput
	// The friendly name of the security config
	Name        pulumi.StringPtrInput
	SamlOptions SecurityConfigSamlConfigOptionsPtrInput
	Type        SecurityConfigTypePtrInput
}

func (SecurityConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConfigArgs)(nil)).Elem()
}

type SecurityConfigInput interface {
	pulumi.Input

	ToSecurityConfigOutput() SecurityConfigOutput
	ToSecurityConfigOutputWithContext(ctx context.Context) SecurityConfigOutput
}

func (*SecurityConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfig)(nil)).Elem()
}

func (i *SecurityConfig) ToSecurityConfigOutput() SecurityConfigOutput {
	return i.ToSecurityConfigOutputWithContext(context.Background())
}

func (i *SecurityConfig) ToSecurityConfigOutputWithContext(ctx context.Context) SecurityConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConfigOutput)
}

type SecurityConfigOutput struct{ *pulumi.OutputState }

func (SecurityConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfig)(nil)).Elem()
}

func (o SecurityConfigOutput) ToSecurityConfigOutput() SecurityConfigOutput {
	return o
}

func (o SecurityConfigOutput) ToSecurityConfigOutputWithContext(ctx context.Context) SecurityConfigOutput {
	return o
}

// The identifier of the security config
func (o SecurityConfigOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConfig) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// Security config description
func (o SecurityConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The friendly name of the security config
func (o SecurityConfigOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConfig) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityConfigOutput) SamlOptions() SecurityConfigSamlConfigOptionsPtrOutput {
	return o.ApplyT(func(v *SecurityConfig) SecurityConfigSamlConfigOptionsPtrOutput { return v.SamlOptions }).(SecurityConfigSamlConfigOptionsPtrOutput)
}

func (o SecurityConfigOutput) Type() SecurityConfigTypePtrOutput {
	return o.ApplyT(func(v *SecurityConfig) SecurityConfigTypePtrOutput { return v.Type }).(SecurityConfigTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigInput)(nil)).Elem(), &SecurityConfig{})
	pulumi.RegisterOutputType(SecurityConfigOutput{})
}
