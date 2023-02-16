// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class DataSourceDynamoDBConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsRegion", required: true)]
        public Input<string> AwsRegion { get; set; } = null!;

        [Input("deltaSyncConfig")]
        public Input<Inputs.DataSourceDeltaSyncConfigArgs>? DeltaSyncConfig { get; set; }

        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        [Input("useCallerCredentials")]
        public Input<bool>? UseCallerCredentials { get; set; }

        [Input("versioned")]
        public Input<bool>? Versioned { get; set; }

        public DataSourceDynamoDBConfigArgs()
        {
        }
        public static new DataSourceDynamoDBConfigArgs Empty => new DataSourceDynamoDBConfigArgs();
    }
}
