// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Outputs
{

    [OutputType]
    public sealed class DeviceFleetAttributes
    {
        public readonly string DeviceFleetName;

        [OutputConstructor]
        private DeviceFleetAttributes(string DeviceFleetName)
        {
            this.DeviceFleetName = DeviceFleetName;
        }
    }
}