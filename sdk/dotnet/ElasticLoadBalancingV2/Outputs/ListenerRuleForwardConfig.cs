// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class ListenerRuleForwardConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-forwardconfig.html#cfn-elasticloadbalancingv2-listenerrule-forwardconfig-targetgroupstickinessconfig
        /// </summary>
        public readonly Outputs.ListenerRuleTargetGroupStickinessConfig? TargetGroupStickinessConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-forwardconfig.html#cfn-elasticloadbalancingv2-listenerrule-forwardconfig-targetgroups
        /// </summary>
        public readonly ImmutableArray<Outputs.ListenerRuleTargetGroupTuple> TargetGroups;

        [OutputConstructor]
        private ListenerRuleForwardConfig(
            Outputs.ListenerRuleTargetGroupStickinessConfig? targetGroupStickinessConfig,

            ImmutableArray<Outputs.ListenerRuleTargetGroupTuple> targetGroups)
        {
            TargetGroupStickinessConfig = targetGroupStickinessConfig;
            TargetGroups = targetGroups;
        }
    }
}
