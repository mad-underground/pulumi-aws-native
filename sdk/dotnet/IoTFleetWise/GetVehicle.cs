// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTFleetWise
{
    public static class GetVehicle
    {
        /// <summary>
        /// Definition of AWS::IoTFleetWise::Vehicle Resource Type
        /// </summary>
        public static Task<GetVehicleResult> InvokeAsync(GetVehicleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVehicleResult>("aws-native:iotfleetwise:getVehicle", args ?? new GetVehicleArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::IoTFleetWise::Vehicle Resource Type
        /// </summary>
        public static Output<GetVehicleResult> Invoke(GetVehicleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVehicleResult>("aws-native:iotfleetwise:getVehicle", args ?? new GetVehicleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVehicleArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetVehicleArgs()
        {
        }
        public static new GetVehicleArgs Empty => new GetVehicleArgs();
    }

    public sealed class GetVehicleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetVehicleInvokeArgs()
        {
        }
        public static new GetVehicleInvokeArgs Empty => new GetVehicleInvokeArgs();
    }


    [OutputType]
    public sealed class GetVehicleResult
    {
        public readonly string? Arn;
        public readonly ImmutableDictionary<string, string>? Attributes;
        public readonly string? CreationTime;
        public readonly string? DecoderManifestArn;
        public readonly string? LastModificationTime;
        public readonly string? ModelManifestArn;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetVehicleResult(
            string? arn,

            ImmutableDictionary<string, string>? attributes,

            string? creationTime,

            string? decoderManifestArn,

            string? lastModificationTime,

            string? modelManifestArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Attributes = attributes;
            CreationTime = creationTime;
            DecoderManifestArn = decoderManifestArn;
            LastModificationTime = lastModificationTime;
            ModelManifestArn = modelManifestArn;
            Tags = tags;
        }
    }
}
