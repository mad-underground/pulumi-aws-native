// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds
{
    /// <summary>
    /// Resource Type definition for AWS::RDS::DBSecurityGroup
    /// </summary>
    [Obsolete(@"DbSecurityGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:rds:DbSecurityGroup")]
    public partial class DbSecurityGroup : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("dbSecurityGroupIngress")]
        public Output<ImmutableArray<Outputs.DbSecurityGroupIngress>> DbSecurityGroupIngress { get; private set; } = null!;

        [Output("ec2VpcId")]
        public Output<string?> Ec2VpcId { get; private set; } = null!;

        [Output("groupDescription")]
        public Output<string> GroupDescription { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DbSecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbSecurityGroup(string name, DbSecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbSecurityGroup", name, args ?? new DbSecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbSecurityGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbSecurityGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "ec2VpcId",
                    "groupDescription",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DbSecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbSecurityGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DbSecurityGroup(name, id, options);
        }
    }

    public sealed class DbSecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("dbSecurityGroupIngress", required: true)]
        private InputList<Inputs.DbSecurityGroupIngressArgs>? _dbSecurityGroupIngress;
        public InputList<Inputs.DbSecurityGroupIngressArgs> DbSecurityGroupIngress
        {
            get => _dbSecurityGroupIngress ?? (_dbSecurityGroupIngress = new InputList<Inputs.DbSecurityGroupIngressArgs>());
            set => _dbSecurityGroupIngress = value;
        }

        [Input("ec2VpcId")]
        public Input<string>? Ec2VpcId { get; set; }

        [Input("groupDescription", required: true)]
        public Input<string> GroupDescription { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public DbSecurityGroupArgs()
        {
        }
        public static new DbSecurityGroupArgs Empty => new DbSecurityGroupArgs();
    }
}
