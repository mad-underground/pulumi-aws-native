// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Specifies a single custom aggregate key for a rate-base rule.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupRateBasedStatementCustomKey
    {
        public readonly Outputs.RuleGroupRateLimitCookie? Cookie;
        public readonly Outputs.RuleGroupRateLimitForwardedIP? ForwardedIP;
        public readonly Outputs.RuleGroupRateLimitHTTPMethod? HTTPMethod;
        public readonly Outputs.RuleGroupRateLimitHeader? Header;
        public readonly Outputs.RuleGroupRateLimitIP? IP;
        public readonly Outputs.RuleGroupRateLimitLabelNamespace? LabelNamespace;
        public readonly Outputs.RuleGroupRateLimitQueryArgument? QueryArgument;
        public readonly Outputs.RuleGroupRateLimitQueryString? QueryString;
        public readonly Outputs.RuleGroupRateLimitUriPath? UriPath;

        [OutputConstructor]
        private RuleGroupRateBasedStatementCustomKey(
            Outputs.RuleGroupRateLimitCookie? cookie,

            Outputs.RuleGroupRateLimitForwardedIP? forwardedIP,

            Outputs.RuleGroupRateLimitHTTPMethod? hTTPMethod,

            Outputs.RuleGroupRateLimitHeader? header,

            Outputs.RuleGroupRateLimitIP? iP,

            Outputs.RuleGroupRateLimitLabelNamespace? labelNamespace,

            Outputs.RuleGroupRateLimitQueryArgument? queryArgument,

            Outputs.RuleGroupRateLimitQueryString? queryString,

            Outputs.RuleGroupRateLimitUriPath? uriPath)
        {
            Cookie = cookie;
            ForwardedIP = forwardedIP;
            HTTPMethod = hTTPMethod;
            Header = header;
            IP = iP;
            LabelNamespace = labelNamespace;
            QueryArgument = queryArgument;
            QueryString = queryString;
            UriPath = uriPath;
        }
    }
}
