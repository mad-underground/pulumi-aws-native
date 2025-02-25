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
    /// &lt;p&gt;The configuration for a streaming session’s upload storage.&lt;/p&gt;
    /// </summary>
    public sealed class LaunchProfileStreamConfigurationSessionStorageArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode", required: true)]
        private InputList<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingSessionStorageMode>? _mode;

        /// <summary>
        /// &lt;p&gt;Allows artists to upload files to their workstations. The only valid option is
        ///                 &lt;code&gt;UPLOAD&lt;/code&gt;.&lt;/p&gt;
        /// </summary>
        public InputList<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingSessionStorageMode> Mode
        {
            get => _mode ?? (_mode = new InputList<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingSessionStorageMode>());
            set => _mode = value;
        }

        [Input("root")]
        public Input<Inputs.LaunchProfileStreamingSessionStorageRootArgs>? Root { get; set; }

        public LaunchProfileStreamConfigurationSessionStorageArgs()
        {
        }
        public static new LaunchProfileStreamConfigurationSessionStorageArgs Empty => new LaunchProfileStreamConfigurationSessionStorageArgs();
    }
}
