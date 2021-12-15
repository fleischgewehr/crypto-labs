import axios from 'axios';

const load = async (id) => {
    return await axios.get(`http://localhost:8080/users/${id}`);
}

export {
    load,
};
