import React from 'react';
import Link from 'next/link';
import RegisterForm from '../components/RegisterForm';

const Register = () => {
  return (
    <div>
      <h1>Create Your Account</h1>
      <RegisterForm />
      <p>
        Already have an account?{' '}
        <Link href="/login">
          <a>Login here</a>
        </Link>
      </p>
    </div>
  );
};

export default Register;
