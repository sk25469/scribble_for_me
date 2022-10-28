import 'dart:convert';

// ignore_for_file: public_member_api_docs, sort_constructors_first

class ClientInfo {
  final String client_id;
  final String room_id;
  final String x;
  final String y;
  final String name;
  ClientInfo({
    required this.client_id,
    required this.room_id,
    required this.x,
    required this.y,
    required this.name,
  });

  ClientInfo copyWith({
    String? client_id,
    String? room_id,
    String? x,
    String? y,
    String? name,
  }) {
    return ClientInfo(
      client_id: client_id ?? this.client_id,
      room_id: room_id ?? this.room_id,
      x: x ?? this.x,
      y: y ?? this.y,
      name: name ?? this.name,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'client_id': client_id,
      'room_id': room_id,
      'x': x,
      'y': y,
      'name': name,
    };
  }

  factory ClientInfo.fromMap(Map<String, dynamic> map) {
    return ClientInfo(
      client_id: map['client_id'] as String,
      room_id: map['room_id'] as String,
      x: map['x'] as String,
      y: map['y'] as String,
      name: map['name'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory ClientInfo.fromJson(String source) =>
      ClientInfo.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() {
    return 'PointInfo(client_id: $client_id, room_id: $room_id, x: $x, y: $y, name: $name)';
  }

  @override
  bool operator ==(covariant ClientInfo other) {
    if (identical(this, other)) return true;

    return other.client_id == client_id &&
        other.room_id == room_id &&
        other.x == x &&
        other.y == y &&
        other.name == name;
  }

  @override
  int get hashCode {
    return client_id.hashCode ^
        room_id.hashCode ^
        x.hashCode ^
        y.hashCode ^
        name.hashCode;
  }
}
