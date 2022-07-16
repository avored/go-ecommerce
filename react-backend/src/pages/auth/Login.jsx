import React, { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

function Login() {
  const navigate = useNavigate()

  function handleOnLoginSubmit (e) {
    e.preventDefault()
    const endpointBaseUrl = 'http://localhost:8080' //import.meta.env.VITE_APP_BACKEND_BASE_URL

    axios({
        method: 'POST',
        url: endpointBaseUrl + "/admin/login",
        data: {email: 'admin@admin.com', password: 'admin123'},
    }).then(response => {
          navigate('/')
        })
        .catch(error => {
            console.error("There was an error!", error);
        });
  }


  return (
    <div className="flex h-screen bg-gray-100 items-center justify-center overflow-hidden">

      {/* Content area */}
      <div className="flex items-center justify-center shadow-md max-w-md bg-white rounded w-full space-y-8 p-12">
  
        <div className="mx-auto  w-full max-w-[550px]">
          <form action="#" onSubmit={handleOnLoginSubmit} method="POST">
           
            <div className="mb-5">
              <label
                htmlFor="email"
                className="mb-3 block text-base font-medium  text-[#07074D]"
              >
                Email Address
              </label>
              <input
                type="email"
                name="email"
                id="email"
                placeholder="example@domain.com"
                className="form-input w-full"
              />
            </div>
            <div className="mb-5">
              <label
                htmlFor="password"
                className="mb-3 block text-base font-medium text-[#07074D]"
              >
                Password
              </label>
              <input
                type="password"
                name="password"
                id="password"
                placeholder="Enter your password"
                className="form-input w-full"
              />
            </div>
            
            <div>
              <button
                className="rounded-md bg-red-500 py-2 px-6 text-base font-semibold text-white outline-none"
              >
                Submit
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

export default Login;