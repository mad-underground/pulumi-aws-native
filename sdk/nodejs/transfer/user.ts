// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Transfer::User
 *
 * @deprecated User is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): User {
        pulumi.log.warn("User is deprecated: User is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new User(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:transfer:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    public readonly homeDirectory!: pulumi.Output<string | undefined>;
    public readonly homeDirectoryMappings!: pulumi.Output<outputs.transfer.UserHomeDirectoryMapEntry[] | undefined>;
    public readonly homeDirectoryType!: pulumi.Output<string | undefined>;
    public readonly policy!: pulumi.Output<string | undefined>;
    public readonly posixProfile!: pulumi.Output<outputs.transfer.UserPosixProfile | undefined>;
    public readonly role!: pulumi.Output<string>;
    public readonly serverId!: pulumi.Output<string>;
    public readonly sshPublicKeys!: pulumi.Output<outputs.transfer.UserSshPublicKey[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated User is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("User is deprecated: User is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.serverId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverId'");
            }
            resourceInputs["homeDirectory"] = args ? args.homeDirectory : undefined;
            resourceInputs["homeDirectoryMappings"] = args ? args.homeDirectoryMappings : undefined;
            resourceInputs["homeDirectoryType"] = args ? args.homeDirectoryType : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["posixProfile"] = args ? args.posixProfile : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["serverId"] = args ? args.serverId : undefined;
            resourceInputs["sshPublicKeys"] = args ? args.sshPublicKeys : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["homeDirectory"] = undefined /*out*/;
            resourceInputs["homeDirectoryMappings"] = undefined /*out*/;
            resourceInputs["homeDirectoryType"] = undefined /*out*/;
            resourceInputs["policy"] = undefined /*out*/;
            resourceInputs["posixProfile"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["serverId"] = undefined /*out*/;
            resourceInputs["sshPublicKeys"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["serverId", "userName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    homeDirectory?: pulumi.Input<string>;
    homeDirectoryMappings?: pulumi.Input<pulumi.Input<inputs.transfer.UserHomeDirectoryMapEntryArgs>[]>;
    homeDirectoryType?: pulumi.Input<string>;
    policy?: pulumi.Input<string>;
    posixProfile?: pulumi.Input<inputs.transfer.UserPosixProfileArgs>;
    role: pulumi.Input<string>;
    serverId: pulumi.Input<string>;
    sshPublicKeys?: pulumi.Input<pulumi.Input<inputs.transfer.UserSshPublicKeyArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    userName?: pulumi.Input<string>;
}
