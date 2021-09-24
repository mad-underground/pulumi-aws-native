// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Outputs
{

    [OutputType]
    public sealed class EndpointS3Settings
    {
        public readonly string? BucketFolder;
        public readonly string? BucketName;
        public readonly string? CompressionType;
        public readonly string? CsvDelimiter;
        public readonly string? CsvRowDelimiter;
        public readonly string? ExternalTableDefinition;
        public readonly string? ServiceAccessRoleArn;

        [OutputConstructor]
        private EndpointS3Settings(
            string? bucketFolder,

            string? bucketName,

            string? compressionType,

            string? csvDelimiter,

            string? csvRowDelimiter,

            string? externalTableDefinition,

            string? serviceAccessRoleArn)
        {
            BucketFolder = bucketFolder;
            BucketName = bucketName;
            CompressionType = compressionType;
            CsvDelimiter = csvDelimiter;
            CsvRowDelimiter = csvRowDelimiter;
            ExternalTableDefinition = externalTableDefinition;
            ServiceAccessRoleArn = serviceAccessRoleArn;
        }
    }
}