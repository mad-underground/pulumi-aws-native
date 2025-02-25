// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetThingGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::IoT::ThingGroup
        /// </summary>
        public static Task<GetThingGroupResult> InvokeAsync(GetThingGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetThingGroupResult>("aws-native:iot:getThingGroup", args ?? new GetThingGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IoT::ThingGroup
        /// </summary>
        public static Output<GetThingGroupResult> Invoke(GetThingGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetThingGroupResult>("aws-native:iot:getThingGroup", args ?? new GetThingGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetThingGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("thingGroupName", required: true)]
        public string ThingGroupName { get; set; } = null!;

        public GetThingGroupArgs()
        {
        }
        public static new GetThingGroupArgs Empty => new GetThingGroupArgs();
    }

    public sealed class GetThingGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("thingGroupName", required: true)]
        public Input<string> ThingGroupName { get; set; } = null!;

        public GetThingGroupInvokeArgs()
        {
        }
        public static new GetThingGroupInvokeArgs Empty => new GetThingGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetThingGroupResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? QueryString;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly Outputs.ThingGroupPropertiesProperties? ThingGroupProperties;

        [OutputConstructor]
        private GetThingGroupResult(
            string? arn,

            string? id,

            string? queryString,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            Outputs.ThingGroupPropertiesProperties? thingGroupProperties)
        {
            Arn = arn;
            Id = id;
            QueryString = queryString;
            Tags = tags;
            ThingGroupProperties = thingGroupProperties;
        }
    }
}
