// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file statemachine/ssl_gc_change.proto (syntax proto2)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Command, CommandJson, Foul, FoulJson, GameState, GameStateJson, Proposal, ProposalJson, RedCard, RedCardJson, State, StateJson, YellowCard, YellowCardJson } from "../state/ssl_gc_state_pb";
import { file_state_ssl_gc_state } from "../state/ssl_gc_state_pb";
import type { Division, DivisionJson, Team, TeamJson } from "../state/ssl_gc_common_pb";
import { file_state_ssl_gc_common } from "../state/ssl_gc_common_pb";
import type { Vector2, Vector2Json } from "../geom/ssl_gc_geometry_pb";
import { file_geom_ssl_gc_geometry } from "../geom/ssl_gc_geometry_pb";
import type { GameEvent, GameEventJson } from "../state/ssl_gc_game_event_pb";
import { file_state_ssl_gc_game_event } from "../state/ssl_gc_game_event_pb";
import type { MatchType, MatchTypeJson, Referee_Stage, Referee_StageJson } from "../state/ssl_gc_referee_message_pb";
import { file_state_ssl_gc_referee_message } from "../state/ssl_gc_referee_message_pb";
import type { BoolValueJson, Int32ValueJson, StringValueJson, Timestamp, TimestampJson, UInt32ValueJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp, file_google_protobuf_wrappers } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file statemachine/ssl_gc_change.proto.
 */
