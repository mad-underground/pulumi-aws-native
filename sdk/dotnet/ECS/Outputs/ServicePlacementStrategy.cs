// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class ServicePlacementStrategy
    {
        public readonly string? Field;
        public readonly Pulumi.AwsNative.ECS.ServicePlacementStrategyType Type;

        [OutputConstructor]
        private ServicePlacementStrategy(
            string? field,

            Pulumi.AwsNative.ECS.ServicePlacementStrategyType type)
        {
            Field = field;
            Type = type;
        }
    }
}
