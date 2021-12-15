import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import { load } from '../../api/profile';

const Profile = () => {
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { id } = useParams();

    useEffect(async () => {
        try {
            const user = await load(id);
            console.log(user);
            setData(user);
        } catch (err) {
            setError(err.response.data);
        } finally {
            setLoading(false);
        }
    }, [id]);

    if (loading) {
        return <p>Loading...</p>
    }
    if (error) {
        return <p>An error occured while loading data: {error}</p>
    }

    return (
        <>
            <h1>{data.username} profile</h1>
            <p>Phone number: {data.phone}</p>
            <p>Address: {data.address}</p>
        </>
    );
}

export default Profile;
