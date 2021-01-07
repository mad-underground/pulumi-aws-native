// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Route53Resolver.Outputs
{

    [OutputType]
    public sealed class ResolverQueryLoggingConfigAttributes
    {
        public readonly string Arn;
        public readonly int AssociationCount;
        public readonly string CreationTime;
        public readonly string CreatorRequestId;
        public readonly string Id;
        public readonly string OwnerId;
        public readonly string ShareStatus;
        public readonly string Status;

        [OutputConstructor]
        private ResolverQueryLoggingConfigAttributes(
            string Arn,

            int AssociationCount,

            string CreationTime,

            string CreatorRequestId,

            string Id,

            string OwnerId,

            string ShareStatus,

            string Status)
        {
            this.Arn = Arn;
            this.AssociationCount = AssociationCount;
            this.CreationTime = CreationTime;
            this.CreatorRequestId = CreatorRequestId;
            this.Id = Id;
            this.OwnerId = OwnerId;
            this.ShareStatus = ShareStatus;
            this.Status = Status;
        }
    }
}
