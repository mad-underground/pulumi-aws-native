// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamDeserializerArgs : global::Pulumi.ResourceArgs
    {
        [Input("hiveJsonSerDe")]
        public Input<Inputs.DeliveryStreamHiveJsonSerDeArgs>? HiveJsonSerDe { get; set; }

        [Input("openXJsonSerDe")]
        public Input<Inputs.DeliveryStreamOpenXJsonSerDeArgs>? OpenXJsonSerDe { get; set; }

        public DeliveryStreamDeserializerArgs()
        {
        }
        public static new DeliveryStreamDeserializerArgs Empty => new DeliveryStreamDeserializerArgs();
    }
}
