import React, { useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';

const styles = {
  container: {
    margin: '20px',
    padding: '20px',
    maxWidth: '400px',
    border: '1px solid #ccc',
    borderRadius: '8px',
    backgroundColor: '#fff',
    boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)',
  },
  heading: {
    fontSize: '24px',
    fontWeight: 'bold',
    marginBottom: '20px',
    textAlign: 'center',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
  },
  inputContainer: {
    margin: '10px 0',
  },
  label: {
    marginBottom: '5px',
  },
  input: {
    padding: '10px',
    border: '1px solid #ccc',
    borderRadius: '5px',
  },
  errorMessage: {
    color: 'red',
    marginBottom: '10px',
  },
  registerButton: {
    padding: '10px',
    backgroundColor: '#007bff',
    color: '#fff',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
};

function RegisterForm() {
  const [formData, setFormData] = useState({ email: '', password: '' });
  const [error, setError] = useState(null);
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post('/api/register', formData);
      router.push('/login');
    } catch (error) {
      setError('Registration failed. Please try again.');
    }
  };

  return (
    <div style={styles.container}>
      <h2 style={styles.heading}>Register</h2>
      {error && <div style={styles.errorMessage}>{error}</div>}
      <form style={styles.form} onSubmit={handleSubmit}>
        <div style={styles.inputContainer}>
          <label style={styles.label} htmlFor="email">
            Email:
          </label>
          <input
            style={styles.input}
            type="email"
            id="email"
            name="email"
            value={formData.email}
            onChange={(e) => setFormData({ ...formData, email: e.target.value })}
            required
          />
        </div>
        <div style={styles.inputContainer}>
          <label style={styles.label} htmlFor="password">
            Password:
          </label>
          <input
            style={styles.input}
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            required
          />
        </div>
        <div>
          <button style={styles.registerButton} type="submit">
            Register
          </button>
        </div>
      </form>
    </div>
  );
}

export default RegisterForm;
