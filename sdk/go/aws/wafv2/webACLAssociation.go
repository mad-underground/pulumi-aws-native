// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html
type WebACLAssociation struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-resourcearn
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-webaclarn
	WebACLArn pulumi.StringOutput `pulumi:"webACLArn"`
}

// NewWebACLAssociation registers a new resource with the given unique name, arguments, and options.
func NewWebACLAssociation(ctx *pulumi.Context,
	name string, args *WebACLAssociationArgs, opts ...pulumi.ResourceOption) (*WebACLAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	if args.WebACLArn == nil {
		return nil, errors.New("invalid value for required argument 'WebACLArn'")
	}
	var resource WebACLAssociation
	err := ctx.RegisterResource("aws-native:wafv2:WebACLAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebACLAssociation gets an existing WebACLAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebACLAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebACLAssociationState, opts ...pulumi.ResourceOption) (*WebACLAssociation, error) {
	var resource WebACLAssociation
	err := ctx.ReadResource("aws-native:wafv2:WebACLAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebACLAssociation resources.
type webACLAssociationState struct {
}

type WebACLAssociationState struct {
}

func (WebACLAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLAssociationState)(nil)).Elem()
}

type webACLAssociationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-resourcearn
	ResourceArn string `pulumi:"resourceArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-webaclarn
	WebACLArn string `pulumi:"webACLArn"`
}

// The set of arguments for constructing a WebACLAssociation resource.
type WebACLAssociationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-resourcearn
	ResourceArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webaclassociation.html#cfn-wafv2-webaclassociation-webaclarn
	WebACLArn pulumi.StringInput
}

func (WebACLAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLAssociationArgs)(nil)).Elem()
}

type WebACLAssociationInput interface {
	pulumi.Input

	ToWebACLAssociationOutput() WebACLAssociationOutput
	ToWebACLAssociationOutputWithContext(ctx context.Context) WebACLAssociationOutput
}

func (*WebACLAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*WebACLAssociation)(nil))
}

func (i *WebACLAssociation) ToWebACLAssociationOutput() WebACLAssociationOutput {
	return i.ToWebACLAssociationOutputWithContext(context.Background())
}

func (i *WebACLAssociation) ToWebACLAssociationOutputWithContext(ctx context.Context) WebACLAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebACLAssociationOutput)
}

type WebACLAssociationOutput struct{ *pulumi.OutputState }

func (WebACLAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebACLAssociation)(nil))
}

func (o WebACLAssociationOutput) ToWebACLAssociationOutput() WebACLAssociationOutput {
	return o
}

func (o WebACLAssociationOutput) ToWebACLAssociationOutputWithContext(ctx context.Context) WebACLAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebACLAssociationOutput{})
}