export const file_statemachine_ssl_gc_change: GenFile = /*@__PURE__*/
  fileDesc("CiBzdGF0ZW1hY2hpbmUvc3NsX2djX2NoYW5nZS5wcm90byKTAQoLU3RhdGVDaGFuZ2USCgoCaWQYASABKAUSGQoJc3RhdGVfcHJlGAIgASgLMgYuU3RhdGUSFQoFc3RhdGUYAyABKAsyBi5TdGF0ZRIXCgZjaGFuZ2UYBCABKAsyBy5DaGFuZ2USLQoJdGltZXN0YW1wGAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcCLUFgoGQ2hhbmdlEg4KBm9yaWdpbhgBIAEoCRISCgpyZXZlcnRpYmxlGBAgASgIEjAKEm5ld19jb21tYW5kX2NoYW5nZRgCIAEoCzISLkNoYW5nZS5OZXdDb21tYW5kSAASMgoTY2hhbmdlX3N0YWdlX2NoYW5nZRgDIAEoCzITLkNoYW5nZS5DaGFuZ2VTdGFnZUgAEkQKHXNldF9iYWxsX3BsYWNlbWVudF9wb3NfY2hhbmdlGAQgASgLMhsuQ2hhbmdlLlNldEJhbGxQbGFjZW1lbnRQb3NIABI3ChZhZGRfeWVsbG93X2NhcmRfY2hhbmdlGAUgASgLMhUuQ2hhbmdlLkFkZFllbGxvd0NhcmRIABIxChNhZGRfcmVkX2NhcmRfY2hhbmdlGAYgASgLMhIuQ2hhbmdlLkFkZFJlZENhcmRIABI5Chd5ZWxsb3dfY2FyZF9vdmVyX2NoYW5nZRgHIAEoCzIWLkNoYW5nZS5ZZWxsb3dDYXJkT3ZlckgAEjUKFWFkZF9nYW1lX2V2ZW50X2NoYW5nZRgIIAEoCzIULkNoYW5nZS5BZGRHYW1lRXZlbnRIABJECh1hZGRfcGFzc2l2ZV9nYW1lX2V2ZW50X2NoYW5nZRgTIAEoCzIbLkNoYW5nZS5BZGRQYXNzaXZlR2FtZUV2ZW50SAASMgoTYWRkX3Byb3Bvc2FsX2NoYW5nZRgJIAEoCzITLkNoYW5nZS5BZGRQcm9wb3NhbEgAEjQKFHVwZGF0ZV9jb25maWdfY2hhbmdlGAwgASgLMhQuQ2hhbmdlLlVwZGF0ZUNvbmZpZ0gAEjsKGHVwZGF0ZV90ZWFtX3N0YXRlX2NoYW5nZRgNIAEoCzIXLkNoYW5nZS5VcGRhdGVUZWFtU3RhdGVIABI0ChRzd2l0Y2hfY29sb3JzX2NoYW5nZRgOIAEoCzIULkNoYW5nZS5Td2l0Y2hDb2xvcnNIABInCg1yZXZlcnRfY2hhbmdlGA8gASgLMg4uQ2hhbmdlLlJldmVydEgAEjUKFW5ld19nYW1lX3N0YXRlX2NoYW5nZRgRIAEoCzIULkNoYW5nZS5OZXdHYW1lU3RhdGVIABJDChxhY2NlcHRfcHJvcG9zYWxfZ3JvdXBfY2hhbmdlGBIgASgLMhsuQ2hhbmdlLkFjY2VwdFByb3Bvc2FsR3JvdXBIABI9ChlzZXRfc3RhdHVzX21lc3NhZ2VfY2hhbmdlGBQgASgLMhguQ2hhbmdlLlNldFN0YXR1c01lc3NhZ2VIABonCgpOZXdDb21tYW5kEhkKB2NvbW1hbmQYASABKAsyCC5Db21tYW5kGjAKC0NoYW5nZVN0YWdlEiEKCW5ld19zdGFnZRgBIAEoDjIOLlJlZmVyZWUuU3RhZ2UaLAoTU2V0QmFsbFBsYWNlbWVudFBvcxIVCgNwb3MYASABKAsyCC5WZWN0b3IyGlIKDUFkZFllbGxvd0NhcmQSFwoIZm9yX3RlYW0YASABKA4yBS5UZWFtEigKFGNhdXNlZF9ieV9nYW1lX2V2ZW50GAIgASgLMgouR2FtZUV2ZW50Gk8KCkFkZFJlZENhcmQSFwoIZm9yX3RlYW0YASABKA4yBS5UZWFtEigKFGNhdXNlZF9ieV9nYW1lX2V2ZW50GAIgASgLMgouR2FtZUV2ZW50GikKDlllbGxvd0NhcmRPdmVyEhcKCGZvcl90ZWFtGAEgASgOMgUuVGVhbRouCgxBZGRHYW1lRXZlbnQSHgoKZ2FtZV9ldmVudBgBIAEoCzIKLkdhbWVFdmVudBo1ChNBZGRQYXNzaXZlR2FtZUV2ZW50Eh4KCmdhbWVfZXZlbnQYASABKAsyCi5HYW1lRXZlbnQaKgoLQWRkUHJvcG9zYWwSGwoIcHJvcG9zYWwYASABKAsyCS5Qcm9wb3NhbBo8ChNBY2NlcHRQcm9wb3NhbEdyb3VwEhAKCGdyb3VwX2lkGAMgASgJEhMKC2FjY2VwdGVkX2J5GAIgASgJGq4BCgxVcGRhdGVDb25maWcSGwoIZGl2aXNpb24YASABKA4yCS5EaXZpc2lvbhIhChJmaXJzdF9raWNrb2ZmX3RlYW0YAiABKA4yBS5UZWFtEh4KCm1hdGNoX3R5cGUYBCABKA4yCi5NYXRjaFR5cGUSOAoTbWF4X3JvYm90c19wZXJfdGVhbRgFIAEoCzIbLmdvb2dsZS5wcm90b2J1Zi5JbnQzMlZhbHVlSgQIAxAEGqYICg9VcGRhdGVUZWFtU3RhdGUSFwoIZm9yX3RlYW0YASABKA4yBS5UZWFtEi8KCXRlYW1fbmFtZRgCIAEoCzIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRIqCgVnb2FscxgDIAEoCzIbLmdvb2dsZS5wcm90b2J1Zi5JbnQzMlZhbHVlEi8KCmdvYWxrZWVwZXIYBCABKAsyGy5nb29nbGUucHJvdG9idWYuSW50MzJWYWx1ZRIyCg10aW1lb3V0c19sZWZ0GAUgASgLMhsuZ29vZ2xlLnByb3RvYnVmLkludDMyVmFsdWUSNwoRdGltZW91dF90aW1lX2xlZnQYBiABKAsyHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSNAoQb25fcG9zaXRpdmVfaGFsZhgHIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5Cb29sVmFsdWUSPAoXYmFsbF9wbGFjZW1lbnRfZmFpbHVyZXMYCCABKAsyGy5nb29nbGUucHJvdG9idWYuSW50MzJWYWx1ZRIyCg5jYW5fcGxhY2VfYmFsbBgJIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5Cb29sVmFsdWUSOQoUY2hhbGxlbmdlX2ZsYWdzX2xlZnQYFSABKAsyGy5nb29nbGUucHJvdG9idWYuSW50MzJWYWx1ZRI7ChZib3Rfc3Vic3RpdHV0aW9uc19sZWZ0GBYgASgLMhsuZ29vZ2xlLnByb3RvYnVmLkludDMyVmFsdWUSPQoZcmVxdWVzdHNfYm90X3N1YnN0aXR1dGlvbhgKIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5Cb29sVmFsdWUSNAoQcmVxdWVzdHNfdGltZW91dBgRIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5Cb29sVmFsdWUSNgoScmVxdWVzdHNfY2hhbGxlbmdlGBIgASgLMhouZ29vZ2xlLnByb3RvYnVmLkJvb2xWYWx1ZRI7ChdyZXF1ZXN0c19lbWVyZ2VuY3lfc3RvcBgTIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5Cb29sVmFsdWUSIAoLeWVsbG93X2NhcmQYFCABKAsyCy5ZZWxsb3dDYXJkEhoKCHJlZF9jYXJkGAwgASgLMgguUmVkQ2FyZBITCgRmb3VsGA0gASgLMgUuRm91bBI4ChJyZW1vdmVfeWVsbG93X2NhcmQYDiABKAsyHC5nb29nbGUucHJvdG9idWYuVUludDMyVmFsdWUSNQoPcmVtb3ZlX3JlZF9jYXJkGA8gASgLMhwuZ29vZ2xlLnByb3RvYnVmLlVJbnQzMlZhbHVlEjEKC3JlbW92ZV9mb3VsGBAgASgLMhwuZ29vZ2xlLnByb3RvYnVmLlVJbnQzMlZhbHVlGg4KDFN3aXRjaENvbG9ycxobCgZSZXZlcnQSEQoJY2hhbmdlX2lkGAEgASgFGi4KDE5ld0dhbWVTdGF0ZRIeCgpnYW1lX3N0YXRlGAEgASgLMgouR2FtZVN0YXRlGioKEFNldFN0YXR1c01lc3NhZ2USFgoOc3RhdHVzX21lc3NhZ2UYASABKAlCCAoGY2hhbmdlQlpCEFNzbEdjQ2hhbmdlUHJvdG9QAVpEZ2l0aHViLmNvbS9Sb2JvQ3VwLVNTTC9zc2wtZ2FtZS1jb250cm9sbGVyL2ludGVybmFsL2FwcC9zdGF0ZW1hY2hpbmU", [file_state_ssl_gc_state, file_state_ssl_gc_common, file_geom_ssl_gc_geometry, file_state_ssl_gc_game_event, file_state_ssl_gc_referee_message, file_google_protobuf_timestamp, file_google_protobuf_wrappers]);

