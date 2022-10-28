import 'dart:async';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:scribble_demo/model/server_response.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class CanvasPage extends StatefulWidget {
  const CanvasPage({super.key});

  @override
  State<CanvasPage> createState() => _CanvasPageState();
}

class _CanvasPageState extends State<CanvasPage> {
  String id = "", joined = "";

  final channel = WebSocketChannel.connect(
    Uri.parse("ws://20.219.210.103:5000/ws"),
  );

  late Stream<ServerResponse> readableStream;
  // late StreamController<ServerResponse> streamController;

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
    readableStream = channel.stream.map((event) {
      String value = event.toString();
      return ServerResponse.fromJson(value);
    }).asBroadcastStream();
    log("converted dynamic to ServerResponse");
    super.initState();
  }

  @override
  void dispose() {
    channel.sink.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.of(context).size.width;
    final height = MediaQuery.of(context).size.height;

    // print(channel.toString());

    return Scaffold(
      appBar: AppBar(
        title: const Text('Scribble Demo'),
        backgroundColor: Colors.black,
        elevation: 0,
      ),
      body: Column(
        children: [
          Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              SizedBox(
                width: width * 0.2,
                height: height * 0.8,
                child: Column(
                  children: [
                    const Text(
                      "Connected User",
                      style: TextStyle(fontSize: 30),
                    ),
                    StreamBuilder<ServerResponse>(
                      stream: readableStream,
                      builder: (context, snapshot) {
                        if (snapshot.hasData) {
                          print(snapshot.data.toString());
                          final responseData = snapshot.data;
                          if (responseData!.response_type == "iam") {
                            id = responseData.id;
                            joined = responseData.connected_clients.length.toString();
                            return Text("New user joined: $id\nTotal joined: $joined");
                          } else if (responseData.response_type == "total") {
                            joined = responseData.connected_clients.length.toString();
                            return Text(
                                "New user joined: ${responseData.id}\nTotal joined: $joined");
                          } else if (responseData.response_type == "dis") {
                            joined = responseData.connected_clients.length.toString();
                            return Text(
                                "User disconnected: ${responseData.id}\nTotal joined: $joined");
                          }
                        }
                        return Text("Total joined: $joined");
                      },
                    ),
                  ],
                ),
              ),
              Container(
                width: width * 0.6,
                height: height * 0.8,
                color: const Color.fromARGB(255, 240, 240, 145),
                child: GestureDetector(
                  onPanStart: onPanStart,
                  onPanEnd: onPanEnd,
                  onPanUpdate: onPanUpdate,
                ),
              ),
              SizedBox(
                width: width * 0.2,
                height: height * 0.8,
                child: StreamBuilder<ServerResponse>(
                  stream: readableStream,
                  builder: (context, snapshot) {
                    print(snapshot.data);
                    if (snapshot.hasError) {
                      log(snapshot.error.toString());
                      return const Text("An error occured");
                    }
                    if (snapshot.hasData) {
                      final responseData = snapshot.data;

                      if (responseData!.response_type == "set") {
                        return Text(
                            "CurrentClientID: $id\nID: ${responseData.client_info.client_id}\ndx: ${responseData.client_info.x}\ndy: ${responseData.client_info.y}");
                      }
                    }

                    return const Text("No movement");
                  },
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
