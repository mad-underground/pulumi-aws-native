// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class SpotFleetMemoryGiBPerVCpuRequest
    {
        public readonly double? Max;
        public readonly double? Min;

        [OutputConstructor]
        private SpotFleetMemoryGiBPerVCpuRequest(
            double? max,

            double? min)
        {
            Max = max;
            Min = min;
        }
    }
}
