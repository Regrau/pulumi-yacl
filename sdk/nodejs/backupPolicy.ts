// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Allows management of [Yandex.Cloud Backup Policy](https://cloud.yandex.ru/docs/backup/concepts/policy).
 *
 * > **NOTE:\_** Cloud Backup Provider must be activated in order to manipulate with policies.
 * Active it either by UI Console or by `yc` command.
 *
 * ## Example Usage
 *
 * A very basic acceptable configuration of Backup Policy consists of the following attributes:
 *
 * - `name`
 * - `reattempts`
 * - `vmSnapshotReattempts`
 * - `scheduling`
 * - `retention`
 *
 * Note, that `reattempts` and `vmSnapshotReattempts` might be empty and will be filled with default values.
 *
 * Please, note the basic example below:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const basicPolicy = new yandex.BackupPolicy("basic_policy", {
 *     reattempts: {},
 *     retention: {
 *         afterBackup: false,
 *     },
 *     scheduling: {
 *         enabled: false,
 *         executeByInterval: 86400,
 *     },
 *     vmSnapshotReattempts: {},
 * });
 * ```
 *
 * For the full policy attributes, take a look at the following example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const myPolicy = new yandex.BackupPolicy("my_policy", {
 *     archiveName: "[Machine Name]-[Plan ID]-[Unique ID]a",
 *     cbt: "USE_IF_ENABLED",
 *     compression: "NORMAL",
 *     fastBackupEnabled: true,
 *     format: "AUTO",
 *     multiVolumeSnapshottingEnabled: true,
 *     performanceWindowEnabled: true,
 *     preserveFileSecuritySettings: true,
 *     quiesceSnapshottingEnabled: true,
 *     reattempts: {
 *         enabled: true,
 *         interval: "1m",
 *         maxAttempts: 10,
 *     },
 *     retention: {
 *         afterBackup: false,
 *         rules: [{
 *             maxAge: "365d",
 *             repeatPeriods: [],
 *         }],
 *     },
 *     scheduling: {
 *         enabled: false,
 *         executeByTimes: [{
 *             includeLastDayOfMonth: true,
 *             monthdays: [],
 *             months: [
 *                 1,
 *                 2,
 *                 3,
 *                 4,
 *                 5,
 *                 6,
 *                 7,
 *                 8,
 *                 9,
 *                 10,
 *                 11,
 *                 12,
 *             ],
 *             repeatAts: ["04:10"],
 *             repeatEvery: "30m",
 *             type: "MONTHLY",
 *             weekdays: [],
 *         }],
 *         maxParallelBackups: 0,
 *         randomMaxDelay: "30m",
 *         scheme: "ALWAYS_INCREMENTAL",
 *         weeklyBackupDay: "MONDAY",
 *     },
 *     silentModeEnabled: true,
 *     splittingBytes: "9223372036854775807",
 *     vmSnapshotReattempts: {
 *         enabled: true,
 *         interval: "1m",
 *         maxAttempts: 10,
 *     },
 *     vssProvider: "NATIVE",
 * });
 * ```
 * ## Defined types
 *
 * ### interval_type 
 *
 * A string type, that accepts values in the format of: `number` + `time type`, where `time type` might be:
 *
 * - `s` — seconds
 * - `m` — minutes
 * - `h` — hours
 * - `d` — days
 * - `w` — weekdays
 * - `M` — months
 *
 * Example of interval value: `"5m", "10d", "2M", "5w"`
 *
 * ### dayType
 *
 * A string type, that accepts the following values:
 * `"ALWAYS_INCREMENTAL"`, `"ALWAYS_FULL"`, `"WEEKLY_FULL_DAILY_INCREMENTAL"`, `'WEEKLY_INCREMENTAL"`.
 */
export class BackupPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BackupPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupPolicyState, opts?: pulumi.CustomResourceOptions): BackupPolicy {
        return new BackupPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/backupPolicy:BackupPolicy';

