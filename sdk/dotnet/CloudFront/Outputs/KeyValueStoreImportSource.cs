// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class KeyValueStoreImportSource
    {
        public readonly string SourceArn;
        public readonly string SourceType;

        [OutputConstructor]
        private KeyValueStoreImportSource(
            string sourceArn,

            string sourceType)
        {
            SourceArn = sourceArn;
            SourceType = sourceType;
        }
    }
}
