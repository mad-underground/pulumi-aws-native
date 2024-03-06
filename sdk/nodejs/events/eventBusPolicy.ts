// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::EventBusPolicy
 *
 * @deprecated EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class EventBusPolicy extends pulumi.CustomResource {
    /**
     * Get an existing EventBusPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventBusPolicy {
        pulumi.log.warn("EventBusPolicy is deprecated: EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new EventBusPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:events:EventBusPolicy';

    /**
     * Returns true if the given object is an instance of EventBusPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventBusPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventBusPolicy.__pulumiType;
    }

    public readonly action!: pulumi.Output<string | undefined>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<outputs.events.EventBusPolicyCondition | undefined>;
    public readonly eventBusName!: pulumi.Output<string | undefined>;
    public readonly principal!: pulumi.Output<string | undefined>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
     */
    public readonly statement!: pulumi.Output<any | undefined>;
    public readonly statementId!: pulumi.Output<string>;

    /**
     * Create a EventBusPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: EventBusPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("EventBusPolicy is deprecated: EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.statementId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statementId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["eventBusName"] = args ? args.eventBusName : undefined;
            resourceInputs["principal"] = args ? args.principal : undefined;
            resourceInputs["statement"] = args ? args.statement : undefined;
            resourceInputs["statementId"] = args ? args.statementId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["condition"] = undefined /*out*/;
            resourceInputs["eventBusName"] = undefined /*out*/;
            resourceInputs["principal"] = undefined /*out*/;
            resourceInputs["statement"] = undefined /*out*/;
            resourceInputs["statementId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["eventBusName", "statementId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EventBusPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventBusPolicy resource.
 */
export interface EventBusPolicyArgs {
    action?: pulumi.Input<string>;
    condition?: pulumi.Input<inputs.events.EventBusPolicyConditionArgs>;
    eventBusName?: pulumi.Input<string>;
    principal?: pulumi.Input<string>;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
     */
    statement?: any;
    statementId: pulumi.Input<string>;
}
