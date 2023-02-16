// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda
{
    public static class GetAccessPoint
    {
        /// <summary>
        /// The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
        /// </summary>
        public static Task<GetAccessPointResult> InvokeAsync(GetAccessPointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessPointResult>("aws-native:s3objectlambda:getAccessPoint", args ?? new GetAccessPointArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
        /// </summary>
        public static Output<GetAccessPointResult> Invoke(GetAccessPointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessPointResult>("aws-native:s3objectlambda:getAccessPoint", args ?? new GetAccessPointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessPointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name you want to assign to this Object lambda Access Point.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetAccessPointArgs()
        {
        }
        public static new GetAccessPointArgs Empty => new GetAccessPointArgs();
    }

    public sealed class GetAccessPointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name you want to assign to this Object lambda Access Point.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetAccessPointInvokeArgs()
        {
        }
        public static new GetAccessPointInvokeArgs Empty => new GetAccessPointInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessPointResult
    {
        public readonly string? Arn;
        /// <summary>
        /// The date and time when the Object lambda Access Point was created.
        /// </summary>
        public readonly string? CreationDate;
        /// <summary>
        /// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
        /// </summary>
        public readonly Outputs.AccessPointObjectLambdaConfiguration? ObjectLambdaConfiguration;
        public readonly Outputs.PolicyStatusProperties? PolicyStatus;
        /// <summary>
        /// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
        /// </summary>
        public readonly Outputs.AccessPointPublicAccessBlockConfiguration? PublicAccessBlockConfiguration;

        [OutputConstructor]
        private GetAccessPointResult(
            string? arn,

            string? creationDate,

            Outputs.AccessPointObjectLambdaConfiguration? objectLambdaConfiguration,

            Outputs.PolicyStatusProperties? policyStatus,

            Outputs.AccessPointPublicAccessBlockConfiguration? publicAccessBlockConfiguration)
        {
            Arn = arn;
            CreationDate = creationDate;
            ObjectLambdaConfiguration = objectLambdaConfiguration;
            PolicyStatus = policyStatus;
            PublicAccessBlockConfiguration = publicAccessBlockConfiguration;
        }
    }
}
