// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ManagedBlockchain.Outputs
{

    [OutputType]
    public sealed class MemberConfiguration
    {
        public readonly string? Description;
        public readonly Outputs.MemberFrameworkConfiguration? MemberFrameworkConfiguration;
        public readonly string Name;

        [OutputConstructor]
        private MemberConfiguration(
            string? description,

            Outputs.MemberFrameworkConfiguration? memberFrameworkConfiguration,

            string name)
        {
            Description = description;
            MemberFrameworkConfiguration = memberFrameworkConfiguration;
            Name = name;
        }
    }
}
