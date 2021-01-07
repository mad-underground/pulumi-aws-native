// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EMR.Outputs
{

    [OutputType]
    public sealed class ClusterScalingTrigger
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html#cfn-elasticmapreduce-cluster-scalingtrigger-cloudwatchalarmdefinition
        /// </summary>
        public readonly Outputs.ClusterCloudWatchAlarmDefinition CloudWatchAlarmDefinition;

        [OutputConstructor]
        private ClusterScalingTrigger(Outputs.ClusterCloudWatchAlarmDefinition CloudWatchAlarmDefinition)
        {
            this.CloudWatchAlarmDefinition = CloudWatchAlarmDefinition;
        }
    }
}