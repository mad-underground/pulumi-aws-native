// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Outputs
{

    [OutputType]
    public sealed class MonitoringScheduleBaselineConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-baselineconfig.html#cfn-sagemaker-monitoringschedule-baselineconfig-constraintsresource
        /// </summary>
        public readonly Outputs.MonitoringScheduleConstraintsResource? ConstraintsResource;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-baselineconfig.html#cfn-sagemaker-monitoringschedule-baselineconfig-statisticsresource
        /// </summary>
        public readonly Outputs.MonitoringScheduleStatisticsResource? StatisticsResource;

        [OutputConstructor]
        private MonitoringScheduleBaselineConfig(
            Outputs.MonitoringScheduleConstraintsResource? ConstraintsResource,

            Outputs.MonitoringScheduleStatisticsResource? StatisticsResource)
        {
            this.ConstraintsResource = ConstraintsResource;
            this.StatisticsResource = StatisticsResource;
        }
    }
}