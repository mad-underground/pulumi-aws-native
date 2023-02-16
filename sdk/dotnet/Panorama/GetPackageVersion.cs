// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Panorama
{
    public static class GetPackageVersion
    {
        /// <summary>
        /// Schema for PackageVersion Resource Type
        /// </summary>
        public static Task<GetPackageVersionResult> InvokeAsync(GetPackageVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPackageVersionResult>("aws-native:panorama:getPackageVersion", args ?? new GetPackageVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Schema for PackageVersion Resource Type
        /// </summary>
        public static Output<GetPackageVersionResult> Invoke(GetPackageVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPackageVersionResult>("aws-native:panorama:getPackageVersion", args ?? new GetPackageVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPackageVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("packageId", required: true)]
        public string PackageId { get; set; } = null!;

        [Input("packageVersion", required: true)]
        public string PackageVersionValue { get; set; } = null!;

        [Input("patchVersion", required: true)]
        public string PatchVersion { get; set; } = null!;

        public GetPackageVersionArgs()
        {
        }
        public static new GetPackageVersionArgs Empty => new GetPackageVersionArgs();
    }

    public sealed class GetPackageVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("packageId", required: true)]
        public Input<string> PackageId { get; set; } = null!;

        [Input("packageVersion", required: true)]
        public Input<string> PackageVersion { get; set; } = null!;

        [Input("patchVersion", required: true)]
        public Input<string> PatchVersion { get; set; } = null!;

        public GetPackageVersionInvokeArgs()
        {
        }
        public static new GetPackageVersionInvokeArgs Empty => new GetPackageVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetPackageVersionResult
    {
        public readonly bool? IsLatestPatch;
        public readonly bool? MarkLatest;
        public readonly string? PackageArn;
        public readonly string? PackageName;
        public readonly int? RegisteredTime;
        public readonly Pulumi.AwsNative.Panorama.PackageVersionStatus? Status;
        public readonly string? StatusDescription;
        public readonly string? UpdatedLatestPatchVersion;

        [OutputConstructor]
        private GetPackageVersionResult(
            bool? isLatestPatch,

            bool? markLatest,

            string? packageArn,

            string? packageName,

            int? registeredTime,

            Pulumi.AwsNative.Panorama.PackageVersionStatus? status,

            string? statusDescription,

            string? updatedLatestPatchVersion)
        {
            IsLatestPatch = isLatestPatch;
            MarkLatest = markLatest;
            PackageArn = packageArn;
            PackageName = packageName;
            RegisteredTime = registeredTime;
            Status = status;
            StatusDescription = statusDescription;
            UpdatedLatestPatchVersion = updatedLatestPatchVersion;
        }
    }
}
