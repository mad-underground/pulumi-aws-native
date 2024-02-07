// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::View
type View struct {
	pulumi.CustomResourceState

	// The actions of the view in an array.
	Actions pulumi.StringArrayOutput `pulumi:"actions"`
	// The description of the view.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The name of the view.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more tags.
	Tags ViewTagArrayOutput `pulumi:"tags"`
	// The template of the view as JSON.
	Template pulumi.AnyOutput `pulumi:"template"`
	// The Amazon Resource Name (ARN) of the view.
	ViewArn pulumi.StringOutput `pulumi:"viewArn"`
	// The view content hash.
	ViewContentSha256 pulumi.StringOutput `pulumi:"viewContentSha256"`
	// The view id of the view.
	ViewId pulumi.StringOutput `pulumi:"viewId"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource View
	err := ctx.RegisterResource("aws-native:connect:View", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:connect:View", name, id, state, &resource, opts...)
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
	// The actions of the view in an array.
	Actions []string `pulumi:"actions"`
	// The description of the view.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn string `pulumi:"instanceArn"`
	// The name of the view.
	Name *string `pulumi:"name"`
	// One or more tags.
	Tags []ViewTag `pulumi:"tags"`
	// The template of the view as JSON.
	Template interface{} `pulumi:"template"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	// The actions of the view in an array.
	Actions pulumi.StringArrayInput
	// The description of the view.
	Description pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn pulumi.StringInput
	// The name of the view.
	Name pulumi.StringPtrInput
	// One or more tags.
	Tags ViewTagArrayInput
	// The template of the view as JSON.
	Template pulumi.Input
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

// The actions of the view in an array.
func (o ViewOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *View) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

// The description of the view.
func (o ViewOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the instance.
func (o ViewOutput) InstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.InstanceArn }).(pulumi.StringOutput)
}

// The name of the view.
func (o ViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// One or more tags.
func (o ViewOutput) Tags() ViewTagArrayOutput {
	return o.ApplyT(func(v *View) ViewTagArrayOutput { return v.Tags }).(ViewTagArrayOutput)
}

// The template of the view as JSON.
func (o ViewOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *View) pulumi.AnyOutput { return v.Template }).(pulumi.AnyOutput)
}

// The Amazon Resource Name (ARN) of the view.
func (o ViewOutput) ViewArn() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ViewArn }).(pulumi.StringOutput)
}

// The view content hash.
func (o ViewOutput) ViewContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ViewContentSha256 }).(pulumi.StringOutput)
}

// The view id of the view.
func (o ViewOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ViewInput)(nil)).Elem(), &View{})
	pulumi.RegisterOutputType(ViewOutput{})
}
