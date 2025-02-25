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
    public sealed class FunctionConfig
    {
        public readonly string Comment;
        public readonly ImmutableArray<Outputs.FunctionKeyValueStoreAssociation> KeyValueStoreAssociations;
        public readonly string Runtime;

        [OutputConstructor]
        private FunctionConfig(
            string comment,

            ImmutableArray<Outputs.FunctionKeyValueStoreAssociation> keyValueStoreAssociations,

            string runtime)
        {
            Comment = comment;
            KeyValueStoreAssociations = keyValueStoreAssociations;
            Runtime = runtime;
        }
    }
}
