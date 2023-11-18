// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    /// <summary>
    /// Packet identifiers map for a given Multiplex program.
    /// </summary>
    public sealed class MultiplexprogramMultiplexProgramPacketIdentifiersMapArgs : global::Pulumi.ResourceArgs
    {
        [Input("audioPids")]
        private InputList<int>? _audioPids;
        public InputList<int> AudioPids
        {
            get => _audioPids ?? (_audioPids = new InputList<int>());
            set => _audioPids = value;
        }

        [Input("dvbSubPids")]
        private InputList<int>? _dvbSubPids;
        public InputList<int> DvbSubPids
        {
            get => _dvbSubPids ?? (_dvbSubPids = new InputList<int>());
            set => _dvbSubPids = value;
        }

        [Input("dvbTeletextPid")]
        public Input<int>? DvbTeletextPid { get; set; }

        [Input("etvPlatformPid")]
        public Input<int>? EtvPlatformPid { get; set; }

        [Input("etvSignalPid")]
        public Input<int>? EtvSignalPid { get; set; }

        [Input("klvDataPids")]
        private InputList<int>? _klvDataPids;
        public InputList<int> KlvDataPids
        {
            get => _klvDataPids ?? (_klvDataPids = new InputList<int>());
            set => _klvDataPids = value;
        }

        [Input("pcrPid")]
        public Input<int>? PcrPid { get; set; }

        [Input("pmtPid")]
        public Input<int>? PmtPid { get; set; }

        [Input("privateMetadataPid")]
        public Input<int>? PrivateMetadataPid { get; set; }

        [Input("scte27Pids")]
        private InputList<int>? _scte27Pids;
        public InputList<int> Scte27Pids
        {
            get => _scte27Pids ?? (_scte27Pids = new InputList<int>());
            set => _scte27Pids = value;
        }

        [Input("scte35Pid")]
        public Input<int>? Scte35Pid { get; set; }

        [Input("timedMetadataPid")]
        public Input<int>? TimedMetadataPid { get; set; }

        [Input("videoPid")]
        public Input<int>? VideoPid { get; set; }

        public MultiplexprogramMultiplexProgramPacketIdentifiersMapArgs()
        {
        }
        public static new MultiplexprogramMultiplexProgramPacketIdentifiersMapArgs Empty => new MultiplexprogramMultiplexProgramPacketIdentifiersMapArgs();
    }
}