// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourceexplorer2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::ResourceExplorer2::View Resource Type
type View struct {
	pulumi.CustomResourceState

	Filters            ViewFiltersPtrOutput            `pulumi:"filters"`
	IncludedProperties ViewIncludedPropertyArrayOutput `pulumi:"includedProperties"`
	Tags               ViewTagMapPtrOutput             `pulumi:"tags"`
	ViewArn            pulumi.StringOutput             `pulumi:"viewArn"`
	ViewName           pulumi.StringOutput             `pulumi:"viewName"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		args = &ViewArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource View
	err := ctx.RegisterResource("aws-native:resourceexplorer2:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetView gets an existing View resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("aws-native:resourceexplorer2:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering View resources.
type viewState struct {
}

type ViewState struct {
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	Filters            *ViewFilters           `pulumi:"filters"`
	IncludedProperties []ViewIncludedProperty `pulumi:"includedProperties"`
	Tags               *ViewTagMap            `pulumi:"tags"`
	ViewName           *string                `pulumi:"viewName"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	Filters            ViewFiltersPtrInput
	IncludedProperties ViewIncludedPropertyArrayInput
	Tags               ViewTagMapPtrInput
	ViewName           pulumi.StringPtrInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

func (o ViewOutput) Filters() ViewFiltersPtrOutput {
	return o.ApplyT(func(v *View) ViewFiltersPtrOutput { return v.Filters }).(ViewFiltersPtrOutput)
}

func (o ViewOutput) IncludedProperties() ViewIncludedPropertyArrayOutput {
	return o.ApplyT(func(v *View) ViewIncludedPropertyArrayOutput { return v.IncludedProperties }).(ViewIncludedPropertyArrayOutput)
}

func (o ViewOutput) Tags() ViewTagMapPtrOutput {
	return o.ApplyT(func(v *View) ViewTagMapPtrOutput { return v.Tags }).(ViewTagMapPtrOutput)
}

func (o ViewOutput) ViewArn() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ViewArn }).(pulumi.StringOutput)
}

func (o ViewOutput) ViewName() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ViewName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ViewInput)(nil)).Elem(), &View{})
	pulumi.RegisterOutputType(ViewOutput{})
}
