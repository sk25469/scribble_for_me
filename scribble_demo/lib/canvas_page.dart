import 'dart:async';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:scribble_demo/model/client_info.dart';
import 'package:scribble_demo/model/client_response.dart';
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
    Uri.parse("ws://localhost:5000/ws"),
  );

  late Stream<ServerResponse> readableStream;
  TextEditingController nameTextController = TextEditingController();
  TextEditingController roomIdTextController = TextEditingController();

  // late StreamController<ServerResponse> streamController;

  // response_type can be "connect-new" or "connect"
  String response_type = "", room_type = "";

  // send the name and room details to the server
  void sendNameAndRoomDetails() {
    channel.sink.add(
      ClientResponse(
        response_type: response_type,
        room_id: roomIdTextController.text,
        client_info: ClientInfo(
          client_id: id,
          name: nameTextController.text,
          room_id: "",
        ),
        room_type: room_type,
      ).toJson(),
    );
    setState(() {});
  }

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
      body: room_type.isEmpty
          ? Center(
              child: SizedBox(
                width: width * 0.5,
                height: height * 0.5,
                child: Column(
                  children: [
                    TextField(
                      controller: nameTextController,
                      decoration: const InputDecoration(hintText: 'Enter your name'),
                    ),
                    SizedBox(
                      height: height * 0.2,
                    ),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        TextButton(
                          onPressed: () {
                            response_type = "connect-new";
                            room_type = "private";
                            sendNameAndRoomDetails();
                          },
                          child: const Text('Create new private room'),
                        ),
                        TextButton(
                          onPressed: () {
                            response_type = "connect-new";
                            room_type = "public";
                            sendNameAndRoomDetails();
                          },
                          child: const Text('Join Public Room'),
                        ),
                        Row(
                          children: [
                            SizedBox(
                              width: width * 0.1,
                              child: TextField(
                                controller: roomIdTextController,
                                decoration: const InputDecoration(
                                  hintText: 'Enter room id',
                                ),
                              ),
                            ),
                            TextButton(
                              onPressed: () {
                                response_type = "connect";
                                room_type = "private";
                                sendNameAndRoomDetails();
                              },
                              child: const Text('Join a private room with id'),
                            ),
                          ],
                        ),
                      ],
                    )
                  ],
                ),
              ),
            )
          : Column(
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
                                if (responseData!.response_type == "total") {
                                  int total = responseData.room_info.grp1.length +
                                      responseData.room_info.grp2.length;
                                  joined = total.toString();
                                  return Text(
                                      "New user joined: ${responseData.client_info.client_id}\nTotal joined: $joined");
                                } else if (responseData.response_type == "dis") {
                                  int total = responseData.room_info.grp1.length +
                                      responseData.room_info.grp2.length;
                                  joined = total.toString();
                                  return Text(
                                      "User disconnected: ${responseData.client_info.client_id}\nTotal joined: $joined");
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
                                  "CurrentClientID: $id\nID: ${responseData.client_info.client_id}\n");
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
