// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Forecast.Outputs
{

    [OutputType]
    public sealed class SchemaProperties
    {
        public readonly ImmutableArray<Outputs.AttributesItemProperties> Attributes;

        [OutputConstructor]
        private SchemaProperties(ImmutableArray<Outputs.AttributesItemProperties> attributes)
        {
            Attributes = attributes;
        }
    }
}
