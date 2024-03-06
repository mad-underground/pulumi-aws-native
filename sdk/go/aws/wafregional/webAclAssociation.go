// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAFRegional::WebACLAssociation
//
// Deprecated: WebAclAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type WebAclAssociation struct {
	pulumi.CustomResourceState

	AwsId       pulumi.StringOutput `pulumi:"awsId"`
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	WebAclId    pulumi.StringOutput `pulumi:"webAclId"`
}

// NewWebAclAssociation registers a new resource with the given unique name, arguments, and options.
func NewWebAclAssociation(ctx *pulumi.Context,
	name string, args *WebAclAssociationArgs, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	if args.WebAclId == nil {
		return nil, errors.New("invalid value for required argument 'WebAclId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"resourceArn",
		"webAclId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebAclAssociation
	err := ctx.RegisterResource("aws-native:wafregional:WebAclAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAclAssociation gets an existing WebAclAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAclAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclAssociationState, opts ...pulumi.ResourceOption) (*WebAclAssociation, error) {
	var resource WebAclAssociation
	err := ctx.ReadResource("aws-native:wafregional:WebAclAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAclAssociation resources.
type webAclAssociationState struct {
}

type WebAclAssociationState struct {
}

func (WebAclAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationState)(nil)).Elem()
}

type webAclAssociationArgs struct {
	ResourceArn string `pulumi:"resourceArn"`
	WebAclId    string `pulumi:"webAclId"`
}

// The set of arguments for constructing a WebAclAssociation resource.
type WebAclAssociationArgs struct {
	ResourceArn pulumi.StringInput
	WebAclId    pulumi.StringInput
}

func (WebAclAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclAssociationArgs)(nil)).Elem()
}

type WebAclAssociationInput interface {
	pulumi.Input

	ToWebAclAssociationOutput() WebAclAssociationOutput
	ToWebAclAssociationOutputWithContext(ctx context.Context) WebAclAssociationOutput
}

func (*WebAclAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclAssociation)(nil)).Elem()
}

func (i *WebAclAssociation) ToWebAclAssociationOutput() WebAclAssociationOutput {
	return i.ToWebAclAssociationOutputWithContext(context.Background())
}

func (i *WebAclAssociation) ToWebAclAssociationOutputWithContext(ctx context.Context) WebAclAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclAssociationOutput)
}

type WebAclAssociationOutput struct{ *pulumi.OutputState }

func (WebAclAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclAssociation)(nil)).Elem()
}

func (o WebAclAssociationOutput) ToWebAclAssociationOutput() WebAclAssociationOutput {
	return o
}

func (o WebAclAssociationOutput) ToWebAclAssociationOutputWithContext(ctx context.Context) WebAclAssociationOutput {
	return o
}

func (o WebAclAssociationOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAclAssociation) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o WebAclAssociationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAclAssociation) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

func (o WebAclAssociationOutput) WebAclId() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAclAssociation) pulumi.StringOutput { return v.WebAclId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclAssociationInput)(nil)).Elem(), &WebAclAssociation{})
	pulumi.RegisterOutputType(WebAclAssociationOutput{})
}
