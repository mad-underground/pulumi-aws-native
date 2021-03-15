// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacelist.html
    /// </summary>
    public sealed class DataSourceConfluenceSpaceListArgs : Pulumi.ResourceArgs
    {
        [Input("confluenceSpaceList")]
        private InputList<string>? _confluenceSpaceList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacelist.html#cfn-kendra-datasource-confluencespacelist-confluencespacelist
        /// </summary>
        public InputList<string> ConfluenceSpaceList
        {
            get => _confluenceSpaceList ?? (_confluenceSpaceList = new InputList<string>());
            set => _confluenceSpaceList = value;
        }

        public DataSourceConfluenceSpaceListArgs()
        {
        }
    }
}
