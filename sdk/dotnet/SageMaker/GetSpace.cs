// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetSpace
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Space
        /// </summary>
        public static Task<GetSpaceResult> InvokeAsync(GetSpaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSpaceResult>("aws-native:sagemaker:getSpace", args ?? new GetSpaceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Space
        /// </summary>
        public static Output<GetSpaceResult> Invoke(GetSpaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSpaceResult>("aws-native:sagemaker:getSpace", args ?? new GetSpaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSpaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the associated Domain.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// A name for the Space.
        /// </summary>
        [Input("spaceName", required: true)]
        public string SpaceName { get; set; } = null!;

        public GetSpaceArgs()
        {
        }
        public static new GetSpaceArgs Empty => new GetSpaceArgs();
    }

    public sealed class GetSpaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the associated Domain.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// A name for the Space.
        /// </summary>
        [Input("spaceName", required: true)]
        public Input<string> SpaceName { get; set; } = null!;

        public GetSpaceInvokeArgs()
        {
        }
        public static new GetSpaceInvokeArgs Empty => new GetSpaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSpaceResult
    {
        /// <summary>
        /// The space Amazon Resource Name (ARN).
        /// </summary>
        public readonly string? SpaceArn;
        public readonly string? SpaceDisplayName;
        public readonly string? Url;

        [OutputConstructor]
        private GetSpaceResult(
            string? spaceArn,

            string? spaceDisplayName,

            string? url)
        {
            SpaceArn = spaceArn;
            SpaceDisplayName = spaceDisplayName;
            Url = url;
        }
    }
}
