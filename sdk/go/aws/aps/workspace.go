// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::APS::Workspace
type Workspace struct {
	pulumi.CustomResourceState

	// The AMP Workspace alert manager definition data
	AlertManagerDefinition pulumi.StringPtrOutput `pulumi:"alertManagerDefinition"`
	// AMP Workspace alias.
	Alias pulumi.StringPtrOutput `pulumi:"alias"`
	// Workspace arn.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// KMS Key ARN used to encrypt and decrypt AMP workspace data.
	KmsKeyArn            pulumi.StringPtrOutput                 `pulumi:"kmsKeyArn"`
	LoggingConfiguration WorkspaceLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// AMP Workspace prometheus endpoint
	PrometheusEndpoint pulumi.StringOutput `pulumi:"prometheusEndpoint"`
	// An array of key-value pairs to apply to this resource.
	Tags WorkspaceTagArrayOutput `pulumi:"tags"`
	// Required to identify a specific APS Workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		args = &WorkspaceArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"kmsKeyArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("aws-native:aps:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("aws-native:aps:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
}

type WorkspaceState struct {
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The AMP Workspace alert manager definition data
	AlertManagerDefinition *string `pulumi:"alertManagerDefinition"`
	// AMP Workspace alias.
	Alias *string `pulumi:"alias"`
	// KMS Key ARN used to encrypt and decrypt AMP workspace data.
	KmsKeyArn            *string                        `pulumi:"kmsKeyArn"`
	LoggingConfiguration *WorkspaceLoggingConfiguration `pulumi:"loggingConfiguration"`
	// An array of key-value pairs to apply to this resource.
	Tags []WorkspaceTag `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The AMP Workspace alert manager definition data
	AlertManagerDefinition pulumi.StringPtrInput
	// AMP Workspace alias.
	Alias pulumi.StringPtrInput
	// KMS Key ARN used to encrypt and decrypt AMP workspace data.
	KmsKeyArn            pulumi.StringPtrInput
	LoggingConfiguration WorkspaceLoggingConfigurationPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags WorkspaceTagArrayInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

// The AMP Workspace alert manager definition data
func (o WorkspaceOutput) AlertManagerDefinition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.AlertManagerDefinition }).(pulumi.StringPtrOutput)
}

// AMP Workspace alias.
func (o WorkspaceOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Alias }).(pulumi.StringPtrOutput)
}

// Workspace arn.
func (o WorkspaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// KMS Key ARN used to encrypt and decrypt AMP workspace data.
func (o WorkspaceOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) LoggingConfiguration() WorkspaceLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceLoggingConfigurationPtrOutput { return v.LoggingConfiguration }).(WorkspaceLoggingConfigurationPtrOutput)
}

// AMP Workspace prometheus endpoint
func (o WorkspaceOutput) PrometheusEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.PrometheusEndpoint }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o WorkspaceOutput) Tags() WorkspaceTagArrayOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceTagArrayOutput { return v.Tags }).(WorkspaceTagArrayOutput)
}

// Required to identify a specific APS Workspace.
func (o WorkspaceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceInput)(nil)).Elem(), &Workspace{})
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
