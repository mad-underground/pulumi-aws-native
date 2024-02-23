// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2.Outputs
{

    [OutputType]
    public sealed class CisScanConfigurationCisTargets
    {
        public readonly ImmutableArray<string> AccountIds;
        public readonly ImmutableDictionary<string, object>? TargetResourceTags;

        [OutputConstructor]
        private CisScanConfigurationCisTargets(
            ImmutableArray<string> accountIds,

            ImmutableDictionary<string, object>? targetResourceTags)
        {
            AccountIds = accountIds;
            TargetResourceTags = targetResourceTags;
        }
    }
}
