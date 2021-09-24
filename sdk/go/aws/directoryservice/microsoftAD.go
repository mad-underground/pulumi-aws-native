// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DirectoryService::MicrosoftAD
//
// Deprecated: MicrosoftAD is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type MicrosoftAD struct {
	pulumi.CustomResourceState

	Alias          pulumi.StringOutput          `pulumi:"alias"`
	CreateAlias    pulumi.BoolPtrOutput         `pulumi:"createAlias"`
	DnsIpAddresses pulumi.StringArrayOutput     `pulumi:"dnsIpAddresses"`
	Edition        pulumi.StringPtrOutput       `pulumi:"edition"`
	EnableSso      pulumi.BoolPtrOutput         `pulumi:"enableSso"`
	Name           pulumi.StringOutput          `pulumi:"name"`
	Password       pulumi.StringOutput          `pulumi:"password"`
	ShortName      pulumi.StringPtrOutput       `pulumi:"shortName"`
	VpcSettings    MicrosoftADVpcSettingsOutput `pulumi:"vpcSettings"`
}

// NewMicrosoftAD registers a new resource with the given unique name, arguments, and options.
func NewMicrosoftAD(ctx *pulumi.Context,
	name string, args *MicrosoftADArgs, opts ...pulumi.ResourceOption) (*MicrosoftAD, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.VpcSettings == nil {
		return nil, errors.New("invalid value for required argument 'VpcSettings'")
	}
	var resource MicrosoftAD
	err := ctx.RegisterResource("aws-native:directoryservice:MicrosoftAD", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMicrosoftAD gets an existing MicrosoftAD resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMicrosoftAD(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicrosoftADState, opts ...pulumi.ResourceOption) (*MicrosoftAD, error) {
	var resource MicrosoftAD
	err := ctx.ReadResource("aws-native:directoryservice:MicrosoftAD", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MicrosoftAD resources.
type microsoftADState struct {
}

type MicrosoftADState struct {
}

func (MicrosoftADState) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftADState)(nil)).Elem()
}

type microsoftADArgs struct {
	CreateAlias *bool                  `pulumi:"createAlias"`
	Edition     *string                `pulumi:"edition"`
	EnableSso   *bool                  `pulumi:"enableSso"`
	Name        string                 `pulumi:"name"`
	Password    string                 `pulumi:"password"`
	ShortName   *string                `pulumi:"shortName"`
	VpcSettings MicrosoftADVpcSettings `pulumi:"vpcSettings"`
}

// The set of arguments for constructing a MicrosoftAD resource.
type MicrosoftADArgs struct {
	CreateAlias pulumi.BoolPtrInput
	Edition     pulumi.StringPtrInput
	EnableSso   pulumi.BoolPtrInput
	Name        pulumi.StringInput
	Password    pulumi.StringInput
	ShortName   pulumi.StringPtrInput
	VpcSettings MicrosoftADVpcSettingsInput
}

func (MicrosoftADArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftADArgs)(nil)).Elem()
}

type MicrosoftADInput interface {
	pulumi.Input

	ToMicrosoftADOutput() MicrosoftADOutput
	ToMicrosoftADOutputWithContext(ctx context.Context) MicrosoftADOutput
}

func (*MicrosoftAD) ElementType() reflect.Type {
	return reflect.TypeOf((*MicrosoftAD)(nil))
}

func (i *MicrosoftAD) ToMicrosoftADOutput() MicrosoftADOutput {
	return i.ToMicrosoftADOutputWithContext(context.Background())
}

func (i *MicrosoftAD) ToMicrosoftADOutputWithContext(ctx context.Context) MicrosoftADOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftADOutput)
}

type MicrosoftADOutput struct{ *pulumi.OutputState }

func (MicrosoftADOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MicrosoftAD)(nil))
}

func (o MicrosoftADOutput) ToMicrosoftADOutput() MicrosoftADOutput {
	return o
}

func (o MicrosoftADOutput) ToMicrosoftADOutputWithContext(ctx context.Context) MicrosoftADOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MicrosoftADOutput{})
}