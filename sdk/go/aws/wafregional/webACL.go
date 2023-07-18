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
// Deprecated: WebACL is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type WebACL struct {
	pulumi.CustomResourceState

	DefaultAction WebACLActionOutput    `pulumi:"defaultAction"`
	MetricName    pulumi.StringOutput   `pulumi:"metricName"`
	Name          pulumi.StringOutput   `pulumi:"name"`
	Rules         WebACLRuleArrayOutput `pulumi:"rules"`
}

// NewWebACL registers a new resource with the given unique name, arguments, and options.
func NewWebACL(ctx *pulumi.Context,
	name string, args *WebACLArgs, opts ...pulumi.ResourceOption) (*WebACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebACL
	err := ctx.RegisterResource("aws-native:wafregional:WebACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebACL gets an existing WebACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebACLState, opts ...pulumi.ResourceOption) (*WebACL, error) {
	var resource WebACL
	err := ctx.ReadResource("aws-native:wafregional:WebACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebACL resources.
type webACLState struct {
}

type WebACLState struct {
}

func (WebACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLState)(nil)).Elem()
}

type webACLArgs struct {
	DefaultAction WebACLAction `pulumi:"defaultAction"`
	MetricName    string       `pulumi:"metricName"`
	Name          *string      `pulumi:"name"`
	Rules         []WebACLRule `pulumi:"rules"`
}

// The set of arguments for constructing a WebACL resource.
type WebACLArgs struct {
	DefaultAction WebACLActionInput
	MetricName    pulumi.StringInput
	Name          pulumi.StringPtrInput
	Rules         WebACLRuleArrayInput
}

func (WebACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLArgs)(nil)).Elem()
}

type WebACLInput interface {
	pulumi.Input

	ToWebACLOutput() WebACLOutput
	ToWebACLOutputWithContext(ctx context.Context) WebACLOutput
}

func (*WebACL) ElementType() reflect.Type {
	return reflect.TypeOf((**WebACL)(nil)).Elem()
}

func (i *WebACL) ToWebACLOutput() WebACLOutput {
	return i.ToWebACLOutputWithContext(context.Background())
}

func (i *WebACL) ToWebACLOutputWithContext(ctx context.Context) WebACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebACLOutput)
}

type WebACLOutput struct{ *pulumi.OutputState }

func (WebACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebACL)(nil)).Elem()
}

func (o WebACLOutput) ToWebACLOutput() WebACLOutput {
	return o
}

func (o WebACLOutput) ToWebACLOutputWithContext(ctx context.Context) WebACLOutput {
	return o
}

func (o WebACLOutput) DefaultAction() WebACLActionOutput {
	return o.ApplyT(func(v *WebACL) WebACLActionOutput { return v.DefaultAction }).(WebACLActionOutput)
}

func (o WebACLOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebACL) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

func (o WebACLOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebACL) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebACLOutput) Rules() WebACLRuleArrayOutput {
	return o.ApplyT(func(v *WebACL) WebACLRuleArrayOutput { return v.Rules }).(WebACLRuleArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebACLInput)(nil)).Elem(), &WebACL{})
	pulumi.RegisterOutputType(WebACLOutput{})
}
