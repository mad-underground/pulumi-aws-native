// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// The configuration details of the Amazon Redshift data source.
    /// </summary>
    public sealed class DataSourceRedshiftRunConfigurationInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data access role included in the configuration details of the Amazon Redshift data source.
        /// </summary>
        [Input("dataAccessRole")]
        public Input<string>? DataAccessRole { get; set; }

        /// <summary>
        /// The details of the credentials required to access an Amazon Redshift cluster.
        /// </summary>
        [Input("redshiftCredentialConfiguration", required: true)]
        public Input<Inputs.DataSourceRedshiftCredentialConfigurationArgs> RedshiftCredentialConfiguration { get; set; } = null!;

        /// <summary>
        /// The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.
        /// </summary>
        [Input("redshiftStorage", required: true)]
        public InputUnion<Inputs.DataSourceRedshiftStorage0PropertiesArgs, Inputs.DataSourceRedshiftStorage1PropertiesArgs> RedshiftStorage { get; set; } = null!;

        [Input("relationalFilterConfigurations", required: true)]
        private InputList<Inputs.DataSourceRelationalFilterConfigurationArgs>? _relationalFilterConfigurations;
        public InputList<Inputs.DataSourceRelationalFilterConfigurationArgs> RelationalFilterConfigurations
        {
            get => _relationalFilterConfigurations ?? (_relationalFilterConfigurations = new InputList<Inputs.DataSourceRelationalFilterConfigurationArgs>());
            set => _relationalFilterConfigurations = value;
        }

        public DataSourceRedshiftRunConfigurationInputArgs()
        {
        }
        public static new DataSourceRedshiftRunConfigurationInputArgs Empty => new DataSourceRedshiftRunConfigurationInputArgs();
    }
}
