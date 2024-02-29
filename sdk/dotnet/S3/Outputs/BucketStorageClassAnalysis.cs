// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.
    /// </summary>
    [OutputType]
    public sealed class BucketStorageClassAnalysis
    {
        /// <summary>
        /// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
        /// </summary>
        public readonly Outputs.BucketDataExport? DataExport;

        [OutputConstructor]
        private BucketStorageClassAnalysis(Outputs.BucketDataExport? dataExport)
        {
            DataExport = dataExport;
        }
    }
}