/**
 * A state change
 *
 * @generated from message StateChange
 */
export type StateChange = Message<"StateChange"> & {
  /**
   * A unique increasing id
   *
   * @generated from field: optional int32 id = 1;
   */
  id: number;

  /**
   * The previous state
   *
   * @generated from field: optional State state_pre = 2;
   */
  statePre?: State;

  /**
   * The state after the change was applied
   *
   * @generated from field: optional State state = 3;
   */
  state?: State;

  /**
   * The change itself
   *
   * @generated from field: optional Change change = 4;
   */
  change?: Change;

  /**
   * The timestamp when the change was triggered
   *
   * @generated from field: optional google.protobuf.Timestamp timestamp = 5;
   */
  timestamp?: Timestamp;
};

/**
 * A state change
 *
 * @generated from message StateChange
 */
export type StateChangeJson = {
  /**
   * A unique increasing id
   *
   * @generated from field: optional int32 id = 1;
   */
  id?: number;

  /**
   * The previous state
   *
   * @generated from field: optional State state_pre = 2;
   */
  statePre?: StateJson;

  /**
   * The state after the change was applied
   *
   * @generated from field: optional State state = 3;
   */
  state?: StateJson;

  /**
   * The change itself
   *
   * @generated from field: optional Change change = 4;
   */
  change?: ChangeJson;

  /**
   * The timestamp when the change was triggered
   *
   * @generated from field: optional google.protobuf.Timestamp timestamp = 5;
   */
  timestamp?: TimestampJson;
};

/**
 * Describes the message StateChange.
 * Use `create(StateChangeSchema)` to create a new message.
 */
export const StateChangeSchema: GenMessage<StateChange, StateChangeJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 0);

/**
 * A certain change
 *
 * @generated from message Change
 */
export type Change = Message<"Change"> & {
  /**
   * An identifier of the origin that triggered the change
   *
   * @generated from field: optional string origin = 1;
   */
  origin: string;

  /**
   * Is this change revertible?
   *
   * @generated from field: optional bool revertible = 16;
   */
  revertible: boolean;

  /**
   * @generated from oneof Change.change
   */
  change: {
    /**
     * @generated from field: Change.NewCommand new_command_change = 2;
     */
    value: Change_NewCommand;
    case: "newCommandChange";
  } | {
    /**
     * @generated from field: Change.ChangeStage change_stage_change = 3;
     */
    value: Change_ChangeStage;
    case: "changeStageChange";
  } | {
    /**
     * @generated from field: Change.SetBallPlacementPos set_ball_placement_pos_change = 4;
     */
    value: Change_SetBallPlacementPos;
    case: "setBallPlacementPosChange";
  } | {
    /**
     * @generated from field: Change.AddYellowCard add_yellow_card_change = 5;
     */
    value: Change_AddYellowCard;
    case: "addYellowCardChange";
  } | {
    /**
     * @generated from field: Change.AddRedCard add_red_card_change = 6;
     */
    value: Change_AddRedCard;
    case: "addRedCardChange";
  } | {
    /**
     * @generated from field: Change.YellowCardOver yellow_card_over_change = 7;
     */
    value: Change_YellowCardOver;
    case: "yellowCardOverChange";
  } | {
    /**
     * @generated from field: Change.AddGameEvent add_game_event_change = 8;
     */
    value: Change_AddGameEvent;
    case: "addGameEventChange";
  } | {
    /**
     * @generated from field: Change.AddPassiveGameEvent add_passive_game_event_change = 19;
     */
    value: Change_AddPassiveGameEvent;
    case: "addPassiveGameEventChange";
  } | {
    /**
     * @generated from field: Change.AddProposal add_proposal_change = 9;
     */
    value: Change_AddProposal;
    case: "addProposalChange";
  } | {
    /**
     * @generated from field: Change.UpdateConfig update_config_change = 12;
     */
    value: Change_UpdateConfig;
    case: "updateConfigChange";
  } | {
    /**
     * @generated from field: Change.UpdateTeamState update_team_state_change = 13;
     */
    value: Change_UpdateTeamState;
    case: "updateTeamStateChange";
  } | {
    /**
     * @generated from field: Change.SwitchColors switch_colors_change = 14;
     */
    value: Change_SwitchColors;
    case: "switchColorsChange";
  } | {
    /**
     * @generated from field: Change.Revert revert_change = 15;
     */
    value: Change_Revert;
    case: "revertChange";
  } | {
    /**
     * @generated from field: Change.NewGameState new_game_state_change = 17;
     */
    value: Change_NewGameState;
    case: "newGameStateChange";
  } | {
    /**
     * @generated from field: Change.AcceptProposalGroup accept_proposal_group_change = 18;
     */
    value: Change_AcceptProposalGroup;
    case: "acceptProposalGroupChange";
  } | {
    /**
     * @generated from field: Change.SetStatusMessage set_status_message_change = 20;
     */
    value: Change_SetStatusMessage;
    case: "setStatusMessageChange";
  } | { case: undefined; value?: undefined };
};

