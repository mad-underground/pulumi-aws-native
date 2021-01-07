// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AutoScalingPlans.Outputs
{

    [OutputType]
    public sealed class ScalingPlanProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-applicationsource
        /// </summary>
        public readonly Outputs.ScalingPlanApplicationSource ApplicationSource;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-scalinginstructions
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingPlanScalingInstruction> ScalingInstructions;

        [OutputConstructor]
        private ScalingPlanProperties(
            Outputs.ScalingPlanApplicationSource ApplicationSource,

            ImmutableArray<Outputs.ScalingPlanScalingInstruction> ScalingInstructions)
        {
            this.ApplicationSource = ApplicationSource;
            this.ScalingInstructions = ScalingInstructions;
        }
    }
}