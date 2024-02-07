// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EC2::NetworkInsightsPath
type NetworkInsightsPath struct {
	pulumi.CustomResourceState

	CreatedDate            pulumi.StringOutput                    `pulumi:"createdDate"`
	Destination            pulumi.StringPtrOutput                 `pulumi:"destination"`
	DestinationArn         pulumi.StringOutput                    `pulumi:"destinationArn"`
	DestinationIp          pulumi.StringPtrOutput                 `pulumi:"destinationIp"`
	DestinationPort        pulumi.IntPtrOutput                    `pulumi:"destinationPort"`
	FilterAtDestination    NetworkInsightsPathPathFilterPtrOutput `pulumi:"filterAtDestination"`
	FilterAtSource         NetworkInsightsPathPathFilterPtrOutput `pulumi:"filterAtSource"`
	NetworkInsightsPathArn pulumi.StringOutput                    `pulumi:"networkInsightsPathArn"`
	NetworkInsightsPathId  pulumi.StringOutput                    `pulumi:"networkInsightsPathId"`
	Protocol               NetworkInsightsPathProtocolOutput      `pulumi:"protocol"`
	Source                 pulumi.StringOutput                    `pulumi:"source"`
	SourceArn              pulumi.StringOutput                    `pulumi:"sourceArn"`
	SourceIp               pulumi.StringPtrOutput                 `pulumi:"sourceIp"`
	Tags                   NetworkInsightsPathTagArrayOutput      `pulumi:"tags"`
}

// NewNetworkInsightsPath registers a new resource with the given unique name, arguments, and options.
func NewNetworkInsightsPath(ctx *pulumi.Context,
	name string, args *NetworkInsightsPathArgs, opts ...pulumi.ResourceOption) (*NetworkInsightsPath, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"destination",
		"destinationIp",
		"destinationPort",
		"filterAtDestination",
		"filterAtSource",
		"protocol",
		"source",
		"sourceIp",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkInsightsPath
	err := ctx.RegisterResource("aws-native:ec2:NetworkInsightsPath", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInsightsPath gets an existing NetworkInsightsPath resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInsightsPath(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInsightsPathState, opts ...pulumi.ResourceOption) (*NetworkInsightsPath, error) {
	var resource NetworkInsightsPath
	err := ctx.ReadResource("aws-native:ec2:NetworkInsightsPath", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInsightsPath resources.
type networkInsightsPathState struct {
}

type NetworkInsightsPathState struct {
}

func (NetworkInsightsPathState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsPathState)(nil)).Elem()
}

type networkInsightsPathArgs struct {
	Destination         *string                        `pulumi:"destination"`
	DestinationIp       *string                        `pulumi:"destinationIp"`
	DestinationPort     *int                           `pulumi:"destinationPort"`
	FilterAtDestination *NetworkInsightsPathPathFilter `pulumi:"filterAtDestination"`
	FilterAtSource      *NetworkInsightsPathPathFilter `pulumi:"filterAtSource"`
	Protocol            NetworkInsightsPathProtocol    `pulumi:"protocol"`
	Source              string                         `pulumi:"source"`
	SourceIp            *string                        `pulumi:"sourceIp"`
	Tags                []NetworkInsightsPathTag       `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInsightsPath resource.
type NetworkInsightsPathArgs struct {
	Destination         pulumi.StringPtrInput
	DestinationIp       pulumi.StringPtrInput
	DestinationPort     pulumi.IntPtrInput
	FilterAtDestination NetworkInsightsPathPathFilterPtrInput
	FilterAtSource      NetworkInsightsPathPathFilterPtrInput
	Protocol            NetworkInsightsPathProtocolInput
	Source              pulumi.StringInput
	SourceIp            pulumi.StringPtrInput
	Tags                NetworkInsightsPathTagArrayInput
}

func (NetworkInsightsPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsPathArgs)(nil)).Elem()
}

type NetworkInsightsPathInput interface {
	pulumi.Input

	ToNetworkInsightsPathOutput() NetworkInsightsPathOutput
	ToNetworkInsightsPathOutputWithContext(ctx context.Context) NetworkInsightsPathOutput
}

func (*NetworkInsightsPath) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsPath)(nil)).Elem()
}

func (i *NetworkInsightsPath) ToNetworkInsightsPathOutput() NetworkInsightsPathOutput {
	return i.ToNetworkInsightsPathOutputWithContext(context.Background())
}

func (i *NetworkInsightsPath) ToNetworkInsightsPathOutputWithContext(ctx context.Context) NetworkInsightsPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsPathOutput)
}

type NetworkInsightsPathOutput struct{ *pulumi.OutputState }

func (NetworkInsightsPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsPath)(nil)).Elem()
}

func (o NetworkInsightsPathOutput) ToNetworkInsightsPathOutput() NetworkInsightsPathOutput {
	return o
}

func (o NetworkInsightsPathOutput) ToNetworkInsightsPathOutputWithContext(ctx context.Context) NetworkInsightsPathOutput {
	return o
}

func (o NetworkInsightsPathOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o NetworkInsightsPathOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringPtrOutput { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o NetworkInsightsPathOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.DestinationArn }).(pulumi.StringOutput)
}

func (o NetworkInsightsPathOutput) DestinationIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringPtrOutput { return v.DestinationIp }).(pulumi.StringPtrOutput)
}

func (o NetworkInsightsPathOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.IntPtrOutput { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o NetworkInsightsPathOutput) FilterAtDestination() NetworkInsightsPathPathFilterPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) NetworkInsightsPathPathFilterPtrOutput { return v.FilterAtDestination }).(NetworkInsightsPathPathFilterPtrOutput)
}

func (o NetworkInsightsPathOutput) FilterAtSource() NetworkInsightsPathPathFilterPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) NetworkInsightsPathPathFilterPtrOutput { return v.FilterAtSource }).(NetworkInsightsPathPathFilterPtrOutput)
}

func (o NetworkInsightsPathOutput) NetworkInsightsPathArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.NetworkInsightsPathArn }).(pulumi.StringOutput)
}

func (o NetworkInsightsPathOutput) NetworkInsightsPathId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.NetworkInsightsPathId }).(pulumi.StringOutput)
}

func (o NetworkInsightsPathOutput) Protocol() NetworkInsightsPathProtocolOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) NetworkInsightsPathProtocolOutput { return v.Protocol }).(NetworkInsightsPathProtocolOutput)
}

func (o NetworkInsightsPathOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o NetworkInsightsPathOutput) SourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.SourceArn }).(pulumi.StringOutput)
}

func (o NetworkInsightsPathOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringPtrOutput { return v.SourceIp }).(pulumi.StringPtrOutput)
}

func (o NetworkInsightsPathOutput) Tags() NetworkInsightsPathTagArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) NetworkInsightsPathTagArrayOutput { return v.Tags }).(NetworkInsightsPathTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsPathInput)(nil)).Elem(), &NetworkInsightsPath{})
	pulumi.RegisterOutputType(NetworkInsightsPathOutput{})
}
