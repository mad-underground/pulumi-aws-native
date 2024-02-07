// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
type Grant struct {
	pulumi.CustomResourceState

	AllowedOperations pulumi.StringArrayOutput `pulumi:"allowedOperations"`
	// Arn of the grant.
	GrantArn pulumi.StringOutput `pulumi:"grantArn"`
	// Name for the created Grant.
	GrantName pulumi.StringPtrOutput `pulumi:"grantName"`
	// Home region for the created grant.
	HomeRegion pulumi.StringPtrOutput `pulumi:"homeRegion"`
	// License Arn for the grant.
	LicenseArn pulumi.StringPtrOutput   `pulumi:"licenseArn"`
	Principals pulumi.StringArrayOutput `pulumi:"principals"`
	Status     pulumi.StringPtrOutput   `pulumi:"status"`
	// The version of the grant.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewGrant registers a new resource with the given unique name, arguments, and options.
func NewGrant(ctx *pulumi.Context,
	name string, args *GrantArgs, opts ...pulumi.ResourceOption) (*Grant, error) {
	if args == nil {
		args = &GrantArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Grant
	err := ctx.RegisterResource("aws-native:licensemanager:Grant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrant gets an existing Grant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantState, opts ...pulumi.ResourceOption) (*Grant, error) {
	var resource Grant
	err := ctx.ReadResource("aws-native:licensemanager:Grant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Grant resources.
type grantState struct {
}

type GrantState struct {
}

func (GrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantState)(nil)).Elem()
}

type grantArgs struct {
	AllowedOperations []string `pulumi:"allowedOperations"`
	// Name for the created Grant.
	GrantName *string `pulumi:"grantName"`
	// Home region for the created grant.
	HomeRegion *string `pulumi:"homeRegion"`
	// License Arn for the grant.
	LicenseArn *string  `pulumi:"licenseArn"`
	Principals []string `pulumi:"principals"`
	Status     *string  `pulumi:"status"`
}

// The set of arguments for constructing a Grant resource.
type GrantArgs struct {
	AllowedOperations pulumi.StringArrayInput
	// Name for the created Grant.
	GrantName pulumi.StringPtrInput
	// Home region for the created grant.
	HomeRegion pulumi.StringPtrInput
	// License Arn for the grant.
	LicenseArn pulumi.StringPtrInput
	Principals pulumi.StringArrayInput
	Status     pulumi.StringPtrInput
}

func (GrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantArgs)(nil)).Elem()
}

type GrantInput interface {
	pulumi.Input

	ToGrantOutput() GrantOutput
	ToGrantOutputWithContext(ctx context.Context) GrantOutput
}

func (*Grant) ElementType() reflect.Type {
	return reflect.TypeOf((**Grant)(nil)).Elem()
}

func (i *Grant) ToGrantOutput() GrantOutput {
	return i.ToGrantOutputWithContext(context.Background())
}

func (i *Grant) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantOutput)
}

type GrantOutput struct{ *pulumi.OutputState }

func (GrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Grant)(nil)).Elem()
}

func (o GrantOutput) ToGrantOutput() GrantOutput {
	return o
}

func (o GrantOutput) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return o
}

func (o GrantOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringArrayOutput { return v.AllowedOperations }).(pulumi.StringArrayOutput)
}

// Arn of the grant.
func (o GrantOutput) GrantArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringOutput { return v.GrantArn }).(pulumi.StringOutput)
}

// Name for the created Grant.
func (o GrantOutput) GrantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringPtrOutput { return v.GrantName }).(pulumi.StringPtrOutput)
}

// Home region for the created grant.
func (o GrantOutput) HomeRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringPtrOutput { return v.HomeRegion }).(pulumi.StringPtrOutput)
}

// License Arn for the grant.
func (o GrantOutput) LicenseArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringPtrOutput { return v.LicenseArn }).(pulumi.StringPtrOutput)
}

func (o GrantOutput) Principals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringArrayOutput { return v.Principals }).(pulumi.StringArrayOutput)
}

func (o GrantOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The version of the grant.
func (o GrantOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrantInput)(nil)).Elem(), &Grant{})
	pulumi.RegisterOutputType(GrantOutput{})
}
