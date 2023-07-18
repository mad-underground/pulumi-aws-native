// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applicationautoscaling

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApplicationAutoScaling::ScalingPolicy
//
// Deprecated: ScalingPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ScalingPolicy struct {
	pulumi.CustomResourceState

	PolicyName                               pulumi.StringOutput                                            `pulumi:"policyName"`
	PolicyType                               pulumi.StringOutput                                            `pulumi:"policyType"`
	ResourceId                               pulumi.StringPtrOutput                                         `pulumi:"resourceId"`
	ScalableDimension                        pulumi.StringPtrOutput                                         `pulumi:"scalableDimension"`
	ScalingTargetId                          pulumi.StringPtrOutput                                         `pulumi:"scalingTargetId"`
	ServiceNamespace                         pulumi.StringPtrOutput                                         `pulumi:"serviceNamespace"`
	StepScalingPolicyConfiguration           ScalingPolicyStepScalingPolicyConfigurationPtrOutput           `pulumi:"stepScalingPolicyConfiguration"`
	TargetTrackingScalingPolicyConfiguration ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// NewScalingPolicy registers a new resource with the given unique name, arguments, and options.
func NewScalingPolicy(ctx *pulumi.Context,
	name string, args *ScalingPolicyArgs, opts ...pulumi.ResourceOption) (*ScalingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScalingPolicy
	err := ctx.RegisterResource("aws-native:applicationautoscaling:ScalingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPolicy gets an existing ScalingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPolicyState, opts ...pulumi.ResourceOption) (*ScalingPolicy, error) {
	var resource ScalingPolicy
	err := ctx.ReadResource("aws-native:applicationautoscaling:ScalingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPolicy resources.
type scalingPolicyState struct {
}

type ScalingPolicyState struct {
}

func (ScalingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPolicyState)(nil)).Elem()
}

type scalingPolicyArgs struct {
	PolicyName                               string                                                 `pulumi:"policyName"`
	PolicyType                               string                                                 `pulumi:"policyType"`
	ResourceId                               *string                                                `pulumi:"resourceId"`
	ScalableDimension                        *string                                                `pulumi:"scalableDimension"`
	ScalingTargetId                          *string                                                `pulumi:"scalingTargetId"`
	ServiceNamespace                         *string                                                `pulumi:"serviceNamespace"`
	StepScalingPolicyConfiguration           *ScalingPolicyStepScalingPolicyConfiguration           `pulumi:"stepScalingPolicyConfiguration"`
	TargetTrackingScalingPolicyConfiguration *ScalingPolicyTargetTrackingScalingPolicyConfiguration `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// The set of arguments for constructing a ScalingPolicy resource.
type ScalingPolicyArgs struct {
	PolicyName                               pulumi.StringInput
	PolicyType                               pulumi.StringInput
	ResourceId                               pulumi.StringPtrInput
	ScalableDimension                        pulumi.StringPtrInput
	ScalingTargetId                          pulumi.StringPtrInput
	ServiceNamespace                         pulumi.StringPtrInput
	StepScalingPolicyConfiguration           ScalingPolicyStepScalingPolicyConfigurationPtrInput
	TargetTrackingScalingPolicyConfiguration ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrInput
}

func (ScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPolicyArgs)(nil)).Elem()
}

type ScalingPolicyInput interface {
	pulumi.Input

	ToScalingPolicyOutput() ScalingPolicyOutput
	ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput
}

func (*ScalingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPolicy)(nil)).Elem()
}

func (i *ScalingPolicy) ToScalingPolicyOutput() ScalingPolicyOutput {
	return i.ToScalingPolicyOutputWithContext(context.Background())
}

func (i *ScalingPolicy) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyOutput)
}

type ScalingPolicyOutput struct{ *pulumi.OutputState }

func (ScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyOutput) ToScalingPolicyOutput() ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

func (o ScalingPolicyOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

func (o ScalingPolicyOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ScalingPolicyOutput) ScalableDimension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ScalableDimension }).(pulumi.StringPtrOutput)
}

func (o ScalingPolicyOutput) ScalingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ScalingTargetId }).(pulumi.StringPtrOutput)
}

func (o ScalingPolicyOutput) ServiceNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) pulumi.StringPtrOutput { return v.ServiceNamespace }).(pulumi.StringPtrOutput)
}

func (o ScalingPolicyOutput) StepScalingPolicyConfiguration() ScalingPolicyStepScalingPolicyConfigurationPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) ScalingPolicyStepScalingPolicyConfigurationPtrOutput {
		return v.StepScalingPolicyConfiguration
	}).(ScalingPolicyStepScalingPolicyConfigurationPtrOutput)
}

func (o ScalingPolicyOutput) TargetTrackingScalingPolicyConfiguration() ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput {
	return o.ApplyT(func(v *ScalingPolicy) ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput {
		return v.TargetTrackingScalingPolicyConfiguration
	}).(ScalingPolicyTargetTrackingScalingPolicyConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyInput)(nil)).Elem(), &ScalingPolicy{})
	pulumi.RegisterOutputType(ScalingPolicyOutput{})
}
