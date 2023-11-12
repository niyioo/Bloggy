import React, { useState } from 'react';
import Link from 'next/link';
import LoginForm from '../components/LoginForm';
import { loginUser } from '../redux/actions/userActions';
import { useRouter } from 'next/router';

const styles = {
  container: {
    margin: '20px auto',
    padding: '20px',
    maxWidth: '400px',
    border: '1px solid #ddd',
    borderRadius: '8px',
    backgroundColor: '#fff',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  },
  heading: {
    fontSize: '28px',
    fontWeight: 'bold',
    marginBottom: '20px',
    textAlign: 'center',
    color: '#333',
  },
  errorMessage: {
    color: 'red',
    marginBottom: '10px',
    textAlign: 'center',
  },
  registerLink: {
    textAlign: 'center',
    marginTop: '10px',
    fontSize: '16px',
    color: '#333',
  },
  registerText: {
    color: '#007bff',
    cursor: 'pointer',
  },
  formContainer: {
    margin: '20px 0',
  },
};

const Login = () => {
  const [error, setError] = useState('');
  const router = useRouter();

  const handleLogin = async (formData) => {
    try {
      await loginUser(formData);
      router.push('/dashboard');
    } catch (error) {
      setError('Login failed. Please check your credentials.');
    }
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>Login to Your Account</h1>
      {error && <p style={styles.errorMessage}>{error}</p>}

      <div style={styles.formContainer}>
        <LoginForm onLogin={handleLogin} />
      </div>

      <p style={styles.registerLink}>
        Don't have an account?{' '}
        <Link href="/register">
          <span style={styles.registerText}>Register here</span>
        </Link>
      </p>
    </div>
  );
};

export default Login;
