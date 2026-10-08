// import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './page/home'
import Reserve from './page/reserve'
import './App.css'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/reserve" element={<Reserve/>}/>
      </Routes>
    </BrowserRouter>
  )
}

export default App
