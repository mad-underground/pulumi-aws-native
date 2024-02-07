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

// Resource Type definition for AWS::WAFRegional::WebACL
//
// Deprecated: WebAcl is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type WebAcl struct {
	pulumi.CustomResourceState

	DefaultAction WebAclActionOutput    `pulumi:"defaultAction"`
	MetricName    pulumi.StringOutput   `pulumi:"metricName"`
	Name          pulumi.StringOutput   `pulumi:"name"`
	Rules         WebAclRuleArrayOutput `pulumi:"rules"`
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
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"metricName",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebAcl
	err := ctx.RegisterResource("aws-native:wafregional:WebAcl", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:wafregional:WebAcl", name, id, state, &resource, opts...)
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
	DefaultAction WebAclAction `pulumi:"defaultAction"`
	MetricName    string       `pulumi:"metricName"`
	Name          *string      `pulumi:"name"`
	Rules         []WebAclRule `pulumi:"rules"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	DefaultAction WebAclActionInput
	MetricName    pulumi.StringInput
	Name          pulumi.StringPtrInput
	Rules         WebAclRuleArrayInput
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

func (o WebAclOutput) DefaultAction() WebAclActionOutput {
	return o.ApplyT(func(v *WebAcl) WebAclActionOutput { return v.DefaultAction }).(WebAclActionOutput)
}

func (o WebAclOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

func (o WebAclOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAcl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAclOutput) Rules() WebAclRuleArrayOutput {
	return o.ApplyT(func(v *WebAcl) WebAclRuleArrayOutput { return v.Rules }).(WebAclRuleArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclInput)(nil)).Elem(), &WebAcl{})
	pulumi.RegisterOutputType(WebAclOutput{})
}
