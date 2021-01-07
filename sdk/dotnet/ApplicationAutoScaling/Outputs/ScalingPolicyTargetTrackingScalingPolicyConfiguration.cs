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
    public sealed class ScalingPolicyTargetTrackingScalingPolicyConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-customizedmetricspecification
        /// </summary>
        public readonly Outputs.ScalingPolicyCustomizedMetricSpecification? CustomizedMetricSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-disablescalein
        /// </summary>
        public readonly bool? DisableScaleIn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-predefinedmetricspecification
        /// </summary>
        public readonly Outputs.ScalingPolicyPredefinedMetricSpecification? PredefinedMetricSpecification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-scaleincooldown
        /// </summary>
        public readonly int? ScaleInCooldown;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-scaleoutcooldown
        /// </summary>
        public readonly int? ScaleOutCooldown;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration-targetvalue
        /// </summary>
        public readonly double TargetValue;

        [OutputConstructor]
        private ScalingPolicyTargetTrackingScalingPolicyConfiguration(
            Outputs.ScalingPolicyCustomizedMetricSpecification? CustomizedMetricSpecification,

            bool? DisableScaleIn,

            Outputs.ScalingPolicyPredefinedMetricSpecification? PredefinedMetricSpecification,

            int? ScaleInCooldown,

            int? ScaleOutCooldown,

            double TargetValue)
        {
            this.CustomizedMetricSpecification = CustomizedMetricSpecification;
            this.DisableScaleIn = DisableScaleIn;
            this.PredefinedMetricSpecification = PredefinedMetricSpecification;
            this.ScaleInCooldown = ScaleInCooldown;
            this.ScaleOutCooldown = ScaleOutCooldown;
            this.TargetValue = TargetValue;
        }
    }
}