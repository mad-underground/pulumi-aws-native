// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::StorageSystem.
type StorageSystem struct {
	pulumi.CustomResourceState

	// The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.
	AgentArns pulumi.StringArrayOutput `pulumi:"agentArns"`
	// The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.
	CloudWatchLogGroupArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogGroupArn"`
	// Indicates whether the DataSync agent can access the on-premises storage system.
	ConnectivityStatus StorageSystemConnectivityStatusOutput `pulumi:"connectivityStatus"`
	// A familiar name for the on-premises storage system.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The ARN of a secret stored by AWS Secrets Manager.
	SecretsManagerArn   pulumi.StringOutput                     `pulumi:"secretsManagerArn"`
	ServerConfiguration StorageSystemServerConfigurationOutput  `pulumi:"serverConfiguration"`
	ServerCredentials   StorageSystemServerCredentialsPtrOutput `pulumi:"serverCredentials"`
	// The ARN of the on-premises storage system added to DataSync Discovery.
	StorageSystemArn pulumi.StringOutput `pulumi:"storageSystemArn"`
	// The type of on-premises storage system that DataSync Discovery will analyze.
	SystemType StorageSystemSystemTypeOutput `pulumi:"systemType"`
	// An array of key-value pairs to apply to this resource.
	Tags StorageSystemTagArrayOutput `pulumi:"tags"`
}

// NewStorageSystem registers a new resource with the given unique name, arguments, and options.
func NewStorageSystem(ctx *pulumi.Context,
	name string, args *StorageSystemArgs, opts ...pulumi.ResourceOption) (*StorageSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentArns == nil {
		return nil, errors.New("invalid value for required argument 'AgentArns'")
	}
	if args.ServerConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'ServerConfiguration'")
	}
	if args.SystemType == nil {
		return nil, errors.New("invalid value for required argument 'SystemType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StorageSystem
	err := ctx.RegisterResource("aws-native:datasync:StorageSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageSystem gets an existing StorageSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageSystemState, opts ...pulumi.ResourceOption) (*StorageSystem, error) {
	var resource StorageSystem
	err := ctx.ReadResource("aws-native:datasync:StorageSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageSystem resources.
type storageSystemState struct {
}

type StorageSystemState struct {
}

func (StorageSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSystemState)(nil)).Elem()
}

type storageSystemArgs struct {
	// The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.
	AgentArns []string `pulumi:"agentArns"`
	// The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.
	CloudWatchLogGroupArn *string `pulumi:"cloudWatchLogGroupArn"`
	// A familiar name for the on-premises storage system.
	Name                *string                          `pulumi:"name"`
	ServerConfiguration StorageSystemServerConfiguration `pulumi:"serverConfiguration"`
	ServerCredentials   *StorageSystemServerCredentials  `pulumi:"serverCredentials"`
	// The type of on-premises storage system that DataSync Discovery will analyze.
	SystemType StorageSystemSystemType `pulumi:"systemType"`
	// An array of key-value pairs to apply to this resource.
	Tags []StorageSystemTag `pulumi:"tags"`
}

// The set of arguments for constructing a StorageSystem resource.
type StorageSystemArgs struct {
	// The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.
	AgentArns pulumi.StringArrayInput
	// The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.
	CloudWatchLogGroupArn pulumi.StringPtrInput
	// A familiar name for the on-premises storage system.
	Name                pulumi.StringPtrInput
	ServerConfiguration StorageSystemServerConfigurationInput
	ServerCredentials   StorageSystemServerCredentialsPtrInput
	// The type of on-premises storage system that DataSync Discovery will analyze.
	SystemType StorageSystemSystemTypeInput
	// An array of key-value pairs to apply to this resource.
	Tags StorageSystemTagArrayInput
}

func (StorageSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSystemArgs)(nil)).Elem()
}

type StorageSystemInput interface {
	pulumi.Input

	ToStorageSystemOutput() StorageSystemOutput
	ToStorageSystemOutputWithContext(ctx context.Context) StorageSystemOutput
}

func (*StorageSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSystem)(nil)).Elem()
}

func (i *StorageSystem) ToStorageSystemOutput() StorageSystemOutput {
	return i.ToStorageSystemOutputWithContext(context.Background())
}

func (i *StorageSystem) ToStorageSystemOutputWithContext(ctx context.Context) StorageSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSystemOutput)
}

type StorageSystemOutput struct{ *pulumi.OutputState }

func (StorageSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSystem)(nil)).Elem()
}

func (o StorageSystemOutput) ToStorageSystemOutput() StorageSystemOutput {
	return o
}

func (o StorageSystemOutput) ToStorageSystemOutputWithContext(ctx context.Context) StorageSystemOutput {
	return o
}

// The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.
func (o StorageSystemOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageSystem) pulumi.StringArrayOutput { return v.AgentArns }).(pulumi.StringArrayOutput)
}

// The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.
func (o StorageSystemOutput) CloudWatchLogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSystem) pulumi.StringPtrOutput { return v.CloudWatchLogGroupArn }).(pulumi.StringPtrOutput)
}

// Indicates whether the DataSync agent can access the on-premises storage system.
func (o StorageSystemOutput) ConnectivityStatus() StorageSystemConnectivityStatusOutput {
	return o.ApplyT(func(v *StorageSystem) StorageSystemConnectivityStatusOutput { return v.ConnectivityStatus }).(StorageSystemConnectivityStatusOutput)
}

// A familiar name for the on-premises storage system.
func (o StorageSystemOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSystem) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The ARN of a secret stored by AWS Secrets Manager.
func (o StorageSystemOutput) SecretsManagerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSystem) pulumi.StringOutput { return v.SecretsManagerArn }).(pulumi.StringOutput)
}

func (o StorageSystemOutput) ServerConfiguration() StorageSystemServerConfigurationOutput {
	return o.ApplyT(func(v *StorageSystem) StorageSystemServerConfigurationOutput { return v.ServerConfiguration }).(StorageSystemServerConfigurationOutput)
}

func (o StorageSystemOutput) ServerCredentials() StorageSystemServerCredentialsPtrOutput {
	return o.ApplyT(func(v *StorageSystem) StorageSystemServerCredentialsPtrOutput { return v.ServerCredentials }).(StorageSystemServerCredentialsPtrOutput)
}

// The ARN of the on-premises storage system added to DataSync Discovery.
func (o StorageSystemOutput) StorageSystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSystem) pulumi.StringOutput { return v.StorageSystemArn }).(pulumi.StringOutput)
}

// The type of on-premises storage system that DataSync Discovery will analyze.
func (o StorageSystemOutput) SystemType() StorageSystemSystemTypeOutput {
	return o.ApplyT(func(v *StorageSystem) StorageSystemSystemTypeOutput { return v.SystemType }).(StorageSystemSystemTypeOutput)
}

// An array of key-value pairs to apply to this resource.
func (o StorageSystemOutput) Tags() StorageSystemTagArrayOutput {
	return o.ApplyT(func(v *StorageSystem) StorageSystemTagArrayOutput { return v.Tags }).(StorageSystemTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageSystemInput)(nil)).Elem(), &StorageSystem{})
	pulumi.RegisterOutputType(StorageSystemOutput{})
}
