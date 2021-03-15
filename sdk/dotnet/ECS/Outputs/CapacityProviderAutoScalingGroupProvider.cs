// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class CapacityProviderAutoScalingGroupProvider
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-autoscalinggrouparn
        /// </summary>
        public readonly string AutoScalingGroupArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-managedscaling
        /// </summary>
        public readonly Outputs.CapacityProviderManagedScaling? ManagedScaling;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-managedterminationprotection
        /// </summary>
        public readonly string? ManagedTerminationProtection;

        [OutputConstructor]
        private CapacityProviderAutoScalingGroupProvider(
            string autoScalingGroupArn,

            Outputs.CapacityProviderManagedScaling? managedScaling,

            string? managedTerminationProtection)
        {
            AutoScalingGroupArn = autoScalingGroupArn;
            ManagedScaling = managedScaling;
            ManagedTerminationProtection = managedTerminationProtection;
        }
    }
}
