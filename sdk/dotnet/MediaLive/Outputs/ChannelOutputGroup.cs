// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelOutputGroup
    {
        public readonly string? Name;
        public readonly Outputs.ChannelOutputGroupSettings? OutputGroupSettings;
        public readonly ImmutableArray<Outputs.ChannelOutput> Outputs;

        [OutputConstructor]
        private ChannelOutputGroup(
            string? name,

            Outputs.ChannelOutputGroupSettings? outputGroupSettings,

            ImmutableArray<Outputs.ChannelOutput> outputs)
        {
            Name = name;
            OutputGroupSettings = outputGroupSettings;
            Outputs = outputs;
        }
    }
}
