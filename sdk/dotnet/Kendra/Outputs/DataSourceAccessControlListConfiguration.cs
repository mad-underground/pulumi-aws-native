// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceAccessControlListConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-accesscontrollistconfiguration.html#cfn-kendra-datasource-accesscontrollistconfiguration-keypath
        /// </summary>
        public readonly string? KeyPath;

        [OutputConstructor]
        private DataSourceAccessControlListConfiguration(string? KeyPath)
        {
            this.KeyPath = KeyPath;
        }
    }
}
