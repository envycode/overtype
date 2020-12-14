import { get } from 'svelte/store';
import {
  username,
  code,
  participantId,
  action,
  currentWordCount,
  wsStore,
  WS_ACTIVE
} from '@/store/index.js';
import { querySerializer } from '@/util/query.js';

export let wsURL;

let wsConnection;
let wsTries = 5;
let timeout = 1000;


wsURL = ENV.WEB_SOCKET_BASE_URL;

export function initWS() {
  if (wsTries <= 0) {
    console.error('unable to estabilish WS after 5 tries!');
    wsConnection = null;
    wsTries = 5;
    WS_ACTIVE.set(false);
    return;
  }
  //Don't open a new websocket if it already exists. Figure out a better way for event filtering #FIXME
  if (wsConnection) {
    return;
  }
  const params = {
    code: get(code),
    participant_id: get(participantId),
    participant_name: get(username)
  };

  wsConnection = new WebSocket(`${wsURL}/api/join-room${querySerializer(params)}`);
  wsConnection.onopen = function () {
    wsConnection.send(
      JSON.stringify(
        Object.assign(params, {
          action: get(action),
          current_word_count: get(currentWordCount)
        })
      )
    );
    WS_ACTIVE.set(true);
    setTimeout(heartbeat, 1000);
  };

  // Log errors
  wsConnection.onerror = function (error) {
    wsTries--;
    console.error('WebSocket Error ', error);
  };

  // Log messages from the server
  wsConnection.onmessage = function (event) {
    try {
      var data = JSON.parse(event.data);
      if (data.room_state === 3) {
        console.error('Finished');
        closeWS();
      }
      wsStore.set(data);
    }
    catch (e) {
      //console.error('got non json data', event.data, e);
    }
  };
  wsConnection.onclose = function (e) {
    if (e.code != 1000) {
      closeWS();
    } else {
      setTimeout(function () {
        initWS();
      }, timeout);
    }
  };
}

export function closeWS() {
  if (wsConnection) {
    wsConnection.onclose = function () {
      wsConnection = null;
    };
    wsConnection.close();
    WS_ACTIVE.set(false);
  }
}

function heartbeat() {
  if (!wsConnection || wsConnection.readyState !== 1) {
    return;
  }
  const params = {
    code: get(code),
    participant_id: get(participantId),
    participant_name: get(username),
    action: get(action),
    current_word_count: get(currentWordCount)
  }

  wsConnection.send(JSON.stringify(params));
  setTimeout(heartbeat, 1000);
}
