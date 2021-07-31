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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iottopicpublish.html
    /// </summary>
    [OutputType]
    public sealed class DetectorModelIotTopicPublish
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iottopicpublish.html#cfn-iotevents-detectormodel-iottopicpublish-mqtttopic
        /// </summary>
        public readonly string? MqttTopic;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iottopicpublish.html#cfn-iotevents-detectormodel-iottopicpublish-payload
        /// </summary>
        public readonly Outputs.DetectorModelPayload? Payload;

        [OutputConstructor]
        private DetectorModelIotTopicPublish(
            string? mqttTopic,

            Outputs.DetectorModelPayload? payload)
        {
            MqttTopic = mqttTopic;
            Payload = payload;
        }
    }
}
