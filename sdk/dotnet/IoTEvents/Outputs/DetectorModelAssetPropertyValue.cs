// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-assetpropertyvalue.html
    /// </summary>
    [OutputType]
    public sealed class DetectorModelAssetPropertyValue
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-assetpropertyvalue.html#cfn-iotevents-detectormodel-assetpropertyvalue-quality
        /// </summary>
        public readonly string? Quality;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-assetpropertyvalue.html#cfn-iotevents-detectormodel-assetpropertyvalue-timestamp
        /// </summary>
        public readonly Outputs.DetectorModelAssetPropertyTimestamp? Timestamp;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-assetpropertyvalue.html#cfn-iotevents-detectormodel-assetpropertyvalue-value
        /// </summary>
        public readonly Outputs.DetectorModelAssetPropertyVariant? Value;

        [OutputConstructor]
        private DetectorModelAssetPropertyValue(
            string? quality,

            Outputs.DetectorModelAssetPropertyTimestamp? timestamp,

            Outputs.DetectorModelAssetPropertyVariant? value)
        {
            Quality = quality;
            Timestamp = timestamp;
            Value = value;
        }
    }
}
