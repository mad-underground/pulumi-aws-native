// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rekognition
{
    public static class GetCollection
    {
        /// <summary>
        /// The AWS::Rekognition::Collection type creates an Amazon Rekognition Collection. A collection is a logical grouping of information about detected faces which can later be referenced for searches on the group
        /// </summary>
        public static Task<GetCollectionResult> InvokeAsync(GetCollectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCollectionResult>("aws-native:rekognition:getCollection", args ?? new GetCollectionArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Rekognition::Collection type creates an Amazon Rekognition Collection. A collection is a logical grouping of information about detected faces which can later be referenced for searches on the group
        /// </summary>
        public static Output<GetCollectionResult> Invoke(GetCollectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCollectionResult>("aws-native:rekognition:getCollection", args ?? new GetCollectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCollectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("collectionId", required: true)]
        public string CollectionId { get; set; } = null!;

        public GetCollectionArgs()
        {
        }
        public static new GetCollectionArgs Empty => new GetCollectionArgs();
    }

    public sealed class GetCollectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("collectionId", required: true)]
        public Input<string> CollectionId { get; set; } = null!;

        public GetCollectionInvokeArgs()
        {
        }
        public static new GetCollectionInvokeArgs Empty => new GetCollectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetCollectionResult
    {
        public readonly string? Arn;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.CollectionTag> Tags;

        [OutputConstructor]
        private GetCollectionResult(
            string? arn,

            ImmutableArray<Outputs.CollectionTag> tags)
        {
            Arn = arn;
            Tags = tags;
        }
    }
}
