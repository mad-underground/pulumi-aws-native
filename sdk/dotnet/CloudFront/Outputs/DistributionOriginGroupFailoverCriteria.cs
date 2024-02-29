// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    /// <summary>
    /// A complex data type that includes information about the failover criteria for an origin group, including the status codes for which CloudFront will failover from the primary origin to the second origin.
    /// </summary>
    [OutputType]
    public sealed class DistributionOriginGroupFailoverCriteria
    {
        /// <summary>
        /// The status codes that, when returned from the primary origin, will trigger CloudFront to failover to the second origin.
        /// </summary>
        public readonly Outputs.DistributionStatusCodes StatusCodes;

        [OutputConstructor]
        private DistributionOriginGroupFailoverCriteria(Outputs.DistributionStatusCodes statusCodes)
        {
            StatusCodes = statusCodes;
        }
    }
}
