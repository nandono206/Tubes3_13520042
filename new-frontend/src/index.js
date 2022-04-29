import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { ChakraProvider } from '@chakra-ui/react'
import { BrowserRouter,Routes,
  Route, } from "react-router-dom";
import { render } from "react-dom";
import SimpleSidebar from './components/SideBar.tsx';
import AddPenyakit from './pages/tambahPenyakit';
import DnaTest from './pages/DnaTest';
import Search from './pages/cekHistory';
import Home from './pages/Home';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <ChakraProvider>
      <BrowserRouter>
      <Routes>
      <Route path="/" exact element={<SimpleSidebar children=<DnaTest /> />} />
      <Route path="/history" element={<SimpleSidebar children=<Search/> />}/>
      <Route path="/addsakit" element={<SimpleSidebar children=<AddPenyakit /> />}/>
      <Route path="/sidebar" element={<SimpleSidebar children=<Home /> />} />
      
    </Routes>
      </BrowserRouter>
      
    </ChakraProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
