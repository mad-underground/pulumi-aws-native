// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Outputs
{

    [OutputType]
    public sealed class PatchBaselinePatchSource
    {
        public readonly string? Configuration;
        public readonly string? Name;
        public readonly ImmutableArray<string> Products;

        [OutputConstructor]
        private PatchBaselinePatchSource(
            string? configuration,

            string? name,

            ImmutableArray<string> products)
        {
            Configuration = configuration;
            Name = name;
            Products = products;
        }
    }
}