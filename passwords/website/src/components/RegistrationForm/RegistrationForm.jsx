import React, { useState } from 'react';

import { register } from '../../api/auth';

const RegistrationForm = () => {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();

    const onSubmit = async (e) => {
        e.preventDefault();
        const resp = await register(username, password);
        if (resp) {
            window.alert('success');
        }
    }

    return (
        <>
            <h1>Registration</h1>
            <form onSubmit={onSubmit}>
                <label>
                    <p>Username</p>
                    <input type="text" onChange={e => setUsername(e.target.value)}/>
                </label>
                <label>
                    <p>Password</p>
                    <input type="password" onChange={e => setPassword(e.target.value)}/>
                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
            </form>
        </>
    );
}

export default RegistrationForm;
