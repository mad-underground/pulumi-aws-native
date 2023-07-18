// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
type StudioSessionMapping struct {
	pulumi.CustomResourceState

	// The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
	IdentityName pulumi.StringOutput `pulumi:"identityName"`
	// Specifies whether the identity to map to the Studio is a user or a group.
	IdentityType StudioSessionMappingIdentityTypeOutput `pulumi:"identityType"`
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
	SessionPolicyArn pulumi.StringOutput `pulumi:"sessionPolicyArn"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId pulumi.StringOutput `pulumi:"studioId"`
}

// NewStudioSessionMapping registers a new resource with the given unique name, arguments, and options.
func NewStudioSessionMapping(ctx *pulumi.Context,
	name string, args *StudioSessionMappingArgs, opts ...pulumi.ResourceOption) (*StudioSessionMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityName == nil {
		return nil, errors.New("invalid value for required argument 'IdentityName'")
	}
	if args.IdentityType == nil {
		return nil, errors.New("invalid value for required argument 'IdentityType'")
	}
	if args.SessionPolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'SessionPolicyArn'")
	}
	if args.StudioId == nil {
		return nil, errors.New("invalid value for required argument 'StudioId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StudioSessionMapping
	err := ctx.RegisterResource("aws-native:emr:StudioSessionMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStudioSessionMapping gets an existing StudioSessionMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStudioSessionMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudioSessionMappingState, opts ...pulumi.ResourceOption) (*StudioSessionMapping, error) {
	var resource StudioSessionMapping
	err := ctx.ReadResource("aws-native:emr:StudioSessionMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StudioSessionMapping resources.
type studioSessionMappingState struct {
}

type StudioSessionMappingState struct {
}

func (StudioSessionMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*studioSessionMappingState)(nil)).Elem()
}

type studioSessionMappingArgs struct {
	// The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
	IdentityName string `pulumi:"identityName"`
	// Specifies whether the identity to map to the Studio is a user or a group.
	IdentityType StudioSessionMappingIdentityType `pulumi:"identityType"`
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
	SessionPolicyArn string `pulumi:"sessionPolicyArn"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId string `pulumi:"studioId"`
}

// The set of arguments for constructing a StudioSessionMapping resource.
type StudioSessionMappingArgs struct {
	// The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
	IdentityName pulumi.StringInput
	// Specifies whether the identity to map to the Studio is a user or a group.
	IdentityType StudioSessionMappingIdentityTypeInput
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
	SessionPolicyArn pulumi.StringInput
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	StudioId pulumi.StringInput
}

func (StudioSessionMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studioSessionMappingArgs)(nil)).Elem()
}

type StudioSessionMappingInput interface {
	pulumi.Input

	ToStudioSessionMappingOutput() StudioSessionMappingOutput
	ToStudioSessionMappingOutputWithContext(ctx context.Context) StudioSessionMappingOutput
}

func (*StudioSessionMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioSessionMapping)(nil)).Elem()
}

func (i *StudioSessionMapping) ToStudioSessionMappingOutput() StudioSessionMappingOutput {
	return i.ToStudioSessionMappingOutputWithContext(context.Background())
}

func (i *StudioSessionMapping) ToStudioSessionMappingOutputWithContext(ctx context.Context) StudioSessionMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioSessionMappingOutput)
}

type StudioSessionMappingOutput struct{ *pulumi.OutputState }

func (StudioSessionMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioSessionMapping)(nil)).Elem()
}

func (o StudioSessionMappingOutput) ToStudioSessionMappingOutput() StudioSessionMappingOutput {
	return o
}

func (o StudioSessionMappingOutput) ToStudioSessionMappingOutputWithContext(ctx context.Context) StudioSessionMappingOutput {
	return o
}

// The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
func (o StudioSessionMappingOutput) IdentityName() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioSessionMapping) pulumi.StringOutput { return v.IdentityName }).(pulumi.StringOutput)
}

// Specifies whether the identity to map to the Studio is a user or a group.
func (o StudioSessionMappingOutput) IdentityType() StudioSessionMappingIdentityTypeOutput {
	return o.ApplyT(func(v *StudioSessionMapping) StudioSessionMappingIdentityTypeOutput { return v.IdentityType }).(StudioSessionMappingIdentityTypeOutput)
}

// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles.
func (o StudioSessionMappingOutput) SessionPolicyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioSessionMapping) pulumi.StringOutput { return v.SessionPolicyArn }).(pulumi.StringOutput)
}

// The ID of the Amazon EMR Studio to which the user or group will be mapped.
func (o StudioSessionMappingOutput) StudioId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioSessionMapping) pulumi.StringOutput { return v.StudioId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StudioSessionMappingInput)(nil)).Elem(), &StudioSessionMapping{})
	pulumi.RegisterOutputType(StudioSessionMappingOutput{})
}
