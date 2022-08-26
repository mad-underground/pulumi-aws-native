// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rolesanywhere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::RolesAnywhere::CRL Resource Type
type CRL struct {
	pulumi.CustomResourceState

	CrlData        pulumi.StringPtrOutput `pulumi:"crlData"`
	CrlId          pulumi.StringOutput    `pulumi:"crlId"`
	Enabled        pulumi.BoolPtrOutput   `pulumi:"enabled"`
	Name           pulumi.StringPtrOutput `pulumi:"name"`
	Tags           CRLTagArrayOutput      `pulumi:"tags"`
	TrustAnchorArn pulumi.StringPtrOutput `pulumi:"trustAnchorArn"`
}

// NewCRL registers a new resource with the given unique name, arguments, and options.
func NewCRL(ctx *pulumi.Context,
	name string, args *CRLArgs, opts ...pulumi.ResourceOption) (*CRL, error) {
	if args == nil {
		args = &CRLArgs{}
	}

	var resource CRL
	err := ctx.RegisterResource("aws-native:rolesanywhere:CRL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCRL gets an existing CRL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCRL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CRLState, opts ...pulumi.ResourceOption) (*CRL, error) {
	var resource CRL
	err := ctx.ReadResource("aws-native:rolesanywhere:CRL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CRL resources.
type crlState struct {
}

type CRLState struct {
}

func (CRLState) ElementType() reflect.Type {
	return reflect.TypeOf((*crlState)(nil)).Elem()
}

type crlArgs struct {
	CrlData        *string  `pulumi:"crlData"`
	Enabled        *bool    `pulumi:"enabled"`
	Name           *string  `pulumi:"name"`
	Tags           []CRLTag `pulumi:"tags"`
	TrustAnchorArn *string  `pulumi:"trustAnchorArn"`
}

// The set of arguments for constructing a CRL resource.
type CRLArgs struct {
	CrlData        pulumi.StringPtrInput
	Enabled        pulumi.BoolPtrInput
	Name           pulumi.StringPtrInput
	Tags           CRLTagArrayInput
	TrustAnchorArn pulumi.StringPtrInput
}

func (CRLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*crlArgs)(nil)).Elem()
}

type CRLInput interface {
	pulumi.Input

	ToCRLOutput() CRLOutput
	ToCRLOutputWithContext(ctx context.Context) CRLOutput
}

func (*CRL) ElementType() reflect.Type {
	return reflect.TypeOf((**CRL)(nil)).Elem()
}

func (i *CRL) ToCRLOutput() CRLOutput {
	return i.ToCRLOutputWithContext(context.Background())
}

func (i *CRL) ToCRLOutputWithContext(ctx context.Context) CRLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CRLOutput)
}

type CRLOutput struct{ *pulumi.OutputState }

func (CRLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CRL)(nil)).Elem()
}

func (o CRLOutput) ToCRLOutput() CRLOutput {
	return o
}

func (o CRLOutput) ToCRLOutputWithContext(ctx context.Context) CRLOutput {
	return o
}

func (o CRLOutput) CrlData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CRL) pulumi.StringPtrOutput { return v.CrlData }).(pulumi.StringPtrOutput)
}

func (o CRLOutput) CrlId() pulumi.StringOutput {
	return o.ApplyT(func(v *CRL) pulumi.StringOutput { return v.CrlId }).(pulumi.StringOutput)
}

func (o CRLOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CRL) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CRLOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CRL) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CRLOutput) Tags() CRLTagArrayOutput {
	return o.ApplyT(func(v *CRL) CRLTagArrayOutput { return v.Tags }).(CRLTagArrayOutput)
}

func (o CRLOutput) TrustAnchorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CRL) pulumi.StringPtrOutput { return v.TrustAnchorArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CRLInput)(nil)).Elem(), &CRL{})
	pulumi.RegisterOutputType(CRLOutput{})
}