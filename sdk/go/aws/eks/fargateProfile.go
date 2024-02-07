// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::EKS::FargateProfile
type FargateProfile struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the Cluster
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Name of FargateProfile
	FargateProfileName pulumi.StringPtrOutput `pulumi:"fargateProfileName"`
	// The IAM policy arn for pods
	PodExecutionRoleArn pulumi.StringOutput               `pulumi:"podExecutionRoleArn"`
	Selectors           FargateProfileSelectorArrayOutput `pulumi:"selectors"`
	Subnets             pulumi.StringArrayOutput          `pulumi:"subnets"`
	// An array of key-value pairs to apply to this resource.
	Tags FargateProfileTagArrayOutput `pulumi:"tags"`
}

// NewFargateProfile registers a new resource with the given unique name, arguments, and options.
func NewFargateProfile(ctx *pulumi.Context,
	name string, args *FargateProfileArgs, opts ...pulumi.ResourceOption) (*FargateProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PodExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'PodExecutionRoleArn'")
	}
	if args.Selectors == nil {
		return nil, errors.New("invalid value for required argument 'Selectors'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clusterName",
		"fargateProfileName",
		"podExecutionRoleArn",
		"selectors[*]",
		"subnets[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FargateProfile
	err := ctx.RegisterResource("aws-native:eks:FargateProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFargateProfile gets an existing FargateProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFargateProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FargateProfileState, opts ...pulumi.ResourceOption) (*FargateProfile, error) {
	var resource FargateProfile
	err := ctx.ReadResource("aws-native:eks:FargateProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FargateProfile resources.
type fargateProfileState struct {
}

type FargateProfileState struct {
}

func (FargateProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateProfileState)(nil)).Elem()
}

type fargateProfileArgs struct {
	// Name of the Cluster
	ClusterName string `pulumi:"clusterName"`
	// Name of FargateProfile
	FargateProfileName *string `pulumi:"fargateProfileName"`
	// The IAM policy arn for pods
	PodExecutionRoleArn string                   `pulumi:"podExecutionRoleArn"`
	Selectors           []FargateProfileSelector `pulumi:"selectors"`
	Subnets             []string                 `pulumi:"subnets"`
	// An array of key-value pairs to apply to this resource.
	Tags []FargateProfileTag `pulumi:"tags"`
}

// The set of arguments for constructing a FargateProfile resource.
type FargateProfileArgs struct {
	// Name of the Cluster
	ClusterName pulumi.StringInput
	// Name of FargateProfile
	FargateProfileName pulumi.StringPtrInput
	// The IAM policy arn for pods
	PodExecutionRoleArn pulumi.StringInput
	Selectors           FargateProfileSelectorArrayInput
	Subnets             pulumi.StringArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags FargateProfileTagArrayInput
}

func (FargateProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateProfileArgs)(nil)).Elem()
}

type FargateProfileInput interface {
	pulumi.Input

	ToFargateProfileOutput() FargateProfileOutput
	ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput
}

func (*FargateProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**FargateProfile)(nil)).Elem()
}

func (i *FargateProfile) ToFargateProfileOutput() FargateProfileOutput {
	return i.ToFargateProfileOutputWithContext(context.Background())
}

func (i *FargateProfile) ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfileOutput)
}

type FargateProfileOutput struct{ *pulumi.OutputState }

func (FargateProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FargateProfile)(nil)).Elem()
}

func (o FargateProfileOutput) ToFargateProfileOutput() FargateProfileOutput {
	return o
}

func (o FargateProfileOutput) ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput {
	return o
}

func (o FargateProfileOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *FargateProfile) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of the Cluster
func (o FargateProfileOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *FargateProfile) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Name of FargateProfile
func (o FargateProfileOutput) FargateProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FargateProfile) pulumi.StringPtrOutput { return v.FargateProfileName }).(pulumi.StringPtrOutput)
}

// The IAM policy arn for pods
func (o FargateProfileOutput) PodExecutionRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FargateProfile) pulumi.StringOutput { return v.PodExecutionRoleArn }).(pulumi.StringOutput)
}

func (o FargateProfileOutput) Selectors() FargateProfileSelectorArrayOutput {
	return o.ApplyT(func(v *FargateProfile) FargateProfileSelectorArrayOutput { return v.Selectors }).(FargateProfileSelectorArrayOutput)
}

func (o FargateProfileOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FargateProfile) pulumi.StringArrayOutput { return v.Subnets }).(pulumi.StringArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o FargateProfileOutput) Tags() FargateProfileTagArrayOutput {
	return o.ApplyT(func(v *FargateProfile) FargateProfileTagArrayOutput { return v.Tags }).(FargateProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FargateProfileInput)(nil)).Elem(), &FargateProfile{})
	pulumi.RegisterOutputType(FargateProfileOutput{})
}
