// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Inputs
{

    /// <summary>
    /// A Schema specifies the expected data model of the table.
    /// </summary>
    public sealed class SchemaPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("compositePartitionKey")]
        private InputList<Inputs.TablePartitionKeyArgs>? _compositePartitionKey;
        public InputList<Inputs.TablePartitionKeyArgs> CompositePartitionKey
        {
            get => _compositePartitionKey ?? (_compositePartitionKey = new InputList<Inputs.TablePartitionKeyArgs>());
            set => _compositePartitionKey = value;
        }

        public SchemaPropertiesArgs()
        {
        }
        public static new SchemaPropertiesArgs Empty => new SchemaPropertiesArgs();
    }
}
