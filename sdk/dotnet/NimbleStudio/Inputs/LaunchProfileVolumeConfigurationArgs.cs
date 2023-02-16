// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    /// <summary>
    /// &lt;p&gt;Custom volume configuration for the root volumes that are attached to streaming
    ///             sessions.&lt;/p&gt;
    ///          &lt;p&gt;This parameter is only allowed when &lt;code&gt;sessionPersistenceMode&lt;/code&gt; is
    ///                 &lt;code&gt;ACTIVATED&lt;/code&gt;.&lt;/p&gt;
    /// </summary>
    public sealed class LaunchProfileVolumeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The number of I/O operations per second for the root volume that is attached to
        ///             streaming session.&lt;/p&gt;
        /// </summary>
        [Input("iops")]
        public Input<double>? Iops { get; set; }

        /// <summary>
        /// &lt;p&gt;The size of the root volume that is attached to the streaming session. The root volume
        ///             size is measured in GiBs.&lt;/p&gt;
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        /// <summary>
        /// &lt;p&gt;The throughput to provision for the root volume that is attached to the streaming
        ///             session. The throughput is measured in MiB/s.&lt;/p&gt;
        /// </summary>
        [Input("throughput")]
        public Input<double>? Throughput { get; set; }

        public LaunchProfileVolumeConfigurationArgs()
        {
        }
        public static new LaunchProfileVolumeConfigurationArgs Empty => new LaunchProfileVolumeConfigurationArgs();
    }
}
