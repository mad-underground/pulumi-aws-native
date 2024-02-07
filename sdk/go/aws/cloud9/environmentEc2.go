// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud9

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cloud9::EnvironmentEC2
//
// Deprecated: EnvironmentEc2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type EnvironmentEc2 struct {
	pulumi.CustomResourceState

	Arn                      pulumi.StringOutput                 `pulumi:"arn"`
	AutomaticStopTimeMinutes pulumi.IntPtrOutput                 `pulumi:"automaticStopTimeMinutes"`
	ConnectionType           pulumi.StringPtrOutput              `pulumi:"connectionType"`
	Description              pulumi.StringPtrOutput              `pulumi:"description"`
	ImageId                  pulumi.StringOutput                 `pulumi:"imageId"`
	InstanceType             pulumi.StringOutput                 `pulumi:"instanceType"`
	Name                     pulumi.StringPtrOutput              `pulumi:"name"`
	OwnerArn                 pulumi.StringPtrOutput              `pulumi:"ownerArn"`
	Repositories             EnvironmentEc2RepositoryArrayOutput `pulumi:"repositories"`
	SubnetId                 pulumi.StringPtrOutput              `pulumi:"subnetId"`
	Tags                     EnvironmentEc2TagArrayOutput        `pulumi:"tags"`
}

// NewEnvironmentEc2 registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentEc2(ctx *pulumi.Context,
	name string, args *EnvironmentEc2Args, opts ...pulumi.ResourceOption) (*EnvironmentEc2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"automaticStopTimeMinutes",
		"connectionType",
		"imageId",
		"instanceType",
		"ownerArn",
		"repositories[*]",
		"subnetId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnvironmentEc2
	err := ctx.RegisterResource("aws-native:cloud9:EnvironmentEc2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentEc2 gets an existing EnvironmentEc2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentEc2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentEc2State, opts ...pulumi.ResourceOption) (*EnvironmentEc2, error) {
	var resource EnvironmentEc2
	err := ctx.ReadResource("aws-native:cloud9:EnvironmentEc2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentEc2 resources.
type environmentEc2State struct {
}

type EnvironmentEc2State struct {
}

func (EnvironmentEc2State) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentEc2State)(nil)).Elem()
}

type environmentEc2Args struct {
	AutomaticStopTimeMinutes *int                       `pulumi:"automaticStopTimeMinutes"`
	ConnectionType           *string                    `pulumi:"connectionType"`
	Description              *string                    `pulumi:"description"`
	ImageId                  string                     `pulumi:"imageId"`
	InstanceType             string                     `pulumi:"instanceType"`
	Name                     *string                    `pulumi:"name"`
	OwnerArn                 *string                    `pulumi:"ownerArn"`
	Repositories             []EnvironmentEc2Repository `pulumi:"repositories"`
	SubnetId                 *string                    `pulumi:"subnetId"`
	Tags                     []EnvironmentEc2Tag        `pulumi:"tags"`
}

// The set of arguments for constructing a EnvironmentEc2 resource.
type EnvironmentEc2Args struct {
	AutomaticStopTimeMinutes pulumi.IntPtrInput
	ConnectionType           pulumi.StringPtrInput
	Description              pulumi.StringPtrInput
	ImageId                  pulumi.StringInput
	InstanceType             pulumi.StringInput
	Name                     pulumi.StringPtrInput
	OwnerArn                 pulumi.StringPtrInput
	Repositories             EnvironmentEc2RepositoryArrayInput
	SubnetId                 pulumi.StringPtrInput
	Tags                     EnvironmentEc2TagArrayInput
}

func (EnvironmentEc2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentEc2Args)(nil)).Elem()
}

type EnvironmentEc2Input interface {
	pulumi.Input

	ToEnvironmentEc2Output() EnvironmentEc2Output
	ToEnvironmentEc2OutputWithContext(ctx context.Context) EnvironmentEc2Output
}

func (*EnvironmentEc2) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentEc2)(nil)).Elem()
}

func (i *EnvironmentEc2) ToEnvironmentEc2Output() EnvironmentEc2Output {
	return i.ToEnvironmentEc2OutputWithContext(context.Background())
}

func (i *EnvironmentEc2) ToEnvironmentEc2OutputWithContext(ctx context.Context) EnvironmentEc2Output {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEc2Output)
}

type EnvironmentEc2Output struct{ *pulumi.OutputState }

func (EnvironmentEc2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentEc2)(nil)).Elem()
}

func (o EnvironmentEc2Output) ToEnvironmentEc2Output() EnvironmentEc2Output {
	return o
}

func (o EnvironmentEc2Output) ToEnvironmentEc2OutputWithContext(ctx context.Context) EnvironmentEc2Output {
	return o
}

func (o EnvironmentEc2Output) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o EnvironmentEc2Output) AutomaticStopTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.IntPtrOutput { return v.AutomaticStopTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o EnvironmentEc2Output) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringPtrOutput { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

func (o EnvironmentEc2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentEc2Output) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

func (o EnvironmentEc2Output) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

func (o EnvironmentEc2Output) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentEc2Output) OwnerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringPtrOutput { return v.OwnerArn }).(pulumi.StringPtrOutput)
}

func (o EnvironmentEc2Output) Repositories() EnvironmentEc2RepositoryArrayOutput {
	return o.ApplyT(func(v *EnvironmentEc2) EnvironmentEc2RepositoryArrayOutput { return v.Repositories }).(EnvironmentEc2RepositoryArrayOutput)
}

func (o EnvironmentEc2Output) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentEc2) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o EnvironmentEc2Output) Tags() EnvironmentEc2TagArrayOutput {
	return o.ApplyT(func(v *EnvironmentEc2) EnvironmentEc2TagArrayOutput { return v.Tags }).(EnvironmentEc2TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEc2Input)(nil)).Elem(), &EnvironmentEc2{})
	pulumi.RegisterOutputType(EnvironmentEc2Output{})
}
