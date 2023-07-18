// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodeBuild::SourceCredential
//
// Deprecated: SourceCredential is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type SourceCredential struct {
	pulumi.CustomResourceState

	AuthType   pulumi.StringOutput    `pulumi:"authType"`
	ServerType pulumi.StringOutput    `pulumi:"serverType"`
	Token      pulumi.StringOutput    `pulumi:"token"`
	Username   pulumi.StringPtrOutput `pulumi:"username"`
}

// NewSourceCredential registers a new resource with the given unique name, arguments, and options.
func NewSourceCredential(ctx *pulumi.Context,
	name string, args *SourceCredentialArgs, opts ...pulumi.ResourceOption) (*SourceCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthType == nil {
		return nil, errors.New("invalid value for required argument 'AuthType'")
	}
	if args.ServerType == nil {
		return nil, errors.New("invalid value for required argument 'ServerType'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SourceCredential
	err := ctx.RegisterResource("aws-native:codebuild:SourceCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCredential gets an existing SourceCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCredentialState, opts ...pulumi.ResourceOption) (*SourceCredential, error) {
	var resource SourceCredential
	err := ctx.ReadResource("aws-native:codebuild:SourceCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCredential resources.
type sourceCredentialState struct {
}

type SourceCredentialState struct {
}

func (SourceCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCredentialState)(nil)).Elem()
}

type sourceCredentialArgs struct {
	AuthType   string  `pulumi:"authType"`
	ServerType string  `pulumi:"serverType"`
	Token      string  `pulumi:"token"`
	Username   *string `pulumi:"username"`
}

// The set of arguments for constructing a SourceCredential resource.
type SourceCredentialArgs struct {
	AuthType   pulumi.StringInput
	ServerType pulumi.StringInput
	Token      pulumi.StringInput
	Username   pulumi.StringPtrInput
}

func (SourceCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCredentialArgs)(nil)).Elem()
}

type SourceCredentialInput interface {
	pulumi.Input

	ToSourceCredentialOutput() SourceCredentialOutput
	ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput
}

func (*SourceCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCredential)(nil)).Elem()
}

func (i *SourceCredential) ToSourceCredentialOutput() SourceCredentialOutput {
	return i.ToSourceCredentialOutputWithContext(context.Background())
}

func (i *SourceCredential) ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCredentialOutput)
}

type SourceCredentialOutput struct{ *pulumi.OutputState }

func (SourceCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCredential)(nil)).Elem()
}

func (o SourceCredentialOutput) ToSourceCredentialOutput() SourceCredentialOutput {
	return o
}

func (o SourceCredentialOutput) ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput {
	return o
}

func (o SourceCredentialOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

func (o SourceCredentialOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

func (o SourceCredentialOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o SourceCredentialOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceCredential) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceCredentialInput)(nil)).Elem(), &SourceCredential{})
	pulumi.RegisterOutputType(SourceCredentialOutput{})
}
