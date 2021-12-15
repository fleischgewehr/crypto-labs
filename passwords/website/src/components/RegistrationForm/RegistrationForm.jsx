import React, { useState } from 'react';

import { register } from '../../api/auth';

const RegistrationForm = () => {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();
    const [phone, setPhone] = useState();
    const [address, setAddress] = useState();
    const [err, setErr] = useState();

    const onSubmit = async (e) => {
        e.preventDefault();
        try {
            await register(username, password, phone, address);
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
            <h1>Registration</h1>
            <form onSubmit={onSubmit}>
                <label>
                    <p>Username</p>
                    <input type="text" onChange={onChangeUsername}/>
                </label>
                <label>
                    <p>Password</p>
                    <input type="password" onChange={onChangePassword}/>
                </label>
                <label>
                    <p>Phone number</p>
                    <input type="text" onChange={(e) => setPhone(e.target.value)}/>
                </label>
                <label>
                    <p>Address</p>
                    <input type="text" onChange={(e) => setAddress(e.target.value)}/>
                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
                <p>{err}</p>
            </form>
        </>
    );
}

export default RegistrationForm;
