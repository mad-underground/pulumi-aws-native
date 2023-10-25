// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::Rule
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:events:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly eventBusName!: pulumi.Output<string | undefined>;
    public readonly eventPattern!: pulumi.Output<any | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly scheduleExpression!: pulumi.Output<string | undefined>;
    public readonly state!: pulumi.Output<string | undefined>;
    public readonly targets!: pulumi.Output<outputs.events.RuleTarget[] | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eventBusName"] = args ? args.eventBusName : undefined;
            resourceInputs["eventPattern"] = args ? args.eventPattern : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["scheduleExpression"] = args ? args.scheduleExpression : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["eventBusName"] = undefined /*out*/;
            resourceInputs["eventPattern"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["scheduleExpression"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["eventBusName", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    description?: pulumi.Input<string>;
    eventBusName?: pulumi.Input<string>;
    eventPattern?: any;
    name?: pulumi.Input<string>;
    roleArn?: pulumi.Input<string>;
    scheduleExpression?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    targets?: pulumi.Input<pulumi.Input<inputs.events.RuleTargetArgs>[]>;
}
