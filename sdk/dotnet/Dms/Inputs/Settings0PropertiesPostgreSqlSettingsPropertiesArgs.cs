// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Inputs
{

    public sealed class Settings0PropertiesPostgreSqlSettingsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        [Input("sslMode")]
        public Input<Pulumi.AwsNative.Dms.DataProviderDmsSslModeValue>? SslMode { get; set; }

        public Settings0PropertiesPostgreSqlSettingsPropertiesArgs()
        {
        }
        public static new Settings0PropertiesPostgreSqlSettingsPropertiesArgs Empty => new Settings0PropertiesPostgreSqlSettingsPropertiesArgs();
    }
}