// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    [OutputType]
    public sealed class EnvironmentBlueprintConfigurationRegionalParameter
    {
        public readonly ImmutableDictionary<string, string>? Parameters;
        public readonly string? Region;

        [OutputConstructor]
        private EnvironmentBlueprintConfigurationRegionalParameter(
            ImmutableDictionary<string, string>? parameters,

            string? region)
        {
            Parameters = parameters;
            Region = region;
        }
    }
}
