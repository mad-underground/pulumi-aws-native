// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetFeatureGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::FeatureGroup
        /// </summary>
        public static Task<GetFeatureGroupResult> InvokeAsync(GetFeatureGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeatureGroupResult>("aws-native:sagemaker:getFeatureGroup", args ?? new GetFeatureGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::FeatureGroup
        /// </summary>
        public static Output<GetFeatureGroupResult> Invoke(GetFeatureGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeatureGroupResult>("aws-native:sagemaker:getFeatureGroup", args ?? new GetFeatureGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the FeatureGroup.
        /// </summary>
        [Input("featureGroupName", required: true)]
        public string FeatureGroupName { get; set; } = null!;

        public GetFeatureGroupArgs()
        {
        }
        public static new GetFeatureGroupArgs Empty => new GetFeatureGroupArgs();
    }

    public sealed class GetFeatureGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the FeatureGroup.
        /// </summary>
        [Input("featureGroupName", required: true)]
        public Input<string> FeatureGroupName { get; set; } = null!;

        public GetFeatureGroupInvokeArgs()
        {
        }
        public static new GetFeatureGroupInvokeArgs Empty => new GetFeatureGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeatureGroupResult
    {
        /// <summary>
        /// A timestamp of FeatureGroup creation time.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// An Array of Feature Definition
        /// </summary>
        public readonly ImmutableArray<Outputs.FeatureGroupFeatureDefinition> FeatureDefinitions;
        /// <summary>
        /// The status of the feature group.
        /// </summary>
        public readonly string? FeatureGroupStatus;
        public readonly Outputs.OnlineStoreConfigProperties? OnlineStoreConfig;
        public readonly Outputs.FeatureGroupThroughputConfig? ThroughputConfig;

        [OutputConstructor]
        private GetFeatureGroupResult(
            string? creationTime,

            ImmutableArray<Outputs.FeatureGroupFeatureDefinition> featureDefinitions,

            string? featureGroupStatus,

            Outputs.OnlineStoreConfigProperties? onlineStoreConfig,

            Outputs.FeatureGroupThroughputConfig? throughputConfig)
        {
            CreationTime = creationTime;
            FeatureDefinitions = featureDefinitions;
            FeatureGroupStatus = featureGroupStatus;
            OnlineStoreConfig = onlineStoreConfig;
            ThroughputConfig = throughputConfig;
        }
    }
}
