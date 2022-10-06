import 'dart:async';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class CanvasPage extends StatefulWidget {
  const CanvasPage({super.key});

  @override
  State<CanvasPage> createState() => _CanvasPageState();
}

class _CanvasPageState extends State<CanvasPage> {
  String id = "";

  final channel = WebSocketChannel.connect(
    Uri.parse("ws://localhost:5000/ws"),
  );

  late Stream<dynamic> readableStream;
  String joined = "";

  // void _onTapHandler(TapDownDetails details) {
  //   Offset tapPosition;
  //   RenderBox? referenceBox = context.findRenderObject() as RenderBox;
  //   setState(() {
  //     tapPosition = referenceBox.globalToLocal(details.globalPosition);
  //     infoBroadcastController.add("${tapPosition.dx} ${tapPosition.dy}");
  //   });
  // }

  void onPanStart(DragStartDetails details) {
    RenderBox? box = context.findRenderObject() as RenderBox;
    Offset point = box.globalToLocal(details.globalPosition);

    channel.sink.add("${point.dx} ${point.dy}");
  }

  void onPanUpdate(DragUpdateDetails details) {
    RenderBox? box = context.findRenderObject() as RenderBox;
    Offset point = box.globalToLocal(details.globalPosition);

    log(details.globalPosition.toString());

    channel.sink.add("${point.dx} ${point.dy}");
  }

  void onPanEnd(DragEndDetails details) {
    // lines = List.from(lines)..add(line);

    // linesStreamController.add(lines);
  }

  // void setClientDetails() {
  //   // channel.stream.listen((event) {
  //   //   print("Listening to connected websocket");
  //   //   final values = event.toString().split(" ");
  //   //   print("request is ${values[0]}");
  //   //   print("id is ${values[1]}");

  //   //   if (values[0] == "iam") {
  //   //     id = values[1];
  //   //   }
  //   // });

  // }

  @override
  void initState() {
    readableStream = channel.stream.asBroadcastStream();
    // setClientDetails();
    super.initState();
  }

  @override
  void dispose() {
    channel.sink.close();
    // infoBroadcastController.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Scribble Demo'),
        backgroundColor: Colors.black,
        elevation: 0,
      ),
      body: Column(
        children: [
          Container(
            width: 200,
            height: 200,
            color: Colors.red,
            child: GestureDetector(
              onPanStart: onPanStart,
              onPanEnd: onPanEnd,
              onPanUpdate: onPanUpdate,
            ),
          ),
          StreamBuilder(
            stream: readableStream,
            builder: (context, snapshot) {
              if (snapshot.hasData) {
                print(snapshot.data);
                print(id);
                final values = snapshot.data.toString().split(" ");
                if (values[0] == "iam") {
                  id = values[1];
                  joined = values[2];
                  return Text("New user joined: $id\nTotal joined: $joined");
                } else if (values[0] == "total") {
                  joined = values[1];
                  return Text("New user joined\nTotal joined: $joined");
                }

                if (values[0] == "set") {
                  return Text(
                      "CurrentId: $id\nID: ${values[1]}\ndx: ${values[2]}\ndy: ${values[3]}\nTotal joined: $joined");
                }

                if (values[0] == "dis") {
                  joined = values[2];
                  return Text("User disconnected: ${values[1]}\nTotal joined: $joined");
                }
              }

              return Text(snapshot.error.toString());
            },
          ),
        ],
      ),
    );
  }
}
