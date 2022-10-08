// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:convert';

class PointInfo {
  final String id;
  final String x;
  final String y;
  PointInfo({
    required this.id,
    required this.x,
    required this.y,
  });

  PointInfo copyWith({
    String? id,
    String? x,
    String? y,
  }) {
    return PointInfo(
      id: id ?? this.id,
      x: x ?? this.x,
      y: y ?? this.y,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'x': x,
      'y': y,
    };
  }

  factory PointInfo.fromMap(Map<String, dynamic> map) {
    return PointInfo(
      id: map['id'] as String,
      x: map['x'] as String,
      y: map['y'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory PointInfo.fromJson(String source) =>
      PointInfo.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() => 'PointInfo(id: $id, x: $x, y: $y)';

  @override
  bool operator ==(covariant PointInfo other) {
    if (identical(this, other)) return true;

    return other.id == id && other.x == x && other.y == y;
  }

  @override
  int get hashCode => id.hashCode ^ x.hashCode ^ y.hashCode;
}
