// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class MultiRegionAccessPointRegion
    {
        public readonly string Bucket;

        [OutputConstructor]
        private MultiRegionAccessPointRegion(string bucket)
        {
            Bucket = bucket;
        }
    }
}