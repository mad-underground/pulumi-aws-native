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
    public sealed class PublicKeyAttributes
    {
        public readonly string CreatedTime;
        public readonly string Id;

        [OutputConstructor]
        private PublicKeyAttributes(
            string CreatedTime,

            string Id)
        {
            this.CreatedTime = CreatedTime;
            this.Id = Id;
        }
    }
}
