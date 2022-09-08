// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances
 */
export function getStackSet(args: GetStackSetArgs, opts?: pulumi.InvokeOptions): Promise<GetStackSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudformation:getStackSet", {
        "stackSetId": args.stackSetId,
    }, opts);
}

export interface GetStackSetArgs {
    /**
     * The ID of the stack set that you're creating.
     */
    stackSetId: string;
}

export interface GetStackSetResult {
    /**
     * The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
     */
    readonly administrationRoleARN?: string;
    /**
     * Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.
     */
    readonly autoDeployment?: outputs.cloudformation.StackSetAutoDeployment;
    /**
     * In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.
     */
    readonly capabilities?: enums.cloudformation.StackSetCapability[];
    /**
     * A description of the stack set. You can use the description to identify the stack set's purpose or other important information.
     */
    readonly description?: string;
    /**
     * The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.
     */
    readonly executionRoleName?: string;
    /**
     * Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
     */
    readonly managedExecution?: outputs.cloudformation.ManagedExecutionProperties;
    /**
     * The input parameters for the stack set template.
     */
    readonly parameters?: outputs.cloudformation.StackSetParameter[];
    /**
     * A group of stack instances with parameters in some specific accounts and regions.
     */
    readonly stackInstancesGroup?: outputs.cloudformation.StackSetStackInstances[];
    /**
     * The ID of the stack set that you're creating.
     */
    readonly stackSetId?: string;
    /**
     * The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
     */
    readonly tags?: outputs.cloudformation.StackSetTag[];
    /**
     * The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
     */
    readonly templateBody?: string;
}

export function getStackSetOutput(args: GetStackSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStackSetResult> {
    return pulumi.output(args).apply(a => getStackSet(a, opts))
}

export interface GetStackSetOutputArgs {
    /**
     * The ID of the stack set that you're creating.
     */
    stackSetId: pulumi.Input<string>;
}
