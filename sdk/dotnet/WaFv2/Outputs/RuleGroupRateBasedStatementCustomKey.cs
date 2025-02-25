// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Specifies a single custom aggregate key for a rate-base rule.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupRateBasedStatementCustomKey
    {
        public readonly Outputs.RuleGroupRateLimitCookie? Cookie;
        public readonly Outputs.RuleGroupRateLimitForwardedIp? ForwardedIp;
        public readonly Outputs.RuleGroupRateLimitHeader? Header;
        public readonly Outputs.RuleGroupRateLimitHttpMethod? HttpMethod;
        public readonly Outputs.RuleGroupRateLimitIp? Ip;
        public readonly Outputs.RuleGroupRateLimitLabelNamespace? LabelNamespace;
        public readonly Outputs.RuleGroupRateLimitQueryArgument? QueryArgument;
        public readonly Outputs.RuleGroupRateLimitQueryString? QueryString;
        public readonly Outputs.RuleGroupRateLimitUriPath? UriPath;

        [OutputConstructor]
        private RuleGroupRateBasedStatementCustomKey(
            Outputs.RuleGroupRateLimitCookie? cookie,

            Outputs.RuleGroupRateLimitForwardedIp? forwardedIp,

            Outputs.RuleGroupRateLimitHeader? header,

            Outputs.RuleGroupRateLimitHttpMethod? httpMethod,

            Outputs.RuleGroupRateLimitIp? ip,

            Outputs.RuleGroupRateLimitLabelNamespace? labelNamespace,

            Outputs.RuleGroupRateLimitQueryArgument? queryArgument,

            Outputs.RuleGroupRateLimitQueryString? queryString,

            Outputs.RuleGroupRateLimitUriPath? uriPath)
        {
            Cookie = cookie;
            ForwardedIp = forwardedIp;
            Header = header;
            HttpMethod = httpMethod;
            Ip = ip;
            LabelNamespace = labelNamespace;
            QueryArgument = queryArgument;
            QueryString = queryString;
            UriPath = uriPath;
        }
    }
}
