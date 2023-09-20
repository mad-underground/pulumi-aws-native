// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Specifies the request's URI Path as an aggregate key for a rate-based rule.
    /// </summary>
    public sealed class RuleGroupRateLimitUriPathArgs : global::Pulumi.ResourceArgs
    {
        [Input("textTransformations", required: true)]
        private InputList<Inputs.RuleGroupTextTransformationArgs>? _textTransformations;
        public InputList<Inputs.RuleGroupTextTransformationArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.RuleGroupTextTransformationArgs>());
            set => _textTransformations = value;
        }

        public RuleGroupRateLimitUriPathArgs()
        {
        }
        public static new RuleGroupRateLimitUriPathArgs Empty => new RuleGroupRateLimitUriPathArgs();
    }
}