// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SecretsManager::SecretTargetAttachment
 *
 * @deprecated SecretTargetAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class SecretTargetAttachment extends pulumi.CustomResource {
    /**
     * Get an existing SecretTargetAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecretTargetAttachment {
        pulumi.log.warn("SecretTargetAttachment is deprecated: SecretTargetAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new SecretTargetAttachment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:secretsmanager:SecretTargetAttachment';

    /**
     * Returns true if the given object is an instance of SecretTargetAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretTargetAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretTargetAttachment.__pulumiType;
    }

    public readonly secretId!: pulumi.Output<string>;
    public readonly targetId!: pulumi.Output<string>;
    public readonly targetType!: pulumi.Output<string>;

    /**
     * Create a SecretTargetAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated SecretTargetAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: SecretTargetAttachmentArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("SecretTargetAttachment is deprecated: SecretTargetAttachment is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.secretId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretId'");
            }
            if ((!args || args.targetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetId'");
            }
            if ((!args || args.targetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetType'");
            }
            inputs["secretId"] = args ? args.secretId : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
            inputs["targetType"] = args ? args.targetType : undefined;
        } else {
            inputs["secretId"] = undefined /*out*/;
            inputs["targetId"] = undefined /*out*/;
            inputs["targetType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecretTargetAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecretTargetAttachment resource.
 */
export interface SecretTargetAttachmentArgs {
    secretId: pulumi.Input<string>;
    targetId: pulumi.Input<string>;
    targetType: pulumi.Input<string>;
}