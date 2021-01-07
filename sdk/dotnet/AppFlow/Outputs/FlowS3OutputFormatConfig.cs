// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowS3OutputFormatConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-aggregationconfig
        /// </summary>
        public readonly Outputs.FlowAggregationConfig? AggregationConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-filetype
        /// </summary>
        public readonly string? FileType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-s3outputformatconfig.html#cfn-appflow-flow-s3outputformatconfig-prefixconfig
        /// </summary>
        public readonly Outputs.FlowPrefixConfig? PrefixConfig;

        [OutputConstructor]
        private FlowS3OutputFormatConfig(
            Outputs.FlowAggregationConfig? AggregationConfig,

            string? FileType,

            Outputs.FlowPrefixConfig? PrefixConfig)
        {
            this.AggregationConfig = AggregationConfig;
            this.FileType = FileType;
            this.PrefixConfig = PrefixConfig;
        }
    }
}