// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location
{
    public static class GetGeofenceCollection
    {
        /// <summary>
        /// Definition of AWS::Location::GeofenceCollection Resource Type
        /// </summary>
        public static Task<GetGeofenceCollectionResult> InvokeAsync(GetGeofenceCollectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGeofenceCollectionResult>("aws-native:location:getGeofenceCollection", args ?? new GetGeofenceCollectionArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Location::GeofenceCollection Resource Type
        /// </summary>
        public static Output<GetGeofenceCollectionResult> Invoke(GetGeofenceCollectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGeofenceCollectionResult>("aws-native:location:getGeofenceCollection", args ?? new GetGeofenceCollectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGeofenceCollectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("collectionName", required: true)]
        public string CollectionName { get; set; } = null!;

        public GetGeofenceCollectionArgs()
        {
        }
        public static new GetGeofenceCollectionArgs Empty => new GetGeofenceCollectionArgs();
    }

    public sealed class GetGeofenceCollectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("collectionName", required: true)]
        public Input<string> CollectionName { get; set; } = null!;

        public GetGeofenceCollectionInvokeArgs()
        {
        }
        public static new GetGeofenceCollectionInvokeArgs Empty => new GetGeofenceCollectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetGeofenceCollectionResult
    {
        public readonly string? Arn;
        public readonly string? CollectionArn;
        public readonly string? CreateTime;
        public readonly string? Description;
        public readonly Pulumi.AwsNative.Location.GeofenceCollectionPricingPlan? PricingPlan;
        public readonly string? PricingPlanDataSource;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private GetGeofenceCollectionResult(
            string? arn,

            string? collectionArn,

            string? createTime,

            string? description,

            Pulumi.AwsNative.Location.GeofenceCollectionPricingPlan? pricingPlan,

            string? pricingPlanDataSource,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? updateTime)
        {
            Arn = arn;
            CollectionArn = collectionArn;
            CreateTime = createTime;
            Description = description;
            PricingPlan = pricingPlan;
            PricingPlanDataSource = pricingPlanDataSource;
            Tags = tags;
            UpdateTime = updateTime;
        }
    }
}