/**
 * A certain change
 *
 * @generated from message Change
 */
export type ChangeJson = {
  /**
   * An identifier of the origin that triggered the change
   *
   * @generated from field: optional string origin = 1;
   */
  origin?: string;

  /**
   * Is this change revertible?
   *
   * @generated from field: optional bool revertible = 16;
   */
  revertible?: boolean;

  /**
   * @generated from field: Change.NewCommand new_command_change = 2;
   */
  newCommandChange?: Change_NewCommandJson;

  /**
   * @generated from field: Change.ChangeStage change_stage_change = 3;
   */
  changeStageChange?: Change_ChangeStageJson;

  /**
   * @generated from field: Change.SetBallPlacementPos set_ball_placement_pos_change = 4;
   */
  setBallPlacementPosChange?: Change_SetBallPlacementPosJson;

  /**
   * @generated from field: Change.AddYellowCard add_yellow_card_change = 5;
   */
  addYellowCardChange?: Change_AddYellowCardJson;

  /**
   * @generated from field: Change.AddRedCard add_red_card_change = 6;
   */
  addRedCardChange?: Change_AddRedCardJson;

  /**
   * @generated from field: Change.YellowCardOver yellow_card_over_change = 7;
   */
  yellowCardOverChange?: Change_YellowCardOverJson;

  /**
   * @generated from field: Change.AddGameEvent add_game_event_change = 8;
   */
  addGameEventChange?: Change_AddGameEventJson;

  /**
   * @generated from field: Change.AddPassiveGameEvent add_passive_game_event_change = 19;
   */
  addPassiveGameEventChange?: Change_AddPassiveGameEventJson;

  /**
   * @generated from field: Change.AddProposal add_proposal_change = 9;
   */
  addProposalChange?: Change_AddProposalJson;

  /**
   * @generated from field: Change.UpdateConfig update_config_change = 12;
   */
  updateConfigChange?: Change_UpdateConfigJson;

  /**
   * @generated from field: Change.UpdateTeamState update_team_state_change = 13;
   */
  updateTeamStateChange?: Change_UpdateTeamStateJson;

  /**
   * @generated from field: Change.SwitchColors switch_colors_change = 14;
   */
  switchColorsChange?: Change_SwitchColorsJson;

  /**
   * @generated from field: Change.Revert revert_change = 15;
   */
  revertChange?: Change_RevertJson;

  /**
   * @generated from field: Change.NewGameState new_game_state_change = 17;
   */
  newGameStateChange?: Change_NewGameStateJson;

  /**
   * @generated from field: Change.AcceptProposalGroup accept_proposal_group_change = 18;
   */
  acceptProposalGroupChange?: Change_AcceptProposalGroupJson;

  /**
   * @generated from field: Change.SetStatusMessage set_status_message_change = 20;
   */
  setStatusMessageChange?: Change_SetStatusMessageJson;
};

/**
 * Describes the message Change.
 * Use `create(ChangeSchema)` to create a new message.
 */
export const ChangeSchema: GenMessage<Change, ChangeJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1);

/**
 * New referee command
 *
 * @generated from message Change.NewCommand
 */
export type Change_NewCommand = Message<"Change.NewCommand"> & {
  /**
   * The command
   *
   * @generated from field: optional Command command = 1;
   */
  command?: Command;
};

/**
 * New referee command
 *
 * @generated from message Change.NewCommand
 */
export type Change_NewCommandJson = {
  /**
   * The command
   *
   * @generated from field: optional Command command = 1;
   */
  command?: CommandJson;
};

/**
 * Describes the message Change.NewCommand.
 * Use `create(Change_NewCommandSchema)` to create a new message.
 */
export const Change_NewCommandSchema: GenMessage<Change_NewCommand, Change_NewCommandJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 0);

/**
 * Switch to a new stage
 *
 * @generated from message Change.ChangeStage
 */
export type Change_ChangeStage = Message<"Change.ChangeStage"> & {
  /**
   * The new stage
   *
   * @generated from field: optional Referee.Stage new_stage = 1;
   */
  newStage: Referee_Stage;
};

/**
 * Switch to a new stage
 *
 * @generated from message Change.ChangeStage
 */
