// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfiguration.html
    /// </summary>
    public sealed class IndexUserTokenConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfiguration.html#cfn-kendra-index-usertokenconfiguration-jsontokentypeconfiguration
        /// </summary>
        [Input("JsonTokenTypeConfiguration")]
        public Input<Inputs.IndexJsonTokenTypeConfigurationArgs>? JsonTokenTypeConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfiguration.html#cfn-kendra-index-usertokenconfiguration-jwttokentypeconfiguration
        /// </summary>
        [Input("JwtTokenTypeConfiguration")]
        public Input<Inputs.IndexJwtTokenTypeConfigurationArgs>? JwtTokenTypeConfiguration { get; set; }

        public IndexUserTokenConfigurationArgs()
        {
        }
    }
}
