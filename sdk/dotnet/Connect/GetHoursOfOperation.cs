// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetHoursOfOperation
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::HoursOfOperation
        /// </summary>
        public static Task<GetHoursOfOperationResult> InvokeAsync(GetHoursOfOperationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHoursOfOperationResult>("aws-native:connect:getHoursOfOperation", args ?? new GetHoursOfOperationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::HoursOfOperation
        /// </summary>
        public static Output<GetHoursOfOperationResult> Invoke(GetHoursOfOperationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHoursOfOperationResult>("aws-native:connect:getHoursOfOperation", args ?? new GetHoursOfOperationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHoursOfOperationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the hours of operation.
        /// </summary>
        [Input("hoursOfOperationArn", required: true)]
        public string HoursOfOperationArn { get; set; } = null!;

        public GetHoursOfOperationArgs()
        {
        }
        public static new GetHoursOfOperationArgs Empty => new GetHoursOfOperationArgs();
    }

    public sealed class GetHoursOfOperationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the hours of operation.
        /// </summary>
        [Input("hoursOfOperationArn", required: true)]
        public Input<string> HoursOfOperationArn { get; set; } = null!;

        public GetHoursOfOperationInvokeArgs()
        {
        }
        public static new GetHoursOfOperationInvokeArgs Empty => new GetHoursOfOperationInvokeArgs();
    }


    [OutputType]
    public sealed class GetHoursOfOperationResult
    {
        /// <summary>
        /// Configuration information for the hours of operation: day, start time, and end time.
        /// </summary>
        public readonly ImmutableArray<Outputs.HoursOfOperationConfig> Config;
        /// <summary>
        /// The description of the hours of operation.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the hours of operation.
        /// </summary>
        public readonly string? HoursOfOperationArn;
        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        public readonly string? InstanceArn;
        /// <summary>
        /// The name of the hours of operation.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// One or more tags.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The time zone of the hours of operation.
        /// </summary>
        public readonly string? TimeZone;

        [OutputConstructor]
        private GetHoursOfOperationResult(
            ImmutableArray<Outputs.HoursOfOperationConfig> config,

            string? description,

            string? hoursOfOperationArn,

            string? instanceArn,

            string? name,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? timeZone)
        {
            Config = config;
            Description = description;
            HoursOfOperationArn = hoursOfOperationArn;
            InstanceArn = instanceArn;
            Name = name;
            Tags = tags;
            TimeZone = timeZone;
        }
    }
}
