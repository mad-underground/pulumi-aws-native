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
    public sealed class ListenerRuleRuleCondition
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-conditions-field
        /// </summary>
        public readonly string? Field;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-hostheaderconfig
        /// </summary>
        public readonly Outputs.ListenerRuleHostHeaderConfig? HostHeaderConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-httpheaderconfig
        /// </summary>
        public readonly Outputs.ListenerRuleHttpHeaderConfig? HttpHeaderConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-httprequestmethodconfig
        /// </summary>
        public readonly Outputs.ListenerRuleHttpRequestMethodConfig? HttpRequestMethodConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-pathpatternconfig
        /// </summary>
        public readonly Outputs.ListenerRulePathPatternConfig? PathPatternConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-querystringconfig
        /// </summary>
        public readonly Outputs.ListenerRuleQueryStringConfig? QueryStringConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-sourceipconfig
        /// </summary>
        public readonly Outputs.ListenerRuleSourceIpConfig? SourceIpConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-conditions.html#cfn-elasticloadbalancingv2-listenerrule-conditions-values
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ListenerRuleRuleCondition(
            string? field,

            Outputs.ListenerRuleHostHeaderConfig? hostHeaderConfig,

            Outputs.ListenerRuleHttpHeaderConfig? httpHeaderConfig,

            Outputs.ListenerRuleHttpRequestMethodConfig? httpRequestMethodConfig,

            Outputs.ListenerRulePathPatternConfig? pathPatternConfig,

            Outputs.ListenerRuleQueryStringConfig? queryStringConfig,

            Outputs.ListenerRuleSourceIpConfig? sourceIpConfig,

            ImmutableArray<string> values)
        {
            Field = field;
            HostHeaderConfig = hostHeaderConfig;
            HttpHeaderConfig = httpHeaderConfig;
            HttpRequestMethodConfig = httpRequestMethodConfig;
            PathPatternConfig = pathPatternConfig;
            QueryStringConfig = queryStringConfig;
            SourceIpConfig = sourceIpConfig;
            Values = values;
        }
    }
}
