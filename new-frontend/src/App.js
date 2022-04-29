import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';
// import Nav from "./components/Nav";
import {BrowserRouter ,
  Routes,
  Route,
  Link} from "react-router-dom";
import AddPenyakit from "./pages/AddPenyakit";
import Home from "./pages/Home";

import SimpleCard from './pages/tambahPenyakit';
import Simple from './pages/cekHistory';
import CardWithIllustration from './pages/DnaTest';
import SimpleSidebar from './components/SideBar.tsx';

function App() {

  
  return (
    <div className="App">
        {/* <BrowserRouter>
            <main className="form-signin">
                <Route path="/" element={ <Home/>}/>
                <Route path="/login" element={ <AddPenyakit />}/>
                <Route path="/register" element={ <TesSakit />}/>
            </main>
        
        </BrowserRouter> */}
    {/* <h1>hello</h1>
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">History</Link>
            </li>
            <li>
              <Link to="/login">AddPenyakit</Link>
            </li>
            <li>
              <Link to="/register">tesSakit</Link>
            </li>
          </ul>
        </nav>

        <Routes>
          <Route path="/" component={<Home />} />
          <Route path="/login" component={<AddPenyakit />} />
          <Route path="/register" component={<TesSakit />} />
          
            
          
          {/* <Route path="/register" exact>
            <TesSakit />
          </Route> */}
          {/* <Route path="/">
            <Home />
          </Route> */}
        {/* </Routes> */}
      {/* </div> */}
    {/* </Router> */} 

        {/* <AddPenyakit mes={mes} setMes={setMes}/> */}

        {/* <main className="form-signin">
          <AddPenyakit />
        </main> */}
        {/* <SimpleSidebar /> */}
        {/* <Simple /> */}
        <SimpleCard />
        {/* <CardWithIllustration /> */}
        

        
      
    
    </div>
  );
}

export default App;
