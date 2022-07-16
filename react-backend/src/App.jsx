import React, { useEffect } from 'react';
import {
  Routes,
  Route
} from 'react-router-dom';

import './css/style.scss';

import './charts/ChartjsConfig';

// Import pages
import Dashboard from './pages/Dashboard';
import Login from './pages/auth/Login';

function App() {

  // const location = useLocation();

  // useEffect(() => {
  //   // document.querySelector('html').style.scrollBehavior = 'auto'
  //   // window.scroll({ top: 0 })
  //   // document.querySelector('html').style.scrollBehavior = ''
  // }, [location.pathname]); // triggered on route change

  return (
    <>
      <Routes>
        <Route exact path="/admin/login" element={<Login />} />
        <Route exact path="/" element={<Dashboard />} />
      </Routes>
    </>
  );
}

export default App;
