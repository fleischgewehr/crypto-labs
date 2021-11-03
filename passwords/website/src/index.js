import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import { render } from 'react-dom';

import App from './App'
import RegistrationForm from "./components/RegistrationForm/RegistrationForm";
import Login from "./components/Login/Login";

const rootElement = document.getElementById("root");
render(
    <BrowserRouter>
        <Link to="/">Homepage</Link>{" "}
        <Link to="/login">Login</Link>{" "}
        <Link to="/register">Register</Link>
        <Routes>
            <Route path="/" element={<App />}></Route>
            <Route path="/login" element={<Login />}></Route>
            <Route path="/register" element={<RegistrationForm />}></Route>
        </Routes>
    </BrowserRouter>,
    rootElement,
);
