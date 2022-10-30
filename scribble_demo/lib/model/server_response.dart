// ignore_for_file: public_member_api_docs, sort_constructors_first

import 'dart:convert';

import 'package:scribble_demo/model/room_info.dart';

import 'client_info.dart';

class ServerResponse {
  // ignore: non_constant_identifier_names
  final String response_type;
  final RoomInfo room_info;
  final ClientInfo client_info;
  ServerResponse({
    required this.response_type,
    required this.room_info,
    required this.client_info,
  });

  ServerResponse copyWith({
    String? response_type,
    RoomInfo? room_info,
    ClientInfo? client_info,
  }) {
    return ServerResponse(
      response_type: response_type ?? this.response_type,
      room_info: room_info ?? this.room_info,
      client_info: client_info ?? this.client_info,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'response_type': response_type,
      'room_info': room_info.toMap(),
      'client_info': client_info.toMap(),
    };
  }

  factory ServerResponse.fromMap(Map<String, dynamic> map) {
    return ServerResponse(
      response_type: map['response_type'] as String,
      room_info: RoomInfo.fromMap(map['room_info'] as Map<String, dynamic>),
      client_info: ClientInfo.fromMap(map['client_info'] as Map<String, dynamic>),
    );
  }

  String toJson() => json.encode(toMap());

  factory ServerResponse.fromJson(String source) =>
      ServerResponse.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() =>
      'ServerResponse(response_type: $response_type, room_info: $room_info, client_info: $client_info)';

  @override
  bool operator ==(covariant ServerResponse other) {
    if (identical(this, other)) return true;

    return other.response_type == response_type &&
        other.room_info == room_info &&
        other.client_info == client_info;
  }

  @override
  int get hashCode => response_type.hashCode ^ room_info.hashCode ^ client_info.hashCode;
}
