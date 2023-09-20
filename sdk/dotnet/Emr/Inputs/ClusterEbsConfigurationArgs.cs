// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr.Inputs
{

    public sealed class ClusterEbsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("ebsBlockDeviceConfigs")]
        private InputList<Inputs.ClusterEbsBlockDeviceConfigArgs>? _ebsBlockDeviceConfigs;
        public InputList<Inputs.ClusterEbsBlockDeviceConfigArgs> EbsBlockDeviceConfigs
        {
            get => _ebsBlockDeviceConfigs ?? (_ebsBlockDeviceConfigs = new InputList<Inputs.ClusterEbsBlockDeviceConfigArgs>());
            set => _ebsBlockDeviceConfigs = value;
        }

        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        public ClusterEbsConfigurationArgs()
        {
        }
        public static new ClusterEbsConfigurationArgs Empty => new ClusterEbsConfigurationArgs();
    }
}