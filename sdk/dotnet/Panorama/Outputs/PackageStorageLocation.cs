// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Panorama.Outputs
{

    [OutputType]
    public sealed class PackageStorageLocation
    {
        public readonly string? BinaryPrefixLocation;
        public readonly string? Bucket;
        public readonly string? GeneratedPrefixLocation;
        public readonly string? ManifestPrefixLocation;
        public readonly string? RepoPrefixLocation;

        [OutputConstructor]
        private PackageStorageLocation(
            string? binaryPrefixLocation,

            string? bucket,

            string? generatedPrefixLocation,

            string? manifestPrefixLocation,

            string? repoPrefixLocation)
        {
            BinaryPrefixLocation = binaryPrefixLocation;
            Bucket = bucket;
            GeneratedPrefixLocation = generatedPrefixLocation;
            ManifestPrefixLocation = manifestPrefixLocation;
            RepoPrefixLocation = repoPrefixLocation;
        }
    }
}