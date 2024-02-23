// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorad

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type
type DirectoryRegistration struct {
	pulumi.CustomResourceState

	DirectoryId              pulumi.StringOutput    `pulumi:"directoryId"`
	DirectoryRegistrationArn pulumi.StringOutput    `pulumi:"directoryRegistrationArn"`
	Tags                     pulumi.StringMapOutput `pulumi:"tags"`
}

// NewDirectoryRegistration registers a new resource with the given unique name, arguments, and options.
func NewDirectoryRegistration(ctx *pulumi.Context,
	name string, args *DirectoryRegistrationArgs, opts ...pulumi.ResourceOption) (*DirectoryRegistration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"directoryId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectoryRegistration
	err := ctx.RegisterResource("aws-native:pcaconnectorad:DirectoryRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryRegistration gets an existing DirectoryRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryRegistrationState, opts ...pulumi.ResourceOption) (*DirectoryRegistration, error) {
	var resource DirectoryRegistration
	err := ctx.ReadResource("aws-native:pcaconnectorad:DirectoryRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryRegistration resources.
type directoryRegistrationState struct {
}

type DirectoryRegistrationState struct {
}

func (DirectoryRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRegistrationState)(nil)).Elem()
}

type directoryRegistrationArgs struct {
	DirectoryId string            `pulumi:"directoryId"`
	Tags        map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DirectoryRegistration resource.
type DirectoryRegistrationArgs struct {
	DirectoryId pulumi.StringInput
	Tags        pulumi.StringMapInput
}

func (DirectoryRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryRegistrationArgs)(nil)).Elem()
}

type DirectoryRegistrationInput interface {
	pulumi.Input

	ToDirectoryRegistrationOutput() DirectoryRegistrationOutput
	ToDirectoryRegistrationOutputWithContext(ctx context.Context) DirectoryRegistrationOutput
}

func (*DirectoryRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRegistration)(nil)).Elem()
}

func (i *DirectoryRegistration) ToDirectoryRegistrationOutput() DirectoryRegistrationOutput {
	return i.ToDirectoryRegistrationOutputWithContext(context.Background())
}

func (i *DirectoryRegistration) ToDirectoryRegistrationOutputWithContext(ctx context.Context) DirectoryRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryRegistrationOutput)
}

type DirectoryRegistrationOutput struct{ *pulumi.OutputState }

func (DirectoryRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryRegistration)(nil)).Elem()
}

func (o DirectoryRegistrationOutput) ToDirectoryRegistrationOutput() DirectoryRegistrationOutput {
	return o
}

func (o DirectoryRegistrationOutput) ToDirectoryRegistrationOutputWithContext(ctx context.Context) DirectoryRegistrationOutput {
	return o
}

func (o DirectoryRegistrationOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRegistration) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

func (o DirectoryRegistrationOutput) DirectoryRegistrationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryRegistration) pulumi.StringOutput { return v.DirectoryRegistrationArn }).(pulumi.StringOutput)
}

func (o DirectoryRegistrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DirectoryRegistration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryRegistrationInput)(nil)).Elem(), &DirectoryRegistration{})
	pulumi.RegisterOutputType(DirectoryRegistrationOutput{})
}
