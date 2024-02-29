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
    /// Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.
    /// </summary>
    [OutputType]
    public sealed class BucketInventoryConfiguration
    {
        /// <summary>
        /// Contains information about where to publish the inventory results.
        /// </summary>
        public readonly Outputs.BucketDestination Destination;
        /// <summary>
        /// Specifies whether the inventory is enabled or disabled. If set to ``True``, an inventory list is generated. If set to ``False``, no inventory list is generated.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID used to identify the inventory configuration.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Object versions to include in the inventory list. If set to ``All``, the list includes all the object versions, which adds the version-related fields ``VersionId``, ``IsLatest``, and ``DeleteMarker`` to the list. If set to ``Current``, the list does not contain these version-related fields.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketInventoryConfigurationIncludedObjectVersions IncludedObjectVersions;
        /// <summary>
        /// Contains the optional fields that are included in the inventory results.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.S3.BucketInventoryConfigurationOptionalFieldsItem> OptionalFields;
        /// <summary>
        /// Specifies the inventory filter prefix.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// Specifies the schedule for generating inventory results.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketInventoryConfigurationScheduleFrequency ScheduleFrequency;

        [OutputConstructor]
        private BucketInventoryConfiguration(
            Outputs.BucketDestination destination,

            bool enabled,

            string id,

            Pulumi.AwsNative.S3.BucketInventoryConfigurationIncludedObjectVersions includedObjectVersions,

            ImmutableArray<Pulumi.AwsNative.S3.BucketInventoryConfigurationOptionalFieldsItem> optionalFields,

            string? prefix,

            Pulumi.AwsNative.S3.BucketInventoryConfigurationScheduleFrequency scheduleFrequency)
        {
            Destination = destination;
            Enabled = enabled;
            Id = id;
            IncludedObjectVersions = includedObjectVersions;
            OptionalFields = optionalFields;
            Prefix = prefix;
            ScheduleFrequency = scheduleFrequency;
        }
    }
}
