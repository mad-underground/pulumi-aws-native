// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Contains information about the phone configuration settings for a user.
    /// </summary>
    [OutputType]
    public sealed class UserPhoneConfig
    {
        public readonly int? AfterContactWorkTimeLimit;
        public readonly bool? AutoAccept;
        public readonly string? DeskPhoneNumber;
        public readonly Pulumi.AwsNative.Connect.UserPhoneType PhoneType;

        [OutputConstructor]
        private UserPhoneConfig(
            int? afterContactWorkTimeLimit,

            bool? autoAccept,

            string? deskPhoneNumber,

            Pulumi.AwsNative.Connect.UserPhoneType phoneType)
        {
            AfterContactWorkTimeLimit = afterContactWorkTimeLimit;
            AutoAccept = autoAccept;
            DeskPhoneNumber = deskPhoneNumber;
            PhoneType = phoneType;
        }
    }
}
