import React, { useState } from 'react';

import { login } from '../../api/auth';

const Login = () => {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();
    const [err, setErr] = useState();

    const onSubmit = async (e) => {
        e.preventDefault();
        try {
            await login(username, password);
        } catch (err) {
            setErr(err.response.data);
            return;
        }
        window.alert('success');
    }

    const onChangeUsername = (e) => {
        setUsername(e.target.value);
        setErr('');
    }

    const onChangePassword = (e) => {
        setPassword(e.target.value);
        setErr('');
    }

    return (
        <>
            <h1>Login</h1>
            <form onSubmit={onSubmit}>
                <label>
                    <p>Username</p>
                    <input type="text" onChange={onChangeUsername}/>
                </label>
                <label>
                    <p>Password</p>
                    <input type="password" onChange={onChangePassword}/>
                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
                <p>{err}</p>
            </form>
        </>
    );
}

export default Login;
