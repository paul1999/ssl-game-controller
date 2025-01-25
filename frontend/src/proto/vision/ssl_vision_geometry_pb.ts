// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file vision/ssl_vision_geometry.proto (syntax proto2)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file vision/ssl_vision_geometry.proto.
 */
export const file_vision_ssl_vision_geometry: GenFile = /*@__PURE__*/
  fileDesc("CiB2aXNpb24vc3NsX3Zpc2lvbl9nZW9tZXRyeS5wcm90byIgCghWZWN0b3IyZhIJCgF4GAEgAigCEgkKAXkYAiACKAIiiAEKFFNTTF9GaWVsZExpbmVTZWdtZW50EgwKBG5hbWUYASACKAkSFQoCcDEYAiACKAsyCS5WZWN0b3IyZhIVCgJwMhgDIAIoCzIJLlZlY3RvcjJmEhEKCXRoaWNrbmVzcxgEIAIoAhIhCgR0eXBlGAUgASgOMhMuU1NMX0ZpZWxkU2hhcGVUeXBlIp0BChRTU0xfRmllbGRDaXJjdWxhckFyYxIMCgRuYW1lGAEgAigJEhkKBmNlbnRlchgCIAIoCzIJLlZlY3RvcjJmEg4KBnJhZGl1cxgDIAIoAhIKCgJhMRgEIAIoAhIKCgJhMhgFIAIoAhIRCgl0aGlja25lc3MYBiACKAISIQoEdHlwZRgHIAEoDjITLlNTTF9GaWVsZFNoYXBlVHlwZSKwAwoVU1NMX0dlb21ldHJ5RmllbGRTaXplEhQKDGZpZWxkX2xlbmd0aBgBIAIoBRITCgtmaWVsZF93aWR0aBgCIAIoBRISCgpnb2FsX3dpZHRoGAMgAigFEhIKCmdvYWxfZGVwdGgYBCACKAUSFgoOYm91bmRhcnlfd2lkdGgYBSACKAUSKgoLZmllbGRfbGluZXMYBiADKAsyFS5TU0xfRmllbGRMaW5lU2VnbWVudBIpCgpmaWVsZF9hcmNzGAcgAygLMhUuU1NMX0ZpZWxkQ2lyY3VsYXJBcmMSGgoScGVuYWx0eV9hcmVhX2RlcHRoGAggASgFEhoKEnBlbmFsdHlfYXJlYV93aWR0aBgJIAEoBRIcChRjZW50ZXJfY2lyY2xlX3JhZGl1cxgKIAEoBRIWCg5saW5lX3RoaWNrbmVzcxgLIAEoBRIjChtnb2FsX2NlbnRlcl90b19wZW5hbHR5X21hcmsYDCABKAUSEwoLZ29hbF9oZWlnaHQYDSABKAUSEwoLYmFsbF9yYWRpdXMYDiABKAISGAoQbWF4X3JvYm90X3JhZGl1cxgPIAEoAiKAAwodU1NMX0dlb21ldHJ5Q2FtZXJhQ2FsaWJyYXRpb24SEQoJY2FtZXJhX2lkGAEgAigNEhQKDGZvY2FsX2xlbmd0aBgCIAIoAhIZChFwcmluY2lwYWxfcG9pbnRfeBgDIAIoAhIZChFwcmluY2lwYWxfcG9pbnRfeRgEIAIoAhISCgpkaXN0b3J0aW9uGAUgAigCEgoKAnEwGAYgAigCEgoKAnExGAcgAigCEgoKAnEyGAggAigCEgoKAnEzGAkgAigCEgoKAnR4GAogAigCEgoKAnR5GAsgAigCEgoKAnR6GAwgAigCEh8KF2Rlcml2ZWRfY2FtZXJhX3dvcmxkX3R4GA0gASgCEh8KF2Rlcml2ZWRfY2FtZXJhX3dvcmxkX3R5GA4gASgCEh8KF2Rlcml2ZWRfY2FtZXJhX3dvcmxkX3R6GA8gASgCEhkKEXBpeGVsX2ltYWdlX3dpZHRoGBAgASgNEhoKEnBpeGVsX2ltYWdlX2hlaWdodBgRIAEoDSJWCh1TU0xfQmFsbE1vZGVsU3RyYWlnaHRUd29QaGFzZRIRCglhY2Nfc2xpZGUYASACKAESEAoIYWNjX3JvbGwYAiACKAESEAoIa19zd2l0Y2gYAyACKAEibAoaU1NMX0JhbGxNb2RlbENoaXBGaXhlZExvc3MSHAoUZGFtcGluZ194eV9maXJzdF9ob3AYASACKAESHQoVZGFtcGluZ194eV9vdGhlcl9ob3BzGAIgAigBEhEKCWRhbXBpbmdfehgDIAIoASKGAQoSU1NMX0dlb21ldHJ5TW9kZWxzEjoKEnN0cmFpZ2h0X3R3b19waGFzZRgBIAEoCzIeLlNTTF9CYWxsTW9kZWxTdHJhaWdodFR3b1BoYXNlEjQKD2NoaXBfZml4ZWRfbG9zcxgCIAEoCzIbLlNTTF9CYWxsTW9kZWxDaGlwRml4ZWRMb3NzIo0BChBTU0xfR2VvbWV0cnlEYXRhEiUKBWZpZWxkGAEgAigLMhYuU1NMX0dlb21ldHJ5RmllbGRTaXplEi0KBWNhbGliGAIgAygLMh4uU1NMX0dlb21ldHJ5Q2FtZXJhQ2FsaWJyYXRpb24SIwoGbW9kZWxzGAMgASgLMhMuU1NMX0dlb21ldHJ5TW9kZWxzKtsCChJTU0xfRmllbGRTaGFwZVR5cGUSDQoJVW5kZWZpbmVkEAASEAoMQ2VudGVyQ2lyY2xlEAESEAoMVG9wVG91Y2hMaW5lEAISEwoPQm90dG9tVG91Y2hMaW5lEAMSEAoMTGVmdEdvYWxMaW5lEAQSEQoNUmlnaHRHb2FsTGluZRAFEg8KC0hhbGZ3YXlMaW5lEAYSDgoKQ2VudGVyTGluZRAHEhYKEkxlZnRQZW5hbHR5U3RyZXRjaBAIEhcKE1JpZ2h0UGVuYWx0eVN0cmV0Y2gQCRIfChtMZWZ0RmllbGRMZWZ0UGVuYWx0eVN0cmV0Y2gQChIgChxMZWZ0RmllbGRSaWdodFBlbmFsdHlTdHJldGNoEAsSIAocUmlnaHRGaWVsZExlZnRQZW5hbHR5U3RyZXRjaBAMEiEKHVJpZ2h0RmllbGRSaWdodFBlbmFsdHlTdHJldGNoEA1CWkIWU3NsVmlzaW9uR2VvbWV0cnlQcm90b1ABWj5naXRodWIuY29tL1JvYm9DdXAtU1NML3NzbC1nYW1lLWNvbnRyb2xsZXIvaW50ZXJuYWwvYXBwL3Zpc2lvbg");

