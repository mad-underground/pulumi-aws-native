// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GuardDuty.Inputs
{

    public sealed class DetectorCfnDataSourceConfigurationsArgs : global::Pulumi.ResourceArgs
    {
        [Input("kubernetes")]
        public Input<Inputs.DetectorCfnKubernetesConfigurationArgs>? Kubernetes { get; set; }

        [Input("malwareProtection")]
        public Input<Inputs.DetectorCfnMalwareProtectionConfigurationArgs>? MalwareProtection { get; set; }

        [Input("s3Logs")]
        public Input<Inputs.DetectorCfns3LogsConfigurationArgs>? S3Logs { get; set; }

        public DetectorCfnDataSourceConfigurationsArgs()
        {
        }
        public static new DetectorCfnDataSourceConfigurationsArgs Empty => new DetectorCfnDataSourceConfigurationsArgs();
    }
}