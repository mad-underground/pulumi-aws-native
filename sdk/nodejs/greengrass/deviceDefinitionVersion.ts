// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::DeviceDefinitionVersion
 *
 * @deprecated DeviceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DeviceDefinitionVersion extends pulumi.CustomResource {
    /**
     * Get an existing DeviceDefinitionVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DeviceDefinitionVersion {
        pulumi.log.warn("DeviceDefinitionVersion is deprecated: DeviceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DeviceDefinitionVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:greengrass:DeviceDefinitionVersion';

    /**
     * Returns true if the given object is an instance of DeviceDefinitionVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceDefinitionVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceDefinitionVersion.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly deviceDefinitionId!: pulumi.Output<string>;
    public readonly devices!: pulumi.Output<outputs.greengrass.DeviceDefinitionVersionDevice[]>;

    /**
     * Create a DeviceDefinitionVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DeviceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: DeviceDefinitionVersionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DeviceDefinitionVersion is deprecated: DeviceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deviceDefinitionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceDefinitionId'");
            }
            if ((!args || args.devices === undefined) && !opts.urn) {
                throw new Error("Missing required property 'devices'");
            }
            resourceInputs["deviceDefinitionId"] = args ? args.deviceDefinitionId : undefined;
            resourceInputs["devices"] = args ? args.devices : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["deviceDefinitionId"] = undefined /*out*/;
            resourceInputs["devices"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["deviceDefinitionId", "devices[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DeviceDefinitionVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DeviceDefinitionVersion resource.
 */
export interface DeviceDefinitionVersionArgs {
    deviceDefinitionId: pulumi.Input<string>;
    devices: pulumi.Input<pulumi.Input<inputs.greengrass.DeviceDefinitionVersionDeviceArgs>[]>;
}
