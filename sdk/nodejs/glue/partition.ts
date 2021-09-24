// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::Partition
 *
 * @deprecated Partition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Partition extends pulumi.CustomResource {
    /**
     * Get an existing Partition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Partition {
        pulumi.log.warn("Partition is deprecated: Partition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Partition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:glue:Partition';

    /**
     * Returns true if the given object is an instance of Partition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Partition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Partition.__pulumiType;
    }

    public readonly catalogId!: pulumi.Output<string>;
    public readonly databaseName!: pulumi.Output<string>;
    public readonly partitionInput!: pulumi.Output<outputs.glue.PartitionPartitionInput>;
    public readonly tableName!: pulumi.Output<string>;

    /**
     * Create a Partition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Partition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: PartitionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Partition is deprecated: Partition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.catalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalogId'");
            }
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.partitionInput === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionInput'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["partitionInput"] = args ? args.partitionInput : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
        } else {
            inputs["catalogId"] = undefined /*out*/;
            inputs["databaseName"] = undefined /*out*/;
            inputs["partitionInput"] = undefined /*out*/;
            inputs["tableName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Partition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Partition resource.
 */
export interface PartitionArgs {
    catalogId: pulumi.Input<string>;
    databaseName: pulumi.Input<string>;
    partitionInput: pulumi.Input<inputs.glue.PartitionPartitionInputArgs>;
    tableName: pulumi.Input<string>;
}