export type Change_ChangeStageJson = {
  /**
   * The new stage
   *
   * @generated from field: optional Referee.Stage new_stage = 1;
   */
  newStage?: Referee_StageJson;
};

/**
 * Describes the message Change.ChangeStage.
 * Use `create(Change_ChangeStageSchema)` to create a new message.
 */
export const Change_ChangeStageSchema: GenMessage<Change_ChangeStage, Change_ChangeStageJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 1);

/**
 * Set the ball placement pos
 *
 * @generated from message Change.SetBallPlacementPos
 */
export type Change_SetBallPlacementPos = Message<"Change.SetBallPlacementPos"> & {
  /**
   * The position in [m]
   *
   * @generated from field: optional Vector2 pos = 1;
   */
  pos?: Vector2;
};

/**
 * Set the ball placement pos
 *
 * @generated from message Change.SetBallPlacementPos
 */
export type Change_SetBallPlacementPosJson = {
  /**
   * The position in [m]
   *
   * @generated from field: optional Vector2 pos = 1;
   */
  pos?: Vector2Json;
};

/**
 * Describes the message Change.SetBallPlacementPos.
 * Use `create(Change_SetBallPlacementPosSchema)` to create a new message.
 */
export const Change_SetBallPlacementPosSchema: GenMessage<Change_SetBallPlacementPos, Change_SetBallPlacementPosJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 2);

/**
 * Add a new yellow card
 *
 * @generated from message Change.AddYellowCard
 */
export type Change_AddYellowCard = Message<"Change.AddYellowCard"> & {
  /**
   * The team that the card is for
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam: Team;

  /**
   * The game event that caused the card
   *
   * @generated from field: optional GameEvent caused_by_game_event = 2;
   */
  causedByGameEvent?: GameEvent;
};

/**
 * Add a new yellow card
 *
 * @generated from message Change.AddYellowCard
 */
export type Change_AddYellowCardJson = {
  /**
   * The team that the card is for
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam?: TeamJson;

  /**
   * The game event that caused the card
   *
   * @generated from field: optional GameEvent caused_by_game_event = 2;
   */
  causedByGameEvent?: GameEventJson;
};

/**
 * Describes the message Change.AddYellowCard.
 * Use `create(Change_AddYellowCardSchema)` to create a new message.
 */
export const Change_AddYellowCardSchema: GenMessage<Change_AddYellowCard, Change_AddYellowCardJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 3);

/**
 * Add a new red card
 *
 * @generated from message Change.AddRedCard
 */
export type Change_AddRedCard = Message<"Change.AddRedCard"> & {
  /**
   * The team that the card is for
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam: Team;

  /**
   * The game event that caused the card
   *
   * @generated from field: optional GameEvent caused_by_game_event = 2;
   */
  causedByGameEvent?: GameEvent;
};

/**
 * Add a new red card
 *
 * @generated from message Change.AddRedCard
 */
export type Change_AddRedCardJson = {
  /**
   * The team that the card is for
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam?: TeamJson;

  /**
   * The game event that caused the card
   *
   * @generated from field: optional GameEvent caused_by_game_event = 2;
   */
  causedByGameEvent?: GameEventJson;
};

/**
 * Describes the message Change.AddRedCard.
 * Use `create(Change_AddRedCardSchema)` to create a new message.
 */
export const Change_AddRedCardSchema: GenMessage<Change_AddRedCard, Change_AddRedCardJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 4);

/**
 * Trigger when a yellow card timed out
 *
 * @generated from message Change.YellowCardOver
 */
export type Change_YellowCardOver = Message<"Change.YellowCardOver"> & {
  /**
   * The team that the card was for
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam: Team;
};

/**
 * Trigger when a yellow card timed out
 *
 * @generated from message Change.YellowCardOver
 */
export type Change_YellowCardOverJson = {
  /**
   * The team that the card was for
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam?: TeamJson;
};

/**
 * Describes the message Change.YellowCardOver.
 * Use `create(Change_YellowCardOverSchema)` to create a new message.
 */
export const Change_YellowCardOverSchema: GenMessage<Change_YellowCardOver, Change_YellowCardOverJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 5);

/**
 * Add a new game event
 *
 * @generated from message Change.AddGameEvent
 */
export type Change_AddGameEvent = Message<"Change.AddGameEvent"> & {
  /**
   * The game event
   *
   * @generated from field: optional GameEvent game_event = 1;
   */
  gameEvent?: GameEvent;
};

/**
 * Add a new game event
 *
 * @generated from message Change.AddGameEvent
 */
export type Change_AddGameEventJson = {
  /**
   * The game event
   *
   * @generated from field: optional GameEvent game_event = 1;
   */
  gameEvent?: GameEventJson;
};

/**
 * Describes the message Change.AddGameEvent.
 * Use `create(Change_AddGameEventSchema)` to create a new message.
 */
export const Change_AddGameEventSchema: GenMessage<Change_AddGameEvent, Change_AddGameEventJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 6);

/**
 * Add a new passive game event (that is only logged, but does not automatically trigger anything)
 *
 * @generated from message Change.AddPassiveGameEvent
 */
export type Change_AddPassiveGameEvent = Message<"Change.AddPassiveGameEvent"> & {
  /**
   * The game event
   *
   * @generated from field: optional GameEvent game_event = 1;
   */
  gameEvent?: GameEvent;
};

