// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SecretsManager::RotationSchedule
 *
 * @deprecated RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class RotationSchedule extends pulumi.CustomResource {
    /**
     * Get an existing RotationSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RotationSchedule {
        pulumi.log.warn("RotationSchedule is deprecated: RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new RotationSchedule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:secretsmanager:RotationSchedule';

    /**
     * Returns true if the given object is an instance of RotationSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RotationSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RotationSchedule.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly hostedRotationLambda!: pulumi.Output<outputs.secretsmanager.RotationScheduleHostedRotationLambda | undefined>;
    public readonly rotateImmediatelyOnUpdate!: pulumi.Output<boolean | undefined>;
    public readonly rotationLambdaArn!: pulumi.Output<string | undefined>;
    public readonly rotationRules!: pulumi.Output<outputs.secretsmanager.RotationScheduleRotationRules | undefined>;
    public readonly secretId!: pulumi.Output<string>;

    /**
     * Create a RotationSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: RotationScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("RotationSchedule is deprecated: RotationSchedule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.secretId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretId'");
            }
            resourceInputs["hostedRotationLambda"] = args ? args.hostedRotationLambda : undefined;
            resourceInputs["rotateImmediatelyOnUpdate"] = args ? args.rotateImmediatelyOnUpdate : undefined;
            resourceInputs["rotationLambdaArn"] = args ? args.rotationLambdaArn : undefined;
            resourceInputs["rotationRules"] = args ? args.rotationRules : undefined;
            resourceInputs["secretId"] = args ? args.secretId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["hostedRotationLambda"] = undefined /*out*/;
            resourceInputs["rotateImmediatelyOnUpdate"] = undefined /*out*/;
            resourceInputs["rotationLambdaArn"] = undefined /*out*/;
            resourceInputs["rotationRules"] = undefined /*out*/;
            resourceInputs["secretId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["secretId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RotationSchedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RotationSchedule resource.
 */
export interface RotationScheduleArgs {
    hostedRotationLambda?: pulumi.Input<inputs.secretsmanager.RotationScheduleHostedRotationLambdaArgs>;
    rotateImmediatelyOnUpdate?: pulumi.Input<boolean>;
    rotationLambdaArn?: pulumi.Input<string>;
    rotationRules?: pulumi.Input<inputs.secretsmanager.RotationScheduleRotationRulesArgs>;
    secretId: pulumi.Input<string>;
}
