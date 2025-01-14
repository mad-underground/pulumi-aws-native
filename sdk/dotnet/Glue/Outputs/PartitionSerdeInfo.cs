// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class PartitionSerdeInfo
    {
        public readonly string? Name;
        public readonly object? Parameters;
        public readonly string? SerializationLibrary;

        [OutputConstructor]
        private PartitionSerdeInfo(
            string? name,

            object? parameters,

            string? serializationLibrary)
        {
            Name = name;
            Parameters = parameters;
            SerializationLibrary = serializationLibrary;
        }
    }
}
