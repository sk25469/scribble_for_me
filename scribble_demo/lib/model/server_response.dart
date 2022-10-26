// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

import 'package:flutter/foundation.dart';

import 'package:scribble_demo/model/point_info.dart';

class ServerResponse {
  // ignore: non_constant_identifier_names
  final String response_type;
  final String id;
  final List<String> connected_clients;
  final PointInfo point_info;
  ServerResponse({
    required this.response_type,
    required this.id,
    required this.connected_clients,
    required this.point_info,
  });

  ServerResponse copyWith({
    String? response_type,
    String? id,
    List<String>? connected_clients,
    PointInfo? point_info,
  }) {
    return ServerResponse(
      response_type: response_type ?? this.response_type,
      id: id ?? this.id,
      connected_clients: connected_clients ?? this.connected_clients,
      point_info: point_info ?? this.point_info,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'response_type': response_type,
      'id': id,
      'connected_clients': connected_clients,
      'point_info': point_info.toMap(),
    };
  }

  factory ServerResponse.fromMap(Map<String, dynamic> map) {
    return ServerResponse(
      response_type: map['response_type'] as String,
      id: map['id'] as String,
      connected_clients: List<String>.from(map['connected_clients']),
      point_info: PointInfo.fromMap(map['point_info'] as Map<String, dynamic>),
    );
  }

  String toJson() => json.encode(toMap());

  factory ServerResponse.fromJson(String source) =>
      ServerResponse.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'ServerResponse(response_type: $response_type, id: $id, connected_clients: $connected_clients, point_info: $point_info)';
  }

  @override
  bool operator ==(covariant ServerResponse other) {
    if (identical(this, other)) return true;

    return other.response_type == response_type &&
        other.id == id &&
        listEquals(other.connected_clients, connected_clients) &&
        other.point_info == point_info;
  }

  @override
  int get hashCode {
    return response_type.hashCode ^
        id.hashCode ^
        connected_clients.hashCode ^
        point_info.hashCode;
  }
}
