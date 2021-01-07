// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppStream.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html
    /// </summary>
    public sealed class DirectoryConfigPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-directoryname
        /// </summary>
        [Input("DirectoryName", required: true)]
        public Input<string> DirectoryName { get; set; } = null!;

        [Input("OrganizationalUnitDistinguishedNames", required: true)]
        private InputList<string>? _OrganizationalUnitDistinguishedNames;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-organizationalunitdistinguishednames
        /// </summary>
        public InputList<string> OrganizationalUnitDistinguishedNames
        {
            get => _OrganizationalUnitDistinguishedNames ?? (_OrganizationalUnitDistinguishedNames = new InputList<string>());
            set => _OrganizationalUnitDistinguishedNames = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-serviceaccountcredentials
        /// </summary>
        [Input("ServiceAccountCredentials", required: true)]
        public Input<Inputs.DirectoryConfigServiceAccountCredentialsArgs> ServiceAccountCredentials { get; set; } = null!;

        public DirectoryConfigPropertiesArgs()
        {
        }
    }
}