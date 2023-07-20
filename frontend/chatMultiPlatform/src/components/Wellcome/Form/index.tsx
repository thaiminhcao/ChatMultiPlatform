import React, { useState } from 'react';

const Index = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSignIn = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    fetch('http://localhost:8888/users/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    })
      .then(response => response.json())
      .then(data => {
        // Xử lý kết quả từ API tại đây
        console.log(data.message);
      })
      .catch(error => {
        // Xử lý lỗi tại đây
        console.error(error);
      });
  };

  return (
    <div className="flex items-center justify-center h-screen">
      <div className="w-full max-w-md">
        <h1 className="text-4xl font-semibold mb-8 text-center">Chat Multiple Platform</h1>
        <div className="bg-white shadow-lg rounded-lg px-12 py-10">
          <form className="space-y-6" onSubmit={handleSignIn}>
            <div>
              <label htmlFor="username" className="block text-sm font-medium text-gray-700">Username</label>
              <input
                type="text"
                id="username"
                value={username}
                onChange={e => setUsername(e.target.value)}
                className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                placeholder="Enter your username"
              />
            </div>
            <div>
              <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
              <input
                type="password"
                id="password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                placeholder="Enter your password"
              />
            </div>
            <button
              type="submit"
              className="w-full py-3 px-6 bg-blue-500 text-white rounded-md hover:bg-blue-600"
            >
              Sign In
            </button>
          </form>
          <div className="text-center mt-4">
            <button
              onClick={() => { window.location.href = '/sign-in'; }}
              className="text-blue-500 hover:underline"
            >
              Don't have an account? Sign up
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Index;
