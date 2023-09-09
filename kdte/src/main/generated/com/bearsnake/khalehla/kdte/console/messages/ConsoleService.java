// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: exec/msg/console.proto

package com.bearsnake.khalehla.kdte.console.messages;

public final class ConsoleService {
  private ConsoleService() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_Empty_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_Empty_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_ClearReadReplyMessageRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_ClearReadReplyMessageRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_SolicitedInputMessage_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_SolicitedInputMessage_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_UnsolicitedInputMessage_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_UnsolicitedInputMessage_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_PollInputMessageResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_PollInputMessageResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_ReadOnlyMessageRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_ReadOnlyMessageRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_ReadReplyMessageRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_ReadReplyMessageRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_msg_StatusMessageRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_msg_StatusMessageRequest_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\026exec/msg/console.proto\022\003msg\"\007\n\005Empty\"1" +
      "\n\034ClearReadReplyMessageRequest\022\021\n\tmessag" +
      "eId\030\001 \001(\r\"8\n\025SolicitedInputMessage\022\021\n\tme" +
      "ssageId\030\001 \001(\r\022\014\n\004text\030\002 \001(\t\"\'\n\027Unsolicit" +
      "edInputMessage\022\014\n\004text\030\001 \001(\t\"\262\001\n\030PollInp" +
      "utMessageResponse\022\020\n\010hasInput\030\001 \001(\010\0226\n\020s" +
      "olicitedMessage\030\002 \001(\0132\032.msg.SolicitedInp" +
      "utMessageH\000\022:\n\022unsolicitedMessage\030\003 \001(\0132" +
      "\034.msg.UnsolicitedInputMessageH\000B\020\n\016input" +
      "Specifier\"6\n\026ReadOnlyMessageRequest\022\016\n\006s" +
      "ender\030\003 \001(\t\022\014\n\004text\030\004 \003(\t\"f\n\027ReadReplyMe" +
      "ssageRequest\022\021\n\tmessageId\030\002 \001(\005\022\016\n\006sende" +
      "r\030\003 \001(\t\022\014\n\004text\030\004 \003(\t\022\032\n\022maxReplyCharact" +
      "ers\030\005 \001(\005\"$\n\024StatusMessageRequest\022\014\n\004tex" +
      "t\030\004 \003(\t2\373\002\n\007Console\022H\n\025ClearReadReplyMes" +
      "sage\022!.msg.ClearReadReplyMessageRequest\032" +
      "\n.msg.Empty\"\000\022?\n\020PollInputMessage\022\n.msg." +
      "Empty\032\035.msg.PollInputMessageResponse\"\000\022!" +
      "\n\005Reset\022\n.msg.Empty\032\n.msg.Empty\"\000\022@\n\023Sen" +
      "dReadOnlyMessage\022\033.msg.ReadOnlyMessageRe" +
      "quest\032\n.msg.Empty\"\000\022B\n\024SendReadReplyMess" +
      "age\022\034.msg.ReadReplyMessageRequest\032\n.msg." +
      "Empty\"\000\022<\n\021SendStatusMessage\022\031.msg.Statu" +
      "sMessageRequest\032\n.msg.Empty\"\000BJ\n,com.bea" +
      "rsnake.khalehla.kdte.console.messagesB\016C" +
      "onsoleServiceP\001Z\010exec/msgb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_msg_Empty_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_msg_Empty_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_Empty_descriptor,
        new java.lang.String[] { });
    internal_static_msg_ClearReadReplyMessageRequest_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_msg_ClearReadReplyMessageRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_ClearReadReplyMessageRequest_descriptor,
        new java.lang.String[] { "MessageId", });
    internal_static_msg_SolicitedInputMessage_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_msg_SolicitedInputMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_SolicitedInputMessage_descriptor,
        new java.lang.String[] { "MessageId", "Text", });
    internal_static_msg_UnsolicitedInputMessage_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_msg_UnsolicitedInputMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_UnsolicitedInputMessage_descriptor,
        new java.lang.String[] { "Text", });
    internal_static_msg_PollInputMessageResponse_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_msg_PollInputMessageResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_PollInputMessageResponse_descriptor,
        new java.lang.String[] { "HasInput", "SolicitedMessage", "UnsolicitedMessage", "InputSpecifier", });
    internal_static_msg_ReadOnlyMessageRequest_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_msg_ReadOnlyMessageRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_ReadOnlyMessageRequest_descriptor,
        new java.lang.String[] { "Sender", "Text", });
    internal_static_msg_ReadReplyMessageRequest_descriptor =
      getDescriptor().getMessageTypes().get(6);
    internal_static_msg_ReadReplyMessageRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_ReadReplyMessageRequest_descriptor,
        new java.lang.String[] { "MessageId", "Sender", "Text", "MaxReplyCharacters", });
    internal_static_msg_StatusMessageRequest_descriptor =
      getDescriptor().getMessageTypes().get(7);
    internal_static_msg_StatusMessageRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_msg_StatusMessageRequest_descriptor,
        new java.lang.String[] { "Text", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
