// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex
{
    public static class GetBot
    {
        /// <summary>
        /// Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
        /// </summary>
        public static Task<GetBotResult> InvokeAsync(GetBotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBotResult>("aws-native:lex:getBot", args ?? new GetBotArgs(), options.WithDefaults());

        /// <summary>
        /// Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
        /// </summary>
        public static Output<GetBotResult> Invoke(GetBotInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBotResult>("aws-native:lex:getBot", args ?? new GetBotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBotArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetBotArgs()
        {
        }
    }

    public sealed class GetBotInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetBotInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBotResult
    {
        public readonly string? Arn;
        /// <summary>
        /// Data privacy setting of the Bot.
        /// </summary>
        public readonly Outputs.DataPrivacyProperties? DataPrivacy;
        public readonly string? Description;
        public readonly string? Id;
        /// <summary>
        /// IdleSessionTTLInSeconds of the resource
        /// </summary>
        public readonly int? IdleSessionTTLInSeconds;
        public readonly string? Name;
        public readonly string? RoleArn;

        [OutputConstructor]
        private GetBotResult(
            string? arn,

            Outputs.DataPrivacyProperties? dataPrivacy,

            string? description,

            string? id,

            int? idleSessionTTLInSeconds,

            string? name,

            string? roleArn)
        {
            Arn = arn;
            DataPrivacy = dataPrivacy;
            Description = description;
            Id = id;
            IdleSessionTTLInSeconds = idleSessionTTLInSeconds;
            Name = name;
            RoleArn = roleArn;
        }
    }
}
