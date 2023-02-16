// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS::SimSpaceWeaver::Simulation resource creates an AWS Simulation.
 */
export class Simulation extends pulumi.CustomResource {
    /**
     * Get an existing Simulation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Simulation {
        return new Simulation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:simspaceweaver:Simulation';

    /**
     * Returns true if the given object is an instance of Simulation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Simulation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Simulation.__pulumiType;
    }

    /**
     * Json object with all simulation details
     */
    public /*out*/ readonly describePayload!: pulumi.Output<string>;
    /**
     * The name of the simulation.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Role ARN.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly schemaS3Location!: pulumi.Output<outputs.simspaceweaver.SimulationS3Location | undefined>;

    /**
     * Create a Simulation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SimulationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["schemaS3Location"] = args ? args.schemaS3Location : undefined;
            resourceInputs["describePayload"] = undefined /*out*/;
        } else {
            resourceInputs["describePayload"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["schemaS3Location"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Simulation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Simulation resource.
 */
export interface SimulationArgs {
    /**
     * The name of the simulation.
     */
    name?: pulumi.Input<string>;
    /**
     * Role ARN.
     */
    roleArn?: pulumi.Input<string>;
    schemaS3Location?: pulumi.Input<inputs.simspaceweaver.SimulationS3LocationArgs>;
}
