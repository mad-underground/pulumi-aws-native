// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class GroupsArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.SegmentDimensionsArgs>? _dimensions;
        public InputList<Inputs.SegmentDimensionsArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.SegmentDimensionsArgs>());
            set => _dimensions = value;
        }

        [Input("sourceSegments")]
        private InputList<Inputs.SegmentSourceSegmentsArgs>? _sourceSegments;
        public InputList<Inputs.SegmentSourceSegmentsArgs> SourceSegments
        {
            get => _sourceSegments ?? (_sourceSegments = new InputList<Inputs.SegmentSourceSegmentsArgs>());
            set => _sourceSegments = value;
        }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public GroupsArgs()
        {
        }
        public static new GroupsArgs Empty => new GroupsArgs();
    }
}
