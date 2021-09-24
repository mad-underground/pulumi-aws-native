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
    public sealed class NodeNodeConfiguration
    {
        public readonly string AvailabilityZone;
        public readonly string InstanceType;

        [OutputConstructor]
        private NodeNodeConfiguration(
            string availabilityZone,

            string instanceType)
        {
            AvailabilityZone = availabilityZone;
            InstanceType = instanceType;
        }
    }
}