// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class ServiceTimeoutConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("idleTimeoutSeconds")]
        public Input<int>? IdleTimeoutSeconds { get; set; }

        [Input("perRequestTimeoutSeconds")]
        public Input<int>? PerRequestTimeoutSeconds { get; set; }

        public ServiceTimeoutConfigurationArgs()
        {
        }
        public static new ServiceTimeoutConfigurationArgs Empty => new ServiceTimeoutConfigurationArgs();
    }
}