/**
 * Add a new passive game event (that is only logged, but does not automatically trigger anything)
 *
 * @generated from message Change.AddPassiveGameEvent
 */
export type Change_AddPassiveGameEventJson = {
  /**
   * The game event
   *
   * @generated from field: optional GameEvent game_event = 1;
   */
  gameEvent?: GameEventJson;
};

/**
 * Describes the message Change.AddPassiveGameEvent.
 * Use `create(Change_AddPassiveGameEventSchema)` to create a new message.
 */
export const Change_AddPassiveGameEventSchema: GenMessage<Change_AddPassiveGameEvent, Change_AddPassiveGameEventJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 7);

/**
 * Add a new proposal (i.e. from an auto referee for majority voting)
 *
 * @generated from message Change.AddProposal
 */
export type Change_AddProposal = Message<"Change.AddProposal"> & {
  /**
   * The proposal
   *
   * @generated from field: optional Proposal proposal = 1;
   */
  proposal?: Proposal;
};

/**
 * Add a new proposal (i.e. from an auto referee for majority voting)
 *
 * @generated from message Change.AddProposal
 */
export type Change_AddProposalJson = {
  /**
   * The proposal
   *
   * @generated from field: optional Proposal proposal = 1;
   */
  proposal?: ProposalJson;
};

/**
 * Describes the message Change.AddProposal.
 * Use `create(Change_AddProposalSchema)` to create a new message.
 */
export const Change_AddProposalSchema: GenMessage<Change_AddProposal, Change_AddProposalJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 8);

/**
 * Accept a proposal group (that contain one or more proposals of the same type)
 *
 * @generated from message Change.AcceptProposalGroup
 */
export type Change_AcceptProposalGroup = Message<"Change.AcceptProposalGroup"> & {
  /**
   * The id of the group
   *
   * @generated from field: optional string group_id = 3;
   */
  groupId: string;

  /**
   * An identifier of the acceptor
   *
   * @generated from field: optional string accepted_by = 2;
   */
  acceptedBy: string;
};

/**
 * Accept a proposal group (that contain one or more proposals of the same type)
 *
 * @generated from message Change.AcceptProposalGroup
 */
export type Change_AcceptProposalGroupJson = {
  /**
   * The id of the group
   *
   * @generated from field: optional string group_id = 3;
   */
  groupId?: string;

  /**
   * An identifier of the acceptor
   *
   * @generated from field: optional string accepted_by = 2;
   */
  acceptedBy?: string;
};

/**
 * Describes the message Change.AcceptProposalGroup.
 * Use `create(Change_AcceptProposalGroupSchema)` to create a new message.
 */
export const Change_AcceptProposalGroupSchema: GenMessage<Change_AcceptProposalGroup, Change_AcceptProposalGroupJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 9);

/**
 * Update some configuration
 *
 * @generated from message Change.UpdateConfig
 */
export type Change_UpdateConfig = Message<"Change.UpdateConfig"> & {
  /**
   * The division to play with
   *
   * @generated from field: optional Division division = 1;
   */
  division: Division;

  /**
   * the team that does/did the first kick off
   *
   * @generated from field: optional Team first_kickoff_team = 2;
   */
  firstKickoffTeam: Team;

  /**
   * The match type
   *
   * @generated from field: optional MatchType match_type = 4;
   */
  matchType: MatchType;

  /**
   * The number of robots per team
   *
   * @generated from field: optional google.protobuf.Int32Value max_robots_per_team = 5;
   */
  maxRobotsPerTeam?: number;
};

/**
 * Update some configuration
 *
 * @generated from message Change.UpdateConfig
 */
export type Change_UpdateConfigJson = {
  /**
   * The division to play with
   *
   * @generated from field: optional Division division = 1;
   */
  division?: DivisionJson;

  /**
   * the team that does/did the first kick off
   *
   * @generated from field: optional Team first_kickoff_team = 2;
   */
  firstKickoffTeam?: TeamJson;

  /**
   * The match type
   *
   * @generated from field: optional MatchType match_type = 4;
   */
  matchType?: MatchTypeJson;

  /**
   * The number of robots per team
   *
   * @generated from field: optional google.protobuf.Int32Value max_robots_per_team = 5;
   */
  maxRobotsPerTeam?: Int32ValueJson;
};

/**
 * Describes the message Change.UpdateConfig.
 * Use `create(Change_UpdateConfigSchema)` to create a new message.
 */
export const Change_UpdateConfigSchema: GenMessage<Change_UpdateConfig, Change_UpdateConfigJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 10);

/**
 * Update the current state of a team (all fields that should be updated are set)
 *
 * @generated from message Change.UpdateTeamState
 */
