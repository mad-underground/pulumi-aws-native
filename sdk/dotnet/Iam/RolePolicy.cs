// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam
{
    /// <summary>
    /// Adds or updates an inline policy document that is embedded in the specified IAM role.
    ///  When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role, using [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html). You can update a role's trust policy using [UpdateAssumeRolePolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html). For information about roles, see [roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html) in the *IAM User Guide*.
    ///  A role can also have a managed policy attached to it. To attach a managed policy to a role, use [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-i
    /// </summary>
    [AwsNativeResourceType("aws-native:iam:RolePolicy")]
    public partial class RolePolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The policy document.
        ///  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
        ///  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
        ///   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
        ///   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
        ///   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::RolePolicy` for more information about the expected schema for this property.
        /// </summary>
        [Output("policyDocument")]
        public Output<object?> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The name of the policy document.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// The name of the role to associate the policy with.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;


        /// <summary>
        /// Create a RolePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RolePolicy(string name, RolePolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iam:RolePolicy", name, args ?? new RolePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RolePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iam:RolePolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "policyName",
                    "roleName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RolePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RolePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RolePolicy(name, id, options);
        }
    }

    public sealed class RolePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document.
        ///  You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
        ///  The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
        ///   +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
        ///   +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
        ///   +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::IAM::RolePolicy` for more information about the expected schema for this property.
        /// </summary>
        [Input("policyDocument")]
        public Input<object>? PolicyDocument { get; set; }

        /// <summary>
        /// The name of the policy document.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the role to associate the policy with.
        ///  This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        public RolePolicyArgs()
        {
        }
        public static new RolePolicyArgs Empty => new RolePolicyArgs();
    }
}
