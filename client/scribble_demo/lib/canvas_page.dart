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
  String id = "", joined = "";

  final channel = WebSocketChannel.connect(
    Uri.parse("ws://localhost:5000/ws"),
  );

  late Stream<dynamic> readableStream;

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
                    // Expanded(
                    //   child: ListView.builder(
                    //     itemBuilder: (context, index) => Text("User $index"),
                    //     itemCount: 3,
                    //   ),
                    // ),
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
                            return Text(
                                "New user joined: ${values[2]}\nTotal joined: $joined");
                          } else if (values[0] == "dis") {
                            joined = values[2];
                            return Text(
                                "User disconnected: ${values[1]}\nTotal joined: $joined");
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
                child: StreamBuilder(
                  stream: readableStream,
                  builder: (context, snapshot) {
                    if (snapshot.hasData) {
                      print(snapshot.data);
                      print(id);
                      final values = snapshot.data.toString().split(" ");
                      // if (values[0] == "iam") {
                      //   id = values[1];
                      //   joined = values[2];
                      //   return Text("New user joined: $id\nTotal joined: $joined");
                      // } else if (values[0] == "total") {
                      //   joined = values[1];
                      //   return Text(
                      //       "New user joined: ${values[2]}\nTotal joined: $joined");
                      // }

                      if (values[0] == "set") {
                        return Text(
                            "CurrentId: $id\nID: ${values[1]}\ndx: ${values[2]}\ndy: ${values[3]}");
                      }

                      // if (values[0] == "dis") {
                      //   joined = values[2];
                      //   return Text(
                      //       "User disconnected: ${values[1]}\nTotal joined: $joined");
                      // }
                    }

                    return Text(snapshot.error.toString());
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
