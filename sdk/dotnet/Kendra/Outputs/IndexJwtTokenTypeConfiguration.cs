// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class IndexJwtTokenTypeConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-claimregex
        /// </summary>
        public readonly string? ClaimRegex;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-groupattributefield
        /// </summary>
        public readonly string? GroupAttributeField;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-issuer
        /// </summary>
        public readonly string? Issuer;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-keylocation
        /// </summary>
        public readonly string KeyLocation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-secretmanagerarn
        /// </summary>
        public readonly string? SecretManagerArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-url
        /// </summary>
        public readonly string? URL;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jwttokentypeconfiguration.html#cfn-kendra-index-jwttokentypeconfiguration-usernameattributefield
        /// </summary>
        public readonly string? UserNameAttributeField;

        [OutputConstructor]
        private IndexJwtTokenTypeConfiguration(
            string? claimRegex,

            string? groupAttributeField,

            string? issuer,

            string keyLocation,

            string? secretManagerArn,

            string? uRL,

            string? userNameAttributeField)
        {
            ClaimRegex = claimRegex;
            GroupAttributeField = groupAttributeField;
            Issuer = issuer;
            KeyLocation = keyLocation;
            SecretManagerArn = secretManagerArn;
            URL = uRL;
            UserNameAttributeField = userNameAttributeField;
        }
    }
}
