// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

import 'package:flutter/foundation.dart';

import 'package:scribble_demo/model/client_info.dart';

class RoomInfo {
  final String room_id;
  final List<ClientInfo> grp1;
  final List<ClientInfo> grp2;
  RoomInfo({
    required this.room_id,
    required this.grp1,
    required this.grp2,
  });

  RoomInfo copyWith({
    String? room_id,
    List<ClientInfo>? grp1,
    List<ClientInfo>? grp2,
  }) {
    return RoomInfo(
      room_id: room_id ?? this.room_id,
      grp1: grp1 ?? this.grp1,
      grp2: grp2 ?? this.grp2,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'room_id': room_id,
      'grp1': grp1.map((x) => x.toMap()).toList(),
      'grp2': grp2.map((x) => x.toMap()).toList(),
    };
  }

  factory RoomInfo.fromMap(Map<String, dynamic> map) {
    return RoomInfo(
      room_id: map['room_id'] as String,
      grp1: List<ClientInfo>.from(
        (map['grp1'] as List<dynamic>).map<ClientInfo>(
          (x) => ClientInfo.fromMap(x as Map<String, dynamic>),
        ),
      ),
      grp2: List<ClientInfo>.from(
        (map['grp2'] as List<dynamic>).map<ClientInfo>(
          (x) => ClientInfo.fromMap(x as Map<String, dynamic>),
        ),
      ),
    );
  }

  String toJson() => json.encode(toMap());

  factory RoomInfo.fromJson(String source) =>
      RoomInfo.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() => 'RoomInfo(room_id: $room_id, grp1: $grp1, grp2: $grp2)';

  @override
  bool operator ==(covariant RoomInfo other) {
    if (identical(this, other)) return true;

    return other.room_id == room_id &&
        listEquals(other.grp1, grp1) &&
        listEquals(other.grp2, grp2);
  }

  @override
  int get hashCode => room_id.hashCode ^ grp1.hashCode ^ grp2.hashCode;
}
