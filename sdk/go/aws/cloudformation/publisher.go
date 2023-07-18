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

// Register as a publisher in the CloudFormation Registry.
type Publisher struct {
	pulumi.CustomResourceState

	// Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf
	AcceptTermsAndConditions pulumi.BoolOutput `pulumi:"acceptTermsAndConditions"`
	// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
	ConnectionArn pulumi.StringPtrOutput `pulumi:"connectionArn"`
	// The type of account used as the identity provider when registering this publisher with CloudFormation.
	IdentityProvider PublisherIdentityProviderOutput `pulumi:"identityProvider"`
	// The publisher id assigned by CloudFormation for publishing in this region.
	PublisherId pulumi.StringOutput `pulumi:"publisherId"`
	// The URL to the publisher's profile with the identity provider.
	PublisherProfile pulumi.StringOutput `pulumi:"publisherProfile"`
	// Whether the publisher is verified.
	PublisherStatus PublisherStatusOutput `pulumi:"publisherStatus"`
}

// NewPublisher registers a new resource with the given unique name, arguments, and options.
func NewPublisher(ctx *pulumi.Context,
	name string, args *PublisherArgs, opts ...pulumi.ResourceOption) (*Publisher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AcceptTermsAndConditions == nil {
		return nil, errors.New("invalid value for required argument 'AcceptTermsAndConditions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Publisher
	err := ctx.RegisterResource("aws-native:cloudformation:Publisher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublisher gets an existing Publisher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublisher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublisherState, opts ...pulumi.ResourceOption) (*Publisher, error) {
	var resource Publisher
	err := ctx.ReadResource("aws-native:cloudformation:Publisher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Publisher resources.
type publisherState struct {
}

type PublisherState struct {
}

func (PublisherState) ElementType() reflect.Type {
	return reflect.TypeOf((*publisherState)(nil)).Elem()
}

type publisherArgs struct {
	// Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf
	AcceptTermsAndConditions bool `pulumi:"acceptTermsAndConditions"`
	// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
	ConnectionArn *string `pulumi:"connectionArn"`
}

// The set of arguments for constructing a Publisher resource.
type PublisherArgs struct {
	// Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf
	AcceptTermsAndConditions pulumi.BoolInput
	// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
	ConnectionArn pulumi.StringPtrInput
}

func (PublisherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publisherArgs)(nil)).Elem()
}

type PublisherInput interface {
	pulumi.Input

	ToPublisherOutput() PublisherOutput
	ToPublisherOutputWithContext(ctx context.Context) PublisherOutput
}

func (*Publisher) ElementType() reflect.Type {
	return reflect.TypeOf((**Publisher)(nil)).Elem()
}

func (i *Publisher) ToPublisherOutput() PublisherOutput {
	return i.ToPublisherOutputWithContext(context.Background())
}

func (i *Publisher) ToPublisherOutputWithContext(ctx context.Context) PublisherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherOutput)
}

type PublisherOutput struct{ *pulumi.OutputState }

func (PublisherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Publisher)(nil)).Elem()
}

func (o PublisherOutput) ToPublisherOutput() PublisherOutput {
	return o
}

func (o PublisherOutput) ToPublisherOutputWithContext(ctx context.Context) PublisherOutput {
	return o
}

// Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf
func (o PublisherOutput) AcceptTermsAndConditions() pulumi.BoolOutput {
	return o.ApplyT(func(v *Publisher) pulumi.BoolOutput { return v.AcceptTermsAndConditions }).(pulumi.BoolOutput)
}

// If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.
func (o PublisherOutput) ConnectionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Publisher) pulumi.StringPtrOutput { return v.ConnectionArn }).(pulumi.StringPtrOutput)
}

// The type of account used as the identity provider when registering this publisher with CloudFormation.
func (o PublisherOutput) IdentityProvider() PublisherIdentityProviderOutput {
	return o.ApplyT(func(v *Publisher) PublisherIdentityProviderOutput { return v.IdentityProvider }).(PublisherIdentityProviderOutput)
}

// The publisher id assigned by CloudFormation for publishing in this region.
func (o PublisherOutput) PublisherId() pulumi.StringOutput {
	return o.ApplyT(func(v *Publisher) pulumi.StringOutput { return v.PublisherId }).(pulumi.StringOutput)
}

// The URL to the publisher's profile with the identity provider.
func (o PublisherOutput) PublisherProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Publisher) pulumi.StringOutput { return v.PublisherProfile }).(pulumi.StringOutput)
}

// Whether the publisher is verified.
func (o PublisherOutput) PublisherStatus() PublisherStatusOutput {
	return o.ApplyT(func(v *Publisher) PublisherStatusOutput { return v.PublisherStatus }).(PublisherStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublisherInput)(nil)).Elem(), &Publisher{})
	pulumi.RegisterOutputType(PublisherOutput{})
}
