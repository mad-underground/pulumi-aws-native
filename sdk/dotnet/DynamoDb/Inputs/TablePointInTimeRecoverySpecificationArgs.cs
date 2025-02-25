// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Inputs
{

    /// <summary>
    /// The settings used to enable point in time recovery.
    /// </summary>
    public sealed class TablePointInTimeRecoverySpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
        /// </summary>
        [Input("pointInTimeRecoveryEnabled")]
        public Input<bool>? PointInTimeRecoveryEnabled { get; set; }

        public TablePointInTimeRecoverySpecificationArgs()
        {
        }
        public static new TablePointInTimeRecoverySpecificationArgs Empty => new TablePointInTimeRecoverySpecificationArgs();
    }
}
