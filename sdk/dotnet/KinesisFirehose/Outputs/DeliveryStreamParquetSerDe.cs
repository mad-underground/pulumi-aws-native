// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamParquetSerDe
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-blocksizebytes
        /// </summary>
        public readonly int? BlockSizeBytes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-compression
        /// </summary>
        public readonly string? Compression;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-enabledictionarycompression
        /// </summary>
        public readonly bool? EnableDictionaryCompression;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-maxpaddingbytes
        /// </summary>
        public readonly int? MaxPaddingBytes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-pagesizebytes
        /// </summary>
        public readonly int? PageSizeBytes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-writerversion
        /// </summary>
        public readonly string? WriterVersion;

        [OutputConstructor]
        private DeliveryStreamParquetSerDe(
            int? blockSizeBytes,

            string? compression,

            bool? enableDictionaryCompression,

            int? maxPaddingBytes,

            int? pageSizeBytes,

            string? writerVersion)
        {
            BlockSizeBytes = blockSizeBytes;
            Compression = compression;
            EnableDictionaryCompression = enableDictionaryCompression;
            MaxPaddingBytes = maxPaddingBytes;
            PageSizeBytes = pageSizeBytes;
            WriterVersion = writerVersion;
        }
    }
}
