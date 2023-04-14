// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    [OutputType]
    public sealed class MissionProfileStreamsKmsKey
    {
        public readonly string? KmsAliasArn;
        public readonly string? KmsKeyArn;

        [OutputConstructor]
        private MissionProfileStreamsKmsKey(
            string? kmsAliasArn,

            string? kmsKeyArn)
        {
            KmsAliasArn = kmsAliasArn;
            KmsKeyArn = kmsKeyArn;
        }
    }
}