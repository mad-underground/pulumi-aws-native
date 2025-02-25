// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The configurations to be set at cluster level.
    /// </summary>
    public sealed class ClusterConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("executeCommandConfiguration")]
        public Input<Inputs.ClusterExecuteCommandConfigurationArgs>? ExecuteCommandConfiguration { get; set; }

        public ClusterConfigurationArgs()
        {
        }
        public static new ClusterConfigurationArgs Empty => new ClusterConfigurationArgs();
    }
}
