import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import './index.css'
import Home from './components/Home.tsx'
import KubePage from './components/KubePage.tsx'

// const basename = '/frontend'; // Set the basename to match the subpath in Ingress

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router 
    // basename={basename}
    >
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/deployments" element={<KubePage typeOption='deployments' />} />
        <Route path="/pods" element={<KubePage typeOption='pods' />} />
        <Route path="/services" element={<KubePage typeOption='services' />} />
        <Route path="/configmaps" element={<KubePage typeOption='configmaps' />} />
        <Route path="/ingresses" element={<KubePage typeOption='ingress' />} />
        <Route path="/pvcs" element={<KubePage typeOption='pvc' />} />
        <Route path="/endpoints" element={<KubePage typeOption='endpoints' />} />
        <Route path="/nodes" element={<KubePage typeOption='nodes' />} />
        <Route path="/namespaces" element={<KubePage typeOption='namespaces' />} />
      </Routes>
    </Router>
  </StrictMode>,
)
