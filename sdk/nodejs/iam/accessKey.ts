// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::AccessKey
 *
 * @deprecated AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class AccessKey extends pulumi.CustomResource {
    /**
     * Get an existing AccessKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccessKey {
        pulumi.log.warn("AccessKey is deprecated: AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new AccessKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iam:AccessKey';

    /**
     * Returns true if the given object is an instance of AccessKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessKey.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public /*out*/ readonly secretAccessKey!: pulumi.Output<string>;
    public readonly serial!: pulumi.Output<number | undefined>;
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a AccessKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: AccessKeyArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("AccessKey is deprecated: AccessKey is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["serial"] = args ? args.serial : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["secretAccessKey"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["secretAccessKey"] = undefined /*out*/;
            resourceInputs["serial"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["serial", "userName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(AccessKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccessKey resource.
 */
export interface AccessKeyArgs {
    serial?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    userName: pulumi.Input<string>;
}
