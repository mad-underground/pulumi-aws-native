// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class ComponentVersionComponentPlatform
    {
        public readonly ImmutableDictionary<string, string>? Attributes;
        public readonly string? Name;

        [OutputConstructor]
        private ComponentVersionComponentPlatform(
            ImmutableDictionary<string, string>? attributes,

            string? name)
        {
            Attributes = attributes;
            Name = name;
        }
    }
}
