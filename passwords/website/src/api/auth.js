import axios from 'axios';

const login = async (username, password) => {
    try {
        return await axios.post('http://localhost:8080/users/login', { username, password });
    } catch (e) {}
}

const register = async (username, password) => {
    try {
        return await axios.post('http://localhost:8080/users', { username, password });
    } catch (e) {}
}

export {
    login,
    register,
};
