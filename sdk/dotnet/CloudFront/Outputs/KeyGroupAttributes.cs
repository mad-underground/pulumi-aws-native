// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFront.Outputs
{

    [OutputType]
    public sealed class KeyGroupAttributes
    {
        public readonly string Id;
        public readonly string LastModifiedTime;

        [OutputConstructor]
        private KeyGroupAttributes(
            string Id,

            string LastModifiedTime)
        {
            this.Id = Id;
            this.LastModifiedTime = LastModifiedTime;
        }
    }
}
