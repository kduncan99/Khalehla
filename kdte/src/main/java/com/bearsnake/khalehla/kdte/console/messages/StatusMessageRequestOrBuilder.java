// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: exec/msg/console.proto

package com.bearsnake.khalehla.kdte.console.messages;

public interface StatusMessageRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:msg.StatusMessageRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated string text = 4;</code>
   * @return A list containing the text.
   */
  java.util.List<java.lang.String>
      getTextList();
  /**
   * <code>repeated string text = 4;</code>
   * @return The count of text.
   */
  int getTextCount();
  /**
   * <code>repeated string text = 4;</code>
   * @param index The index of the element to return.
   * @return The text at the given index.
   */
  java.lang.String getText(int index);
  /**
   * <code>repeated string text = 4;</code>
   * @param index The index of the value to return.
   * @return The bytes of the text at the given index.
   */
  com.google.protobuf.ByteString
      getTextBytes(int index);
}
