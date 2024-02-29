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
    /// A filter that you can specify for selection for modifications on replicas.
    /// </summary>
    [OutputType]
    public sealed class BucketReplicaModifications
    {
        /// <summary>
        /// Specifies whether Amazon S3 replicates modifications on replicas.
        ///  *Allowed values*: ``Enabled`` | ``Disabled``
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketReplicaModificationsStatus Status;

        [OutputConstructor]
        private BucketReplicaModifications(Pulumi.AwsNative.S3.BucketReplicaModificationsStatus status)
        {
            Status = status;
        }
    }
}
