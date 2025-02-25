// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSM::MaintenanceWindowTarget
 *
 * @deprecated MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class MaintenanceWindowTarget extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindowTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MaintenanceWindowTarget {
        pulumi.log.warn("MaintenanceWindowTarget is deprecated: MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new MaintenanceWindowTarget(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ssm:MaintenanceWindowTarget';

    /**
     * Returns true if the given object is an instance of MaintenanceWindowTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindowTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindowTarget.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly ownerInformation!: pulumi.Output<string | undefined>;
    public readonly resourceType!: pulumi.Output<string>;
    public readonly targets!: pulumi.Output<outputs.ssm.MaintenanceWindowTargetTargets[]>;
    public readonly windowId!: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindowTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: MaintenanceWindowTargetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("MaintenanceWindowTarget is deprecated: MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.targets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targets'");
            }
            if ((!args || args.windowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'windowId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerInformation"] = args ? args.ownerInformation : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["windowId"] = args ? args.windowId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["ownerInformation"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
            resourceInputs["windowId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["windowId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(MaintenanceWindowTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MaintenanceWindowTarget resource.
 */
export interface MaintenanceWindowTargetArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ownerInformation?: pulumi.Input<string>;
    resourceType: pulumi.Input<string>;
    targets: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTargetTargetsArgs>[]>;
    windowId: pulumi.Input<string>;
}
