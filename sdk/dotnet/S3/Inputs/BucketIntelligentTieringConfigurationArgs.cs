// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.S3.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html
    /// </summary>
    public sealed class BucketIntelligentTieringConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-id
        /// </summary>
        [Input("Id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-prefix
        /// </summary>
        [Input("Prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-status
        /// </summary>
        [Input("Status", required: true)]
        public Input<string> Status { get; set; } = null!;

        [Input("TagFilters")]
        private InputList<Inputs.BucketTagFilterArgs>? _TagFilters;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-tagfilters
        /// </summary>
        public InputList<Inputs.BucketTagFilterArgs> TagFilters
        {
            get => _TagFilters ?? (_TagFilters = new InputList<Inputs.BucketTagFilterArgs>());
            set => _TagFilters = value;
        }

        [Input("Tierings", required: true)]
        private InputList<Inputs.BucketTieringArgs>? _Tierings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-intelligenttieringconfiguration.html#cfn-s3-bucket-intelligenttieringconfiguration-tierings
        /// </summary>
        public InputList<Inputs.BucketTieringArgs> Tierings
        {
            get => _Tierings ?? (_Tierings = new InputList<Inputs.BucketTieringArgs>());
            set => _Tierings = value;
        }

        public BucketIntelligentTieringConfigurationArgs()
        {
        }
    }
}
