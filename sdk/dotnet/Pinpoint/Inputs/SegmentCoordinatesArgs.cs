// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class SegmentCoordinatesArgs : global::Pulumi.ResourceArgs
    {
        [Input("latitude", required: true)]
        public Input<double> Latitude { get; set; } = null!;

        [Input("longitude", required: true)]
        public Input<double> Longitude { get; set; } = null!;

        public SegmentCoordinatesArgs()
        {
        }
        public static new SegmentCoordinatesArgs Empty => new SegmentCoordinatesArgs();
    }
}
