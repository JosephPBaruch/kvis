import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import './index.css'
import Home from './components/Home.tsx'
import KubePage from './components/KubePage.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/deployments" element={<KubePage typeOption='deployments' />} />
        <Route path="/pods" element={<KubePage typeOption='pods' />} />
      </Routes>
    </Router>
  </StrictMode>,
)
