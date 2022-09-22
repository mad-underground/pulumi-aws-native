// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetDomain
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Domain
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws-native:sagemaker:getDomain", args ?? new GetDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::Domain
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainResult>("aws-native:sagemaker:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        public GetDomainArgs()
        {
        }
        public static new GetDomainArgs Empty => new GetDomainArgs();
    }

    public sealed class GetDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        public GetDomainInvokeArgs()
        {
        }
        public static new GetDomainInvokeArgs Empty => new GetDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// The default user settings.
        /// </summary>
        public readonly Outputs.DomainUserSettings? DefaultUserSettings;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the created domain.
        /// </summary>
        public readonly string? DomainArn;
        /// <summary>
        /// The domain name.
        /// </summary>
        public readonly string? DomainId;
        public readonly Outputs.DomainSettings? DomainSettings;
        /// <summary>
        /// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
        /// </summary>
        public readonly string? HomeEfsFileSystemId;
        /// <summary>
        /// The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.
        /// </summary>
        public readonly string? SecurityGroupIdForDomainBoundary;
        /// <summary>
        /// The SSO managed application instance ID.
        /// </summary>
        public readonly string? SingleSignOnManagedApplicationInstanceId;
        /// <summary>
        /// The URL to the created domain.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private GetDomainResult(
            Outputs.DomainUserSettings? defaultUserSettings,

            string? domainArn,

            string? domainId,

            Outputs.DomainSettings? domainSettings,

            string? homeEfsFileSystemId,

            string? securityGroupIdForDomainBoundary,

            string? singleSignOnManagedApplicationInstanceId,

            string? url)
        {
            DefaultUserSettings = defaultUserSettings;
            DomainArn = domainArn;
            DomainId = domainId;
            DomainSettings = domainSettings;
            HomeEfsFileSystemId = homeEfsFileSystemId;
            SecurityGroupIdForDomainBoundary = securityGroupIdForDomainBoundary;
            SingleSignOnManagedApplicationInstanceId = singleSignOnManagedApplicationInstanceId;
            Url = url;
        }
    }
}