/**
 * A 2D float vector.
 *
 * @generated from message Vector2f
 */
export type Vector2f = Message<"Vector2f"> & {
  /**
   * @generated from field: required float x = 1;
   */
  x: number;

  /**
   * @generated from field: required float y = 2;
   */
  y: number;
};

/**
 * A 2D float vector.
 *
 * @generated from message Vector2f
 */
export type Vector2fJson = {
  /**
   * @generated from field: required float x = 1;
   */
  x?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float y = 2;
   */
  y?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message Vector2f.
 * Use `create(Vector2fSchema)` to create a new message.
 */
export const Vector2fSchema: GenMessage<Vector2f, Vector2fJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 0);

/**
 * Represents a field marking as a line segment represented by a start point p1,
 * and end point p2, and a line thickness. The start and end points are along
 * the center of the line, so the thickness of the line extends by thickness / 2
 * on either side of the line.
 *
 * @generated from message SSL_FieldLineSegment
 */
export type SSL_FieldLineSegment = Message<"SSL_FieldLineSegment"> & {
  /**
   * Name of this field marking.
   *
   * @generated from field: required string name = 1;
   */
  name: string;

  /**
   * Start point of the line segment.
   *
   * @generated from field: required Vector2f p1 = 2;
   */
  p1?: Vector2f;

  /**
   * End point of the line segment.
   *
   * @generated from field: required Vector2f p2 = 3;
   */
  p2?: Vector2f;

  /**
   * Thickness of the line segment.
   *
   * @generated from field: required float thickness = 4;
   */
  thickness: number;

  /**
   * The type of this shape
   *
   * @generated from field: optional SSL_FieldShapeType type = 5;
   */
  type: SSL_FieldShapeType;
};

