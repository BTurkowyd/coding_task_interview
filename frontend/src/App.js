import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Bikes from './components/Bikes/Bikes';
import Login from './components/Login/Login';

function App() {

    return (
        <div className="App">
            <h1>Bike renting system</h1>
            <BrowserRouter>
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/bikes" element={<Bikes />} />
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export default App;