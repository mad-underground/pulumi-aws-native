// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::RDS::DBClusterParameterGroup resource creates a new Amazon RDS DB cluster parameter group. For more information, see Managing an Amazon Aurora DB Cluster in the Amazon Aurora User Guide.
 */
export class DBClusterParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing DBClusterParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DBClusterParameterGroup {
        return new DBClusterParameterGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DBClusterParameterGroup';

    /**
     * Returns true if the given object is an instance of DBClusterParameterGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DBClusterParameterGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DBClusterParameterGroup.__pulumiType;
    }

    public /*out*/ readonly dBClusterParameterGroupName!: pulumi.Output<string>;
    /**
     * A friendly description for this DB cluster parameter group.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
     */
    public readonly parameters!: pulumi.Output<any>;
    /**
     * The list of tags for the cluster parameter group.
     */
    public readonly tags!: pulumi.Output<outputs.rds.DBClusterParameterGroupTag[] | undefined>;

    /**
     * Create a DBClusterParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DBClusterParameterGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.family === undefined) && !opts.urn) {
                throw new Error("Missing required property 'family'");
            }
            if ((!args || args.parameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameters'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["family"] = args ? args.family : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["dBClusterParameterGroupName"] = undefined /*out*/;
        } else {
            resourceInputs["dBClusterParameterGroupName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["family"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DBClusterParameterGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DBClusterParameterGroup resource.
 */
export interface DBClusterParameterGroupArgs {
    /**
     * A friendly description for this DB cluster parameter group.
     */
    description: pulumi.Input<string>;
    /**
     * The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
     */
    family: pulumi.Input<string>;
    /**
     * An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
     */
    parameters: any;
    /**
     * The list of tags for the cluster parameter group.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.rds.DBClusterParameterGroupTagArgs>[]>;
}