/**
 * Represents a field marking as a line segment represented by a start point p1,
 * and end point p2, and a line thickness. The start and end points are along
 * the center of the line, so the thickness of the line extends by thickness / 2
 * on either side of the line.
 *
 * @generated from message SSL_FieldLineSegment
 */
export type SSL_FieldLineSegmentJson = {
  /**
   * Name of this field marking.
   *
   * @generated from field: required string name = 1;
   */
  name?: string;

  /**
   * Start point of the line segment.
   *
   * @generated from field: required Vector2f p1 = 2;
   */
  p1?: Vector2fJson;

  /**
   * End point of the line segment.
   *
   * @generated from field: required Vector2f p2 = 3;
   */
  p2?: Vector2fJson;

  /**
   * Thickness of the line segment.
   *
   * @generated from field: required float thickness = 4;
   */
  thickness?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The type of this shape
   *
   * @generated from field: optional SSL_FieldShapeType type = 5;
   */
  type?: SSL_FieldShapeTypeJson;
};

/**
 * Describes the message SSL_FieldLineSegment.
 * Use `create(SSL_FieldLineSegmentSchema)` to create a new message.
 */
export const SSL_FieldLineSegmentSchema: GenMessage<SSL_FieldLineSegment, SSL_FieldLineSegmentJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 1);

/**
 * Represents a field marking as a circular arc segment represented by center point, a
 * start angle, an end angle, and an arc thickness.
 *
 * @generated from message SSL_FieldCircularArc
 */
export type SSL_FieldCircularArc = Message<"SSL_FieldCircularArc"> & {
  /**
   * Name of this field marking.
   *
   * @generated from field: required string name = 1;
   */
  name: string;

  /**
   * Center point of the circular arc.
   *
   * @generated from field: required Vector2f center = 2;
   */
  center?: Vector2f;

  /**
   * Radius of the arc.
   *
   * @generated from field: required float radius = 3;
   */
  radius: number;

  /**
   * Start angle in counter-clockwise order.
   *
   * @generated from field: required float a1 = 4;
   */
  a1: number;

  /**
   * End angle in counter-clockwise order.
   *
   * @generated from field: required float a2 = 5;
   */
  a2: number;

  /**
   * Thickness of the arc.
   *
   * @generated from field: required float thickness = 6;
   */
  thickness: number;

  /**
   * The type of this shape
   *
   * @generated from field: optional SSL_FieldShapeType type = 7;
   */
  type: SSL_FieldShapeType;
};

/**
 * Represents a field marking as a circular arc segment represented by center point, a
 * start angle, an end angle, and an arc thickness.
 *
 * @generated from message SSL_FieldCircularArc
 */
export type SSL_FieldCircularArcJson = {
  /**
   * Name of this field marking.
   *
   * @generated from field: required string name = 1;
   */
  name?: string;

  /**
   * Center point of the circular arc.
   *
   * @generated from field: required Vector2f center = 2;
   */
  center?: Vector2fJson;

  /**
   * Radius of the arc.
   *
   * @generated from field: required float radius = 3;
   */
  radius?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Start angle in counter-clockwise order.
   *
   * @generated from field: required float a1 = 4;
   */
  a1?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * End angle in counter-clockwise order.
   *
   * @generated from field: required float a2 = 5;
   */
  a2?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Thickness of the arc.
   *
   * @generated from field: required float thickness = 6;
   */
  thickness?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The type of this shape
   *
   * @generated from field: optional SSL_FieldShapeType type = 7;
   */
  type?: SSL_FieldShapeTypeJson;
};

/**
 * Describes the message SSL_FieldCircularArc.
 * Use `create(SSL_FieldCircularArcSchema)` to create a new message.
 */
