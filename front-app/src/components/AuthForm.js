import React, { useRef } from 'react';
import axios from 'axios';

function AuthForm() {
  const loginRef = useRef();
  const passRef = useRef();

  const btnHandler = async () => {
    const res = await axios.post(
      'http://localhost:8081/v1/auth',
      {
        action: 'auth',
        auth: {
          email: loginRef.current.value,
          password: passRef.current.value,
        },
      },
    );

    console.log(res, loginRef.current);
  };

  return (
    <div>
      <input ref={loginRef} type="text" placeholder="email" />
      <input ref={passRef} type="password" placeholder="password" />
      <button type="button" onClick={btnHandler}> Send login data </button>
    </div>
  );
}

export default AuthForm;
