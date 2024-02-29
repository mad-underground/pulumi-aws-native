// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb
{
    /// <summary>
    /// The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
    ///  You should be aware of the following behaviors when working with DDB tables:
    ///   +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).
    /// 
    ///    Our guidance is to use the latest schema documented here for y
    /// </summary>
    [AwsNativeResourceType("aws-native:dynamodb:Table")]
    public partial class Table : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A list of attributes that describe the key schema for the table and indexes.
        ///  This property is required to create a DDB table.
        ///  Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
        /// </summary>
        [Output("attributeDefinitions")]
        public Output<ImmutableArray<Outputs.TableAttributeDefinition>> AttributeDefinitions { get; private set; } = null!;

        /// <summary>
        /// Specify how you are charged for read and write throughput and how you manage capacity.
        ///  Valid values include:
        ///   +   ``PROVISIONED`` - We recommend using ``PROVISIONED`` for predictable workloads. ``PROVISIONED`` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
        ///   +   ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for unpredictable workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
        ///   
        ///  If not specified, the default is ``PROVISIONED``.
        /// </summary>
        [Output("billingMode")]
        public Output<string?> BillingMode { get; private set; } = null!;

        /// <summary>
        /// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
        /// </summary>
        [Output("contributorInsightsSpecification")]
        public Output<Outputs.TableContributorInsightsSpecification?> ContributorInsightsSpecification { get; private set; } = null!;

        /// <summary>
        /// Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
        /// </summary>
        [Output("deletionProtectionEnabled")]
        public Output<bool?> DeletionProtectionEnabled { get; private set; } = null!;

        /// <summary>
        /// Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
        ///   If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
        ///  If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. 
        ///  Updates are not supported. The following are exceptions:
        ///   +  If you update either the contributor insights specification or the provisioned throughput value
        /// </summary>
        [Output("globalSecondaryIndexes")]
        public Output<ImmutableArray<Outputs.TableGlobalSecondaryIndex>> GlobalSecondaryIndexes { get; private set; } = null!;

        /// <summary>
        /// Specifies the properties of data being imported from the S3 bucket source to the table.
        ///   If you specify the ``ImportSourceSpecification`` property, and also specify either the ``StreamSpecification``, the ``TableClass`` property, or the ``DeletionProtectionEnabled`` property, the IAM entity creating/updating stack must have ``UpdateTable`` permission.
        /// </summary>
        [Output("importSourceSpecification")]
        public Output<Outputs.TableImportSourceSpecification?> ImportSourceSpecification { get; private set; } = null!;

        /// <summary>
        /// Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
        /// </summary>
        [Output("keySchema")]
        public Output<Union<ImmutableArray<Outputs.TableKeySchema>, object>> KeySchema { get; private set; } = null!;

        /// <summary>
        /// The Kinesis Data Streams configuration for the specified table.
        /// </summary>
        [Output("kinesisStreamSpecification")]
        public Output<Outputs.TableKinesisStreamSpecification?> KinesisStreamSpecification { get; private set; } = null!;

        /// <summary>
        /// Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
        /// </summary>
        [Output("localSecondaryIndexes")]
        public Output<ImmutableArray<Outputs.TableLocalSecondaryIndex>> LocalSecondaryIndexes { get; private set; } = null!;

        /// <summary>
        /// The settings used to enable point in time recovery.
        /// </summary>
        [Output("pointInTimeRecoverySpecification")]
        public Output<Outputs.TablePointInTimeRecoverySpecification?> PointInTimeRecoverySpecification { get; private set; } = null!;

        /// <summary>
        /// Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). 
        ///  If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
        /// </summary>
        [Output("provisionedThroughput")]
        public Output<Outputs.TableProvisionedThroughput?> ProvisionedThroughput { get; private set; } = null!;

        /// <summary>
        /// Specifies the settings to enable server-side encryption.
        /// </summary>
        [Output("sseSpecification")]
        public Output<Outputs.TableSseSpecification?> SseSpecification { get; private set; } = null!;

        [Output("streamArn")]
        public Output<string> StreamArn { get; private set; } = null!;

        /// <summary>
        /// The settings for the DDB table stream, which capture changes to items stored in the table.
        /// </summary>
        [Output("streamSpecification")]
        public Output<Outputs.TableStreamSpecification?> StreamSpecification { get; private set; } = null!;

        /// <summary>
        /// The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
        /// </summary>
        [Output("tableClass")]
        public Output<string?> TableClass { get; private set; } = null!;

        /// <summary>
        /// A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        ///  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Time to Live (TTL) settings for the table.
        ///   For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
        /// </summary>
        [Output("timeToLiveSpecification")]
        public Output<Outputs.TableTimeToLiveSpecification?> TimeToLiveSpecification { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("aws-native:dynamodb:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:dynamodb:Table", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "importSourceSpecification",
                    "tableName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Table(name, id, options);
        }
    }

    public sealed class TableArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributeDefinitions")]
        private InputList<Inputs.TableAttributeDefinitionArgs>? _attributeDefinitions;

        /// <summary>
        /// A list of attributes that describe the key schema for the table and indexes.
        ///  This property is required to create a DDB table.
        ///  Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
        /// </summary>
        public InputList<Inputs.TableAttributeDefinitionArgs> AttributeDefinitions
        {
            get => _attributeDefinitions ?? (_attributeDefinitions = new InputList<Inputs.TableAttributeDefinitionArgs>());
            set => _attributeDefinitions = value;
        }

        /// <summary>
        /// Specify how you are charged for read and write throughput and how you manage capacity.
        ///  Valid values include:
        ///   +   ``PROVISIONED`` - We recommend using ``PROVISIONED`` for predictable workloads. ``PROVISIONED`` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
        ///   +   ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for unpredictable workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
        ///   
        ///  If not specified, the default is ``PROVISIONED``.
        /// </summary>
        [Input("billingMode")]
        public Input<string>? BillingMode { get; set; }

        /// <summary>
        /// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
        /// </summary>
        [Input("contributorInsightsSpecification")]
        public Input<Inputs.TableContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        /// <summary>
        /// Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
        /// </summary>
        [Input("deletionProtectionEnabled")]
        public Input<bool>? DeletionProtectionEnabled { get; set; }

        [Input("globalSecondaryIndexes")]
        private InputList<Inputs.TableGlobalSecondaryIndexArgs>? _globalSecondaryIndexes;

        /// <summary>
        /// Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
        ///   If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
        ///  If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. 
        ///  Updates are not supported. The following are exceptions:
        ///   +  If you update either the contributor insights specification or the provisioned throughput value
        /// </summary>
        public InputList<Inputs.TableGlobalSecondaryIndexArgs> GlobalSecondaryIndexes
        {
            get => _globalSecondaryIndexes ?? (_globalSecondaryIndexes = new InputList<Inputs.TableGlobalSecondaryIndexArgs>());
            set => _globalSecondaryIndexes = value;
        }

        /// <summary>
        /// Specifies the properties of data being imported from the S3 bucket source to the table.
        ///   If you specify the ``ImportSourceSpecification`` property, and also specify either the ``StreamSpecification``, the ``TableClass`` property, or the ``DeletionProtectionEnabled`` property, the IAM entity creating/updating stack must have ``UpdateTable`` permission.
        /// </summary>
        [Input("importSourceSpecification")]
        public Input<Inputs.TableImportSourceSpecificationArgs>? ImportSourceSpecification { get; set; }

        /// <summary>
        /// Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
        /// </summary>
        [Input("keySchema", required: true)]
        public InputUnion<ImmutableArray<Inputs.TableKeySchemaArgs>, object> KeySchema { get; set; } = null!;

        /// <summary>
        /// The Kinesis Data Streams configuration for the specified table.
        /// </summary>
        [Input("kinesisStreamSpecification")]
        public Input<Inputs.TableKinesisStreamSpecificationArgs>? KinesisStreamSpecification { get; set; }

        [Input("localSecondaryIndexes")]
        private InputList<Inputs.TableLocalSecondaryIndexArgs>? _localSecondaryIndexes;

        /// <summary>
        /// Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
        /// </summary>
        public InputList<Inputs.TableLocalSecondaryIndexArgs> LocalSecondaryIndexes
        {
            get => _localSecondaryIndexes ?? (_localSecondaryIndexes = new InputList<Inputs.TableLocalSecondaryIndexArgs>());
            set => _localSecondaryIndexes = value;
        }

        /// <summary>
        /// The settings used to enable point in time recovery.
        /// </summary>
        [Input("pointInTimeRecoverySpecification")]
        public Input<Inputs.TablePointInTimeRecoverySpecificationArgs>? PointInTimeRecoverySpecification { get; set; }

        /// <summary>
        /// Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). 
        ///  If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
        /// </summary>
        [Input("provisionedThroughput")]
        public Input<Inputs.TableProvisionedThroughputArgs>? ProvisionedThroughput { get; set; }

        /// <summary>
        /// Specifies the settings to enable server-side encryption.
        /// </summary>
        [Input("sseSpecification")]
        public Input<Inputs.TableSseSpecificationArgs>? SseSpecification { get; set; }

        /// <summary>
        /// The settings for the DDB table stream, which capture changes to items stored in the table.
        /// </summary>
        [Input("streamSpecification")]
        public Input<Inputs.TableStreamSpecificationArgs>? StreamSpecification { get; set; }

        /// <summary>
        /// The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
        /// </summary>
        [Input("tableClass")]
        public Input<string>? TableClass { get; set; }

        /// <summary>
        /// A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        ///  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Time to Live (TTL) settings for the table.
        ///   For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
        /// </summary>
        [Input("timeToLiveSpecification")]
        public Input<Inputs.TableTimeToLiveSpecificationArgs>? TimeToLiveSpecification { get; set; }

        public TableArgs()
        {
        }
        public static new TableArgs Empty => new TableArgs();
    }
}
