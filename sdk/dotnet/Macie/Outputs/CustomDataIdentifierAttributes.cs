// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Macie.Outputs
{

    [OutputType]
    public sealed class CustomDataIdentifierAttributes
    {
        public readonly string Arn;
        public readonly string CreatedAt;
        public readonly bool Deleted;
        public readonly string Id;

        [OutputConstructor]
        private CustomDataIdentifierAttributes(
            string Arn,

            string CreatedAt,

            bool Deleted,

            string Id)
        {
            this.Arn = Arn;
            this.CreatedAt = CreatedAt;
            this.Deleted = Deleted;
            this.Id = Id;
        }
    }
}