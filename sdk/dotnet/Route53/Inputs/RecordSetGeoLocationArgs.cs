// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    public sealed class RecordSetGeoLocationArgs : global::Pulumi.ResourceArgs
    {
        [Input("continentCode")]
        public Input<string>? ContinentCode { get; set; }

        [Input("countryCode")]
        public Input<string>? CountryCode { get; set; }

        [Input("subdivisionCode")]
        public Input<string>? SubdivisionCode { get; set; }

        public RecordSetGeoLocationArgs()
        {
        }
        public static new RecordSetGeoLocationArgs Empty => new RecordSetGeoLocationArgs();
    }
}
