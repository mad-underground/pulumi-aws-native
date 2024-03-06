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
    /// Includes cookies of a web request.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupCookies
    {
        public readonly Outputs.RuleGroupCookieMatchPattern MatchPattern;
        public readonly Pulumi.AwsNative.WaFv2.RuleGroupMapMatchScope MatchScope;
        public readonly Pulumi.AwsNative.WaFv2.RuleGroupOversizeHandling OversizeHandling;

        [OutputConstructor]
        private RuleGroupCookies(
            Outputs.RuleGroupCookieMatchPattern matchPattern,

            Pulumi.AwsNative.WaFv2.RuleGroupMapMatchScope matchScope,

            Pulumi.AwsNative.WaFv2.RuleGroupOversizeHandling oversizeHandling)
        {
            MatchPattern = matchPattern;
            MatchScope = matchScope;
            OversizeHandling = oversizeHandling;
        }
    }
}
