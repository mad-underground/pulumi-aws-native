// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a runtime environment that can run migrated mainframe applications.
type Environment struct {
	pulumi.CustomResourceState

	// The description of the environment.
	Description pulumi.StringPtrOutput      `pulumi:"description"`
	EngineType  EnvironmentEngineTypeOutput `pulumi:"engineType"`
	// The version of the runtime engine for the environment.
	EngineVersion pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) of the runtime environment.
	EnvironmentArn pulumi.StringOutput `pulumi:"environmentArn"`
	// The unique identifier of the environment.
	EnvironmentId          pulumi.StringOutput                        `pulumi:"environmentId"`
	HighAvailabilityConfig EnvironmentHighAvailabilityConfigPtrOutput `pulumi:"highAvailabilityConfig"`
	// The type of instance underlying the environment.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow pulumi.StringPtrOutput `pulumi:"preferredMaintenanceWindow"`
	// Specifies whether the environment is publicly accessible.
	PubliclyAccessible pulumi.BoolPtrOutput `pulumi:"publiclyAccessible"`
	// The list of security groups for the VPC associated with this environment.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The storage configurations defined for the runtime environment.
	StorageConfigurations EnvironmentStorageConfigurationArrayOutput `pulumi:"storageConfigurations"`
	// The unique identifiers of the subnets assigned to this runtime environment.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Tags associated to this environment.
	Tags EnvironmentTagMapPtrOutput `pulumi:"tags"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EngineType == nil {
		return nil, errors.New("invalid value for required argument 'EngineType'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource Environment
	err := ctx.RegisterResource("aws-native:m2:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("aws-native:m2:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// The description of the environment.
	Description *string               `pulumi:"description"`
	EngineType  EnvironmentEngineType `pulumi:"engineType"`
	// The version of the runtime engine for the environment.
	EngineVersion          *string                            `pulumi:"engineVersion"`
	HighAvailabilityConfig *EnvironmentHighAvailabilityConfig `pulumi:"highAvailabilityConfig"`
	// The type of instance underlying the environment.
	InstanceType string `pulumi:"instanceType"`
	// The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Specifies whether the environment is publicly accessible.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The list of security groups for the VPC associated with this environment.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The storage configurations defined for the runtime environment.
	StorageConfigurations []EnvironmentStorageConfiguration `pulumi:"storageConfigurations"`
	// The unique identifiers of the subnets assigned to this runtime environment.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tags associated to this environment.
	Tags *EnvironmentTagMap `pulumi:"tags"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// The description of the environment.
	Description pulumi.StringPtrInput
	EngineType  EnvironmentEngineTypeInput
	// The version of the runtime engine for the environment.
	EngineVersion          pulumi.StringPtrInput
	HighAvailabilityConfig EnvironmentHighAvailabilityConfigPtrInput
	// The type of instance underlying the environment.
	InstanceType pulumi.StringInput
	// The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
	KmsKeyId pulumi.StringPtrInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Specifies whether the environment is publicly accessible.
	PubliclyAccessible pulumi.BoolPtrInput
	// The list of security groups for the VPC associated with this environment.
	SecurityGroupIds pulumi.StringArrayInput
	// The storage configurations defined for the runtime environment.
	StorageConfigurations EnvironmentStorageConfigurationArrayInput
	// The unique identifiers of the subnets assigned to this runtime environment.
	SubnetIds pulumi.StringArrayInput
	// Tags associated to this environment.
	Tags EnvironmentTagMapPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// The description of the environment.
func (o EnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentOutput) EngineType() EnvironmentEngineTypeOutput {
	return o.ApplyT(func(v *Environment) EnvironmentEngineTypeOutput { return v.EngineType }).(EnvironmentEngineTypeOutput)
}

// The version of the runtime engine for the environment.
func (o EnvironmentOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the runtime environment.
func (o EnvironmentOutput) EnvironmentArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.EnvironmentArn }).(pulumi.StringOutput)
}

// The unique identifier of the environment.
func (o EnvironmentOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) HighAvailabilityConfig() EnvironmentHighAvailabilityConfigPtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentHighAvailabilityConfigPtrOutput { return v.HighAvailabilityConfig }).(EnvironmentHighAvailabilityConfigPtrOutput)
}

// The type of instance underlying the environment.
func (o EnvironmentOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
func (o EnvironmentOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The name of the environment.
func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
func (o EnvironmentOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

// Specifies whether the environment is publicly accessible.
func (o EnvironmentOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.BoolPtrOutput { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

// The list of security groups for the VPC associated with this environment.
func (o EnvironmentOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The storage configurations defined for the runtime environment.
func (o EnvironmentOutput) StorageConfigurations() EnvironmentStorageConfigurationArrayOutput {
	return o.ApplyT(func(v *Environment) EnvironmentStorageConfigurationArrayOutput { return v.StorageConfigurations }).(EnvironmentStorageConfigurationArrayOutput)
}

// The unique identifiers of the subnets assigned to this runtime environment.
func (o EnvironmentOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Tags associated to this environment.
func (o EnvironmentOutput) Tags() EnvironmentTagMapPtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentTagMapPtrOutput { return v.Tags }).(EnvironmentTagMapPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
}
