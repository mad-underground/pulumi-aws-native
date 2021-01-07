// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelKeyProviderSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-keyprovidersettings.html#cfn-medialive-channel-keyprovidersettings-statickeysettings
        /// </summary>
        public readonly Outputs.ChannelStaticKeySettings? StaticKeySettings;

        [OutputConstructor]
        private ChannelKeyProviderSettings(Outputs.ChannelStaticKeySettings? StaticKeySettings)
        {
            this.StaticKeySettings = StaticKeySettings;
        }
    }
}
