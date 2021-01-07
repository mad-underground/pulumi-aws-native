// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamProcessingConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-processors
        /// </summary>
        public readonly ImmutableArray<Outputs.DeliveryStreamProcessor> Processors;

        [OutputConstructor]
        private DeliveryStreamProcessingConfiguration(
            bool? Enabled,

            ImmutableArray<Outputs.DeliveryStreamProcessor> Processors)
        {
            this.Enabled = Enabled;
            this.Processors = Processors;
        }
    }
}