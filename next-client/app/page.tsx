"use client";

import React, { useEffect, useRef, useState } from "react";
import SelectProcess from "@/components/selectProcess";

export default function Home() {
  const [video, setVideo] = useState<File | null>(null);
  const socket = useRef<WebSocket | null>(null);
  const arr = useRef<Array<number> | null>(null);
  const [uploaded, setUploaded] = useState(false);

  const changeVideo = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setVideo(e.target.files[0]);
    }
  };

  type ChunkResponse = {
    ChunkId: number;
    Uploaded: boolean;
  };

  const uploadVideo = () => {
    setUploaded(false);

    if (video != null && socket != null) {
      const reader = new FileReader();
      reader.onload = () => {
        const arrayBuffer = reader.result as ArrayBuffer;
        const unit8Array = new Uint8Array(arrayBuffer);
        arr.current = Array(Math.ceil(unit8Array.length / (1024 * 1024))).fill(0);

        let start = 0;
        let end = 0;
        let current = 0;
        while (end < unit8Array.length) {
          start = end;
          end = end + 1024 * 1024 < unit8Array.length ? end + 1024 * 1024 : unit8Array.length;
          const slicedArray = unit8Array.slice(start, end);
          const metadata = {
            id: current,
            name: video.name,
            videoType: video.type,
            totalChunks: arr.current.length,
          };
          current++;

          const encoder = new TextEncoder();
          const metaDataBytes = encoder.encode(JSON.stringify(metadata));

          const metaDataBytesLength = new Uint8Array(4);
          new DataView(metaDataBytesLength.buffer).setUint32(0, metaDataBytes.length, true);

          const combinedBuffer = new Uint8Array(
            metaDataBytesLength.byteLength + metaDataBytes.byteLength + slicedArray.byteLength
          );

          combinedBuffer.set(metaDataBytesLength, 0);
          combinedBuffer.set(metaDataBytes, 4);
          combinedBuffer.set(slicedArray, 4 + metaDataBytes.byteLength);

          socket.current?.send(combinedBuffer.buffer);
        }
      };

      reader.onerror = (err) => {
        console.log("Error occurred during reading: ", err);
      };

      reader.readAsArrayBuffer(video);
    }
  };

  const readRecievingData = (data: Blob) => {
    const reader = new FileReader();
    reader.onload = () => {
      const text = reader.result;
      const response: ChunkResponse = JSON.parse(text as string);

      if (response.Uploaded && arr.current != null) {
        arr.current[response.ChunkId] = 1;

        const newArr = arr.current.filter((e) => e == 1);
        if (newArr.length == arr.current.length) {
          setUploaded(true);
        }
      }
      console.log(response);
    };

    reader.readAsText(data);
  };

  useEffect(() => {
    const webSocket = new WebSocket("ws://localhost:8080/api/v1/ws");
    socket.current = webSocket;

    const handleOpen = () => {
      console.log("websocket connection established");
      socket.current?.send("Hello Server!");
    };

    const handleMessage = (e: MessageEvent) => {
      readRecievingData(e.data);
    };

    const handleClose = (e: CloseEvent) => {
      console.log("webocket closed. Reason: ", e.code, e.reason);
    };

    const handleError = (e: Event) => {
      console.log("websocket error: ", e);
    };

    webSocket.addEventListener("open", handleOpen);
    webSocket.addEventListener("message", handleMessage);
    webSocket.addEventListener("close", handleClose);
    webSocket.addEventListener("error", handleError);

    return () => {
      webSocket.removeEventListener("open", handleOpen);
      webSocket.removeEventListener("message", handleMessage);
      webSocket.removeEventListener("close", handleClose);
      webSocket.removeEventListener("error", handleError);
      webSocket.close();
    };
  }, []);

  return (
    <>
      <input type="file" accept="video/*" onChange={changeVideo} />
      <button onClick={uploadVideo}>click</button>
      {uploaded && (
        <div>
          <div>Uplaoded</div>
          <SelectProcess />
        </div>
      )}
    </>
  );
}
