// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KMS.Outputs
{

    [OutputType]
    public sealed class KeyAttributes
    {
        public readonly string Arn;
        public readonly string KeyId;

        [OutputConstructor]
        private KeyAttributes(
            string Arn,

            string KeyId)
        {
            this.Arn = Arn;
            this.KeyId = KeyId;
        }
    }
}
