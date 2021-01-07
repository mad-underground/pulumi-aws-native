// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Pinpoint.Outputs
{

    [OutputType]
    public sealed class APNSChannelProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-applicationid
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-bundleid
        /// </summary>
        public readonly string? BundleId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-certificate
        /// </summary>
        public readonly string? Certificate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-defaultauthenticationmethod
        /// </summary>
        public readonly string? DefaultAuthenticationMethod;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-privatekey
        /// </summary>
        public readonly string? PrivateKey;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-teamid
        /// </summary>
        public readonly string? TeamId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-tokenkey
        /// </summary>
        public readonly string? TokenKey;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-apnschannel.html#cfn-pinpoint-apnschannel-tokenkeyid
        /// </summary>
        public readonly string? TokenKeyId;

        [OutputConstructor]
        private APNSChannelProperties(
            string ApplicationId,

            string? BundleId,

            string? Certificate,

            string? DefaultAuthenticationMethod,

            bool? Enabled,

            string? PrivateKey,

            string? TeamId,

            string? TokenKey,

            string? TokenKeyId)
        {
            this.ApplicationId = ApplicationId;
            this.BundleId = BundleId;
            this.Certificate = Certificate;
            this.DefaultAuthenticationMethod = DefaultAuthenticationMethod;
            this.Enabled = Enabled;
            this.PrivateKey = PrivateKey;
            this.TeamId = TeamId;
            this.TokenKey = TokenKey;
            this.TokenKeyId = TokenKeyId;
        }
    }
}