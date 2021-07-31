// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html
    /// </summary>
    [OutputType]
    public sealed class IndexRelevance
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-duration
        /// </summary>
        public readonly string? Duration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-freshness
        /// </summary>
        public readonly bool? Freshness;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-importance
        /// </summary>
        public readonly int? Importance;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-rankorder
        /// </summary>
        public readonly string? RankOrder;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-valueimportanceitems
        /// </summary>
        public readonly Outputs.IndexValueImportanceItems? ValueImportanceItems;

        [OutputConstructor]
        private IndexRelevance(
            string? duration,

            bool? freshness,

            int? importance,

            string? rankOrder,

            Outputs.IndexValueImportanceItems? valueImportanceItems)
        {
            Duration = duration;
            Freshness = freshness;
            Importance = importance;
            RankOrder = rankOrder;
            ValueImportanceItems = valueImportanceItems;
        }
    }
}
