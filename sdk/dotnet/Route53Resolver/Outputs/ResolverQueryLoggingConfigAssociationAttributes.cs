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
    public sealed class ResolverQueryLoggingConfigAssociationAttributes
    {
        public readonly string CreationTime;
        public readonly string Error;
        public readonly string ErrorMessage;
        public readonly string Id;
        public readonly string Status;

        [OutputConstructor]
        private ResolverQueryLoggingConfigAssociationAttributes(
            string CreationTime,

            string Error,

            string ErrorMessage,

            string Id,

            string Status)
        {
            this.CreationTime = CreationTime;
            this.Error = Error;
            this.ErrorMessage = ErrorMessage;
            this.Id = Id;
            this.Status = Status;
        }
    }
}
