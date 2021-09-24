// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Workteam
 *
 * @deprecated Workteam is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Workteam extends pulumi.CustomResource {
    /**
     * Get an existing Workteam resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workteam {
        pulumi.log.warn("Workteam is deprecated: Workteam is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Workteam(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:Workteam';

    /**
     * Returns true if the given object is an instance of Workteam.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workteam {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workteam.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public readonly memberDefinitions!: pulumi.Output<outputs.sagemaker.WorkteamMemberDefinition[] | undefined>;
    public readonly notificationConfiguration!: pulumi.Output<outputs.sagemaker.WorkteamNotificationConfiguration | undefined>;
    public readonly tags!: pulumi.Output<outputs.sagemaker.WorkteamTag[] | undefined>;
    public readonly workteamName!: pulumi.Output<string | undefined>;

    /**
     * Create a Workteam resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Workteam is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: WorkteamArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Workteam is deprecated: Workteam is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["description"] = args ? args.description : undefined;
            inputs["memberDefinitions"] = args ? args.memberDefinitions : undefined;
            inputs["notificationConfiguration"] = args ? args.notificationConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workteamName"] = args ? args.workteamName : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["memberDefinitions"] = undefined /*out*/;
            inputs["notificationConfiguration"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["workteamName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Workteam.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workteam resource.
 */
export interface WorkteamArgs {
    description?: pulumi.Input<string>;
    memberDefinitions?: pulumi.Input<pulumi.Input<inputs.sagemaker.WorkteamMemberDefinitionArgs>[]>;
    notificationConfiguration?: pulumi.Input<inputs.sagemaker.WorkteamNotificationConfigurationArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.WorkteamTagArgs>[]>;
    workteamName?: pulumi.Input<string>;
}