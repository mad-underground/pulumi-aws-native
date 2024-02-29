// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
    /// </summary>
    public sealed class BucketDataExportArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The place to store the data for an analysis.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.BucketDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// The version of the output schema to use when exporting data. Must be ``V_1``.
        /// </summary>
        [Input("outputSchemaVersion", required: true)]
        public Input<string> OutputSchemaVersion { get; set; } = null!;

        public BucketDataExportArgs()
        {
        }
        public static new BucketDataExportArgs Empty => new BucketDataExportArgs();
    }
}