export const SSL_FieldCircularArcSchema: GenMessage<SSL_FieldCircularArc, SSL_FieldCircularArcJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 2);

/**
 * @generated from message SSL_GeometryFieldSize
 */
export type SSL_GeometryFieldSize = Message<"SSL_GeometryFieldSize"> & {
  /**
   * @generated from field: required int32 field_length = 1;
   */
  fieldLength: number;

  /**
   * @generated from field: required int32 field_width = 2;
   */
  fieldWidth: number;

  /**
   * @generated from field: required int32 goal_width = 3;
   */
  goalWidth: number;

  /**
   * @generated from field: required int32 goal_depth = 4;
   */
  goalDepth: number;

  /**
   * @generated from field: required int32 boundary_width = 5;
   */
  boundaryWidth: number;

  /**
   * @generated from field: repeated SSL_FieldLineSegment field_lines = 6;
   */
  fieldLines: SSL_FieldLineSegment[];

  /**
   * @generated from field: repeated SSL_FieldCircularArc field_arcs = 7;
   */
  fieldArcs: SSL_FieldCircularArc[];

  /**
   * @generated from field: optional int32 penalty_area_depth = 8;
   */
  penaltyAreaDepth: number;

  /**
   * @generated from field: optional int32 penalty_area_width = 9;
   */
  penaltyAreaWidth: number;

  /**
   * @generated from field: optional int32 center_circle_radius = 10;
   */
  centerCircleRadius: number;

  /**
   * @generated from field: optional int32 line_thickness = 11;
   */
  lineThickness: number;

  /**
   * @generated from field: optional int32 goal_center_to_penalty_mark = 12;
   */
  goalCenterToPenaltyMark: number;

  /**
   * @generated from field: optional int32 goal_height = 13;
   */
  goalHeight: number;

  /**
   * @generated from field: optional float ball_radius = 14;
   */
  ballRadius: number;

  /**
   * @generated from field: optional float max_robot_radius = 15;
   */
  maxRobotRadius: number;
};

/**
 * @generated from message SSL_GeometryFieldSize
 */
export type SSL_GeometryFieldSizeJson = {
  /**
   * @generated from field: required int32 field_length = 1;
   */
  fieldLength?: number;

  /**
   * @generated from field: required int32 field_width = 2;
   */
  fieldWidth?: number;

  /**
   * @generated from field: required int32 goal_width = 3;
   */
  goalWidth?: number;

  /**
   * @generated from field: required int32 goal_depth = 4;
   */
  goalDepth?: number;

  /**
   * @generated from field: required int32 boundary_width = 5;
   */
  boundaryWidth?: number;

  /**
   * @generated from field: repeated SSL_FieldLineSegment field_lines = 6;
   */
  fieldLines?: SSL_FieldLineSegmentJson[];

  /**
   * @generated from field: repeated SSL_FieldCircularArc field_arcs = 7;
   */
  fieldArcs?: SSL_FieldCircularArcJson[];

  /**
   * @generated from field: optional int32 penalty_area_depth = 8;
   */
  penaltyAreaDepth?: number;

  /**
   * @generated from field: optional int32 penalty_area_width = 9;
   */
  penaltyAreaWidth?: number;

  /**
   * @generated from field: optional int32 center_circle_radius = 10;
   */
  centerCircleRadius?: number;

  /**
   * @generated from field: optional int32 line_thickness = 11;
   */
  lineThickness?: number;

  /**
   * @generated from field: optional int32 goal_center_to_penalty_mark = 12;
   */
  goalCenterToPenaltyMark?: number;

  /**
   * @generated from field: optional int32 goal_height = 13;
   */
  goalHeight?: number;

  /**
   * @generated from field: optional float ball_radius = 14;
   */
  ballRadius?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float max_robot_radius = 15;
   */
  maxRobotRadius?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message SSL_GeometryFieldSize.
 * Use `create(SSL_GeometryFieldSizeSchema)` to create a new message.
 */
export const SSL_GeometryFieldSizeSchema: GenMessage<SSL_GeometryFieldSize, SSL_GeometryFieldSizeJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 3);

/**
 * @generated from message SSL_GeometryCameraCalibration
 */
