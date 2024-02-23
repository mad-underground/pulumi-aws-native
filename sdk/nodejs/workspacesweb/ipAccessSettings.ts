// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::WorkSpacesWeb::IpAccessSettings Resource Type
 */
export class IpAccessSettings extends pulumi.CustomResource {
    /**
     * Get an existing IpAccessSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IpAccessSettings {
        return new IpAccessSettings(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:workspacesweb:IpAccessSettings';

    /**
     * Returns true if the given object is an instance of IpAccessSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpAccessSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpAccessSettings.__pulumiType;
    }

    public readonly additionalEncryptionContext!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly associatedPortalArns!: pulumi.Output<string[]>;
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    public readonly customerManagedKey!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly displayName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly ipAccessSettingsArn!: pulumi.Output<string>;
    public readonly ipRules!: pulumi.Output<outputs.workspacesweb.IpAccessSettingsIpRule[]>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a IpAccessSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpAccessSettingsArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.ipRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipRules'");
            }
            resourceInputs["additionalEncryptionContext"] = args ? args.additionalEncryptionContext : undefined;
            resourceInputs["customerManagedKey"] = args ? args.customerManagedKey : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["ipRules"] = args ? args.ipRules : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["associatedPortalArns"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["ipAccessSettingsArn"] = undefined /*out*/;
        } else {
            resourceInputs["additionalEncryptionContext"] = undefined /*out*/;
            resourceInputs["associatedPortalArns"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["customerManagedKey"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["ipAccessSettingsArn"] = undefined /*out*/;
            resourceInputs["ipRules"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["additionalEncryptionContext.*", "customerManagedKey"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(IpAccessSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IpAccessSettings resource.
 */
export interface IpAccessSettingsArgs {
    additionalEncryptionContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    customerManagedKey?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    displayName?: pulumi.Input<string>;
    ipRules: pulumi.Input<pulumi.Input<inputs.workspacesweb.IpAccessSettingsIpRuleArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
