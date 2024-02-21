// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NeptuneGraph.Inputs
{

    /// <summary>
    /// The vector search configuration.
    /// </summary>
    public sealed class GraphVectorSearchConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The vector search dimension
        /// </summary>
        [Input("vectorSearchDimension", required: true)]
        public Input<int> VectorSearchDimension { get; set; } = null!;

        public GraphVectorSearchConfigurationArgs()
        {
        }
        public static new GraphVectorSearchConfigurationArgs Empty => new GraphVectorSearchConfigurationArgs();
    }
}