export type SSL_GeometryCameraCalibration = Message<"SSL_GeometryCameraCalibration"> & {
  /**
   * @generated from field: required uint32 camera_id = 1;
   */
  cameraId: number;

  /**
   * @generated from field: required float focal_length = 2;
   */
  focalLength: number;

  /**
   * @generated from field: required float principal_point_x = 3;
   */
  principalPointX: number;

  /**
   * @generated from field: required float principal_point_y = 4;
   */
  principalPointY: number;

  /**
   * @generated from field: required float distortion = 5;
   */
  distortion: number;

  /**
   * @generated from field: required float q0 = 6;
   */
  q0: number;

  /**
   * @generated from field: required float q1 = 7;
   */
  q1: number;

  /**
   * @generated from field: required float q2 = 8;
   */
  q2: number;

  /**
   * @generated from field: required float q3 = 9;
   */
  q3: number;

  /**
   * @generated from field: required float tx = 10;
   */
  tx: number;

  /**
   * @generated from field: required float ty = 11;
   */
  ty: number;

  /**
   * @generated from field: required float tz = 12;
   */
  tz: number;

  /**
   * @generated from field: optional float derived_camera_world_tx = 13;
   */
  derivedCameraWorldTx: number;

  /**
   * @generated from field: optional float derived_camera_world_ty = 14;
   */
  derivedCameraWorldTy: number;

  /**
   * @generated from field: optional float derived_camera_world_tz = 15;
   */
  derivedCameraWorldTz: number;

  /**
   * @generated from field: optional uint32 pixel_image_width = 16;
   */
  pixelImageWidth: number;

  /**
   * @generated from field: optional uint32 pixel_image_height = 17;
   */
  pixelImageHeight: number;
};

/**
 * @generated from message SSL_GeometryCameraCalibration
 */
export type SSL_GeometryCameraCalibrationJson = {
  /**
   * @generated from field: required uint32 camera_id = 1;
   */
  cameraId?: number;

  /**
   * @generated from field: required float focal_length = 2;
   */
  focalLength?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float principal_point_x = 3;
   */
  principalPointX?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float principal_point_y = 4;
   */
  principalPointY?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float distortion = 5;
   */
  distortion?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float q0 = 6;
   */
  q0?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float q1 = 7;
   */
  q1?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float q2 = 8;
   */
  q2?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float q3 = 9;
   */
  q3?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float tx = 10;
   */
  tx?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float ty = 11;
   */
  ty?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float tz = 12;
   */
  tz?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float derived_camera_world_tx = 13;
   */
  derivedCameraWorldTx?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float derived_camera_world_ty = 14;
   */
  derivedCameraWorldTy?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float derived_camera_world_tz = 15;
   */
  derivedCameraWorldTz?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional uint32 pixel_image_width = 16;
   */
  pixelImageWidth?: number;

  /**
   * @generated from field: optional uint32 pixel_image_height = 17;
   */
  pixelImageHeight?: number;
};

/**
 * Describes the message SSL_GeometryCameraCalibration.
 * Use `create(SSL_GeometryCameraCalibrationSchema)` to create a new message.
 */
export const SSL_GeometryCameraCalibrationSchema: GenMessage<SSL_GeometryCameraCalibration, SSL_GeometryCameraCalibrationJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 4);

/**
 * Two-Phase model for straight-kicked balls.
 * There are two phases with different accelerations during the ball kicks:
 * 1. Sliding
 * 2. Rolling
 * The full model is described in the TDP of ER-Force from 2016, which can be found here:
 * https://ssl.robocup.org/wp-content/uploads/2019/01/2016_ETDP_ER-Force.pdf
 *
 * @generated from message SSL_BallModelStraightTwoPhase
 */
export type SSL_BallModelStraightTwoPhase = Message<"SSL_BallModelStraightTwoPhase"> & {
  /**
   * Ball sliding acceleration [m/s^2] (should be negative)
   *
   * @generated from field: required double acc_slide = 1;
   */
  accSlide: number;

  /**
   * Ball rolling acceleration [m/s^2] (should be negative)
   *
   * @generated from field: required double acc_roll = 2;
   */
  accRoll: number;

  /**
   * Fraction of the initial velocity where the ball starts to roll
   *
   * @generated from field: required double k_switch = 3;
   */
  kSwitch: number;
};

