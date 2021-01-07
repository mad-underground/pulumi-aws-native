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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-valueimportanceitems.html
    /// </summary>
    public sealed class IndexValueImportanceItemsArgs : Pulumi.ResourceArgs
    {
        [Input("ValueImportanceItems")]
        private InputList<Inputs.IndexValueImportanceItemArgs>? _ValueImportanceItems;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-valueimportanceitems.html#cfn-kendra-index-valueimportanceitems-valueimportanceitems
        /// </summary>
        public InputList<Inputs.IndexValueImportanceItemArgs> ValueImportanceItems
        {
            get => _ValueImportanceItems ?? (_ValueImportanceItems = new InputList<Inputs.IndexValueImportanceItemArgs>());
            set => _ValueImportanceItems = value;
        }

        public IndexValueImportanceItemsArgs()
        {
        }
    }
}