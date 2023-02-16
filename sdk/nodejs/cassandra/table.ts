// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Cassandra::Table
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cassandra:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    public readonly billingMode!: pulumi.Output<outputs.cassandra.TableBillingMode | undefined>;
    /**
     * Clustering key columns of the table
     */
    public readonly clusteringKeyColumns!: pulumi.Output<outputs.cassandra.TableClusteringKeyColumn[] | undefined>;
    /**
     * Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
     */
    public readonly defaultTimeToLive!: pulumi.Output<number | undefined>;
    public readonly encryptionSpecification!: pulumi.Output<outputs.cassandra.TableEncryptionSpecification | undefined>;
    /**
     * Name for Cassandra keyspace
     */
    public readonly keyspaceName!: pulumi.Output<string>;
    /**
     * Partition key columns of the table
     */
    public readonly partitionKeyColumns!: pulumi.Output<outputs.cassandra.TableColumn[]>;
    /**
     * Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
     */
    public readonly pointInTimeRecoveryEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Non-key columns of the table
     */
    public readonly regularColumns!: pulumi.Output<outputs.cassandra.TableColumn[] | undefined>;
    /**
     * Name for Cassandra table
     */
    public readonly tableName!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource
     */
    public readonly tags!: pulumi.Output<outputs.cassandra.TableTag[] | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keyspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyspaceName'");
            }
            if ((!args || args.partitionKeyColumns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionKeyColumns'");
            }
            resourceInputs["billingMode"] = args ? args.billingMode : undefined;
            resourceInputs["clusteringKeyColumns"] = args ? args.clusteringKeyColumns : undefined;
            resourceInputs["defaultTimeToLive"] = args ? args.defaultTimeToLive : undefined;
            resourceInputs["encryptionSpecification"] = args ? args.encryptionSpecification : undefined;
            resourceInputs["keyspaceName"] = args ? args.keyspaceName : undefined;
            resourceInputs["partitionKeyColumns"] = args ? args.partitionKeyColumns : undefined;
            resourceInputs["pointInTimeRecoveryEnabled"] = args ? args.pointInTimeRecoveryEnabled : undefined;
            resourceInputs["regularColumns"] = args ? args.regularColumns : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["billingMode"] = undefined /*out*/;
            resourceInputs["clusteringKeyColumns"] = undefined /*out*/;
            resourceInputs["defaultTimeToLive"] = undefined /*out*/;
            resourceInputs["encryptionSpecification"] = undefined /*out*/;
            resourceInputs["keyspaceName"] = undefined /*out*/;
            resourceInputs["partitionKeyColumns"] = undefined /*out*/;
            resourceInputs["pointInTimeRecoveryEnabled"] = undefined /*out*/;
            resourceInputs["regularColumns"] = undefined /*out*/;
            resourceInputs["tableName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    billingMode?: pulumi.Input<inputs.cassandra.TableBillingModeArgs>;
    /**
     * Clustering key columns of the table
     */
    clusteringKeyColumns?: pulumi.Input<pulumi.Input<inputs.cassandra.TableClusteringKeyColumnArgs>[]>;
    /**
     * Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
     */
    defaultTimeToLive?: pulumi.Input<number>;
    encryptionSpecification?: pulumi.Input<inputs.cassandra.TableEncryptionSpecificationArgs>;
    /**
     * Name for Cassandra keyspace
     */
    keyspaceName: pulumi.Input<string>;
    /**
     * Partition key columns of the table
     */
    partitionKeyColumns: pulumi.Input<pulumi.Input<inputs.cassandra.TableColumnArgs>[]>;
    /**
     * Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
     */
    pointInTimeRecoveryEnabled?: pulumi.Input<boolean>;
    /**
     * Non-key columns of the table
     */
    regularColumns?: pulumi.Input<pulumi.Input<inputs.cassandra.TableColumnArgs>[]>;
    /**
     * Name for Cassandra table
     */
    tableName?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.cassandra.TableTagArgs>[]>;
}