/**
 * Two-Phase model for straight-kicked balls.
 * There are two phases with different accelerations during the ball kicks:
 * 1. Sliding
 * 2. Rolling
 * The full model is described in the TDP of ER-Force from 2016, which can be found here:
 * https://ssl.robocup.org/wp-content/uploads/2019/01/2016_ETDP_ER-Force.pdf
 *
 * @generated from message SSL_BallModelStraightTwoPhase
 */
export type SSL_BallModelStraightTwoPhaseJson = {
  /**
   * Ball sliding acceleration [m/s^2] (should be negative)
   *
   * @generated from field: required double acc_slide = 1;
   */
  accSlide?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Ball rolling acceleration [m/s^2] (should be negative)
   *
   * @generated from field: required double acc_roll = 2;
   */
  accRoll?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Fraction of the initial velocity where the ball starts to roll
   *
   * @generated from field: required double k_switch = 3;
   */
  kSwitch?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message SSL_BallModelStraightTwoPhase.
 * Use `create(SSL_BallModelStraightTwoPhaseSchema)` to create a new message.
 */
export const SSL_BallModelStraightTwoPhaseSchema: GenMessage<SSL_BallModelStraightTwoPhase, SSL_BallModelStraightTwoPhaseJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 5);

/**
 * Fixed-Loss model for chipped balls.
 * Uses fixed damping factors for xy and z direction per hop.
 *
 * @generated from message SSL_BallModelChipFixedLoss
 */
export type SSL_BallModelChipFixedLoss = Message<"SSL_BallModelChipFixedLoss"> & {
  /**
   * Chip kick velocity damping factor in XY direction for the first hop
   *
   * @generated from field: required double damping_xy_first_hop = 1;
   */
  dampingXyFirstHop: number;

  /**
   * Chip kick velocity damping factor in XY direction for all following hops
   *
   * @generated from field: required double damping_xy_other_hops = 2;
   */
  dampingXyOtherHops: number;

  /**
   * Chip kick velocity damping factor in Z direction for all hops
   *
   * @generated from field: required double damping_z = 3;
   */
  dampingZ: number;
};

/**
 * Fixed-Loss model for chipped balls.
 * Uses fixed damping factors for xy and z direction per hop.
 *
 * @generated from message SSL_BallModelChipFixedLoss
 */
export type SSL_BallModelChipFixedLossJson = {
  /**
   * Chip kick velocity damping factor in XY direction for the first hop
   *
   * @generated from field: required double damping_xy_first_hop = 1;
   */
  dampingXyFirstHop?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Chip kick velocity damping factor in XY direction for all following hops
   *
   * @generated from field: required double damping_xy_other_hops = 2;
   */
  dampingXyOtherHops?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Chip kick velocity damping factor in Z direction for all hops
   *
   * @generated from field: required double damping_z = 3;
   */
  dampingZ?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message SSL_BallModelChipFixedLoss.
 * Use `create(SSL_BallModelChipFixedLossSchema)` to create a new message.
 */
export const SSL_BallModelChipFixedLossSchema: GenMessage<SSL_BallModelChipFixedLoss, SSL_BallModelChipFixedLossJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 6);

/**
 * @generated from message SSL_GeometryModels
 */
export type SSL_GeometryModels = Message<"SSL_GeometryModels"> & {
  /**
   * @generated from field: optional SSL_BallModelStraightTwoPhase straight_two_phase = 1;
   */
  straightTwoPhase?: SSL_BallModelStraightTwoPhase;

  /**
   * @generated from field: optional SSL_BallModelChipFixedLoss chip_fixed_loss = 2;
   */
  chipFixedLoss?: SSL_BallModelChipFixedLoss;
};

/**
 * @generated from message SSL_GeometryModels
 */
export type SSL_GeometryModelsJson = {
  /**
   * @generated from field: optional SSL_BallModelStraightTwoPhase straight_two_phase = 1;
   */
  straightTwoPhase?: SSL_BallModelStraightTwoPhaseJson;

  /**
   * @generated from field: optional SSL_BallModelChipFixedLoss chip_fixed_loss = 2;
   */
  chipFixedLoss?: SSL_BallModelChipFixedLossJson;
};

