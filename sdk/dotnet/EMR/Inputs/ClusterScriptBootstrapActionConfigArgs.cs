// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EMR.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html
    /// </summary>
    public sealed class ClusterScriptBootstrapActionConfigArgs : Pulumi.ResourceArgs
    {
        [Input("Args")]
        private InputList<string>? _Args;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html#cfn-elasticmapreduce-cluster-scriptbootstrapactionconfig-args
        /// </summary>
        public InputList<string> Args
        {
            get => _Args ?? (_Args = new InputList<string>());
            set => _Args = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scriptbootstrapactionconfig.html#cfn-elasticmapreduce-cluster-scriptbootstrapactionconfig-path
        /// </summary>
        [Input("Path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public ClusterScriptBootstrapActionConfigArgs()
        {
        }
    }
}