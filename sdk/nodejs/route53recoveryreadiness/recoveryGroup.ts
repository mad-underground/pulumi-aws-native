// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS Route53 Recovery Readiness Recovery Group Schema and API specifications.
 */
export class RecoveryGroup extends pulumi.CustomResource {
    /**
     * Get an existing RecoveryGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RecoveryGroup {
        return new RecoveryGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53recoveryreadiness:RecoveryGroup';

    /**
     * Returns true if the given object is an instance of RecoveryGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecoveryGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecoveryGroup.__pulumiType;
    }

    /**
     * A list of the cell Amazon Resource Names (ARNs) in the recovery group.
     */
    public readonly cells!: pulumi.Output<string[] | undefined>;
    /**
     * A collection of tags associated with a resource.
     */
    public /*out*/ readonly recoveryGroupArn!: pulumi.Output<string>;
    /**
     * The name of the recovery group to create.
     */
    public readonly recoveryGroupName!: pulumi.Output<string | undefined>;
    /**
     * A collection of tags associated with a resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a RecoveryGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RecoveryGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["cells"] = args ? args.cells : undefined;
            resourceInputs["recoveryGroupName"] = args ? args.recoveryGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["recoveryGroupArn"] = undefined /*out*/;
        } else {
            resourceInputs["cells"] = undefined /*out*/;
            resourceInputs["recoveryGroupArn"] = undefined /*out*/;
            resourceInputs["recoveryGroupName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["recoveryGroupName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RecoveryGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RecoveryGroup resource.
 */
export interface RecoveryGroupArgs {
    /**
     * A list of the cell Amazon Resource Names (ARNs) in the recovery group.
     */
    cells?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the recovery group to create.
     */
    recoveryGroupName?: pulumi.Input<string>;
    /**
     * A collection of tags associated with a resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