/**
 * Describes the message SSL_GeometryModels.
 * Use `create(SSL_GeometryModelsSchema)` to create a new message.
 */
export const SSL_GeometryModelsSchema: GenMessage<SSL_GeometryModels, SSL_GeometryModelsJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 7);

/**
 * @generated from message SSL_GeometryData
 */
export type SSL_GeometryData = Message<"SSL_GeometryData"> & {
  /**
   * @generated from field: required SSL_GeometryFieldSize field = 1;
   */
  field?: SSL_GeometryFieldSize;

  /**
   * @generated from field: repeated SSL_GeometryCameraCalibration calib = 2;
   */
  calib: SSL_GeometryCameraCalibration[];

  /**
   * @generated from field: optional SSL_GeometryModels models = 3;
   */
  models?: SSL_GeometryModels;
};

/**
 * @generated from message SSL_GeometryData
 */
export type SSL_GeometryDataJson = {
  /**
   * @generated from field: required SSL_GeometryFieldSize field = 1;
   */
  field?: SSL_GeometryFieldSizeJson;

  /**
   * @generated from field: repeated SSL_GeometryCameraCalibration calib = 2;
   */
  calib?: SSL_GeometryCameraCalibrationJson[];

  /**
   * @generated from field: optional SSL_GeometryModels models = 3;
   */
  models?: SSL_GeometryModelsJson;
};

/**
 * Describes the message SSL_GeometryData.
 * Use `create(SSL_GeometryDataSchema)` to create a new message.
 */
export const SSL_GeometryDataSchema: GenMessage<SSL_GeometryData, SSL_GeometryDataJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_geometry, 8);

/**
 * @generated from enum SSL_FieldShapeType
 */
export enum SSL_FieldShapeType {
  /**
   * @generated from enum value: Undefined = 0;
   */
  Undefined = 0,

  /**
   * @generated from enum value: CenterCircle = 1;
   */
  CenterCircle = 1,

  /**
   * @generated from enum value: TopTouchLine = 2;
   */
  TopTouchLine = 2,

  /**
   * @generated from enum value: BottomTouchLine = 3;
   */
  BottomTouchLine = 3,

  /**
   * @generated from enum value: LeftGoalLine = 4;
   */
  LeftGoalLine = 4,

  /**
   * @generated from enum value: RightGoalLine = 5;
   */
  RightGoalLine = 5,

  /**
   * @generated from enum value: HalfwayLine = 6;
   */
  HalfwayLine = 6,

  /**
   * @generated from enum value: CenterLine = 7;
   */
  CenterLine = 7,

  /**
   * @generated from enum value: LeftPenaltyStretch = 8;
   */
  LeftPenaltyStretch = 8,

  /**
   * @generated from enum value: RightPenaltyStretch = 9;
   */
  RightPenaltyStretch = 9,

  /**
   * @generated from enum value: LeftFieldLeftPenaltyStretch = 10;
   */
  LeftFieldLeftPenaltyStretch = 10,

  /**
   * @generated from enum value: LeftFieldRightPenaltyStretch = 11;
   */
  LeftFieldRightPenaltyStretch = 11,

  /**
   * @generated from enum value: RightFieldLeftPenaltyStretch = 12;
   */
  RightFieldLeftPenaltyStretch = 12,

  /**
   * @generated from enum value: RightFieldRightPenaltyStretch = 13;
   */
  RightFieldRightPenaltyStretch = 13,
}

/**
 * @generated from enum SSL_FieldShapeType
 */
export type SSL_FieldShapeTypeJson = "Undefined" | "CenterCircle" | "TopTouchLine" | "BottomTouchLine" | "LeftGoalLine" | "RightGoalLine" | "HalfwayLine" | "CenterLine" | "LeftPenaltyStretch" | "RightPenaltyStretch" | "LeftFieldLeftPenaltyStretch" | "LeftFieldRightPenaltyStretch" | "RightFieldLeftPenaltyStretch" | "RightFieldRightPenaltyStretch";

/**
 * Describes the enum SSL_FieldShapeType.
 */
export const SSL_FieldShapeTypeSchema: GenEnum<SSL_FieldShapeType, SSL_FieldShapeTypeJson> = /*@__PURE__*/
  enumDesc(file_vision_ssl_vision_geometry, 0);

