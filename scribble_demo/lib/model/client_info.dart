import 'dart:convert';

// ignore_for_file: public_member_api_docs, sort_constructors_first

class ClientInfo {
  final String client_id;
  final String room_id;
  final String name;
  ClientInfo({
    required this.client_id,
    required this.room_id,
    required this.name,
  });

  ClientInfo copyWith({
    String? client_id,
    String? room_id,
    String? name,
  }) {
    return ClientInfo(
      client_id: client_id ?? this.client_id,
      room_id: room_id ?? this.room_id,
      name: name ?? this.name,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'client_id': client_id,
      'room_id': room_id,
      'name': name,
    };
  }

  factory ClientInfo.fromMap(Map<String, dynamic> map) {
    return ClientInfo(
      client_id: map['client_id'] as String,
      room_id: map['room_id'] as String,
      name: map['name'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory ClientInfo.fromJson(String source) =>
      ClientInfo.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() =>
      'ClientInfo(client_id: $client_id, room_id: $room_id, name: $name)';

  @override
  bool operator ==(covariant ClientInfo other) {
    if (identical(this, other)) return true;

    return other.client_id == client_id && other.room_id == room_id && other.name == name;
  }

  @override
  int get hashCode => client_id.hashCode ^ room_id.hashCode ^ name.hashCode;
}
