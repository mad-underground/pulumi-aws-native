// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    public sealed class RecordSetGroupGeoProximityLocationArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsRegion")]
        public Input<string>? AwsRegion { get; set; }

        [Input("bias")]
        public Input<int>? Bias { get; set; }

        [Input("coordinates")]
        public Input<Inputs.RecordSetGroupCoordinatesArgs>? Coordinates { get; set; }

        [Input("localZoneGroup")]
        public Input<string>? LocalZoneGroup { get; set; }

        public RecordSetGroupGeoProximityLocationArgs()
        {
        }
        public static new RecordSetGroupGeoProximityLocationArgs Empty => new RecordSetGroupGeoProximityLocationArgs();
    }
}