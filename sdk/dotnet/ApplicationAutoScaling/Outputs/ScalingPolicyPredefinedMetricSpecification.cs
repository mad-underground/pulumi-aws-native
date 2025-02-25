// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationAutoScaling.Outputs
{

    /// <summary>
    /// Represents a predefined metric for a target tracking scaling policy to use with Application Auto Scaling.
    /// </summary>
    [OutputType]
    public sealed class ScalingPolicyPredefinedMetricSpecification
    {
        /// <summary>
        /// The metric type. The ALBRequestCountPerTarget metric type applies only to Spot Fleets and ECS services.
        /// </summary>
        public readonly string PredefinedMetricType;
        /// <summary>
        /// Identifies the resource associated with the metric type. You can't specify a resource label unless the metric type is ALBRequestCountPerTarget and there is a target group attached to the Spot Fleet or ECS service.
        /// </summary>
        public readonly string? ResourceLabel;

        [OutputConstructor]
        private ScalingPolicyPredefinedMetricSpecification(
            string predefinedMetricType,

            string? resourceLabel)
        {
            PredefinedMetricType = predefinedMetricType;
            ResourceLabel = resourceLabel;
        }
    }
}
