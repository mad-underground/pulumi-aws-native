// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SES::EmailIdentity
type EmailIdentity struct {
	pulumi.CustomResourceState

	ConfigurationSetAttributes EmailIdentityConfigurationSetAttributesPtrOutput `pulumi:"configurationSetAttributes"`
	DkimAttributes             EmailIdentityDkimAttributesPtrOutput             `pulumi:"dkimAttributes"`
	DkimDNSTokenName1          pulumi.StringOutput                              `pulumi:"dkimDNSTokenName1"`
	DkimDNSTokenName2          pulumi.StringOutput                              `pulumi:"dkimDNSTokenName2"`
	DkimDNSTokenName3          pulumi.StringOutput                              `pulumi:"dkimDNSTokenName3"`
	DkimDNSTokenValue1         pulumi.StringOutput                              `pulumi:"dkimDNSTokenValue1"`
	DkimDNSTokenValue2         pulumi.StringOutput                              `pulumi:"dkimDNSTokenValue2"`
	DkimDNSTokenValue3         pulumi.StringOutput                              `pulumi:"dkimDNSTokenValue3"`
	DkimSigningAttributes      EmailIdentityDkimSigningAttributesPtrOutput      `pulumi:"dkimSigningAttributes"`
	// The email address or domain to verify.
	EmailIdentity      pulumi.StringOutput                      `pulumi:"emailIdentity"`
	FeedbackAttributes EmailIdentityFeedbackAttributesPtrOutput `pulumi:"feedbackAttributes"`
	MailFromAttributes EmailIdentityMailFromAttributesPtrOutput `pulumi:"mailFromAttributes"`
}

// NewEmailIdentity registers a new resource with the given unique name, arguments, and options.
func NewEmailIdentity(ctx *pulumi.Context,
	name string, args *EmailIdentityArgs, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EmailIdentity == nil {
		return nil, errors.New("invalid value for required argument 'EmailIdentity'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EmailIdentity
	err := ctx.RegisterResource("aws-native:ses:EmailIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailIdentity gets an existing EmailIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailIdentityState, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	var resource EmailIdentity
	err := ctx.ReadResource("aws-native:ses:EmailIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailIdentity resources.
type emailIdentityState struct {
}

type EmailIdentityState struct {
}

func (EmailIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityState)(nil)).Elem()
}

type emailIdentityArgs struct {
	ConfigurationSetAttributes *EmailIdentityConfigurationSetAttributes `pulumi:"configurationSetAttributes"`
	DkimAttributes             *EmailIdentityDkimAttributes             `pulumi:"dkimAttributes"`
	DkimSigningAttributes      *EmailIdentityDkimSigningAttributes      `pulumi:"dkimSigningAttributes"`
	// The email address or domain to verify.
	EmailIdentity      string                           `pulumi:"emailIdentity"`
	FeedbackAttributes *EmailIdentityFeedbackAttributes `pulumi:"feedbackAttributes"`
	MailFromAttributes *EmailIdentityMailFromAttributes `pulumi:"mailFromAttributes"`
}

// The set of arguments for constructing a EmailIdentity resource.
type EmailIdentityArgs struct {
	ConfigurationSetAttributes EmailIdentityConfigurationSetAttributesPtrInput
	DkimAttributes             EmailIdentityDkimAttributesPtrInput
	DkimSigningAttributes      EmailIdentityDkimSigningAttributesPtrInput
	// The email address or domain to verify.
	EmailIdentity      pulumi.StringInput
	FeedbackAttributes EmailIdentityFeedbackAttributesPtrInput
	MailFromAttributes EmailIdentityMailFromAttributesPtrInput
}

func (EmailIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityArgs)(nil)).Elem()
}

type EmailIdentityInput interface {
	pulumi.Input

	ToEmailIdentityOutput() EmailIdentityOutput
	ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput
}

func (*EmailIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailIdentity)(nil)).Elem()
}

func (i *EmailIdentity) ToEmailIdentityOutput() EmailIdentityOutput {
	return i.ToEmailIdentityOutputWithContext(context.Background())
}

func (i *EmailIdentity) ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityOutput)
}

type EmailIdentityOutput struct{ *pulumi.OutputState }

func (EmailIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailIdentity)(nil)).Elem()
}

func (o EmailIdentityOutput) ToEmailIdentityOutput() EmailIdentityOutput {
	return o
}

func (o EmailIdentityOutput) ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput {
	return o
}

func (o EmailIdentityOutput) ConfigurationSetAttributes() EmailIdentityConfigurationSetAttributesPtrOutput {
	return o.ApplyT(func(v *EmailIdentity) EmailIdentityConfigurationSetAttributesPtrOutput {
		return v.ConfigurationSetAttributes
	}).(EmailIdentityConfigurationSetAttributesPtrOutput)
}

func (o EmailIdentityOutput) DkimAttributes() EmailIdentityDkimAttributesPtrOutput {
	return o.ApplyT(func(v *EmailIdentity) EmailIdentityDkimAttributesPtrOutput { return v.DkimAttributes }).(EmailIdentityDkimAttributesPtrOutput)
}

func (o EmailIdentityOutput) DkimDNSTokenName1() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.DkimDNSTokenName1 }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) DkimDNSTokenName2() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.DkimDNSTokenName2 }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) DkimDNSTokenName3() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.DkimDNSTokenName3 }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) DkimDNSTokenValue1() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.DkimDNSTokenValue1 }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) DkimDNSTokenValue2() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.DkimDNSTokenValue2 }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) DkimDNSTokenValue3() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.DkimDNSTokenValue3 }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) DkimSigningAttributes() EmailIdentityDkimSigningAttributesPtrOutput {
	return o.ApplyT(func(v *EmailIdentity) EmailIdentityDkimSigningAttributesPtrOutput { return v.DkimSigningAttributes }).(EmailIdentityDkimSigningAttributesPtrOutput)
}

// The email address or domain to verify.
func (o EmailIdentityOutput) EmailIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailIdentity) pulumi.StringOutput { return v.EmailIdentity }).(pulumi.StringOutput)
}

func (o EmailIdentityOutput) FeedbackAttributes() EmailIdentityFeedbackAttributesPtrOutput {
	return o.ApplyT(func(v *EmailIdentity) EmailIdentityFeedbackAttributesPtrOutput { return v.FeedbackAttributes }).(EmailIdentityFeedbackAttributesPtrOutput)
}

func (o EmailIdentityOutput) MailFromAttributes() EmailIdentityMailFromAttributesPtrOutput {
	return o.ApplyT(func(v *EmailIdentity) EmailIdentityMailFromAttributesPtrOutput { return v.MailFromAttributes }).(EmailIdentityMailFromAttributesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailIdentityInput)(nil)).Elem(), &EmailIdentity{})
	pulumi.RegisterOutputType(EmailIdentityOutput{})
}
