import React, { useState, useEffect, useRef } from 'react';
import Input from './Input';
import Textarea from './Textarea';

const Bbs = () => {
    const [ comment, setComment ] = useState("");
    const [ comments, setComments ] = useState(new Array());
  
    const socket = useRef<WebSocket | null>(null);

    useEffect(() => {
      socket.current = new WebSocket("ws://localhost:8765");
      socket.current.onopen = () => {
        if (socket.current) {
          socket.current.send(JSON.stringify({ name: "listen comments", data: "" }));
        }
      }
      socket.current.onmessage = function (e) {
        const msg = JSON.parse(e.data);
        switch (msg.name) {
          case 'add comment':
            console.log(msg.ok);
            break;
          case 'listen comments':
            if (msg.data == null) {
              return;
            }
            setComments(msg.data);
            break;
        }
      }
    }, []);

    useEffect(() => {
      return () => {
        if (socket.current) {
          socket.current.close();
        }
      }
    }, [socket]);

    const onChangeHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
      setComment(e.target.value);
    };
    
    const submitHandler = (e: React.MouseEvent) => {
      e.preventDefault();
      if (socket.current) {
        socket.current.send(JSON.stringify({ name: "add comment", data: comment }));
      }
      setComment('');
    };

    return (
        <>
        <h2>Realtime BBS</h2>
        <form>
          <div>
            <p><Textarea comments={comments}/></p>
          </div>
          <div>
            <p>
                <Input value={comment} onChange={onChangeHandler} />
                <button onClick={submitHandler}>Send</button>
            </p>
          </div>
        </form>
      </>
    );
};

export default Bbs;
  