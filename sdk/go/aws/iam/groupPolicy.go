// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds or updates an inline policy document that is embedded in the specified IAM group.
//
//	A group can also have managed policies attached to it. To attach a managed policy to a group, use [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.
//	For information about the maximum number of inline policies that you can embed in a group, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.
type GroupPolicy struct {
	pulumi.CustomResourceState

	// The name of the group to associate the policy with.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The policy document.
	//  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
	//  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
	//   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
	//   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
	//   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::GroupPolicy` for more information about the expected schema for this property.
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// The name of the policy document.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
}

// NewGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicy(ctx *pulumi.Context,
	name string, args *GroupPolicyArgs, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"groupName",
		"policyName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupPolicy
	err := ctx.RegisterResource("aws-native:iam:GroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicy gets an existing GroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyState, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	var resource GroupPolicy
	err := ctx.ReadResource("aws-native:iam:GroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicy resources.
type groupPolicyState struct {
}

type GroupPolicyState struct {
}

func (GroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyState)(nil)).Elem()
}

type groupPolicyArgs struct {
	// The name of the group to associate the policy with.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
	GroupName string `pulumi:"groupName"`
	// The policy document.
	//  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
	//  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
	//   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
	//   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
	//   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::GroupPolicy` for more information about the expected schema for this property.
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// The name of the policy document.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	PolicyName string `pulumi:"policyName"`
}

// The set of arguments for constructing a GroupPolicy resource.
type GroupPolicyArgs struct {
	// The name of the group to associate the policy with.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
	GroupName pulumi.StringInput
	// The policy document.
	//  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
	//  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
	//   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
	//   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
	//   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::GroupPolicy` for more information about the expected schema for this property.
	PolicyDocument pulumi.Input
	// The name of the policy document.
	//  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	PolicyName pulumi.StringInput
}

func (GroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyArgs)(nil)).Elem()
}

type GroupPolicyInput interface {
	pulumi.Input

	ToGroupPolicyOutput() GroupPolicyOutput
	ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput
}

func (*GroupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicy)(nil)).Elem()
}

func (i *GroupPolicy) ToGroupPolicyOutput() GroupPolicyOutput {
	return i.ToGroupPolicyOutputWithContext(context.Background())
}

func (i *GroupPolicy) ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyOutput)
}

type GroupPolicyOutput struct{ *pulumi.OutputState }

func (GroupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicy)(nil)).Elem()
}

func (o GroupPolicyOutput) ToGroupPolicyOutput() GroupPolicyOutput {
	return o
}

func (o GroupPolicyOutput) ToGroupPolicyOutputWithContext(ctx context.Context) GroupPolicyOutput {
	return o
}

// The name of the group to associate the policy with.
//
//	This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.
func (o GroupPolicyOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicy) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The policy document.
//
//	You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
//	The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
//	 +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
//	 +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
//	 +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::GroupPolicy` for more information about the expected schema for this property.
func (o GroupPolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *GroupPolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

// The name of the policy document.
//
//	This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
func (o GroupPolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyInput)(nil)).Elem(), &GroupPolicy{})
	pulumi.RegisterOutputType(GroupPolicyOutput{})
}