    /**
     * Returns true if the given object is an instance of BackupPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPolicy.__pulumiType;
    }

    /**
     * — The name of generated archives.
     */
    public readonly archiveName!: pulumi.Output<string | undefined>;
    /**
     * — Configuration of Changed Block Tracking.
     * Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
     */
    public readonly cbt!: pulumi.Output<string | undefined>;
    /**
     * — Archive compression level. Affects CPU.
     * Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
     */
    public readonly compression!: pulumi.Output<string | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * — enables or disables scheduling.
     */
    public /*out*/ readonly enabled!: pulumi.Output<boolean>;
    public readonly fastBackupEnabled!: pulumi.Output<boolean | undefined>;
    public readonly folderId!: pulumi.Output<string>;
    /**
     * — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`.
     * Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * — If true, snapshots of multiple volumes will be taken simultaneously.
     */
    public readonly multiVolumeSnapshottingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * — Name of the policy
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * — Time windows for performance limitations of backup.
     */
    public readonly performanceWindowEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * — Preserves file security settings. It's better to set this option to true.
     */
    public readonly preserveFileSecuritySettings!: pulumi.Output<boolean | undefined>;
    /**
     * — If true, a quiesced snapshot of the virtual machine will be taken.
     */
    public readonly quiesceSnapshottingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * — Amount of reattempts that should be performed while trying to make backup at the host.
     * This attribute consists of the following parameters:
     */
    public readonly reattempts!: pulumi.Output<outputs.BackupPolicyReattempts>;
    /**
     * — Retention policy for backups. Allows to setup backups lifecycle.
     * This attribute consists of the following parameters:
     */
    public readonly retention!: pulumi.Output<outputs.BackupPolicyRetention>;
    /**
     * — Schedule settings for creating backups on the host.
     */
    public readonly scheduling!: pulumi.Output<outputs.BackupPolicyScheduling>;
    /**
     * — if true, a user interaction will be avoided when possible.
     */
    public readonly silentModeEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * — determines the size to split backups. It's better to leave this option unchanged.
     */
    public readonly splittingBytes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * — Amount of reattempts that should be performed while trying to make snapshot.
     * This attribute consists of the following parameters:
     */
    public readonly vmSnapshotReattempts!: pulumi.Output<outputs.BackupPolicyVmSnapshotReattempts>;
    /**
     * — Settings for the volume shadow copy service.
     * Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
     */
    public readonly vssProvider!: pulumi.Output<string | undefined>;

