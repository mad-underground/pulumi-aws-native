// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deserializer.html
    /// </summary>
    [OutputType]
    public sealed class DeliveryStreamDeserializer
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deserializer.html#cfn-kinesisfirehose-deliverystream-deserializer-hivejsonserde
        /// </summary>
        public readonly Outputs.DeliveryStreamHiveJsonSerDe? HiveJsonSerDe;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deserializer.html#cfn-kinesisfirehose-deliverystream-deserializer-openxjsonserde
        /// </summary>
        public readonly Outputs.DeliveryStreamOpenXJsonSerDe? OpenXJsonSerDe;

        [OutputConstructor]
        private DeliveryStreamDeserializer(
            Outputs.DeliveryStreamHiveJsonSerDe? hiveJsonSerDe,

            Outputs.DeliveryStreamOpenXJsonSerDe? openXJsonSerDe)
        {
            HiveJsonSerDe = hiveJsonSerDe;
            OpenXJsonSerDe = openXJsonSerDe;
        }
    }
}
