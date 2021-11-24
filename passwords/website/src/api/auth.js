import axios from 'axios';

const login = async (username, password) => {
    return await axios.post('http://localhost:8080/users/login', { username, password });
}

const register = async (username, password) => {
    return await axios.post('http://localhost:8080/users', { username, password });
}

export {
    login,
    register,
};
