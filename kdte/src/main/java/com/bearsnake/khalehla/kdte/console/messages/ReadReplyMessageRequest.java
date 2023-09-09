// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: exec/msg/console.proto

package com.bearsnake.khalehla.kdte.console.messages;

/**
 * Protobuf type {@code msg.ReadReplyMessageRequest}
 */
public final class ReadReplyMessageRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:msg.ReadReplyMessageRequest)
    ReadReplyMessageRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ReadReplyMessageRequest.newBuilder() to construct.
  private ReadReplyMessageRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ReadReplyMessageRequest() {
    sender_ = "";
    text_ =
        com.google.protobuf.LazyStringArrayList.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new ReadReplyMessageRequest();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.bearsnake.khalehla.kdte.console.messages.ConsoleService.internal_static_msg_ReadReplyMessageRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.bearsnake.khalehla.kdte.console.messages.ConsoleService.internal_static_msg_ReadReplyMessageRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.class, com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.Builder.class);
  }

  public static final int MESSAGEID_FIELD_NUMBER = 2;
  private int messageId_ = 0;
  /**
   * <code>int32 messageId = 2;</code>
   * @return The messageId.
   */
  @java.lang.Override
  public int getMessageId() {
    return messageId_;
  }

  public static final int SENDER_FIELD_NUMBER = 3;
  @SuppressWarnings("serial")
  private volatile java.lang.Object sender_ = "";
  /**
   * <code>string sender = 3;</code>
   * @return The sender.
   */
  @java.lang.Override
  public java.lang.String getSender() {
    java.lang.Object ref = sender_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      sender_ = s;
      return s;
    }
  }
  /**
   * <code>string sender = 3;</code>
   * @return The bytes for sender.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getSenderBytes() {
    java.lang.Object ref = sender_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      sender_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int TEXT_FIELD_NUMBER = 4;
  @SuppressWarnings("serial")
  private com.google.protobuf.LazyStringArrayList text_ =
      com.google.protobuf.LazyStringArrayList.emptyList();
  /**
   * <code>repeated string text = 4;</code>
   * @return A list containing the text.
   */
  public com.google.protobuf.ProtocolStringList
      getTextList() {
    return text_;
  }
  /**
   * <code>repeated string text = 4;</code>
   * @return The count of text.
   */
  public int getTextCount() {
    return text_.size();
  }
  /**
   * <code>repeated string text = 4;</code>
   * @param index The index of the element to return.
   * @return The text at the given index.
   */
  public java.lang.String getText(int index) {
    return text_.get(index);
  }
  /**
   * <code>repeated string text = 4;</code>
   * @param index The index of the value to return.
   * @return The bytes of the text at the given index.
   */
  public com.google.protobuf.ByteString
      getTextBytes(int index) {
    return text_.getByteString(index);
  }

  public static final int MAXREPLYCHARACTERS_FIELD_NUMBER = 5;
  private int maxReplyCharacters_ = 0;
  /**
   * <code>int32 maxReplyCharacters = 5;</code>
   * @return The maxReplyCharacters.
   */
  @java.lang.Override
  public int getMaxReplyCharacters() {
    return maxReplyCharacters_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (messageId_ != 0) {
      output.writeInt32(2, messageId_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(sender_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, sender_);
    }
    for (int i = 0; i < text_.size(); i++) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, text_.getRaw(i));
    }
    if (maxReplyCharacters_ != 0) {
      output.writeInt32(5, maxReplyCharacters_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (messageId_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, messageId_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(sender_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, sender_);
    }
    {
      int dataSize = 0;
      for (int i = 0; i < text_.size(); i++) {
        dataSize += computeStringSizeNoTag(text_.getRaw(i));
      }
      size += dataSize;
      size += 1 * getTextList().size();
    }
    if (maxReplyCharacters_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(5, maxReplyCharacters_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest)) {
      return super.equals(obj);
    }
    com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest other = (com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest) obj;

    if (getMessageId()
        != other.getMessageId()) return false;
    if (!getSender()
        .equals(other.getSender())) return false;
    if (!getTextList()
        .equals(other.getTextList())) return false;
    if (getMaxReplyCharacters()
        != other.getMaxReplyCharacters()) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + MESSAGEID_FIELD_NUMBER;
    hash = (53 * hash) + getMessageId();
    hash = (37 * hash) + SENDER_FIELD_NUMBER;
    hash = (53 * hash) + getSender().hashCode();
    if (getTextCount() > 0) {
      hash = (37 * hash) + TEXT_FIELD_NUMBER;
      hash = (53 * hash) + getTextList().hashCode();
    }
    hash = (37 * hash) + MAXREPLYCHARACTERS_FIELD_NUMBER;
    hash = (53 * hash) + getMaxReplyCharacters();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code msg.ReadReplyMessageRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:msg.ReadReplyMessageRequest)
      com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.bearsnake.khalehla.kdte.console.messages.ConsoleService.internal_static_msg_ReadReplyMessageRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.bearsnake.khalehla.kdte.console.messages.ConsoleService.internal_static_msg_ReadReplyMessageRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.class, com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.Builder.class);
    }

    // Construct using com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      messageId_ = 0;
      sender_ = "";
      text_ =
          com.google.protobuf.LazyStringArrayList.emptyList();
      maxReplyCharacters_ = 0;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.bearsnake.khalehla.kdte.console.messages.ConsoleService.internal_static_msg_ReadReplyMessageRequest_descriptor;
    }

    @java.lang.Override
    public com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest getDefaultInstanceForType() {
      return com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.getDefaultInstance();
    }

    @java.lang.Override
    public com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest build() {
      com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest buildPartial() {
      com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest result = new com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.messageId_ = messageId_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.sender_ = sender_;
      }
      if (((from_bitField0_ & 0x00000004) != 0)) {
        text_.makeImmutable();
        result.text_ = text_;
      }
      if (((from_bitField0_ & 0x00000008) != 0)) {
        result.maxReplyCharacters_ = maxReplyCharacters_;
      }
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest) {
        return mergeFrom((com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest other) {
      if (other == com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest.getDefaultInstance()) return this;
      if (other.getMessageId() != 0) {
        setMessageId(other.getMessageId());
      }
      if (!other.getSender().isEmpty()) {
        sender_ = other.sender_;
        bitField0_ |= 0x00000002;
        onChanged();
      }
      if (!other.text_.isEmpty()) {
        if (text_.isEmpty()) {
          text_ = other.text_;
          bitField0_ |= 0x00000004;
        } else {
          ensureTextIsMutable();
          text_.addAll(other.text_);
        }
        onChanged();
      }
      if (other.getMaxReplyCharacters() != 0) {
        setMaxReplyCharacters(other.getMaxReplyCharacters());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 16: {
              messageId_ = input.readInt32();
              bitField0_ |= 0x00000001;
              break;
            } // case 16
            case 26: {
              sender_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000002;
              break;
            } // case 26
            case 34: {
              java.lang.String s = input.readStringRequireUtf8();
              ensureTextIsMutable();
              text_.add(s);
              break;
            } // case 34
            case 40: {
              maxReplyCharacters_ = input.readInt32();
              bitField0_ |= 0x00000008;
              break;
            } // case 40
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private int messageId_ ;
    /**
     * <code>int32 messageId = 2;</code>
     * @return The messageId.
     */
    @java.lang.Override
    public int getMessageId() {
      return messageId_;
    }
    /**
     * <code>int32 messageId = 2;</code>
     * @param value The messageId to set.
     * @return This builder for chaining.
     */
    public Builder setMessageId(int value) {

      messageId_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>int32 messageId = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearMessageId() {
      bitField0_ = (bitField0_ & ~0x00000001);
      messageId_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object sender_ = "";
    /**
     * <code>string sender = 3;</code>
     * @return The sender.
     */
    public java.lang.String getSender() {
      java.lang.Object ref = sender_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        sender_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string sender = 3;</code>
     * @return The bytes for sender.
     */
    public com.google.protobuf.ByteString
        getSenderBytes() {
      java.lang.Object ref = sender_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        sender_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string sender = 3;</code>
     * @param value The sender to set.
     * @return This builder for chaining.
     */
    public Builder setSender(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      sender_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>string sender = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearSender() {
      sender_ = getDefaultInstance().getSender();
      bitField0_ = (bitField0_ & ~0x00000002);
      onChanged();
      return this;
    }
    /**
     * <code>string sender = 3;</code>
     * @param value The bytes for sender to set.
     * @return This builder for chaining.
     */
    public Builder setSenderBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      sender_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }

    private com.google.protobuf.LazyStringArrayList text_ =
        com.google.protobuf.LazyStringArrayList.emptyList();
    private void ensureTextIsMutable() {
      if (!text_.isModifiable()) {
        text_ = new com.google.protobuf.LazyStringArrayList(text_);
      }
      bitField0_ |= 0x00000004;
    }
    /**
     * <code>repeated string text = 4;</code>
     * @return A list containing the text.
     */
    public com.google.protobuf.ProtocolStringList
        getTextList() {
      text_.makeImmutable();
      return text_;
    }
    /**
     * <code>repeated string text = 4;</code>
     * @return The count of text.
     */
    public int getTextCount() {
      return text_.size();
    }
    /**
     * <code>repeated string text = 4;</code>
     * @param index The index of the element to return.
     * @return The text at the given index.
     */
    public java.lang.String getText(int index) {
      return text_.get(index);
    }
    /**
     * <code>repeated string text = 4;</code>
     * @param index The index of the value to return.
     * @return The bytes of the text at the given index.
     */
    public com.google.protobuf.ByteString
        getTextBytes(int index) {
      return text_.getByteString(index);
    }
    /**
     * <code>repeated string text = 4;</code>
     * @param index The index to set the value at.
     * @param value The text to set.
     * @return This builder for chaining.
     */
    public Builder setText(
        int index, java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      ensureTextIsMutable();
      text_.set(index, value);
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>repeated string text = 4;</code>
     * @param value The text to add.
     * @return This builder for chaining.
     */
    public Builder addText(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      ensureTextIsMutable();
      text_.add(value);
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>repeated string text = 4;</code>
     * @param values The text to add.
     * @return This builder for chaining.
     */
    public Builder addAllText(
        java.lang.Iterable<java.lang.String> values) {
      ensureTextIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, text_);
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>repeated string text = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearText() {
      text_ =
        com.google.protobuf.LazyStringArrayList.emptyList();
      bitField0_ = (bitField0_ & ~0x00000004);;
      onChanged();
      return this;
    }
    /**
     * <code>repeated string text = 4;</code>
     * @param value The bytes of the text to add.
     * @return This builder for chaining.
     */
    public Builder addTextBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      ensureTextIsMutable();
      text_.add(value);
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }

    private int maxReplyCharacters_ ;
    /**
     * <code>int32 maxReplyCharacters = 5;</code>
     * @return The maxReplyCharacters.
     */
    @java.lang.Override
    public int getMaxReplyCharacters() {
      return maxReplyCharacters_;
    }
    /**
     * <code>int32 maxReplyCharacters = 5;</code>
     * @param value The maxReplyCharacters to set.
     * @return This builder for chaining.
     */
    public Builder setMaxReplyCharacters(int value) {

      maxReplyCharacters_ = value;
      bitField0_ |= 0x00000008;
      onChanged();
      return this;
    }
    /**
     * <code>int32 maxReplyCharacters = 5;</code>
     * @return This builder for chaining.
     */
    public Builder clearMaxReplyCharacters() {
      bitField0_ = (bitField0_ & ~0x00000008);
      maxReplyCharacters_ = 0;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:msg.ReadReplyMessageRequest)
  }

  // @@protoc_insertion_point(class_scope:msg.ReadReplyMessageRequest)
  private static final com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest();
  }

  public static com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ReadReplyMessageRequest>
      PARSER = new com.google.protobuf.AbstractParser<ReadReplyMessageRequest>() {
    @java.lang.Override
    public ReadReplyMessageRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<ReadReplyMessageRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ReadReplyMessageRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.bearsnake.khalehla.kdte.console.messages.ReadReplyMessageRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