export type Change_UpdateTeamState = Message<"Change.UpdateTeamState"> & {
  /**
   * The team
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam: Team;

  /**
   * Change the name of the team
   *
   * @generated from field: optional google.protobuf.StringValue team_name = 2;
   */
  teamName?: string;

  /**
   * Change the number of goals that the teams has at the moment
   *
   * @generated from field: optional google.protobuf.Int32Value goals = 3;
   */
  goals?: number;

  /**
   * The id of the goal keeper
   *
   * @generated from field: optional google.protobuf.Int32Value goalkeeper = 4;
   */
  goalkeeper?: number;

  /**
   * The number of timeouts that the team has left
   *
   * @generated from field: optional google.protobuf.Int32Value timeouts_left = 5;
   */
  timeoutsLeft?: number;

  /**
   * The timeout time that the team has left
   *
   * @generated from field: optional google.protobuf.StringValue timeout_time_left = 6;
   */
  timeoutTimeLeft?: string;

  /**
   * Does the team play on the positive or the negative half (in ssl-vision coordinates)?
   *
   * @generated from field: optional google.protobuf.BoolValue on_positive_half = 7;
   */
  onPositiveHalf?: boolean;

  /**
   * The number of ball placement failures
   *
   * @generated from field: optional google.protobuf.Int32Value ball_placement_failures = 8;
   */
  ballPlacementFailures?: number;

  /**
   * Can the team place the ball, or is ball placement for this team disabled and should be skipped?
   *
   * @generated from field: optional google.protobuf.BoolValue can_place_ball = 9;
   */
  canPlaceBall?: boolean;

  /**
   * The number of challenge flags that the team has left
   *
   * @generated from field: optional google.protobuf.Int32Value challenge_flags_left = 21;
   */
  challengeFlagsLeft?: number;

  /**
   * The number of bot substitutions left by the team in this halftime
   *
   * @generated from field: optional google.protobuf.Int32Value bot_substitutions_left = 22;
   */
  botSubstitutionsLeft?: number;

  /**
   * Does the team want to substitute a robot in the next possible situation?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_bot_substitution = 10;
   */
  requestsBotSubstitution?: boolean;

  /**
   * Does the team want to take a timeout in the next possible situation?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_timeout = 17;
   */
  requestsTimeout?: boolean;

  /**
   * Does the team want to challenge a recent decision of the referee?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_challenge = 18;
   */
  requestsChallenge?: boolean;

  /**
   * Does the team want to request an emergency stop?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_emergency_stop = 19;
   */
  requestsEmergencyStop?: boolean;

  /**
   * Update a certain yellow card of the team
   *
   * @generated from field: optional YellowCard yellow_card = 20;
   */
  yellowCard?: YellowCard;

  /**
   * Update a certain red card of the team
   *
   * @generated from field: optional RedCard red_card = 12;
   */
  redCard?: RedCard;

  /**
   * Update a certain foul of the team
   *
   * @generated from field: optional Foul foul = 13;
   */
  foul?: Foul;

  /**
   * Remove the yellow card with this id
   *
   * @generated from field: optional google.protobuf.UInt32Value remove_yellow_card = 14;
   */
  removeYellowCard?: number;

  /**
   * Remove the red card with this id
   *
   * @generated from field: optional google.protobuf.UInt32Value remove_red_card = 15;
   */
  removeRedCard?: number;

  /**
   * Remove the foul with this id
   *
   * @generated from field: optional google.protobuf.UInt32Value remove_foul = 16;
   */
  removeFoul?: number;
};

/**
 * Update the current state of a team (all fields that should be updated are set)
 *
 * @generated from message Change.UpdateTeamState
 */
