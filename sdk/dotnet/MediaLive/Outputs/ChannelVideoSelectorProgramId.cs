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
    public sealed class ChannelVideoSelectorProgramId
    {
        public readonly int? ProgramId;

        [OutputConstructor]
        private ChannelVideoSelectorProgramId(int? programId)
        {
            ProgramId = programId;
        }
    }
}