    /**
     * Create a BackupPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupPolicyArgs | BackupPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupPolicyState | undefined;
            resourceInputs["archiveName"] = state ? state.archiveName : undefined;
            resourceInputs["cbt"] = state ? state.cbt : undefined;
            resourceInputs["compression"] = state ? state.compression : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["fastBackupEnabled"] = state ? state.fastBackupEnabled : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["multiVolumeSnapshottingEnabled"] = state ? state.multiVolumeSnapshottingEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["performanceWindowEnabled"] = state ? state.performanceWindowEnabled : undefined;
            resourceInputs["preserveFileSecuritySettings"] = state ? state.preserveFileSecuritySettings : undefined;
            resourceInputs["quiesceSnapshottingEnabled"] = state ? state.quiesceSnapshottingEnabled : undefined;
            resourceInputs["reattempts"] = state ? state.reattempts : undefined;
            resourceInputs["retention"] = state ? state.retention : undefined;
            resourceInputs["scheduling"] = state ? state.scheduling : undefined;
            resourceInputs["silentModeEnabled"] = state ? state.silentModeEnabled : undefined;
            resourceInputs["splittingBytes"] = state ? state.splittingBytes : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["vmSnapshotReattempts"] = state ? state.vmSnapshotReattempts : undefined;
            resourceInputs["vssProvider"] = state ? state.vssProvider : undefined;
        } else {
            const args = argsOrState as BackupPolicyArgs | undefined;
            if ((!args || args.reattempts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'reattempts'");
            }
            if ((!args || args.retention === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retention'");
            }
            if ((!args || args.scheduling === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduling'");
            }
            if ((!args || args.vmSnapshotReattempts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmSnapshotReattempts'");
            }
            resourceInputs["archiveName"] = args ? args.archiveName : undefined;
            resourceInputs["cbt"] = args ? args.cbt : undefined;
            resourceInputs["compression"] = args ? args.compression : undefined;
            resourceInputs["fastBackupEnabled"] = args ? args.fastBackupEnabled : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["multiVolumeSnapshottingEnabled"] = args ? args.multiVolumeSnapshottingEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["performanceWindowEnabled"] = args ? args.performanceWindowEnabled : undefined;
            resourceInputs["preserveFileSecuritySettings"] = args ? args.preserveFileSecuritySettings : undefined;
            resourceInputs["quiesceSnapshottingEnabled"] = args ? args.quiesceSnapshottingEnabled : undefined;
            resourceInputs["reattempts"] = args ? args.reattempts : undefined;
            resourceInputs["retention"] = args ? args.retention : undefined;
            resourceInputs["scheduling"] = args ? args.scheduling : undefined;
            resourceInputs["silentModeEnabled"] = args ? args.silentModeEnabled : undefined;
            resourceInputs["splittingBytes"] = args ? args.splittingBytes : undefined;
            resourceInputs["vmSnapshotReattempts"] = args ? args.vmSnapshotReattempts : undefined;
            resourceInputs["vssProvider"] = args ? args.vssProvider : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupPolicy resources.
 */
export interface BackupPolicyState {
    /**
     * — The name of generated archives.
     */
    archiveName?: pulumi.Input<string>;
    /**
     * — Configuration of Changed Block Tracking.
     * Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
     */
    cbt?: pulumi.Input<string>;
    /**
     * — Archive compression level. Affects CPU.
     * Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
     */
    compression?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    /**
     * — enables or disables scheduling.
     */
    enabled?: pulumi.Input<boolean>;
    fastBackupEnabled?: pulumi.Input<boolean>;
    folderId?: pulumi.Input<string>;
    /**
     * — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`.
     * Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
     */
    format?: pulumi.Input<string>;
    /**
     * — If true, snapshots of multiple volumes will be taken simultaneously.
     */
    multiVolumeSnapshottingEnabled?: pulumi.Input<boolean>;
    /**
     * — Name of the policy
     */
    name?: pulumi.Input<string>;
    /**
     * — Time windows for performance limitations of backup.
     */
    performanceWindowEnabled?: pulumi.Input<boolean>;
    /**
     * — Preserves file security settings. It's better to set this option to true.
     */
    preserveFileSecuritySettings?: pulumi.Input<boolean>;
    /**
     * — If true, a quiesced snapshot of the virtual machine will be taken.
     */
    quiesceSnapshottingEnabled?: pulumi.Input<boolean>;
    /**
     * — Amount of reattempts that should be performed while trying to make backup at the host.
     * This attribute consists of the following parameters:
     */
    reattempts?: pulumi.Input<inputs.BackupPolicyReattempts>;
    /**
     * — Retention policy for backups. Allows to setup backups lifecycle.
     * This attribute consists of the following parameters:
     */
    retention?: pulumi.Input<inputs.BackupPolicyRetention>;
    /**
     * — Schedule settings for creating backups on the host.
     */
    scheduling?: pulumi.Input<inputs.BackupPolicyScheduling>;
    /**
     * — if true, a user interaction will be avoided when possible.
     */
    silentModeEnabled?: pulumi.Input<boolean>;
    /**
     * — determines the size to split backups. It's better to leave this option unchanged.
     */
    splittingBytes?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<string>;
    /**
     * — Amount of reattempts that should be performed while trying to make snapshot.
     * This attribute consists of the following parameters:
     */
    vmSnapshotReattempts?: pulumi.Input<inputs.BackupPolicyVmSnapshotReattempts>;
    /**
     * — Settings for the volume shadow copy service.
     * Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
     */
    vssProvider?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupPolicy resource.
 */
export interface BackupPolicyArgs {
    /**
     * — The name of generated archives.
     */
    archiveName?: pulumi.Input<string>;
    /**
     * — Configuration of Changed Block Tracking.
     * Available values are: `"USE_IF_ENABLED"`, `"ENABLED_AND_USE"`, `"DO_NOT_USE"`.
     */
    cbt?: pulumi.Input<string>;
    /**
     * — Archive compression level. Affects CPU.
     * Available values: `"NORMAL"`, `"HIGH"`, `"MAX"`, `"OFF"`.
     */
    compression?: pulumi.Input<string>;
    fastBackupEnabled?: pulumi.Input<boolean>;
    folderId?: pulumi.Input<string>;
    /**
     * — Format of the backup. It's strongly recommend to leave this option empty or `"AUTO"`.
     * Available values: `"AUTO"`, `"VERSION_11"`, `"VERSION_12"`.
     */
    format?: pulumi.Input<string>;
    /**
     * — If true, snapshots of multiple volumes will be taken simultaneously.
     */
    multiVolumeSnapshottingEnabled?: pulumi.Input<boolean>;
    /**
     * — Name of the policy
     */
    name?: pulumi.Input<string>;
    /**
     * — Time windows for performance limitations of backup.
     */
    performanceWindowEnabled?: pulumi.Input<boolean>;
    /**
     * — Preserves file security settings. It's better to set this option to true.
     */
    preserveFileSecuritySettings?: pulumi.Input<boolean>;
    /**
     * — If true, a quiesced snapshot of the virtual machine will be taken.
     */
    quiesceSnapshottingEnabled?: pulumi.Input<boolean>;
    /**
     * — Amount of reattempts that should be performed while trying to make backup at the host.
     * This attribute consists of the following parameters:
     */
    reattempts: pulumi.Input<inputs.BackupPolicyReattempts>;
    /**
     * — Retention policy for backups. Allows to setup backups lifecycle.
     * This attribute consists of the following parameters:
     */
    retention: pulumi.Input<inputs.BackupPolicyRetention>;
    /**
     * — Schedule settings for creating backups on the host.
     */
    scheduling: pulumi.Input<inputs.BackupPolicyScheduling>;
    /**
     * — if true, a user interaction will be avoided when possible.
     */
    silentModeEnabled?: pulumi.Input<boolean>;
    /**
     * — determines the size to split backups. It's better to leave this option unchanged.
     */
    splittingBytes?: pulumi.Input<string>;
    /**
     * — Amount of reattempts that should be performed while trying to make snapshot.
     * This attribute consists of the following parameters:
     */
    vmSnapshotReattempts: pulumi.Input<inputs.BackupPolicyVmSnapshotReattempts>;
    /**
     * — Settings for the volume shadow copy service.
     * Available values are: "`NATIVE`", `"TARGET_SYSTEM_DEFINED"`
     */
    vssProvider?: pulumi.Input<string>;
}
