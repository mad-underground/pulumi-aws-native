// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Outputs
{

    /// <summary>
    /// Represents the settings used to enable or disable Time to Live (TTL) for the specified table.
    /// </summary>
    [OutputType]
    public sealed class TableTimeToLiveSpecification
    {
        /// <summary>
        /// The name of the TTL attribute used to store the expiration time for items in the table.
        ///    + The ``AttributeName`` property is required when enabling the TTL, or when TTL is already enabled.
        ///   +  To update this property, you must first disable TTL and then enable TTL with the new attribute name.
        /// </summary>
        public readonly string? AttributeName;
        /// <summary>
        /// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private TableTimeToLiveSpecification(
            string? attributeName,

            bool enabled)
        {
            AttributeName = attributeName;
            Enabled = enabled;
        }
    }
}
