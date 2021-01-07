// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.LakeFormation.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-columnwildcard.html
    /// </summary>
    public sealed class PermissionsColumnWildcardArgs : Pulumi.ResourceArgs
    {
        [Input("ExcludedColumnNames")]
        private InputList<string>? _ExcludedColumnNames;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-columnwildcard.html#cfn-lakeformation-permissions-columnwildcard-excludedcolumnnames
        /// </summary>
        public InputList<string> ExcludedColumnNames
        {
            get => _ExcludedColumnNames ?? (_ExcludedColumnNames = new InputList<string>());
            set => _ExcludedColumnNames = value;
        }

        public PermissionsColumnWildcardArgs()
        {
        }
    }
}