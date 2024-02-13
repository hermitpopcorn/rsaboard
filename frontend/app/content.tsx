"use client";

import { useState } from "react";
import { postMessage, readMessage } from "./actions";
import MenuTab from "@components/menutab";
import PostForm from "@components/post-form";
import ReadForm from "@components/read-form";
import CodeAndKeyDisplay from "@components/code-key-display";
import MessageDisplay from "@components/message-display";

const options = [
  { id: "post", text: "Post new message" },
  { id: "read", text: "Read existing message" },
];

export default function PageContent({ backendUrl }: { backendUrl: string }) {
  const [activeTabId, setActiveTabId] = useState(options[0].id);
  const activeTabIdSwitcher = (to: string): void => {
    setActiveTabId(to);
  };

  const [postedCode, setPostedCode] = useState("");
  const [postedPrivateKey, setPostedPrivateKey] = useState("");
  const [message, setMessage] = useState("");

  const [isWorking, setIsWorking] = useState(false);

  const postSubmitHandler = (params): void => {
    if (isWorking) {
      return;
    }
    setIsWorking(true);

    postMessage(backendUrl, params)
      .then((res) => {
        displayCodeAndPrivateKey(res.code, res.privateKey);
      })
      .catch((err) => alert((err as Error).message))
      .finally(() => setIsWorking(false));
  };

  const displayCodeAndPrivateKey = (code: string, privateKey: string) => {
    setPostedCode(code);
    setPostedPrivateKey(privateKey);
    activeTabIdSwitcher("posted");
  };

  const readSubmitHandler = (params): void => {
    if (isWorking) {
      return;
    }
    setIsWorking(true);

    readMessage(backendUrl, params)
      .then((res) => {
        console.log(res);
        displayMessage(res.message);
      })
      .catch((err) => alert((err as Error).message))
      .finally(() => setIsWorking(false));
  };

  const displayMessage = (message: string) => {
    setMessage(message);
    activeTabIdSwitcher("message");
  };

  return (
    <div className="flex flex-col gap-2">
      <MenuTab
        options={options}
        activeId={activeTabId}
        activeIdSwitcher={activeTabIdSwitcher}
      />
      {activeTabId == "post" ? (
        <PostForm submitHandler={postSubmitHandler} formLock={isWorking} />
      ) : (
        <></>
      )}
      {activeTabId == "posted" ? (
        <CodeAndKeyDisplay code={postedCode} privateKey={postedPrivateKey} />
      ) : (
        <></>
      )}
      {activeTabId == "read" ? (
        <ReadForm submitHandler={readSubmitHandler} formLock={isWorking} />
      ) : (
        <></>
      )}
      {activeTabId == "message" ? <MessageDisplay message={message} /> : <></>}
    </div>
  );
}
