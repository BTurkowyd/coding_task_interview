import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Bikes from './components/Bikes/Bikes';
import Bike from './components/Bike/Bike';
import Login from './components/Login/Login';
import Logout from './components/Logout/Logout';
import Return from './components/Return/Return';

function App() {

    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/logout" element={<Logout />} />
                    <Route path="/bikes" element={<Bikes />} />
                    <Route path="/bike/:id" element={<Bike />} />
                    <Route path="/return/:id" element={<Return />} />
                </Routes>
            </BrowserRouter>
        </div>
    );
}

export default App;
