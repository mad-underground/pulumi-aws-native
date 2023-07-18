// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::DevEndpoint
//
// Deprecated: DevEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DevEndpoint struct {
	pulumi.CustomResourceState

	Arguments             pulumi.AnyOutput         `pulumi:"arguments"`
	EndpointName          pulumi.StringPtrOutput   `pulumi:"endpointName"`
	ExtraJarsS3Path       pulumi.StringPtrOutput   `pulumi:"extraJarsS3Path"`
	ExtraPythonLibsS3Path pulumi.StringPtrOutput   `pulumi:"extraPythonLibsS3Path"`
	GlueVersion           pulumi.StringPtrOutput   `pulumi:"glueVersion"`
	NumberOfNodes         pulumi.IntPtrOutput      `pulumi:"numberOfNodes"`
	NumberOfWorkers       pulumi.IntPtrOutput      `pulumi:"numberOfWorkers"`
	PublicKey             pulumi.StringPtrOutput   `pulumi:"publicKey"`
	PublicKeys            pulumi.StringArrayOutput `pulumi:"publicKeys"`
	RoleArn               pulumi.StringOutput      `pulumi:"roleArn"`
	SecurityConfiguration pulumi.StringPtrOutput   `pulumi:"securityConfiguration"`
	SecurityGroupIds      pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	SubnetId              pulumi.StringPtrOutput   `pulumi:"subnetId"`
	Tags                  pulumi.AnyOutput         `pulumi:"tags"`
	WorkerType            pulumi.StringPtrOutput   `pulumi:"workerType"`
}

// NewDevEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDevEndpoint(ctx *pulumi.Context,
	name string, args *DevEndpointArgs, opts ...pulumi.ResourceOption) (*DevEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DevEndpoint
	err := ctx.RegisterResource("aws-native:glue:DevEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevEndpoint gets an existing DevEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevEndpointState, opts ...pulumi.ResourceOption) (*DevEndpoint, error) {
	var resource DevEndpoint
	err := ctx.ReadResource("aws-native:glue:DevEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevEndpoint resources.
type devEndpointState struct {
}

type DevEndpointState struct {
}

func (DevEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*devEndpointState)(nil)).Elem()
}

type devEndpointArgs struct {
	Arguments             interface{} `pulumi:"arguments"`
	EndpointName          *string     `pulumi:"endpointName"`
	ExtraJarsS3Path       *string     `pulumi:"extraJarsS3Path"`
	ExtraPythonLibsS3Path *string     `pulumi:"extraPythonLibsS3Path"`
	GlueVersion           *string     `pulumi:"glueVersion"`
	NumberOfNodes         *int        `pulumi:"numberOfNodes"`
	NumberOfWorkers       *int        `pulumi:"numberOfWorkers"`
	PublicKey             *string     `pulumi:"publicKey"`
	PublicKeys            []string    `pulumi:"publicKeys"`
	RoleArn               string      `pulumi:"roleArn"`
	SecurityConfiguration *string     `pulumi:"securityConfiguration"`
	SecurityGroupIds      []string    `pulumi:"securityGroupIds"`
	SubnetId              *string     `pulumi:"subnetId"`
	Tags                  interface{} `pulumi:"tags"`
	WorkerType            *string     `pulumi:"workerType"`
}

// The set of arguments for constructing a DevEndpoint resource.
type DevEndpointArgs struct {
	Arguments             pulumi.Input
	EndpointName          pulumi.StringPtrInput
	ExtraJarsS3Path       pulumi.StringPtrInput
	ExtraPythonLibsS3Path pulumi.StringPtrInput
	GlueVersion           pulumi.StringPtrInput
	NumberOfNodes         pulumi.IntPtrInput
	NumberOfWorkers       pulumi.IntPtrInput
	PublicKey             pulumi.StringPtrInput
	PublicKeys            pulumi.StringArrayInput
	RoleArn               pulumi.StringInput
	SecurityConfiguration pulumi.StringPtrInput
	SecurityGroupIds      pulumi.StringArrayInput
	SubnetId              pulumi.StringPtrInput
	Tags                  pulumi.Input
	WorkerType            pulumi.StringPtrInput
}

func (DevEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devEndpointArgs)(nil)).Elem()
}

type DevEndpointInput interface {
	pulumi.Input

	ToDevEndpointOutput() DevEndpointOutput
	ToDevEndpointOutputWithContext(ctx context.Context) DevEndpointOutput
}

func (*DevEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**DevEndpoint)(nil)).Elem()
}

func (i *DevEndpoint) ToDevEndpointOutput() DevEndpointOutput {
	return i.ToDevEndpointOutputWithContext(context.Background())
}

func (i *DevEndpoint) ToDevEndpointOutputWithContext(ctx context.Context) DevEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevEndpointOutput)
}

type DevEndpointOutput struct{ *pulumi.OutputState }

func (DevEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevEndpoint)(nil)).Elem()
}

func (o DevEndpointOutput) ToDevEndpointOutput() DevEndpointOutput {
	return o
}

func (o DevEndpointOutput) ToDevEndpointOutputWithContext(ctx context.Context) DevEndpointOutput {
	return o
}

func (o DevEndpointOutput) Arguments() pulumi.AnyOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.AnyOutput { return v.Arguments }).(pulumi.AnyOutput)
}

func (o DevEndpointOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) ExtraJarsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.ExtraJarsS3Path }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) ExtraPythonLibsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.ExtraPythonLibsS3Path }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) NumberOfNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.IntPtrOutput { return v.NumberOfNodes }).(pulumi.IntPtrOutput)
}

func (o DevEndpointOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.IntPtrOutput { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o DevEndpointOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) PublicKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringArrayOutput { return v.PublicKeys }).(pulumi.StringArrayOutput)
}

func (o DevEndpointOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o DevEndpointOutput) SecurityConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.SecurityConfiguration }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o DevEndpointOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o DevEndpointOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o DevEndpointOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevEndpoint) pulumi.StringPtrOutput { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevEndpointInput)(nil)).Elem(), &DevEndpoint{})
	pulumi.RegisterOutputType(DevEndpointOutput{})
}
