// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

import 'package:scribble_demo/model/client_info.dart';

class ClientResponse {
  final String response_type;
  final String room_id;
  final ClientInfo client_info;
  final String room_type;
  ClientResponse({
    required this.response_type,
    required this.room_id,
    required this.client_info,
    required this.room_type,
  });

  ClientResponse copyWith({
    String? response_type,
    String? room_id,
    ClientInfo? client_info,
    String? room_type,
  }) {
    return ClientResponse(
      response_type: response_type ?? this.response_type,
      room_id: room_id ?? this.room_id,
      client_info: client_info ?? this.client_info,
      room_type: room_type ?? this.room_type,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'response_type': response_type,
      'room_id': room_id,
      'client_info': client_info.toMap(),
      'room_type': room_type,
    };
  }

  factory ClientResponse.fromMap(Map<String, dynamic> map) {
    return ClientResponse(
      response_type: map['response_type'] as String,
      room_id: map['room_id'] as String,
      client_info: ClientInfo.fromMap(map['client_info'] as Map<String, dynamic>),
      room_type: map['room_type'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory ClientResponse.fromJson(String source) =>
      ClientResponse.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'ClientResponse(response_type: $response_type, room_id: $room_id, client_info: $client_info, room_type: $room_type)';
  }

  @override
  bool operator ==(covariant ClientResponse other) {
    if (identical(this, other)) return true;

    return other.response_type == response_type &&
        other.room_id == room_id &&
        other.client_info == client_info &&
        other.room_type == room_type;
  }

  @override
  int get hashCode {
    return response_type.hashCode ^
        room_id.hashCode ^
        client_info.hashCode ^
        room_type.hashCode;
  }
}
