import React from 'react';
import Link from 'next/link';
import LoginForm from '../components/LoginForm';

const Login = () => {
  return (
    <div>
      <h1>Login to Your Account</h1>
      <LoginForm />

      <p>
        Don't have an account?{' '}
        <Link href="/register">
          <a>Register here</a>
        </Link>
      </p>
    </div>
  );
};

export default Login;