export type Change_UpdateTeamStateJson = {
  /**
   * The team
   *
   * @generated from field: optional Team for_team = 1;
   */
  forTeam?: TeamJson;

  /**
   * Change the name of the team
   *
   * @generated from field: optional google.protobuf.StringValue team_name = 2;
   */
  teamName?: StringValueJson;

  /**
   * Change the number of goals that the teams has at the moment
   *
   * @generated from field: optional google.protobuf.Int32Value goals = 3;
   */
  goals?: Int32ValueJson;

  /**
   * The id of the goal keeper
   *
   * @generated from field: optional google.protobuf.Int32Value goalkeeper = 4;
   */
  goalkeeper?: Int32ValueJson;

  /**
   * The number of timeouts that the team has left
   *
   * @generated from field: optional google.protobuf.Int32Value timeouts_left = 5;
   */
  timeoutsLeft?: Int32ValueJson;

  /**
   * The timeout time that the team has left
   *
   * @generated from field: optional google.protobuf.StringValue timeout_time_left = 6;
   */
  timeoutTimeLeft?: StringValueJson;

  /**
   * Does the team play on the positive or the negative half (in ssl-vision coordinates)?
   *
   * @generated from field: optional google.protobuf.BoolValue on_positive_half = 7;
   */
  onPositiveHalf?: BoolValueJson;

  /**
   * The number of ball placement failures
   *
   * @generated from field: optional google.protobuf.Int32Value ball_placement_failures = 8;
   */
  ballPlacementFailures?: Int32ValueJson;

  /**
   * Can the team place the ball, or is ball placement for this team disabled and should be skipped?
   *
   * @generated from field: optional google.protobuf.BoolValue can_place_ball = 9;
   */
  canPlaceBall?: BoolValueJson;

  /**
   * The number of challenge flags that the team has left
   *
   * @generated from field: optional google.protobuf.Int32Value challenge_flags_left = 21;
   */
  challengeFlagsLeft?: Int32ValueJson;

  /**
   * The number of bot substitutions left by the team in this halftime
   *
   * @generated from field: optional google.protobuf.Int32Value bot_substitutions_left = 22;
   */
  botSubstitutionsLeft?: Int32ValueJson;

  /**
   * Does the team want to substitute a robot in the next possible situation?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_bot_substitution = 10;
   */
  requestsBotSubstitution?: BoolValueJson;

  /**
   * Does the team want to take a timeout in the next possible situation?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_timeout = 17;
   */
  requestsTimeout?: BoolValueJson;

  /**
   * Does the team want to challenge a recent decision of the referee?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_challenge = 18;
   */
  requestsChallenge?: BoolValueJson;

  /**
   * Does the team want to request an emergency stop?
   *
   * @generated from field: optional google.protobuf.BoolValue requests_emergency_stop = 19;
   */
  requestsEmergencyStop?: BoolValueJson;

  /**
   * Update a certain yellow card of the team
   *
   * @generated from field: optional YellowCard yellow_card = 20;
   */
  yellowCard?: YellowCardJson;

  /**
   * Update a certain red card of the team
   *
   * @generated from field: optional RedCard red_card = 12;
   */
  redCard?: RedCardJson;

  /**
   * Update a certain foul of the team
   *
   * @generated from field: optional Foul foul = 13;
   */
  foul?: FoulJson;

  /**
   * Remove the yellow card with this id
   *
   * @generated from field: optional google.protobuf.UInt32Value remove_yellow_card = 14;
   */
  removeYellowCard?: UInt32ValueJson;

  /**
   * Remove the red card with this id
   *
   * @generated from field: optional google.protobuf.UInt32Value remove_red_card = 15;
   */
  removeRedCard?: UInt32ValueJson;

  /**
   * Remove the foul with this id
   *
   * @generated from field: optional google.protobuf.UInt32Value remove_foul = 16;
   */
  removeFoul?: UInt32ValueJson;
};

/**
 * Describes the message Change.UpdateTeamState.
 * Use `create(Change_UpdateTeamStateSchema)` to create a new message.
 */
export const Change_UpdateTeamStateSchema: GenMessage<Change_UpdateTeamState, Change_UpdateTeamStateJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 11);

/**
 * Switch the team colors
 *
 * @generated from message Change.SwitchColors
 */
export type Change_SwitchColors = Message<"Change.SwitchColors"> & {
};

/**
 * Switch the team colors
 *
 * @generated from message Change.SwitchColors
 */
export type Change_SwitchColorsJson = {
};

/**
 * Describes the message Change.SwitchColors.
 * Use `create(Change_SwitchColorsSchema)` to create a new message.
 */
export const Change_SwitchColorsSchema: GenMessage<Change_SwitchColors, Change_SwitchColorsJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 12);

/**
 * Revert a certain change
 *
 * @generated from message Change.Revert
 */
export type Change_Revert = Message<"Change.Revert"> & {
  /**
   * The id of the change
   *
   * @generated from field: optional int32 change_id = 1;
   */
  changeId: number;
};

/**
 * Revert a certain change
 *
 * @generated from message Change.Revert
 */
export type Change_RevertJson = {
  /**
   * The id of the change
   *
   * @generated from field: optional int32 change_id = 1;
   */
  changeId?: number;
};

/**
 * Describes the message Change.Revert.
 * Use `create(Change_RevertSchema)` to create a new message.
 */
export const Change_RevertSchema: GenMessage<Change_Revert, Change_RevertJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 13);

/**
 * Change the current game state
 *
 * @generated from message Change.NewGameState
 */
export type Change_NewGameState = Message<"Change.NewGameState"> & {
  /**
   * The new game state
   *
   * @generated from field: optional GameState game_state = 1;
   */
  gameState?: GameState;
};

/**
 * Change the current game state
 *
 * @generated from message Change.NewGameState
 */
export type Change_NewGameStateJson = {
  /**
   * The new game state
   *
   * @generated from field: optional GameState game_state = 1;
   */
  gameState?: GameStateJson;
};

/**
 * Describes the message Change.NewGameState.
 * Use `create(Change_NewGameStateSchema)` to create a new message.
 */
export const Change_NewGameStateSchema: GenMessage<Change_NewGameState, Change_NewGameStateJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 14);

/**
 * @generated from message Change.SetStatusMessage
 */
export type Change_SetStatusMessage = Message<"Change.SetStatusMessage"> & {
  /**
   * The new status message
   *
   * @generated from field: optional string status_message = 1;
   */
  statusMessage: string;
};

/**
 * @generated from message Change.SetStatusMessage
 */
export type Change_SetStatusMessageJson = {
  /**
   * The new status message
   *
   * @generated from field: optional string status_message = 1;
   */
  statusMessage?: string;
};

/**
 * Describes the message Change.SetStatusMessage.
 * Use `create(Change_SetStatusMessageSchema)` to create a new message.
 */
export const Change_SetStatusMessageSchema: GenMessage<Change_SetStatusMessage, Change_SetStatusMessageJson> = /*@__PURE__*/
  messageDesc(file_statemachine_ssl_gc_change, 1, 15);

