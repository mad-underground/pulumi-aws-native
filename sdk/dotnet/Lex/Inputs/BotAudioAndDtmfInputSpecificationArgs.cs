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
    /// Specifies the audio and DTMF input specification.
    /// </summary>
    public sealed class BotAudioAndDtmfInputSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("audioSpecification")]
        public Input<Inputs.BotAudioSpecificationArgs>? AudioSpecification { get; set; }

        [Input("dtmfSpecification")]
        public Input<Inputs.BotDtmfSpecificationArgs>? DtmfSpecification { get; set; }

        /// <summary>
        /// Time for which a bot waits before assuming that the customer isn't going to speak or press a key. This timeout is shared between Audio and DTMF inputs.
        /// </summary>
        [Input("startTimeoutMs", required: true)]
        public Input<int> StartTimeoutMs { get; set; } = null!;

        public BotAudioAndDtmfInputSpecificationArgs()
        {
        }
        public static new BotAudioAndDtmfInputSpecificationArgs Empty => new BotAudioAndDtmfInputSpecificationArgs();
    }
}
