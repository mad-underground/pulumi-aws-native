// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks
{
    public static class GetVolume
    {
        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::Volume
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("aws-native:opsworks:getVolume", args ?? new GetVolumeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::Volume
        /// </summary>
        public static Output<GetVolumeResult> Invoke(GetVolumeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVolumeResult>("aws-native:opsworks:getVolume", args ?? new GetVolumeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVolumeArgs()
        {
        }
    }

    public sealed class GetVolumeInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVolumeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        public readonly string? Id;
        public readonly string? MountPoint;
        public readonly string? Name;

        [OutputConstructor]
        private GetVolumeResult(
            string? id,

            string? mountPoint,

            string? name)
        {
            Id = id;
            MountPoint = mountPoint;
            Name = name;
        }
    }
}
