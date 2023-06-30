// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class SegmentGroups
    {
        public readonly ImmutableArray<Outputs.Groups> Groups;
        public readonly string? Include;

        [OutputConstructor]
        private SegmentGroups(
            ImmutableArray<Outputs.Groups> groups,

            string? include)
        {
            Groups = groups;
            Include = include;
        }
    }
}
