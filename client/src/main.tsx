import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import KubePage from './KubePage.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/deployments" element={<KubePage typeOption='deployments' />} />
      </Routes>
    </Router>
  </StrictMode>,
)
