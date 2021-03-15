// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html
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
    public static readonly __pulumiType = 'aws-native:Cassandra:Table';

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

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-billingmode
     */
    public readonly billingMode!: pulumi.Output<outputs.Cassandra.TableBillingMode | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-clusteringkeycolumns
     */
    public readonly clusteringKeyColumns!: pulumi.Output<outputs.Cassandra.TableClusteringKeyColumn[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-keyspacename
     */
    public readonly keyspaceName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-partitionkeycolumns
     */
    public readonly partitionKeyColumns!: pulumi.Output<outputs.Cassandra.TableColumn[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-regularcolumns
     */
    public readonly regularColumns!: pulumi.Output<outputs.Cassandra.TableColumn[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-tablename
     */
    public readonly tableName!: pulumi.Output<string | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.keyspaceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'keyspaceName'");
            }
            if ((!args || args.partitionKeyColumns === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'partitionKeyColumns'");
            }
            inputs["billingMode"] = args ? args.billingMode : undefined;
            inputs["clusteringKeyColumns"] = args ? args.clusteringKeyColumns : undefined;
            inputs["keyspaceName"] = args ? args.keyspaceName : undefined;
            inputs["partitionKeyColumns"] = args ? args.partitionKeyColumns : undefined;
            inputs["regularColumns"] = args ? args.regularColumns : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
        } else {
            inputs["billingMode"] = undefined /*out*/;
            inputs["clusteringKeyColumns"] = undefined /*out*/;
            inputs["keyspaceName"] = undefined /*out*/;
            inputs["partitionKeyColumns"] = undefined /*out*/;
            inputs["regularColumns"] = undefined /*out*/;
            inputs["tableName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Table.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-billingmode
     */
    readonly billingMode?: pulumi.Input<inputs.Cassandra.TableBillingMode>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-clusteringkeycolumns
     */
    readonly clusteringKeyColumns?: pulumi.Input<pulumi.Input<inputs.Cassandra.TableClusteringKeyColumn>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-keyspacename
     */
    readonly keyspaceName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-partitionkeycolumns
     */
    readonly partitionKeyColumns: pulumi.Input<pulumi.Input<inputs.Cassandra.TableColumn>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-regularcolumns
     */
    readonly regularColumns?: pulumi.Input<pulumi.Input<inputs.Cassandra.TableColumn>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-table.html#cfn-cassandra-table-tablename
     */
    readonly tableName?: pulumi.Input<string>;
}
