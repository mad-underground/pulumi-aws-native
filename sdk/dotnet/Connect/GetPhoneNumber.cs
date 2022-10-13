// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetPhoneNumber
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::PhoneNumber
        /// </summary>
        public static Task<GetPhoneNumberResult> InvokeAsync(GetPhoneNumberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPhoneNumberResult>("aws-native:connect:getPhoneNumber", args ?? new GetPhoneNumberArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::PhoneNumber
        /// </summary>
        public static Output<GetPhoneNumberResult> Invoke(GetPhoneNumberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPhoneNumberResult>("aws-native:connect:getPhoneNumber", args ?? new GetPhoneNumberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPhoneNumberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The phone number ARN
        /// </summary>
        [Input("phoneNumberArn", required: true)]
        public string PhoneNumberArn { get; set; } = null!;

        public GetPhoneNumberArgs()
        {
        }
        public static new GetPhoneNumberArgs Empty => new GetPhoneNumberArgs();
    }

    public sealed class GetPhoneNumberInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The phone number ARN
        /// </summary>
        [Input("phoneNumberArn", required: true)]
        public Input<string> PhoneNumberArn { get; set; } = null!;

        public GetPhoneNumberInvokeArgs()
        {
        }
        public static new GetPhoneNumberInvokeArgs Empty => new GetPhoneNumberInvokeArgs();
    }


    [OutputType]
    public sealed class GetPhoneNumberResult
    {
        /// <summary>
        /// The phone number e164 address.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// The phone number ARN
        /// </summary>
        public readonly string? PhoneNumberArn;
        /// <summary>
        /// One or more tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.PhoneNumberTag> Tags;
        /// <summary>
        /// The ARN of the target the phone number is claimed to.
        /// </summary>
        public readonly string? TargetArn;

        [OutputConstructor]
        private GetPhoneNumberResult(
            string? address,

            string? phoneNumberArn,

            ImmutableArray<Outputs.PhoneNumberTag> tags,

            string? targetArn)
        {
            Address = address;
            PhoneNumberArn = phoneNumberArn;
            Tags = tags;
            TargetArn = targetArn;
        }
    }
}
