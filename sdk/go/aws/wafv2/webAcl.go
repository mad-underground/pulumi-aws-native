// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
type WebAcl struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput                 `pulumi:"arn"`
	AssociationConfig    WebAclAssociationConfigPtrOutput    `pulumi:"associationConfig"`
	Capacity             pulumi.IntOutput                    `pulumi:"capacity"`
	CaptchaConfig        WebAclCaptchaConfigPtrOutput        `pulumi:"captchaConfig"`
	ChallengeConfig      WebAclChallengeConfigPtrOutput      `pulumi:"challengeConfig"`
	CustomResponseBodies WebAclCustomResponseBodiesPtrOutput `pulumi:"customResponseBodies"`
	DefaultAction        WebAclDefaultActionOutput           `pulumi:"defaultAction"`
	Description          pulumi.StringPtrOutput              `pulumi:"description"`
	LabelNamespace       pulumi.StringOutput                 `pulumi:"labelNamespace"`
	Name                 pulumi.StringPtrOutput              `pulumi:"name"`
	// Collection of Rules.
	Rules            WebAclRuleArrayOutput        `pulumi:"rules"`
	Scope            WebAclScopeOutput            `pulumi:"scope"`
	Tags             WebAclTagArrayOutput         `pulumi:"tags"`
	TokenDomains     pulumi.StringArrayOutput     `pulumi:"tokenDomains"`
	VisibilityConfig WebAclVisibilityConfigOutput `pulumi:"visibilityConfig"`
}

// NewWebAcl registers a new resource with the given unique name, arguments, and options.
func NewWebAcl(ctx *pulumi.Context,
	name string, args *WebAclArgs, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.VisibilityConfig == nil {
		return nil, errors.New("invalid value for required argument 'VisibilityConfig'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"scope",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebAcl
	err := ctx.RegisterResource("aws-native:wafv2:WebAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAcl gets an existing WebAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclState, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	var resource WebAcl
	err := ctx.ReadResource("aws-native:wafv2:WebAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type webAclState struct {
}

type WebAclState struct {
}

func (WebAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclState)(nil)).Elem()
}

type webAclArgs struct {
	AssociationConfig    *WebAclAssociationConfig    `pulumi:"associationConfig"`
	CaptchaConfig        *WebAclCaptchaConfig        `pulumi:"captchaConfig"`
	ChallengeConfig      *WebAclChallengeConfig      `pulumi:"challengeConfig"`
	CustomResponseBodies *WebAclCustomResponseBodies `pulumi:"customResponseBodies"`
	DefaultAction        WebAclDefaultAction         `pulumi:"defaultAction"`
	Description          *string                     `pulumi:"description"`
	Name                 *string                     `pulumi:"name"`
	// Collection of Rules.
	Rules            []WebAclRule           `pulumi:"rules"`
	Scope            WebAclScope            `pulumi:"scope"`
	Tags             []WebAclTag            `pulumi:"tags"`
	TokenDomains     []string               `pulumi:"tokenDomains"`
	VisibilityConfig WebAclVisibilityConfig `pulumi:"visibilityConfig"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	AssociationConfig    WebAclAssociationConfigPtrInput
	CaptchaConfig        WebAclCaptchaConfigPtrInput
	ChallengeConfig      WebAclChallengeConfigPtrInput
	CustomResponseBodies WebAclCustomResponseBodiesPtrInput
	DefaultAction        WebAclDefaultActionInput
	Description          pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	// Collection of Rules.
	Rules            WebAclRuleArrayInput
	Scope            WebAclScopeInput
	Tags             WebAclTagArrayInput
	TokenDomains     pulumi.StringArrayInput
	VisibilityConfig WebAclVisibilityConfigInput
}

func (WebAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclArgs)(nil)).Elem()
}

type WebAclInput interface {
	pulumi.Input

	ToWebAclOutput() WebAclOutput
	ToWebAclOutputWithContext(ctx context.Context) WebAclOutput
}

func (*WebAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAcl)(nil)).Elem()
}

func (i *WebAcl) ToWebAclOutput() WebAclOutput {
	return i.ToWebAclOutputWithContext(context.Background())
}

func (i *WebAcl) ToWebAclOutputWithContext(ctx context.Context) WebAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclOutput)
}

type WebAclOutput struct{ *pulumi.OutputState }

func (WebAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAcl)(nil)).Elem()
}

func (o WebAclOutput) ToWebAclOutput() WebAclOutput {
	return o
}

func (o WebAclOutput) ToWebAclOutputWithContext(ctx context.Context) WebAclOutput {
	return o
}

func (o WebAclOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o WebAclOutput) AssociationConfig() WebAclAssociationConfigPtrOutput {
	return o.ApplyT(func(v *WebAcl) WebAclAssociationConfigPtrOutput { return v.AssociationConfig }).(WebAclAssociationConfigPtrOutput)
}

func (o WebAclOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.IntOutput { return v.Capacity }).(pulumi.IntOutput)
}

func (o WebAclOutput) CaptchaConfig() WebAclCaptchaConfigPtrOutput {
	return o.ApplyT(func(v *WebAcl) WebAclCaptchaConfigPtrOutput { return v.CaptchaConfig }).(WebAclCaptchaConfigPtrOutput)
}

func (o WebAclOutput) ChallengeConfig() WebAclChallengeConfigPtrOutput {
	return o.ApplyT(func(v *WebAcl) WebAclChallengeConfigPtrOutput { return v.ChallengeConfig }).(WebAclChallengeConfigPtrOutput)
}

func (o WebAclOutput) CustomResponseBodies() WebAclCustomResponseBodiesPtrOutput {
	return o.ApplyT(func(v *WebAcl) WebAclCustomResponseBodiesPtrOutput { return v.CustomResponseBodies }).(WebAclCustomResponseBodiesPtrOutput)
}

func (o WebAclOutput) DefaultAction() WebAclDefaultActionOutput {
	return o.ApplyT(func(v *WebAcl) WebAclDefaultActionOutput { return v.DefaultAction }).(WebAclDefaultActionOutput)
}

func (o WebAclOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebAclOutput) LabelNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringOutput { return v.LabelNamespace }).(pulumi.StringOutput)
}

func (o WebAclOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Collection of Rules.
func (o WebAclOutput) Rules() WebAclRuleArrayOutput {
	return o.ApplyT(func(v *WebAcl) WebAclRuleArrayOutput { return v.Rules }).(WebAclRuleArrayOutput)
}

func (o WebAclOutput) Scope() WebAclScopeOutput {
	return o.ApplyT(func(v *WebAcl) WebAclScopeOutput { return v.Scope }).(WebAclScopeOutput)
}

func (o WebAclOutput) Tags() WebAclTagArrayOutput {
	return o.ApplyT(func(v *WebAcl) WebAclTagArrayOutput { return v.Tags }).(WebAclTagArrayOutput)
}

func (o WebAclOutput) TokenDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringArrayOutput { return v.TokenDomains }).(pulumi.StringArrayOutput)
}

func (o WebAclOutput) VisibilityConfig() WebAclVisibilityConfigOutput {
	return o.ApplyT(func(v *WebAcl) WebAclVisibilityConfigOutput { return v.VisibilityConfig }).(WebAclVisibilityConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclInput)(nil)).Elem(), &WebAcl{})
	pulumi.RegisterOutputType(WebAclOutput{})
}
