// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Deployment
//
// Deprecated: Deployment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Deployment struct {
	pulumi.CustomResourceState

	ApplicationId          pulumi.StringOutput       `pulumi:"applicationId"`
	ConfigurationProfileId pulumi.StringOutput       `pulumi:"configurationProfileId"`
	ConfigurationVersion   pulumi.StringOutput       `pulumi:"configurationVersion"`
	DeploymentStrategyId   pulumi.StringOutput       `pulumi:"deploymentStrategyId"`
	Description            pulumi.StringPtrOutput    `pulumi:"description"`
	EnvironmentId          pulumi.StringOutput       `pulumi:"environmentId"`
	KmsKeyIdentifier       pulumi.StringPtrOutput    `pulumi:"kmsKeyIdentifier"`
	Tags                   DeploymentTagsArrayOutput `pulumi:"tags"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.ConfigurationProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationProfileId'")
	}
	if args.ConfigurationVersion == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationVersion'")
	}
	if args.DeploymentStrategyId == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentStrategyId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Deployment
	err := ctx.RegisterResource("aws-native:appconfig:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("aws-native:appconfig:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
}

type DeploymentState struct {
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	ApplicationId          string           `pulumi:"applicationId"`
	ConfigurationProfileId string           `pulumi:"configurationProfileId"`
	ConfigurationVersion   string           `pulumi:"configurationVersion"`
	DeploymentStrategyId   string           `pulumi:"deploymentStrategyId"`
	Description            *string          `pulumi:"description"`
	EnvironmentId          string           `pulumi:"environmentId"`
	KmsKeyIdentifier       *string          `pulumi:"kmsKeyIdentifier"`
	Tags                   []DeploymentTags `pulumi:"tags"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	ApplicationId          pulumi.StringInput
	ConfigurationProfileId pulumi.StringInput
	ConfigurationVersion   pulumi.StringInput
	DeploymentStrategyId   pulumi.StringInput
	Description            pulumi.StringPtrInput
	EnvironmentId          pulumi.StringInput
	KmsKeyIdentifier       pulumi.StringPtrInput
	Tags                   DeploymentTagsArrayInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func (o DeploymentOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o DeploymentOutput) ConfigurationProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ConfigurationProfileId }).(pulumi.StringOutput)
}

func (o DeploymentOutput) ConfigurationVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.ConfigurationVersion }).(pulumi.StringOutput)
}

func (o DeploymentOutput) DeploymentStrategyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.DeploymentStrategyId }).(pulumi.StringOutput)
}

func (o DeploymentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DeploymentOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

func (o DeploymentOutput) KmsKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) pulumi.StringPtrOutput { return v.KmsKeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o DeploymentOutput) Tags() DeploymentTagsArrayOutput {
	return o.ApplyT(func(v *Deployment) DeploymentTagsArrayOutput { return v.Tags }).(DeploymentTagsArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentInput)(nil)).Elem(), &Deployment{})
	pulumi.RegisterOutputType(DeploymentOutput{})
}
