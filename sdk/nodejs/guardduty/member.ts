// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GuardDuty::Member
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Member {
        return new Member(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:guardduty:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    public readonly detectorId!: pulumi.Output<string | undefined>;
    public readonly disableEmailNotification!: pulumi.Output<boolean | undefined>;
    public readonly email!: pulumi.Output<string>;
    public readonly memberId!: pulumi.Output<string | undefined>;
    public readonly message!: pulumi.Output<string | undefined>;
    public readonly status!: pulumi.Output<string | undefined>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            resourceInputs["detectorId"] = args ? args.detectorId : undefined;
            resourceInputs["disableEmailNotification"] = args ? args.disableEmailNotification : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["memberId"] = args ? args.memberId : undefined;
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
        } else {
            resourceInputs["detectorId"] = undefined /*out*/;
            resourceInputs["disableEmailNotification"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["memberId"] = undefined /*out*/;
            resourceInputs["message"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["detectorId", "memberId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Member.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    detectorId?: pulumi.Input<string>;
    disableEmailNotification?: pulumi.Input<boolean>;
    email: pulumi.Input<string>;
    memberId?: pulumi.Input<string>;
    message?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
}
