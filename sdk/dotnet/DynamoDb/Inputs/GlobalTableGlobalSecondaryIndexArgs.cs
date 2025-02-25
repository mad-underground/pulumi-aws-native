// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Inputs
{

    public sealed class GlobalTableGlobalSecondaryIndexArgs : global::Pulumi.ResourceArgs
    {
        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        [Input("keySchema", required: true)]
        private InputList<Inputs.GlobalTableKeySchemaArgs>? _keySchema;
        public InputList<Inputs.GlobalTableKeySchemaArgs> KeySchema
        {
            get => _keySchema ?? (_keySchema = new InputList<Inputs.GlobalTableKeySchemaArgs>());
            set => _keySchema = value;
        }

        [Input("projection", required: true)]
        public Input<Inputs.GlobalTableProjectionArgs> Projection { get; set; } = null!;

        [Input("writeProvisionedThroughputSettings")]
        public Input<Inputs.GlobalTableWriteProvisionedThroughputSettingsArgs>? WriteProvisionedThroughputSettings { get; set; }

        public GlobalTableGlobalSecondaryIndexArgs()
        {
        }
        public static new GlobalTableGlobalSecondaryIndexArgs Empty => new GlobalTableGlobalSecondaryIndexArgs();
    }
}
