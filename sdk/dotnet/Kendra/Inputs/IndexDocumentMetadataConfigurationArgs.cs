// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-documentmetadataconfiguration.html
    /// </summary>
    public sealed class IndexDocumentMetadataConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-documentmetadataconfiguration.html#cfn-kendra-index-documentmetadataconfiguration-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-documentmetadataconfiguration.html#cfn-kendra-index-documentmetadataconfiguration-relevance
        /// </summary>
        [Input("Relevance")]
        public Input<Inputs.IndexRelevanceArgs>? Relevance { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-documentmetadataconfiguration.html#cfn-kendra-index-documentmetadataconfiguration-search
        /// </summary>
        [Input("Search")]
        public Input<Inputs.IndexSearchArgs>? Search { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-documentmetadataconfiguration.html#cfn-kendra-index-documentmetadataconfiguration-type
        /// </summary>
        [Input("Type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IndexDocumentMetadataConfigurationArgs()
        {
        }
    }
}
