// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolDomain
type UserPoolDomain struct {
	pulumi.CustomResourceState

	AwsId                  pulumi.StringOutput                           `pulumi:"awsId"`
	CloudFrontDistribution pulumi.StringOutput                           `pulumi:"cloudFrontDistribution"`
	CustomDomainConfig     UserPoolDomainCustomDomainConfigTypePtrOutput `pulumi:"customDomainConfig"`
	Domain                 pulumi.StringOutput                           `pulumi:"domain"`
	UserPoolId             pulumi.StringOutput                           `pulumi:"userPoolId"`
}

// NewUserPoolDomain registers a new resource with the given unique name, arguments, and options.
func NewUserPoolDomain(ctx *pulumi.Context,
	name string, args *UserPoolDomainArgs, opts ...pulumi.ResourceOption) (*UserPoolDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domain",
		"userPoolId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPoolDomain
	err := ctx.RegisterResource("aws-native:cognito:UserPoolDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPoolDomain gets an existing UserPoolDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPoolDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPoolDomainState, opts ...pulumi.ResourceOption) (*UserPoolDomain, error) {
	var resource UserPoolDomain
	err := ctx.ReadResource("aws-native:cognito:UserPoolDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPoolDomain resources.
type userPoolDomainState struct {
}

type UserPoolDomainState struct {
}

func (UserPoolDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolDomainState)(nil)).Elem()
}

type userPoolDomainArgs struct {
	CustomDomainConfig *UserPoolDomainCustomDomainConfigType `pulumi:"customDomainConfig"`
	Domain             string                                `pulumi:"domain"`
	UserPoolId         string                                `pulumi:"userPoolId"`
}

// The set of arguments for constructing a UserPoolDomain resource.
type UserPoolDomainArgs struct {
	CustomDomainConfig UserPoolDomainCustomDomainConfigTypePtrInput
	Domain             pulumi.StringInput
	UserPoolId         pulumi.StringInput
}

func (UserPoolDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolDomainArgs)(nil)).Elem()
}

type UserPoolDomainInput interface {
	pulumi.Input

	ToUserPoolDomainOutput() UserPoolDomainOutput
	ToUserPoolDomainOutputWithContext(ctx context.Context) UserPoolDomainOutput
}

func (*UserPoolDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPoolDomain)(nil)).Elem()
}

func (i *UserPoolDomain) ToUserPoolDomainOutput() UserPoolDomainOutput {
	return i.ToUserPoolDomainOutputWithContext(context.Background())
}

func (i *UserPoolDomain) ToUserPoolDomainOutputWithContext(ctx context.Context) UserPoolDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolDomainOutput)
}

type UserPoolDomainOutput struct{ *pulumi.OutputState }

func (UserPoolDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPoolDomain)(nil)).Elem()
}

func (o UserPoolDomainOutput) ToUserPoolDomainOutput() UserPoolDomainOutput {
	return o
}

func (o UserPoolDomainOutput) ToUserPoolDomainOutputWithContext(ctx context.Context) UserPoolDomainOutput {
	return o
}

func (o UserPoolDomainOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPoolDomain) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o UserPoolDomainOutput) CloudFrontDistribution() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPoolDomain) pulumi.StringOutput { return v.CloudFrontDistribution }).(pulumi.StringOutput)
}

func (o UserPoolDomainOutput) CustomDomainConfig() UserPoolDomainCustomDomainConfigTypePtrOutput {
	return o.ApplyT(func(v *UserPoolDomain) UserPoolDomainCustomDomainConfigTypePtrOutput { return v.CustomDomainConfig }).(UserPoolDomainCustomDomainConfigTypePtrOutput)
}

func (o UserPoolDomainOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPoolDomain) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

func (o UserPoolDomainOutput) UserPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPoolDomain) pulumi.StringOutput { return v.UserPoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPoolDomainInput)(nil)).Elem(), &UserPoolDomain{})
	pulumi.RegisterOutputType(UserPoolDomainOutput{})
}
