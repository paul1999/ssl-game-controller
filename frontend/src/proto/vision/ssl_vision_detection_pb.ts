// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file vision/ssl_vision_detection.proto (syntax proto2)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file vision/ssl_vision_detection.proto.
 */
export const file_vision_ssl_vision_detection: GenFile = /*@__PURE__*/
  fileDesc("CiF2aXNpb24vc3NsX3Zpc2lvbl9kZXRlY3Rpb24ucHJvdG8ieAoRU1NMX0RldGVjdGlvbkJhbGwSEgoKY29uZmlkZW5jZRgBIAIoAhIMCgRhcmVhGAIgASgNEgkKAXgYAyACKAISCQoBeRgEIAIoAhIJCgF6GAUgASgCEg8KB3BpeGVsX3gYBiACKAISDwoHcGl4ZWxfeRgHIAIoAiKXAQoSU1NMX0RldGVjdGlvblJvYm90EhIKCmNvbmZpZGVuY2UYASACKAISEAoIcm9ib3RfaWQYAiABKA0SCQoBeBgDIAIoAhIJCgF5GAQgAigCEhMKC29yaWVudGF0aW9uGAUgASgCEg8KB3BpeGVsX3gYBiACKAISDwoHcGl4ZWxfeRgHIAIoAhIOCgZoZWlnaHQYCCABKAIi2QEKElNTTF9EZXRlY3Rpb25GcmFtZRIUCgxmcmFtZV9udW1iZXIYASACKA0SEQoJdF9jYXB0dXJlGAIgAigBEg4KBnRfc2VudBgDIAIoARIRCgljYW1lcmFfaWQYBCACKA0SIQoFYmFsbHMYBSADKAsyEi5TU0xfRGV0ZWN0aW9uQmFsbBIqCg1yb2JvdHNfeWVsbG93GAYgAygLMhMuU1NMX0RldGVjdGlvblJvYm90EigKC3JvYm90c19ibHVlGAcgAygLMhMuU1NMX0RldGVjdGlvblJvYm90QltCF1NzbFZpc2lvbkRldGVjdGlvblByb3RvUAFaPmdpdGh1Yi5jb20vUm9ib0N1cC1TU0wvc3NsLWdhbWUtY29udHJvbGxlci9pbnRlcm5hbC9hcHAvdmlzaW9u");

/**
 * @generated from message SSL_DetectionBall
 */
export type SSL_DetectionBall = Message<"SSL_DetectionBall"> & {
  /**
   * @generated from field: required float confidence = 1;
   */
  confidence: number;

  /**
   * @generated from field: optional uint32 area = 2;
   */
  area: number;

  /**
   * @generated from field: required float x = 3;
   */
  x: number;

  /**
   * @generated from field: required float y = 4;
   */
  y: number;

  /**
   * @generated from field: optional float z = 5;
   */
  z: number;

  /**
   * @generated from field: required float pixel_x = 6;
   */
  pixelX: number;

  /**
   * @generated from field: required float pixel_y = 7;
   */
  pixelY: number;
};

/**
 * @generated from message SSL_DetectionBall
 */
export type SSL_DetectionBallJson = {
  /**
   * @generated from field: required float confidence = 1;
   */
  confidence?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional uint32 area = 2;
   */
  area?: number;

  /**
   * @generated from field: required float x = 3;
   */
  x?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float y = 4;
   */
  y?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float z = 5;
   */
  z?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float pixel_x = 6;
   */
  pixelX?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float pixel_y = 7;
   */
  pixelY?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message SSL_DetectionBall.
 * Use `create(SSL_DetectionBallSchema)` to create a new message.
 */
export const SSL_DetectionBallSchema: GenMessage<SSL_DetectionBall, SSL_DetectionBallJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_detection, 0);

/**
 * @generated from message SSL_DetectionRobot
 */
export type SSL_DetectionRobot = Message<"SSL_DetectionRobot"> & {
  /**
   * @generated from field: required float confidence = 1;
   */
  confidence: number;

  /**
   * @generated from field: optional uint32 robot_id = 2;
   */
  robotId: number;

  /**
   * @generated from field: required float x = 3;
   */
  x: number;

  /**
   * @generated from field: required float y = 4;
   */
  y: number;

  /**
   * @generated from field: optional float orientation = 5;
   */
  orientation: number;

  /**
   * @generated from field: required float pixel_x = 6;
   */
  pixelX: number;

  /**
   * @generated from field: required float pixel_y = 7;
   */
  pixelY: number;

  /**
   * @generated from field: optional float height = 8;
   */
  height: number;
};

/**
 * @generated from message SSL_DetectionRobot
 */
export type SSL_DetectionRobotJson = {
  /**
   * @generated from field: required float confidence = 1;
   */
  confidence?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional uint32 robot_id = 2;
   */
  robotId?: number;

  /**
   * @generated from field: required float x = 3;
   */
  x?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float y = 4;
   */
  y?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float orientation = 5;
   */
  orientation?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float pixel_x = 6;
   */
  pixelX?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required float pixel_y = 7;
   */
  pixelY?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: optional float height = 8;
   */
  height?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message SSL_DetectionRobot.
 * Use `create(SSL_DetectionRobotSchema)` to create a new message.
 */
export const SSL_DetectionRobotSchema: GenMessage<SSL_DetectionRobot, SSL_DetectionRobotJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_detection, 1);

/**
 * @generated from message SSL_DetectionFrame
 */
export type SSL_DetectionFrame = Message<"SSL_DetectionFrame"> & {
  /**
   * @generated from field: required uint32 frame_number = 1;
   */
  frameNumber: number;

  /**
   * @generated from field: required double t_capture = 2;
   */
  tCapture: number;

  /**
   * @generated from field: required double t_sent = 3;
   */
  tSent: number;

  /**
   * @generated from field: required uint32 camera_id = 4;
   */
  cameraId: number;

  /**
   * @generated from field: repeated SSL_DetectionBall balls = 5;
   */
  balls: SSL_DetectionBall[];

  /**
   * @generated from field: repeated SSL_DetectionRobot robots_yellow = 6;
   */
  robotsYellow: SSL_DetectionRobot[];

  /**
   * @generated from field: repeated SSL_DetectionRobot robots_blue = 7;
   */
  robotsBlue: SSL_DetectionRobot[];
};

/**
 * @generated from message SSL_DetectionFrame
 */
export type SSL_DetectionFrameJson = {
  /**
   * @generated from field: required uint32 frame_number = 1;
   */
  frameNumber?: number;

  /**
   * @generated from field: required double t_capture = 2;
   */
  tCapture?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required double t_sent = 3;
   */
  tSent?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * @generated from field: required uint32 camera_id = 4;
   */
  cameraId?: number;

  /**
   * @generated from field: repeated SSL_DetectionBall balls = 5;
   */
  balls?: SSL_DetectionBallJson[];

  /**
   * @generated from field: repeated SSL_DetectionRobot robots_yellow = 6;
   */
  robotsYellow?: SSL_DetectionRobotJson[];

  /**
   * @generated from field: repeated SSL_DetectionRobot robots_blue = 7;
   */
  robotsBlue?: SSL_DetectionRobotJson[];
};

/**
 * Describes the message SSL_DetectionFrame.
 * Use `create(SSL_DetectionFrameSchema)` to create a new message.
 */
export const SSL_DetectionFrameSchema: GenMessage<SSL_DetectionFrame, SSL_DetectionFrameJson> = /*@__PURE__*/
  messageDesc(file_vision_ssl_vision_detection, 2);

