// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.M2.Inputs
{

    /// <summary>
    /// Defines the details of a high availability configuration.
    /// </summary>
    public sealed class EnvironmentHighAvailabilityConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("desiredCapacity", required: true)]
        public Input<int> DesiredCapacity { get; set; } = null!;

        public EnvironmentHighAvailabilityConfigArgs()
        {
        }
        public static new EnvironmentHighAvailabilityConfigArgs Empty => new EnvironmentHighAvailabilityConfigArgs();
    }
}
