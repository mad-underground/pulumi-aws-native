// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// Specifies a query argument in the request as an aggregate key for a rate-based rule.
    /// </summary>
    public sealed class RuleGroupRateLimitQueryArgumentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the query argument to use.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("textTransformations", required: true)]
        private InputList<Inputs.RuleGroupTextTransformationArgs>? _textTransformations;
        public InputList<Inputs.RuleGroupTextTransformationArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.RuleGroupTextTransformationArgs>());
            set => _textTransformations = value;
        }

        public RuleGroupRateLimitQueryArgumentArgs()
        {
        }
        public static new RuleGroupRateLimitQueryArgumentArgs Empty => new RuleGroupRateLimitQueryArgumentArgs();
    }
}
