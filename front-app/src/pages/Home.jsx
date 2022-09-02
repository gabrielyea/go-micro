import axios from 'axios';
import React, { useState } from 'react';

export default function Home() {
  const [data, setData] = useState();

  const buttonHandler = async () => {
    const res = await axios.get('http://localhost:8081/v1/');
    const {
      data: { msg },
    } = res;
    setData(msg);
    console.log(msg);
  };

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%',
        flexDirection: 'column',
      }}
    >
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          flexDirection: 'column',
          width: '100%',
        }}
      >
        <h1> Test Microservices </h1>
        <textarea name="" id="" cols="30" rows="10" />
      </div>

      <div
        style={{
          display: 'flex',
          width: '100%',
          justifyContent: 'space-around',
        }}
      >
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <h2>Sent</h2>
          <button type="button" onClick={buttonHandler}>
            click me
          </button>
          <textarea
            value={data}
            placeholder="nothing sent"
            name=""
            id=""
            cols="30"
            rows="10"
          />
        </div>
        <div
          style={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <h2>Received</h2>
          <textarea
            name=""
            id=""
            cols="30"
            rows="10"
            placeholder="nothing received"
          />
        </div>
      </div>
    </div>
  );
}
