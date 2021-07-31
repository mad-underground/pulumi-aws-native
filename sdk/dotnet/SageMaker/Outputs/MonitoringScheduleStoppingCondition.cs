// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-stoppingcondition.html
    /// </summary>
    [OutputType]
    public sealed class MonitoringScheduleStoppingCondition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-stoppingcondition.html#cfn-sagemaker-monitoringschedule-stoppingcondition-maxruntimeinseconds
        /// </summary>
        public readonly int MaxRuntimeInSeconds;

        [OutputConstructor]
        private MonitoringScheduleStoppingCondition(int maxRuntimeInSeconds)
        {
            MaxRuntimeInSeconds = maxRuntimeInSeconds;
        }
    }
}
