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
    public sealed class ListenerRuleQueryStringConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-querystringconfig.html#cfn-elasticloadbalancingv2-listenerrule-querystringconfig-values
        /// </summary>
        public readonly ImmutableArray<Outputs.ListenerRuleQueryStringKeyValue> Values;

        [OutputConstructor]
        private ListenerRuleQueryStringConfig(ImmutableArray<Outputs.ListenerRuleQueryStringKeyValue> values)
        {
            Values = values;
        }
    }
}
