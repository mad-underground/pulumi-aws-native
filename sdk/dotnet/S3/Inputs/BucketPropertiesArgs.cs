// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.S3.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
    /// </summary>
    public sealed class BucketPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accelerateconfiguration
        /// </summary>
        [Input("AccelerateConfiguration")]
        public Input<Inputs.BucketAccelerateConfigurationArgs>? AccelerateConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
        /// </summary>
        [Input("AccessControl")]
        public Input<string>? AccessControl { get; set; }

        [Input("AnalyticsConfigurations")]
        private InputList<Inputs.BucketAnalyticsConfigurationArgs>? _AnalyticsConfigurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-analyticsconfigurations
        /// </summary>
        public InputList<Inputs.BucketAnalyticsConfigurationArgs> AnalyticsConfigurations
        {
            get => _AnalyticsConfigurations ?? (_AnalyticsConfigurations = new InputList<Inputs.BucketAnalyticsConfigurationArgs>());
            set => _AnalyticsConfigurations = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-bucketencryption
        /// </summary>
        [Input("BucketEncryption")]
        public Input<Inputs.BucketBucketEncryptionArgs>? BucketEncryption { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
        /// </summary>
        [Input("BucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
        /// </summary>
        [Input("CorsConfiguration")]
        public Input<Inputs.BucketCorsConfigurationArgs>? CorsConfiguration { get; set; }

        [Input("IntelligentTieringConfigurations")]
        private InputList<Inputs.BucketIntelligentTieringConfigurationArgs>? _IntelligentTieringConfigurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-intelligenttieringconfigurations
        /// </summary>
        public InputList<Inputs.BucketIntelligentTieringConfigurationArgs> IntelligentTieringConfigurations
        {
            get => _IntelligentTieringConfigurations ?? (_IntelligentTieringConfigurations = new InputList<Inputs.BucketIntelligentTieringConfigurationArgs>());
            set => _IntelligentTieringConfigurations = value;
        }

        [Input("InventoryConfigurations")]
        private InputList<Inputs.BucketInventoryConfigurationArgs>? _InventoryConfigurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-inventoryconfigurations
        /// </summary>
        public InputList<Inputs.BucketInventoryConfigurationArgs> InventoryConfigurations
        {
            get => _InventoryConfigurations ?? (_InventoryConfigurations = new InputList<Inputs.BucketInventoryConfigurationArgs>());
            set => _InventoryConfigurations = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
        /// </summary>
        [Input("LifecycleConfiguration")]
        public Input<Inputs.BucketLifecycleConfigurationArgs>? LifecycleConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
        /// </summary>
        [Input("LoggingConfiguration")]
        public Input<Inputs.BucketLoggingConfigurationArgs>? LoggingConfiguration { get; set; }

        [Input("MetricsConfigurations")]
        private InputList<Inputs.BucketMetricsConfigurationArgs>? _MetricsConfigurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-metricsconfigurations
        /// </summary>
        public InputList<Inputs.BucketMetricsConfigurationArgs> MetricsConfigurations
        {
            get => _MetricsConfigurations ?? (_MetricsConfigurations = new InputList<Inputs.BucketMetricsConfigurationArgs>());
            set => _MetricsConfigurations = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
        /// </summary>
        [Input("NotificationConfiguration")]
        public Input<Inputs.BucketNotificationConfigurationArgs>? NotificationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockconfiguration
        /// </summary>
        [Input("ObjectLockConfiguration")]
        public Input<Inputs.BucketObjectLockConfigurationArgs>? ObjectLockConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockenabled
        /// </summary>
        [Input("ObjectLockEnabled")]
        public Input<bool>? ObjectLockEnabled { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-ownershipcontrols
        /// </summary>
        [Input("OwnershipControls")]
        public Input<Inputs.BucketOwnershipControlsArgs>? OwnershipControls { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-publicaccessblockconfiguration
        /// </summary>
        [Input("PublicAccessBlockConfiguration")]
        public Input<Inputs.BucketPublicAccessBlockConfigurationArgs>? PublicAccessBlockConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
        /// </summary>
        [Input("ReplicationConfiguration")]
        public Input<Inputs.BucketReplicationConfigurationArgs>? ReplicationConfiguration { get; set; }

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
        /// </summary>
        [Input("VersioningConfiguration")]
        public Input<Inputs.BucketVersioningConfigurationArgs>? VersioningConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
        /// </summary>
        [Input("WebsiteConfiguration")]
        public Input<Inputs.BucketWebsiteConfigurationArgs>? WebsiteConfiguration { get; set; }

        public BucketPropertiesArgs()
        {
        }
    }
}
