// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// The configuration of a Kinesis Data Analytics Studio notebook.
    /// </summary>
    public sealed class ApplicationZeppelinApplicationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics Studio notebook.
        /// </summary>
        [Input("catalogConfiguration")]
        public Input<Inputs.ApplicationCatalogConfigurationArgs>? CatalogConfiguration { get; set; }

        [Input("customArtifactsConfiguration")]
        private InputList<Inputs.ApplicationCustomArtifactConfigurationArgs>? _customArtifactsConfiguration;

        /// <summary>
        /// A list of CustomArtifactConfiguration objects.
        /// </summary>
        public InputList<Inputs.ApplicationCustomArtifactConfigurationArgs> CustomArtifactsConfiguration
        {
            get => _customArtifactsConfiguration ?? (_customArtifactsConfiguration = new InputList<Inputs.ApplicationCustomArtifactConfigurationArgs>());
            set => _customArtifactsConfiguration = value;
        }

        /// <summary>
        /// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
        /// </summary>
        [Input("deployAsApplicationConfiguration")]
        public Input<Inputs.ApplicationDeployAsApplicationConfigurationArgs>? DeployAsApplicationConfiguration { get; set; }

        /// <summary>
        /// The monitoring configuration of a Kinesis Data Analytics Studio notebook.
        /// </summary>
        [Input("monitoringConfiguration")]
        public Input<Inputs.ApplicationZeppelinMonitoringConfigurationArgs>? MonitoringConfiguration { get; set; }

        public ApplicationZeppelinApplicationConfigurationArgs()
        {
        }
    }
}
