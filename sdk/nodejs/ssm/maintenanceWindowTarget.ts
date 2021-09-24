// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
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
        let inputs: pulumi.Inputs = {};
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
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ownerInformation"] = args ? args.ownerInformation : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["windowId"] = args ? args.windowId : undefined;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["ownerInformation"] = undefined /*out*/;
            inputs["resourceType"] = undefined /*out*/;
            inputs["targets"] = undefined /*out*/;
            inputs["windowId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MaintenanceWindowTarget.__pulumiType, name, inputs, opts);
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