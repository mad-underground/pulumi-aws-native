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
    /// Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
    /// </summary>
    [OutputType]
    public sealed class TableProjection
    {
        /// <summary>
        /// Represents the non-key attribute names which will be projected into the index.
        ///  For local secondary indexes, the total count of ``NonKeyAttributes`` summed across all of the local secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.
        /// </summary>
        public readonly ImmutableArray<string> NonKeyAttributes;
        /// <summary>
        /// The set of attributes that are projected into the index:
        ///   +   ``KEYS_ONLY`` - Only the index and primary keys are projected into the index.
        ///   +   ``INCLUDE`` - In addition to the attributes described in ``KEYS_ONLY``, the secondary index will include other non-key attributes that you specify.
        ///   +   ``ALL`` - All of the table attributes are projected into the index.
        ///   
        ///  When using the DynamoDB console, ``ALL`` is selected by default.
        /// </summary>
        public readonly string? ProjectionType;

        [OutputConstructor]
        private TableProjection(
            ImmutableArray<string> nonKeyAttributes,

            string? projectionType)
        {
            NonKeyAttributes = nonKeyAttributes;
            ProjectionType = projectionType;
        }
    }
}
