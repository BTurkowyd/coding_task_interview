import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Bikes from './components/Bikes/bikes';
import Login from './components/Login/login';
import Logout from './components/Logout/logout';

function App() {

  return (
    <div className="App">
      <h1>Bike renting system</h1>
      <BrowserRouter>
        <Routes>
            <Route path="/login" element={<Login/>} />
            <Route path="/bikes" element={<Bikes/>} />
            <Route path='/logout' element={<Logout/>} />

        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
