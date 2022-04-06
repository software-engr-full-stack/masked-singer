import React, { useState, useCallback, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import humps from 'humps';

function parse(data) {
  return humps.camelizeKeys(JSON.parse(data));
}

function WebSocketDemo() {
  const socketUrl = 'ws://localhost:8082/get-votes?competition_name=gotham-2022-04-05';
  // const socketUrl = 'wss://demo.piesocket.com/v3/channel_1?api_key=oCdCMcMPQpbvNjUIzqtvF1d2X2okWpDQj4AwARJuAgtjhzKxVEjQU6IdCjwm&notify_self';

  const [messageHistory, setMessageHistory] = useState([]);

  const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl);

  useEffect(() => {
    if (lastMessage !== null) {
      setMessageHistory((prev) => prev.concat(lastMessage));
    }
  }, [lastMessage, setMessageHistory]);

  const handleClickSendMessage = useCallback(() => sendMessage('Hello'), []); // eslint-disable-line react-hooks/exhaustive-deps

  const connectionStatus = {
    [ReadyState.CONNECTING]: 'Connecting',
    [ReadyState.OPEN]: 'Open',
    [ReadyState.CLOSING]: 'Closing',
    [ReadyState.CLOSED]: 'Closed',
    [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
  }[readyState];

  return (
    <div>
      <button
        type="button"
        onClick={handleClickSendMessage}
        disabled={readyState !== ReadyState.OPEN}
      >
        Click Me to send &apos;Hello&apos;
      </button>
      <div>The WebSocket is currently {connectionStatus}</div>
      {lastMessage ? <span>Last message: {lastMessage.data}</span> : null}
      <ul>
        {
          messageHistory.map((message) => {
            const data = message ? parse(message.data) : null;

            return (
              <div
                key={Math.random()}
              >
                Competition: {data.topic}, singer: {data.key}, vote UUID: {data.value}
              </div>
            );
          })
        }
      </ul>
    </div>
  );
}

function App() {
  return <WebSocketDemo />;
}

export default App;
