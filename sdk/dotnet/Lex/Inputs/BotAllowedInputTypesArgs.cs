// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// Specifies the allowed input types.
    /// </summary>
    public sealed class BotAllowedInputTypesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether audio input is allowed.
        /// </summary>
        [Input("allowAudioInput", required: true)]
        public Input<bool> AllowAudioInput { get; set; } = null!;

        /// <summary>
        /// Indicates whether DTMF input is allowed.
        /// </summary>
        [Input("allowDtmfInput", required: true)]
        public Input<bool> AllowDtmfInput { get; set; } = null!;

        public BotAllowedInputTypesArgs()
        {
        }
        public static new BotAllowedInputTypesArgs Empty => new BotAllowedInputTypesArgs();
    }
}
