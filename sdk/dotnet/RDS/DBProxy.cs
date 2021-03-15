// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html
    /// </summary>
    [AwsNativeResourceType("aws-native:RDS:DBProxy")]
    public partial class DBProxy : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-auth
        /// </summary>
        [Output("auth")]
        public Output<ImmutableArray<Outputs.DBProxyAuthFormat>> Auth { get; private set; } = null!;

        [Output("dBProxyArn")]
        public Output<string> DBProxyArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-dbproxyname
        /// </summary>
        [Output("dBProxyName")]
        public Output<string> DBProxyName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-debuglogging
        /// </summary>
        [Output("debugLogging")]
        public Output<bool?> DebugLogging { get; private set; } = null!;

        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-enginefamily
        /// </summary>
        [Output("engineFamily")]
        public Output<string> EngineFamily { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-idleclienttimeout
        /// </summary>
        [Output("idleClientTimeout")]
        public Output<int?> IdleClientTimeout { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-requiretls
        /// </summary>
        [Output("requireTLS")]
        public Output<bool?> RequireTLS { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-rolearn
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DBProxyTagFormat>> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsecuritygroupids
        /// </summary>
        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsubnetids
        /// </summary>
        [Output("vpcSubnetIds")]
        public Output<ImmutableArray<string>> VpcSubnetIds { get; private set; } = null!;


        /// <summary>
        /// Create a DBProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DBProxy(string name, DBProxyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:RDS:DBProxy", name, args ?? new DBProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DBProxy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:RDS:DBProxy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DBProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DBProxy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DBProxy(name, id, options);
        }
    }

    public sealed class DBProxyArgs : Pulumi.ResourceArgs
    {
        [Input("auth", required: true)]
        private InputList<Inputs.DBProxyAuthFormatArgs>? _auth;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-auth
        /// </summary>
        public InputList<Inputs.DBProxyAuthFormatArgs> Auth
        {
            get => _auth ?? (_auth = new InputList<Inputs.DBProxyAuthFormatArgs>());
            set => _auth = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-dbproxyname
        /// </summary>
        [Input("dBProxyName", required: true)]
        public Input<string> DBProxyName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-debuglogging
        /// </summary>
        [Input("debugLogging")]
        public Input<bool>? DebugLogging { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-enginefamily
        /// </summary>
        [Input("engineFamily", required: true)]
        public Input<string> EngineFamily { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-idleclienttimeout
        /// </summary>
        [Input("idleClientTimeout")]
        public Input<int>? IdleClientTimeout { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-requiretls
        /// </summary>
        [Input("requireTLS")]
        public Input<bool>? RequireTLS { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-rolearn
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.DBProxyTagFormatArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-tags
        /// </summary>
        public InputList<Inputs.DBProxyTagFormatArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DBProxyTagFormatArgs>());
            set => _tags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsecuritygroupids
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        [Input("vpcSubnetIds", required: true)]
        private InputList<string>? _vpcSubnetIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#cfn-rds-dbproxy-vpcsubnetids
        /// </summary>
        public InputList<string> VpcSubnetIds
        {
            get => _vpcSubnetIds ?? (_vpcSubnetIds = new InputList<string>());
            set => _vpcSubnetIds = value;
        }

        public DBProxyArgs()
        {
        }
    }
}
