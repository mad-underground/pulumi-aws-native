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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html
    /// </summary>
    [AwsNativeResourceType("aws-native:RDS:DBProxyTargetGroup")]
    public partial class DBProxyTargetGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
        /// </summary>
        [Output("connectionPoolConfigurationInfo")]
        public Output<Outputs.DBProxyTargetGroupConnectionPoolConfigurationInfoFormat?> ConnectionPoolConfigurationInfo { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
        /// </summary>
        [Output("dBClusterIdentifiers")]
        public Output<ImmutableArray<string>> DBClusterIdentifiers { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
        /// </summary>
        [Output("dBInstanceIdentifiers")]
        public Output<ImmutableArray<string>> DBInstanceIdentifiers { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
        /// </summary>
        [Output("dBProxyName")]
        public Output<string> DBProxyName { get; private set; } = null!;

        [Output("targetGroupArn")]
        public Output<string> TargetGroupArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
        /// </summary>
        [Output("targetGroupName")]
        public Output<string> TargetGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a DBProxyTargetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DBProxyTargetGroup(string name, DBProxyTargetGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:RDS:DBProxyTargetGroup", name, args ?? new DBProxyTargetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DBProxyTargetGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:RDS:DBProxyTargetGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DBProxyTargetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DBProxyTargetGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DBProxyTargetGroup(name, id, options);
        }
    }

    public sealed class DBProxyTargetGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
        /// </summary>
        [Input("connectionPoolConfigurationInfo")]
        public Input<Inputs.DBProxyTargetGroupConnectionPoolConfigurationInfoFormatArgs>? ConnectionPoolConfigurationInfo { get; set; }

        [Input("dBClusterIdentifiers")]
        private InputList<string>? _dBClusterIdentifiers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
        /// </summary>
        public InputList<string> DBClusterIdentifiers
        {
            get => _dBClusterIdentifiers ?? (_dBClusterIdentifiers = new InputList<string>());
            set => _dBClusterIdentifiers = value;
        }

        [Input("dBInstanceIdentifiers")]
        private InputList<string>? _dBInstanceIdentifiers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
        /// </summary>
        public InputList<string> DBInstanceIdentifiers
        {
            get => _dBInstanceIdentifiers ?? (_dBInstanceIdentifiers = new InputList<string>());
            set => _dBInstanceIdentifiers = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
        /// </summary>
        [Input("dBProxyName", required: true)]
        public Input<string> DBProxyName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
        /// </summary>
        [Input("targetGroupName", required: true)]
        public Input<string> TargetGroupName { get; set; } = null!;

        public DBProxyTargetGroupArgs()
        {
        }
    }
}
