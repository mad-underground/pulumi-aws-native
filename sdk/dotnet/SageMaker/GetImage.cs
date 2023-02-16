// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetImage
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Image
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("aws-native:sagemaker:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Image
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("aws-native:sagemaker:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        [Input("imageArn", required: true)]
        public string ImageArn { get; set; } = null!;

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("imageArn", required: true)]
        public Input<string> ImageArn { get; set; } = null!;

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        public readonly string? ImageArn;
        public readonly string? ImageDescription;
        public readonly string? ImageDisplayName;
        public readonly string? ImageRoleArn;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImageTag> Tags;

        [OutputConstructor]
        private GetImageResult(
            string? imageArn,

            string? imageDescription,

            string? imageDisplayName,

            string? imageRoleArn,

            ImmutableArray<Outputs.ImageTag> tags)
        {
            ImageArn = imageArn;
            ImageDescription = imageDescription;
            ImageDisplayName = imageDisplayName;
            ImageRoleArn = imageRoleArn;
            Tags = tags;
        }
    }
}
