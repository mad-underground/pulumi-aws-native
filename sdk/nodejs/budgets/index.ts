// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BudgetArgs } from "./budget";
export type Budget = import("./budget").Budget;
export const Budget: typeof import("./budget").Budget = null as any;
utilities.lazyLoad(exports, ["Budget"], () => require("./budget"));

export { BudgetsActionArgs } from "./budgetsAction";
export type BudgetsAction = import("./budgetsAction").BudgetsAction;
export const BudgetsAction: typeof import("./budgetsAction").BudgetsAction = null as any;
utilities.lazyLoad(exports, ["BudgetsAction"], () => require("./budgetsAction"));

export { GetBudgetArgs, GetBudgetResult, GetBudgetOutputArgs } from "./getBudget";
export const getBudget: typeof import("./getBudget").getBudget = null as any;
export const getBudgetOutput: typeof import("./getBudget").getBudgetOutput = null as any;
utilities.lazyLoad(exports, ["getBudget","getBudgetOutput"], () => require("./getBudget"));

export { GetBudgetsActionArgs, GetBudgetsActionResult, GetBudgetsActionOutputArgs } from "./getBudgetsAction";
export const getBudgetsAction: typeof import("./getBudgetsAction").getBudgetsAction = null as any;
export const getBudgetsActionOutput: typeof import("./getBudgetsAction").getBudgetsActionOutput = null as any;
utilities.lazyLoad(exports, ["getBudgetsAction","getBudgetsActionOutput"], () => require("./getBudgetsAction"));


// Export enums:
export * from "../types/enums/budgets";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:budgets:Budget":
                return new Budget(name, <any>undefined, { urn })
            case "aws-native:budgets:BudgetsAction":
                return new BudgetsAction(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "budgets", _module)
