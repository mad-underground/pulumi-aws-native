// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TopicSemanticEntityType
    {
        public readonly string? SubTypeName;
        public readonly string? TypeName;
        public readonly ImmutableDictionary<string, string>? TypeParameters;

        [OutputConstructor]
        private TopicSemanticEntityType(
            string? subTypeName,

            string? typeName,

            ImmutableDictionary<string, string>? typeParameters)
        {
            SubTypeName = subTypeName;
            TypeName = typeName;
            TypeParameters = typeParameters;
        }
    }
}
