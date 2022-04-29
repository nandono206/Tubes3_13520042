import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';
import Nav from "./components/Nav";
import {BrowserRouter, Route} from "react-router-dom";
import AddPenyakit from "./pages/AddPenyakit";
import Home from "./pages/Home";
import TesSakit from "./pages/TesSakit";

function App() {

  const [mes, setMes] = useState('');
  return (
    <div className="App">
        {/* <BrowserRouter>
            <main className="form-signin">
                <Route path="/" element={ <Home/>}/>
                <Route path="/login" element={ <Login mes={mes} setMes={setMes}/>}/>
                <Route path="/register" element={ <Register mes={mes} setMes={setMes}/>}/>
            </main>
        
        </BrowserRouter> */}

        <AddPenyakit mes={mes} setMes={setMes}/>
        {/* <TesSakit mes={mes} setMes={setMes}/> */}
      
    
    </div>
  );
}

export default App;
