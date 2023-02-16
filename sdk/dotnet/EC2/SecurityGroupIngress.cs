// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::SecurityGroupIngress
    /// </summary>
    [Obsolete(@"SecurityGroupIngress is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:SecurityGroupIngress")]
    public partial class SecurityGroupIngress : global::Pulumi.CustomResource
    {
        [Output("cidrIp")]
        public Output<string?> CidrIp { get; private set; } = null!;

        [Output("cidrIpv6")]
        public Output<string?> CidrIpv6 { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("fromPort")]
        public Output<int?> FromPort { get; private set; } = null!;

        [Output("groupId")]
        public Output<string?> GroupId { get; private set; } = null!;

        [Output("groupName")]
        public Output<string?> GroupName { get; private set; } = null!;

        [Output("ipProtocol")]
        public Output<string> IpProtocol { get; private set; } = null!;

        [Output("sourcePrefixListId")]
        public Output<string?> SourcePrefixListId { get; private set; } = null!;

        [Output("sourceSecurityGroupId")]
        public Output<string?> SourceSecurityGroupId { get; private set; } = null!;

        [Output("sourceSecurityGroupName")]
        public Output<string?> SourceSecurityGroupName { get; private set; } = null!;

        [Output("sourceSecurityGroupOwnerId")]
        public Output<string?> SourceSecurityGroupOwnerId { get; private set; } = null!;

        [Output("toPort")]
        public Output<int?> ToPort { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupIngress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupIngress(string name, SecurityGroupIngressArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:SecurityGroupIngress", name, args ?? new SecurityGroupIngressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupIngress(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:SecurityGroupIngress", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityGroupIngress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupIngress Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityGroupIngress(name, id, options);
        }
    }

    public sealed class SecurityGroupIngressArgs : global::Pulumi.ResourceArgs
    {
        [Input("cidrIp")]
        public Input<string>? CidrIp { get; set; }

        [Input("cidrIpv6")]
        public Input<string>? CidrIpv6 { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        [Input("sourcePrefixListId")]
        public Input<string>? SourcePrefixListId { get; set; }

        [Input("sourceSecurityGroupId")]
        public Input<string>? SourceSecurityGroupId { get; set; }

        [Input("sourceSecurityGroupName")]
        public Input<string>? SourceSecurityGroupName { get; set; }

        [Input("sourceSecurityGroupOwnerId")]
        public Input<string>? SourceSecurityGroupOwnerId { get; set; }

        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public SecurityGroupIngressArgs()
        {
        }
        public static new SecurityGroupIngressArgs Empty => new SecurityGroupIngressArgs();
    }
}
