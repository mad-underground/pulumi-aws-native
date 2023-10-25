// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::SQS::QueuePolicy`` type applies a policy to SQS queues. For an example snippet, see [Declaring an policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sqs-policy) in the *User Guide*.
 *
 * @deprecated QueuePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class QueuePolicy extends pulumi.CustomResource {
    /**
     * Get an existing QueuePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): QueuePolicy {
        pulumi.log.warn("QueuePolicy is deprecated: QueuePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new QueuePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sqs:QueuePolicy';

    /**
     * Returns true if the given object is an instance of QueuePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueuePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueuePolicy.__pulumiType;
    }

    /**
     * A policy document that contains the permissions for the specified SQS queues. For more information about SQS policies, see [Using custom policies with the access policy language](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html) in the *Developer Guide*.
     */
    public readonly policyDocument!: pulumi.Output<any>;
    /**
     * The URLs of the queues to which you want to add the policy. You can use the ``Ref`` function to specify an ``AWS::SQS::Queue`` resource.
     */
    public readonly queues!: pulumi.Output<string[]>;

    /**
     * Create a QueuePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated QueuePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: QueuePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("QueuePolicy is deprecated: QueuePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.policyDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyDocument'");
            }
            if ((!args || args.queues === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queues'");
            }
            resourceInputs["policyDocument"] = args ? args.policyDocument : undefined;
            resourceInputs["queues"] = args ? args.queues : undefined;
        } else {
            resourceInputs["policyDocument"] = undefined /*out*/;
            resourceInputs["queues"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QueuePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a QueuePolicy resource.
 */
export interface QueuePolicyArgs {
    /**
     * A policy document that contains the permissions for the specified SQS queues. For more information about SQS policies, see [Using custom policies with the access policy language](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html) in the *Developer Guide*.
     */
    policyDocument: any;
    /**
     * The URLs of the queues to which you want to add the policy. You can use the ``Ref`` function to specify an ``AWS::SQS::Queue`` resource.
     */
    queues: pulumi.Input<pulumi.Input<string>[]>;
}
