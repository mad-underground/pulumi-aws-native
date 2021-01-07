// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApplicationAutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policyname
        /// </summary>
        public readonly string PolicyName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policytype
        /// </summary>
        public readonly string PolicyType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-resourceid
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalabledimension
        /// </summary>
        public readonly string? ScalableDimension;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalingtargetid
        /// </summary>
        public readonly string? ScalingTargetId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-servicenamespace
        /// </summary>
        public readonly string? ServiceNamespace;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration
        /// </summary>
        public readonly Outputs.ScalingPolicyStepScalingPolicyConfiguration? StepScalingPolicyConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration
        /// </summary>
        public readonly Outputs.ScalingPolicyTargetTrackingScalingPolicyConfiguration? TargetTrackingScalingPolicyConfiguration;

        [OutputConstructor]
        private ScalingPolicyProperties(
            string PolicyName,

            string PolicyType,

            string? ResourceId,

            string? ScalableDimension,

            string? ScalingTargetId,

            string? ServiceNamespace,

            Outputs.ScalingPolicyStepScalingPolicyConfiguration? StepScalingPolicyConfiguration,

            Outputs.ScalingPolicyTargetTrackingScalingPolicyConfiguration? TargetTrackingScalingPolicyConfiguration)
        {
            this.PolicyName = PolicyName;
            this.PolicyType = PolicyType;
            this.ResourceId = ResourceId;
            this.ScalableDimension = ScalableDimension;
            this.ScalingTargetId = ScalingTargetId;
            this.ServiceNamespace = ServiceNamespace;
            this.StepScalingPolicyConfiguration = StepScalingPolicyConfiguration;
            this.TargetTrackingScalingPolicyConfiguration = TargetTrackingScalingPolicyConfiguration;
        }
    }
}