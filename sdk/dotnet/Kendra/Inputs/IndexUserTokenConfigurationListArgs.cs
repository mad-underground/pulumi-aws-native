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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfigurationlist.html
    /// </summary>
    public sealed class IndexUserTokenConfigurationListArgs : Pulumi.ResourceArgs
    {
        [Input("UserTokenConfigurationList")]
        private InputList<Inputs.IndexUserTokenConfigurationArgs>? _UserTokenConfigurationList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-usertokenconfigurationlist.html#cfn-kendra-index-usertokenconfigurationlist-usertokenconfigurationlist
        /// </summary>
        public InputList<Inputs.IndexUserTokenConfigurationArgs> UserTokenConfigurationList
        {
            get => _UserTokenConfigurationList ?? (_UserTokenConfigurationList = new InputList<Inputs.IndexUserTokenConfigurationArgs>());
            set => _UserTokenConfigurationList = value;
        }

        public IndexUserTokenConfigurationListArgs()
        {
        }
    }
}
