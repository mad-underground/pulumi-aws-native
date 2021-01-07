// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElastiCache.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
    /// </summary>
    public sealed class ParameterGroupPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-cacheparametergroupfamily
        /// </summary>
        [Input("CacheParameterGroupFamily", required: true)]
        public Input<string> CacheParameterGroupFamily { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-description
        /// </summary>
        [Input("Description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("Properties")]
        private InputMap<string>? _Properties;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-properties
        /// </summary>
        public InputMap<string> Properties
        {
            get => _Properties ?? (_Properties = new InputMap<string>());
            set => _Properties = value;
        }

        public ParameterGroupPropertiesArgs()
        {
        }
    }
}