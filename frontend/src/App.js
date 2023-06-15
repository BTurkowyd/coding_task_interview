import './App.css';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";

import Bikes from './components/Bikes/Bikes';
import Bike from './components/Bike/Bike';
import Login from './components/Login/Login';
import Logout from './components/Logout/Logout';
import Register from './components/Register/Register';

function App() {

    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Navigate to="/login" />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/logout" element={<Logout />} />
                    <Route path="/bikes" element={<Bikes />} />
                    <Route path="/bike/:id" element={<Bike />} />
                    <Route path="/register" element={<Register />} />
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export default App;
